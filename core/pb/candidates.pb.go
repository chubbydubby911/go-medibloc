// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: candidates.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	candidates.proto

It has these top-level messages:
	Candidate
*/
package corepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Candidate struct {
	Address    []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Collateral []byte `protobuf:"bytes,2,opt,name=collateral,proto3" json:"collateral,omitempty"`
}

func (m *Candidate) Reset()                    { *m = Candidate{} }
func (m *Candidate) String() string            { return proto.CompactTextString(m) }
func (*Candidate) ProtoMessage()               {}
func (*Candidate) Descriptor() ([]byte, []int) { return fileDescriptorCandidates, []int{0} }

func (m *Candidate) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Candidate) GetCollateral() []byte {
	if m != nil {
		return m.Collateral
	}
	return nil
}

func init() {
	proto.RegisterType((*Candidate)(nil), "corepb.Candidate")
}

func init() { proto.RegisterFile("candidates.proto", fileDescriptorCandidates) }

var fileDescriptorCandidates = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0xcc, 0x4b,
	0xc9, 0x4c, 0x49, 0x2c, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce,
	0x2f, 0x4a, 0x2d, 0x48, 0x52, 0x72, 0xe5, 0xe2, 0x74, 0x86, 0xc9, 0x09, 0x49, 0x70, 0xb1, 0x27,
	0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8, 0x42,
	0x72, 0x5c, 0x5c, 0xc9, 0xf9, 0x39, 0x39, 0x89, 0x25, 0xa9, 0x45, 0x89, 0x39, 0x12, 0x4c, 0x60,
	0x49, 0x24, 0x91, 0x24, 0x36, 0xb0, 0xa9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0xc0,
	0xea, 0x79, 0x69, 0x00, 0x00, 0x00,
}
