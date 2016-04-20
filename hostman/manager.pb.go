// Code generated by protoc-gen-go.
// source: manager.proto
// DO NOT EDIT!

/*
Package hostman is a generated protocol buffer package.

It is generated from these files:
	manager.proto

It has these top-level messages:
	InstanceInfo
	InitReply
	InstanceKey
	StopReply
*/
package hostman

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "github.com/leveros/grpc-go"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type InstanceInfo struct {
	Environment       string `protobuf:"bytes,1,opt,name=environment" json:"environment,omitempty"`
	Service           string `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	InstanceID        string `protobuf:"bytes,3,opt,name=instanceID" json:"instanceID,omitempty"`
	ContainerID       string `protobuf:"bytes,4,opt,name=containerID" json:"containerID,omitempty"`
	ServingID         string `protobuf:"bytes,5,opt,name=servingID" json:"servingID,omitempty"`
	LevInstResourceID string `protobuf:"bytes,6,opt,name=levInstResourceID" json:"levInstResourceID,omitempty"`
	LevInstSessionID  string `protobuf:"bytes,7,opt,name=levInstSessionID" json:"levInstSessionID,omitempty"`
}

func (m *InstanceInfo) Reset()                    { *m = InstanceInfo{} }
func (m *InstanceInfo) String() string            { return proto.CompactTextString(m) }
func (*InstanceInfo) ProtoMessage()               {}
func (*InstanceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type InitReply struct {
}

func (m *InitReply) Reset()                    { *m = InitReply{} }
func (m *InitReply) String() string            { return proto.CompactTextString(m) }
func (*InitReply) ProtoMessage()               {}
func (*InitReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type InstanceKey struct {
	Environment string `protobuf:"bytes,1,opt,name=environment" json:"environment,omitempty"`
	Service     string `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	InstanceID  string `protobuf:"bytes,3,opt,name=instanceID" json:"instanceID,omitempty"`
	ServingID   string `protobuf:"bytes,4,opt,name=servingID" json:"servingID,omitempty"`
}

