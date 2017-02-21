// Code generated by protoc-gen-go.
// source: IOTRegistry.proto
// DO NOT EDIT!

/*
Package IOTRegistry is a generated protocol buffer package.

It is generated from these files:
	IOTRegistry.proto

It has these top-level messages:
	RegisterThingTX
	RegisterIdentityTx
*/
package IOTRegistry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegisterThingTX struct {
	Nonce      []byte   `protobuf:"bytes,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Identities []string `protobuf:"bytes,2,rep,name=Identities" json:"Identities,omitempty"`
	OwnerName  string   `protobuf:"bytes,3,opt,name=OwnerName" json:"OwnerName,omitempty"`
	Signature  []byte   `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (m *RegisterThingTX) Reset()         { *m = RegisterThingTX{} }
func (m *RegisterThingTX) String() string { return proto.CompactTextString(m) }
func (*RegisterThingTX) ProtoMessage()    {}

type RegisterIdentityTx struct {
	OwnerName string `protobuf:"bytes,1,opt,name=OwnerName" json:"OwnerName,omitempty"`
	PubKey    []byte `protobuf:"bytes,2,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	Data      string `protobuf:"bytes,3,opt,name=Data" json:"Data,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (m *RegisterIdentityTx) Reset()         { *m = RegisterIdentityTx{} }
func (m *RegisterIdentityTx) String() string { return proto.CompactTextString(m) }
func (*RegisterIdentityTx) ProtoMessage()    {}
