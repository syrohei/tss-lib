// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package crypto

import (
	"bytes"
	"crypto/elliptic"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"

	"github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/tss"
)

// ECPoint convenience helper
type ECPoint struct {
	curve        elliptic.Curve
	coords       [2]*big.Int
	onCurveKnown bool
}

// Creates a new ECPoint and checks that the given coordinates are on the elliptic curve.
func NewECPoint(curve elliptic.Curve, X, Y *big.Int) (*ECPoint, error) {
	if !isOnCurve(curve, X, Y) {
		return nil, fmt.Errorf("NewECPoint: the given point is not on the elliptic curve")
	}
	return &ECPoint{curve, [2]*big.Int{X, Y}, true}, nil
}

// Creates a new ECPoint without checking that the coordinates are on the elliptic curve.
// Only use this function when you are completely sure that the point is already on the curve.
func NewECPointNoCurveCheck(curve elliptic.Curve, X, Y *big.Int) *ECPoint {
	return &ECPoint{curve, [2]*big.Int{X, Y}, false}
}

func (p *ECPoint) X() *big.Int {
	return new(big.Int).Set(p.coords[0])
}

func (p *ECPoint) Y() *big.Int {
	return new(big.Int).Set(p.coords[1])
}

func (p *ECPoint) Add(p1 *ECPoint) (*ECPoint, error) {
	x, y := p.curve.Add(p.X(), p.Y(), p1.X(), p1.Y())
	return NewECPoint(p.curve, x, y)
}

func (p *ECPoint) ScalarMult(k *big.Int) *ECPoint {
	x, y := p.curve.ScalarMult(p.X(), p.Y(), k.Bytes())
	newP, _ := NewECPoint(p.curve, x, y) // it must be on the curve, no need to check.
	return newP
}

func (p *ECPoint) IsOnCurve() bool {
	return isOnCurve(p.curve, p.coords[0], p.coords[1])
}

func (p *ECPoint) Equals(p2 *ECPoint) bool {
	if p == nil || p2 == nil {
		return false
	}
	return p.X().Cmp(p2.X()) == 0 && p.Y().Cmp(p2.Y()) == 0
}

func (p *ECPoint) SetCurve(curve elliptic.Curve) *ECPoint {
	p.curve = curve
	return p
}

func (p *ECPoint) ValidateBasic() bool {
	res := p != nil && p.coords[0] != nil && p.coords[1] != nil && (p.onCurveKnown || p.IsOnCurve())
	if res && !p.onCurveKnown {
		p.onCurveKnown = true
	}
	return res
}

func (p *ECPoint) ToProtobufPoint() *common.ECPoint {
	return &common.ECPoint{
		X: p.X().Bytes(),
		Y: p.Y().Bytes(),
	}
}

// ----- //

func isOnCurve(c elliptic.Curve, x, y *big.Int) bool {
	if x == nil || y == nil {
		return false
	}
	return c.IsOnCurve(x, y)
}

func ScalarBaseMult(curve elliptic.Curve, k *big.Int) *ECPoint {
	x, y := curve.ScalarBaseMult(k.Bytes())
	p, _ := NewECPoint(curve, x, y) // it must be on the curve, no need to check.
	return p
}

func DecompressPoint(curve elliptic.Curve, x *big.Int, sign byte) (*ECPoint, error) {
	if curve == nil || x == nil {
		return nil, errors.New("DecompressPoint() received one or more nil args")
	}
	switch curve {
	case btcec.S256():
		return decompressPoint_Secp256k1(curve, x, sign)
	case elliptic.P256():
		return decompressPoint_P256(curve, x, sign)
	default:
		return nil, fmt.Errorf("DecompressPoint() unsupported curve provided; please implement DecompressPoint for that curve")
	}
}

func decompressPoint_Secp256k1(curve elliptic.Curve, x *big.Int, sign byte) (*ECPoint, error) {
	params := curve.Params()
	modP := common.ModInt(params.P)

	// secp256k1: y^2 = x^3 + 7
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)

	y2 := x3.Add(x3, big.NewInt(7))
	// y2.Mod(y2, params.P)

	// find the sq root mod P
	y := modP.Sqrt(y2)
	if y == nil {
		return nil, errors.New("DecompressPoint() invalid point")
	}
	if y.Bit(0) != uint(sign)&1 {
		y = modP.Neg(y)
	}
	return &ECPoint{
		curve:  curve,
		coords: [2]*big.Int{x, y},
	}, nil
}

