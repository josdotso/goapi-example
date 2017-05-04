// Code generated by protoc-gen-go.
// source: example/pb/app.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	example/pb/app.proto

It has these top-level messages:
	KeyValMessage
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyValMessage struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValMessage) Reset()                    { *m = KeyValMessage{} }
func (m *KeyValMessage) String() string            { return proto.CompactTextString(m) }
func (*KeyValMessage) ProtoMessage()               {}
func (*KeyValMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyValMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyValMessage)(nil), "pb.KeyValMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KeyVal service

type KeyValClient interface {
	KeyValCreate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
	KeyValRead(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
	KeyValUpdate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
	KeyValDelete(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error)
}

type keyValClient struct {
	cc *grpc.ClientConn
}

func NewKeyValClient(cc *grpc.ClientConn) KeyValClient {
	return &keyValClient{cc}
}

func (c *keyValClient) KeyValCreate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/pb.KeyVal/KeyValCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValClient) KeyValRead(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/pb.KeyVal/KeyValRead", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValClient) KeyValUpdate(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/pb.KeyVal/KeyValUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValClient) KeyValDelete(ctx context.Context, in *KeyValMessage, opts ...grpc.CallOption) (*KeyValMessage, error) {
	out := new(KeyValMessage)
	err := grpc.Invoke(ctx, "/pb.KeyVal/KeyValDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyVal service

type KeyValServer interface {
	KeyValCreate(context.Context, *KeyValMessage) (*KeyValMessage, error)
	KeyValRead(context.Context, *KeyValMessage) (*KeyValMessage, error)
	KeyValUpdate(context.Context, *KeyValMessage) (*KeyValMessage, error)
	KeyValDelete(context.Context, *KeyValMessage) (*KeyValMessage, error)
}

func RegisterKeyValServer(s *grpc.Server, srv KeyValServer) {
	s.RegisterService(&_KeyVal_serviceDesc, srv)
}

func _KeyVal_KeyValCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValServer).KeyValCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KeyVal/KeyValCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValServer).KeyValCreate(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyVal_KeyValRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValServer).KeyValRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KeyVal/KeyValRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValServer).KeyValRead(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyVal_KeyValUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValServer).KeyValUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KeyVal/KeyValUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValServer).KeyValUpdate(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyVal_KeyValDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValServer).KeyValDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.KeyVal/KeyValDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValServer).KeyValDelete(ctx, req.(*KeyValMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyVal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.KeyVal",
	HandlerType: (*KeyValServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeyValCreate",
			Handler:    _KeyVal_KeyValCreate_Handler,
		},
		{
			MethodName: "KeyValRead",
			Handler:    _KeyVal_KeyValRead_Handler,
		},
		{
			MethodName: "KeyValUpdate",
			Handler:    _KeyVal_KeyValUpdate_Handler,
		},
		{
			MethodName: "KeyValDelete",
			Handler:    _KeyVal_KeyValDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/pb/app.proto",
}

func init() { proto.RegisterFile("example/pb/app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0x48, 0xd2, 0x4f, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2a, 0x48, 0x92, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8,
	0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x50, 0x32,
	0xe7, 0xe2, 0xf5, 0x4e, 0xad, 0x0c, 0x4b, 0xcc, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15,
	0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85,
	0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0xd1, 0x19,
	0x26, 0x2e, 0x36, 0x88, 0x4e, 0xa1, 0x40, 0x2e, 0x1e, 0x08, 0xcb, 0xb9, 0x28, 0x35, 0xb1, 0x24,
	0x55, 0x48, 0x50, 0xaf, 0x20, 0x49, 0x0f, 0xc5, 0x54, 0x29, 0x4c, 0x21, 0x25, 0xe9, 0xa6, 0xcb,
	0x4f, 0x26, 0x33, 0x89, 0x4a, 0x09, 0xe8, 0x97, 0x19, 0xea, 0x67, 0xa7, 0x56, 0x96, 0x25, 0xe6,
	0xe8, 0x57, 0x67, 0xa7, 0x56, 0xd6, 0x5a, 0x31, 0x6a, 0x09, 0xf9, 0x70, 0x71, 0x41, 0x54, 0x07,
	0xa5, 0x26, 0xa6, 0x10, 0x69, 0xa0, 0x04, 0xd8, 0x40, 0x21, 0x21, 0x0c, 0x03, 0x11, 0x0e, 0x0c,
	0x2d, 0x48, 0x21, 0xd9, 0x81, 0x4a, 0x58, 0x1d, 0xe8, 0x07, 0x33, 0xd2, 0x25, 0x35, 0x27, 0x95,
	0x68, 0x23, 0xa1, 0x4e, 0xd4, 0xc2, 0x30, 0x32, 0x89, 0x0d, 0x1c, 0x1d, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe4, 0x04, 0xf2, 0x70, 0xc8, 0x01, 0x00, 0x00,
}