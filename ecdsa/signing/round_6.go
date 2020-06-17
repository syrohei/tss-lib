// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package signing

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/hashicorp/go-multierror"

	"github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/crypto"
	"github.com/binance-chain/tss-lib/crypto/zkp"
	"github.com/binance-chain/tss-lib/tss"
)

func (round *round6) Start() *tss.Error {
	if round.started {
		return round.WrapError(errors.New("round already started"))
	}
	round.number = 6
	round.started = true
	round.resetOK()

	Pi := round.PartyID()
	i := Pi.Index

	// sigma_i can be discarded here; its only use is to be multiplied with r to become s_i = m*k + r*sigma_i
	sigmaI := round.temp.sigmaI
	round.temp.sigmaI = zero

	// bigR is stored as bytes for the OneRoundData protobuf struct
	bigRX, bigRY := new(big.Int).SetBytes(round.temp.BigR.GetX()), new(big.Int).SetBytes(round.temp.BigR.GetY())
	bigR := crypto.NewECPointNoCurveCheck(tss.EC(), bigRX, bigRY)

	errs := make(map[*tss.PartyID]error)
	bigRBarJProducts := (*crypto.ECPoint)(nil)
	BigRBarJ := make(map[string]*common.ECPoint, len(round.temp.signRound5Messages))
	for j, msg := range round.temp.signRound5Messages {
		Pj := round.Parties().IDs()[j]
		r5msg := msg.Content().(*SignRound5Message)
		bigRBarJ, err := r5msg.UnmarshalRI()
		if err != nil {
			errs[Pj] = err
			continue
		}
		BigRBarJ[Pj.Id] = bigRBarJ.ToProtobufPoint()

		// find products of all Rdash_i to ensure it equals the G point of the curve
		if bigRBarJProducts == nil {
			bigRBarJProducts = bigRBarJ
			continue
		}
		if bigRBarJProducts, err = bigRBarJProducts.Add(bigRBarJ); err != nil {
			errs[Pj] = err
			continue
		}

		if j == i {
			continue
		}
		// verify ZK proof of consistency between R_i and E_i(k_i)
		// ported from: https://git.io/Jf69a
		pdlWSlackPf, err := r5msg.UnmarshalPDLwSlackProof()
		if err != nil {
			errs[Pj] = err
			continue
		}
		r1msg1 := round.temp.signRound1Message1s[j].Content().(*SignRound1Message1)
		pdlWSlackStatement := zkp.PDLwSlackStatement{
			PK:         round.key.PaillierPKs[Pj.Index],
			CipherText: new(big.Int).SetBytes(r1msg1.GetC()),
			Q:          bigRBarJ,
			G:          bigR,
			H1:         round.key.H1j[Pj.Index],
			H2:         round.key.H2j[Pj.Index],
			NTilde:     round.key.NTildej[Pj.Index], // maybe i
		}
		if !pdlWSlackPf.Verify(pdlWSlackStatement) {
			errs[Pj] = fmt.Errorf("failed to verify ZK proof of consistency between R_i and E_i(k_i) for P %d", j)
		}
	}
	if 0 < len(errs) {
		var multiErr error
		culprits := make([]*tss.PartyID, 0, len(errs))
		for Pj, err := range errs {
			multiErr = multierror.Append(multiErr, err)
			culprits = append(culprits, Pj)
		}
		return round.WrapError(multiErr, culprits...)
	}
	{
		ec := tss.EC()
		gX, gY := ec.Params().Gx, ec.Params().Gy
		if bigRBarJProducts.X().Cmp(gX) != 0 || bigRBarJProducts.Y().Cmp(gY) != 0 {
			round.aborting = true
			common.Logger.Warnf("round 6: consistency check failed: g != R products, entering identify abort mode")

			r6msg := NewSignRound6MessageAbort(Pi, &round.temp.r5AbortData)
			round.temp.signRound6Messages[i] = r6msg
			round.out <- r6msg
			return nil
		}
	}
	round.temp.BigRBarJ = BigRBarJ

	TI, lI := round.temp.TI, round.temp.lI
	bigSI := bigR.ScalarMult(sigmaI)

	h, err := crypto.ECBasePoint2(tss.EC())
	if err != nil {
		return round.WrapError(err, Pi)
	}
	stProof, err := zkp.NewSTProof(TI, bigR, h, sigmaI, lI)
	if err != nil {
		return round.WrapError(err, Pi)
	}

	r6msg := NewSignRound6MessageSuccess(Pi, bigSI, stProof)
	round.temp.signRound6Messages[i] = r6msg
	round.out <- r6msg
	return nil
}

func (round *round6) Update() (bool, *tss.Error) {
	for j, msg := range round.temp.signRound6Messages {
		if round.ok[j] {
			continue
		}
		if msg == nil || !round.CanAccept(msg) {
			return false, nil
		}
		round.ok[j] = true
	}
	return true, nil
}

func (round *round6) CanAccept(msg tss.ParsedMessage) bool {
	if _, ok := msg.Content().(*SignRound6Message); ok {
		return msg.IsBroadcast()
	}
	return false
}

func (round *round6) NextRound() tss.Round {
	round.started = false
	return &round7{round}
}
