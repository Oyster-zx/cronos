// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cronos/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("cronos/query.proto", fileDescriptor_d4ed0fd688c48372) }

var fileDescriptor_d4ed0fd688c48372 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xce, 0x31, 0x0e, 0x82, 0x30,
	0x14, 0x06, 0x60, 0x18, 0xc4, 0x84, 0x91, 0x91, 0x98, 0x1e, 0xc0, 0x04, 0x5e, 0xd0, 0x1b, 0xb8,
	0x3a, 0xb9, 0xba, 0xb5, 0x4d, 0x53, 0x9a, 0x48, 0x5f, 0x6d, 0x1f, 0x46, 0x6e, 0xe1, 0xb1, 0x1c,
	0x19, 0x1d, 0x0d, 0x5c, 0xc4, 0x48, 0x75, 0x7b, 0x79, 0xf9, 0xfe, 0x3f, 0x7f, 0x5e, 0x48, 0x8f,
	0x16, 0x03, 0x5c, 0x7b, 0xe5, 0x87, 0xda, 0x79, 0x24, 0x2c, 0xb2, 0xf8, 0x2b, 0x37, 0x1a, 0x51,
	0x5f, 0x14, 0x70, 0x67, 0x80, 0x5b, 0x8b, 0xc4, 0xc9, 0xa0, 0x0d, 0x51, 0x95, 0x5b, 0x89, 0xa1,
	0xc3, 0x00, 0x82, 0x07, 0x15, 0xe3, 0x70, 0x6b, 0x84, 0x22, 0xde, 0x80, 0xe3, 0xda, 0xd8, 0x05,
	0x47, 0xbb, 0x5b, 0xe7, 0xab, 0xd3, 0x57, 0x1c, 0x8e, 0xcf, 0x89, 0xa5, 0xe3, 0xc4, 0xd2, 0xf7,
	0xc4, 0xd2, 0xc7, 0xcc, 0x92, 0x71, 0x66, 0xc9, 0x6b, 0x66, 0xc9, 0xb9, 0xd1, 0x86, 0xda, 0x5e,
	0xd4, 0x12, 0x3b, 0x90, 0x7e, 0x70, 0x84, 0x15, 0x7a, 0x5d, 0xc9, 0x96, 0x1b, 0x0b, 0xbf, 0x91,
	0xf7, 0xff, 0x41, 0x83, 0x53, 0x41, 0x64, 0x4b, 0xf9, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0x8b, 0xc0, 0x32, 0xc4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cronos.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cronos/query.proto",
}
