// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: protob/ecdsa-keygen.proto

package keygen

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
// Represents a BROADCAST message sent during Round 1 of the ECDSA TSS keygen protocol.
type KGRound1Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commitment []byte   `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	PaillierN  []byte   `protobuf:"bytes,2,opt,name=paillier_n,json=paillierN,proto3" json:"paillier_n,omitempty"`
	NTilde     []byte   `protobuf:"bytes,3,opt,name=n_tilde,json=nTilde,proto3" json:"n_tilde,omitempty"`
	H1         []byte   `protobuf:"bytes,4,opt,name=h1,proto3" json:"h1,omitempty"`
	H2         []byte   `protobuf:"bytes,5,opt,name=h2,proto3" json:"h2,omitempty"`
	Dlnproof_1 [][]byte `protobuf:"bytes,6,rep,name=dlnproof_1,json=dlnproof1,proto3" json:"dlnproof_1,omitempty"`
	Dlnproof_2 [][]byte `protobuf:"bytes,7,rep,name=dlnproof_2,json=dlnproof2,proto3" json:"dlnproof_2,omitempty"`
}

func (x *KGRound1Message) Reset() {
	*x = KGRound1Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_keygen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound1Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound1Message) ProtoMessage() {}

func (x *KGRound1Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_keygen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound1Message.ProtoReflect.Descriptor instead.
func (*KGRound1Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_keygen_proto_rawDescGZIP(), []int{0}
}

func (x *KGRound1Message) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

func (x *KGRound1Message) GetPaillierN() []byte {
	if x != nil {
		return x.PaillierN
	}
	return nil
}

func (x *KGRound1Message) GetNTilde() []byte {
	if x != nil {
		return x.NTilde
	}
	return nil
}

func (x *KGRound1Message) GetH1() []byte {
	if x != nil {
		return x.H1
	}
	return nil
}

func (x *KGRound1Message) GetH2() []byte {
	if x != nil {
		return x.H2
	}
	return nil
}

func (x *KGRound1Message) GetDlnproof_1() [][]byte {
	if x != nil {
		return x.Dlnproof_1
	}
	return nil
}

func (x *KGRound1Message) GetDlnproof_2() [][]byte {
	if x != nil {
		return x.Dlnproof_2
	}
	return nil
}

//
// Represents a P2P message sent to each party during Round 2 of the ECDSA TSS keygen protocol.
type KGRound2Message1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Share []byte `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
}

func (x *KGRound2Message1) Reset() {
	*x = KGRound2Message1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_keygen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound2Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound2Message1) ProtoMessage() {}

func (x *KGRound2Message1) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_keygen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound2Message1.ProtoReflect.Descriptor instead.
func (*KGRound2Message1) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_keygen_proto_rawDescGZIP(), []int{1}
}

func (x *KGRound2Message1) GetShare() []byte {
	if x != nil {
		return x.Share
	}
	return nil
}

//
// Represents a BROADCAST message sent to each party during Round 2 of the ECDSA TSS keygen protocol.
type KGRound2Message2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeCommitment [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
}

func (x *KGRound2Message2) Reset() {
	*x = KGRound2Message2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_keygen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound2Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound2Message2) ProtoMessage() {}

func (x *KGRound2Message2) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_keygen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound2Message2.ProtoReflect.Descriptor instead.
func (*KGRound2Message2) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_keygen_proto_rawDescGZIP(), []int{2}
}

func (x *KGRound2Message2) GetDeCommitment() [][]byte {
	if x != nil {
		return x.DeCommitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to each party during Round 3 of the ECDSA TSS keygen protocol.
type KGRound3Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaillierProof [][]byte `protobuf:"bytes,1,rep,name=paillier_proof,json=paillierProof,proto3" json:"paillier_proof,omitempty"`
}

func (x *KGRound3Message) Reset() {
	*x = KGRound3Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_keygen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound3Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound3Message) ProtoMessage() {}

func (x *KGRound3Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_keygen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound3Message.ProtoReflect.Descriptor instead.
func (*KGRound3Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_keygen_proto_rawDescGZIP(), []int{3}
}

func (x *KGRound3Message) GetPaillierProof() [][]byte {
	if x != nil {
		return x.PaillierProof
	}
	return nil
}

var File_protob_ecdsa_keygen_proto protoreflect.FileDescriptor

var file_protob_ecdsa_keygen_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2f, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2d, 0x6b,
	0x65, 0x79, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x0f,
	0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x31, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x4e, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x5f, 0x74, 0x69, 0x6c, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x6e, 0x54, 0x69, 0x6c, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x31, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x68, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x32, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x68, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6c, 0x6e, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x5f, 0x31, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x6c, 0x6e,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x31, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6c, 0x6e, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x5f, 0x32, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x6c, 0x6e, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x32, 0x22, 0x28, 0x0a, 0x10, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x22,
	0x37, 0x0a, 0x10, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x38, 0x0a, 0x0f, 0x4b, 0x47, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x42, 0x0e, 0x5a, 0x0c, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2f, 0x6b, 0x65, 0x79, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protob_ecdsa_keygen_proto_rawDescOnce sync.Once
	file_protob_ecdsa_keygen_proto_rawDescData = file_protob_ecdsa_keygen_proto_rawDesc
)

func file_protob_ecdsa_keygen_proto_rawDescGZIP() []byte {
	file_protob_ecdsa_keygen_proto_rawDescOnce.Do(func() {
		file_protob_ecdsa_keygen_proto_rawDescData = protoimpl.X.CompressGZIP(file_protob_ecdsa_keygen_proto_rawDescData)
	})
	return file_protob_ecdsa_keygen_proto_rawDescData
}

var file_protob_ecdsa_keygen_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protob_ecdsa_keygen_proto_goTypes = []interface{}{
	(*KGRound1Message)(nil),  // 0: KGRound1Message
	(*KGRound2Message1)(nil), // 1: KGRound2Message1
	(*KGRound2Message2)(nil), // 2: KGRound2Message2
	(*KGRound3Message)(nil),  // 3: KGRound3Message
}
var file_protob_ecdsa_keygen_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protob_ecdsa_keygen_proto_init() }
func file_protob_ecdsa_keygen_proto_init() {
	if File_protob_ecdsa_keygen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protob_ecdsa_keygen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGRound1Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protob_ecdsa_keygen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGRound2Message1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protob_ecdsa_keygen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGRound2Message2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protob_ecdsa_keygen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KGRound3Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protob_ecdsa_keygen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protob_ecdsa_keygen_proto_goTypes,
		DependencyIndexes: file_protob_ecdsa_keygen_proto_depIdxs,
		MessageInfos:      file_protob_ecdsa_keygen_proto_msgTypes,
	}.Build()
	File_protob_ecdsa_keygen_proto = out.File
	file_protob_ecdsa_keygen_proto_rawDesc = nil
	file_protob_ecdsa_keygen_proto_goTypes = nil
	file_protob_ecdsa_keygen_proto_depIdxs = nil
}