func (m *InstanceKey) Reset()                    { *m = InstanceKey{} }
func (m *InstanceKey) String() string            { return proto.CompactTextString(m) }
func (*InstanceKey) ProtoMessage()               {}
func (*InstanceKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StopReply struct {
}

func (m *StopReply) Reset()                    { *m = StopReply{} }
func (m *StopReply) String() string            { return proto.CompactTextString(m) }
func (*StopReply) ProtoMessage()               {}
func (*StopReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*InstanceInfo)(nil), "hostman.InstanceInfo")
	proto.RegisterType((*InitReply)(nil), "hostman.InitReply")
	proto.RegisterType((*InstanceKey)(nil), "hostman.InstanceKey")
	proto.RegisterType((*StopReply)(nil), "hostman.StopReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Manager service

type ManagerClient interface {
	InitializeInstance(ctx context.Context, in *InstanceInfo, opts ...grpc.CallOption) (*InitReply, error)
	StopInstance(ctx context.Context, in *InstanceKey, opts ...grpc.CallOption) (*StopReply, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) InitializeInstance(ctx context.Context, in *InstanceInfo, opts ...grpc.CallOption) (*InitReply, error) {
	out := new(InitReply)
	err := grpc.Invoke(ctx, "/hostman.Manager/InitializeInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) StopInstance(ctx context.Context, in *InstanceKey, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := grpc.Invoke(ctx, "/hostman.Manager/StopInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Manager service

type ManagerServer interface {
	InitializeInstance(context.Context, *InstanceInfo) (*InitReply, error)
	StopInstance(context.Context, *InstanceKey) (*StopReply, error)
}

func RegisterManagerServer(s *grpc.Server, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_InitializeInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).InitializeInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hostman.Manager/InitializeInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).InitializeInstance(ctx, req.(*InstanceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_StopInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).StopInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hostman.Manager/StopInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).StopInstance(ctx, req.(*InstanceKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hostman.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeInstance",
			Handler:    _Manager_InitializeInstance_Handler,
		},
		{
			MethodName: "StopInstance",
			Handler:    _Manager_StopInstance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x9b, 0xfe, 0xfb, 0x6f, 0xc8, 0xa4, 0x82, 0x0e, 0x0a, 0x41, 0x44, 0x4a, 0x4e, 0x22,
	0x92, 0x83, 0xde, 0xbc, 0x09, 0x5e, 0x82, 0x78, 0x69, 0x3f, 0x41, 0x0c, 0x63, 0x5d, 0x48, 0x67,
	0xc3, 0xee, 0x1a, 0xa8, 0x67, 0x41, 0x3f, 0xb6, 0x9b, 0x6d, 0xb7, 0xae, 0xe6, 0xec, 0x31, 0xbf,
	0x37, 0xf3, 0xf2, 0x66, 0x66, 0xe1, 0x60, 0x5d, 0x71, 0xb5, 0x22, 0x55, 0xb4, 0x4a, 0x1a, 0x89,
	0xf1, 0x8b, 0xd4, 0xc6, 0xa2, 0xfc, 0x7d, 0x0c, 0xb3, 0x92, 0xb5, 0xa9, 0xb8, 0xa6, 0x92, 0x9f,
	0x25, 0xce, 0x21, 0x25, 0xee, 0x84, 0x92, 0xbc, 0x26, 0x36, 0x59, 0x34, 0x8f, 0x2e, 0x92, 0x45,
	0x88, 0x30, 0x83, 0x58, 0x93, 0xea, 0x44, 0x4d, 0xd9, 0xd8, 0xa9, 0xfe, 0x13, 0xcf, 0x01, 0x84,
	0xf7, 0xba, 0xcf, 0xfe, 0x39, 0x31, 0x20, 0xbd, 0x77, 0x2d, 0xd9, 0x54, 0x82, 0x49, 0xd9, 0x82,
	0xc9, 0xd6, 0x3b, 0x40, 0x78, 0x06, 0x89, 0x33, 0xe3, 0x95, 0xd5, 0xff, 0x3b, 0xfd, 0x1b, 0xe0,
	0x15, 0x1c, 0x35, 0xd4, 0xf5, 0x71, 0x17, 0xa4, 0xe5, 0xab, 0x72, 0xbf, 0x99, 0xba, 0xaa, 0xa1,
	0x80, 0x97, 0x70, 0xb8, 0x83, 0x4b, 0xd2, 0x5a, 0x48, 0xb6, 0xc5, 0xb1, 0x2b, 0x1e, 0xf0, 0x3c,
	0x85, 0xa4, 0x64, 0x61, 0xbb, 0xdb, 0x66, 0x93, 0x7f, 0x44, 0x90, 0xfa, 0x9d, 0x3c, 0xd0, 0xe6,
	0x4f, 0x57, 0xf2, 0x63, 0xe0, 0xc9, 0xaf, 0x81, 0xfb, 0x58, 0x4b, 0x23, 0x5b, 0x17, 0xeb, 0xfa,
	0x33, 0x82, 0xf8, 0x71, 0x7b, 0x45, 0xbc, 0x03, 0xec, 0xf3, 0x8a, 0xaa, 0x11, 0x6f, 0xe4, 0xb3,
	0xe2, 0x49, 0xb1, 0x3b, 0x6b, 0x11, 0x9e, 0xf4, 0x14, 0x03, 0xec, 0x67, 0x1c, 0xe1, 0x2d, 0xcc,
	0x7a, 0xef, 0x7d, 0xf3, 0xf1, 0xa0, 0xd9, 0xce, 0x1e, 0xf4, 0xee, 0x83, 0xe4, 0xa3, 0xa7, 0xa9,
	0x7b, 0x45, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x1e, 0xa8, 0x96, 0x56, 0x02, 0x00,
	0x00,
}
