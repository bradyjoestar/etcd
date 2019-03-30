// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: watchpd.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		watchpd.proto

	It has these top-level messages:
*/
package pb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	etcdserverpb "go.etcd.io/etcd/etcdserver/etcdserverpb"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Watch service

type WatchClient interface {
	// Sends a greeting
	Watch(ctx context.Context, opts ...grpc.CallOption) (Watch_WatchClient, error)
}

type watchClient struct {
	cc *grpc.ClientConn
}

func NewWatchClient(cc *grpc.ClientConn) WatchClient {
	return &watchClient{cc}
}

func (c *watchClient) Watch(ctx context.Context, opts ...grpc.CallOption) (Watch_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Watch_serviceDesc.Streams[0], c.cc, "/pb.Watch/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchWatchClient{stream}
	return x, nil
}

type Watch_WatchClient interface {
	Send(*etcdserverpb.WatchRequest) error
	Recv() (*etcdserverpb.WatchResponse, error)
	grpc.ClientStream
}

type watchWatchClient struct {
	grpc.ClientStream
}

func (x *watchWatchClient) Send(m *etcdserverpb.WatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *watchWatchClient) Recv() (*etcdserverpb.WatchResponse, error) {
	m := new(etcdserverpb.WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Watch service

type WatchServer interface {
	// Sends a greeting
	Watch(Watch_WatchServer) error
}

func RegisterWatchServer(s *grpc.Server, srv WatchServer) {
	s.RegisterService(&_Watch_serviceDesc, srv)
}

func _Watch_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WatchServer).Watch(&watchWatchServer{stream})
}

type Watch_WatchServer interface {
	Send(*etcdserverpb.WatchResponse) error
	Recv() (*etcdserverpb.WatchRequest, error)
	grpc.ServerStream
}

type watchWatchServer struct {
	grpc.ServerStream
}

func (x *watchWatchServer) Send(m *etcdserverpb.WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *watchWatchServer) Recv() (*etcdserverpb.WatchRequest, error) {
	m := new(etcdserverpb.WatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Watch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Watch",
	HandlerType: (*WatchServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Watch_Watch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "watchpd.proto",
}

func init() { proto.RegisterFile("watchpd.proto", fileDescriptorPd) }

var fileDescriptorPd = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x48, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x73,
	0xf5, 0x41, 0x2c, 0x88, 0x8c, 0x94, 0x5a, 0x6a, 0x49, 0x72, 0x8a, 0x3e, 0x88, 0x28, 0x4e, 0x2d,
	0x2a, 0x4b, 0x2d, 0x42, 0x62, 0x16, 0x24, 0xe9, 0x17, 0x15, 0x24, 0x43, 0xd4, 0x19, 0xf9, 0x73,
	0xb1, 0x86, 0x27, 0x96, 0x24, 0x67, 0x08, 0xb9, 0xc1, 0x18, 0x52, 0x7a, 0xc8, 0x4a, 0xf5, 0xc0,
	0x82, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x52, 0xd2, 0x58, 0xe5, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x95, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x9d, 0x24, 0x4e, 0x3c, 0x94, 0x63, 0xb8, 0xf0,
	0x50, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x6c, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x32, 0xfa,
	0xfc, 0x65, 0xbf, 0x00, 0x00, 0x00,
}