// Adapted from IsOnCurve from the stdlib: https://golang.org/src/crypto/elliptic/elliptic.go?s=2055:2110#L45
// With an extra modular square root to recover the Y co-ord
// It's only implemented for secp256k1, secp256r1 and P256 curves for now (ECDSA only)
func decompressPoint_P256(curve elliptic.Curve, x *big.Int, sign byte) (*ECPoint, error) {
	params := curve.Params()
	modP := common.ModInt(params.P)
	three := big.NewInt(3)

	// P-256/secp256r1/prime256v1: y^2 = x^3 - 3x + b
	x3 := modP.Exp(x, three)
	threeX := modP.Mul(x, three)

	// x^3 - 3x
	y2 := new(big.Int).Sub(x3, threeX)
	// .. + b mod P
	y2 = modP.Add(y2, params.B)

	// find the sq root mod P
	y := modP.Sqrt(y2)
	if y == nil {
		return nil, errors.New("DecompressPoint() invalid point")
	}
	if y.Bit(0) != uint(sign)&1 {
		y = modP.Neg(y)
	}
	return &ECPoint{
		curve:  curve,
		coords: [2]*big.Int{x, y},
	}, nil
}

// ----- //

func FlattenECPoints(in []*ECPoint) ([]*big.Int, error) {
	if in == nil {
		return nil, errors.New("FlattenECPoints encountered a nil in slice")
	}
	flat := make([]*big.Int, 0, len(in)*2)
	for _, point := range in {
		if point == nil || point.coords[0] == nil || point.coords[1] == nil {
			return nil, errors.New("FlattenECPoints found nil point/coordinate")
		}
		flat = append(flat, point.coords[0])
		flat = append(flat, point.coords[1])
	}
	return flat, nil
}

func UnFlattenECPoints(curve elliptic.Curve, in []*big.Int, noCurveCheck ...bool) ([]*ECPoint, error) {
	if in == nil || len(in)%2 != 0 {
		return nil, errors.New("UnFlattenECPoints expected an in len divisible by 2")
	}
	var err error
	unFlat := make([]*ECPoint, len(in)/2)
	for i, j := 0, 0; i < len(in); i, j = i+2, j+1 {
		if len(noCurveCheck) == 0 || !noCurveCheck[0] {
			unFlat[j], err = NewECPoint(curve, in[i], in[i+1])
			if err != nil {
				return nil, err
			}
		} else {
			unFlat[j] = NewECPointNoCurveCheck(curve, in[i], in[i+1])
		}
	}
	for _, point := range unFlat {
		if point.coords[0] == nil || point.coords[1] == nil {
			return nil, errors.New("UnFlattenECPoints found nil coordinate after unpack")
		}
	}
	return unFlat, nil
}

// ----- //
// Gob helpers for if you choose to encode messages with Gob.

func (p *ECPoint) GobEncode() ([]byte, error) {
	buf := &bytes.Buffer{}
	x, err := p.coords[0].GobEncode()
	if err != nil {
		return nil, err
	}
	y, err := p.coords[1].GobEncode()
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, uint32(len(x)))
	if err != nil {
		return nil, err
	}
	buf.Write(x)
	err = binary.Write(buf, binary.LittleEndian, uint32(len(y)))
	if err != nil {
		return nil, err
	}
	buf.Write(y)

	return buf.Bytes(), nil
}

func (p *ECPoint) GobDecode(buf []byte) error {
	reader := bytes.NewReader(buf)
	var length uint32
	if err := binary.Read(reader, binary.LittleEndian, &length); err != nil {
		return err
	}
	x := make([]byte, length)
	n, err := reader.Read(x)
	if n != int(length) || err != nil {
		return fmt.Errorf("gob decode failed: %v", err)
	}
	if err := binary.Read(reader, binary.LittleEndian, &length); err != nil {
		return err
	}
	y := make([]byte, length)
	n, err = reader.Read(y)
	if n != int(length) || err != nil {
		return fmt.Errorf("gob decode failed: %v", err)
	}

	X := new(big.Int)
	if err := X.GobDecode(x); err != nil {
		return err
	}
	Y := new(big.Int)
	if err := Y.GobDecode(y); err != nil {
		return err
	}
	p.curve = tss.EC()
	p.coords = [2]*big.Int{X, Y}
	if !p.IsOnCurve() {
		return errors.New("ECPoint.UnmarshalJSON: the point is not on the elliptic curve")
	}
	return nil
}

// ----- //

// crypto.ECPoint is not inherently json marshal-able
func (p *ECPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Coords [2]*big.Int
	}{
		Coords: p.coords,
	})
}

func (p *ECPoint) UnmarshalJSON(payload []byte) error {
	aux := &struct {
		Coords [2]*big.Int
	}{}
	if err := json.Unmarshal(payload, &aux); err != nil {
		return err
	}
	p.curve = tss.EC()
	p.coords = [2]*big.Int{aux.Coords[0], aux.Coords[1]}
	if !p.IsOnCurve() {
		return errors.New("ECPoint.UnmarshalJSON: the point is not on the elliptic curve")
	}
	return nil
}
