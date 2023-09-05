// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/cotaparser/v1/cota.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CotaClient is the client API for Cota service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CotaClient interface {
	GetCotaEntries(ctx context.Context, in *GetCotaEntriesRequest, opts ...grpc.CallOption) (Cota_GetCotaEntriesClient, error)
	GetClassMetadata(ctx context.Context, in *GetClassMetadataRequest, opts ...grpc.CallOption) (Cota_GetClassMetadataClient, error)
	GetIssuerMetadata(ctx context.Context, in *GetIssuerMetadataRequest, opts ...grpc.CallOption) (Cota_GetIssuerMetadataClient, error)
}

type cotaClient struct {
	cc grpc.ClientConnInterface
}

func NewCotaClient(cc grpc.ClientConnInterface) CotaClient {
	return &cotaClient{cc}
}

func (c *cotaClient) GetCotaEntries(ctx context.Context, in *GetCotaEntriesRequest, opts ...grpc.CallOption) (Cota_GetCotaEntriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cota_ServiceDesc.Streams[0], "/cotaparser.api.v1.Cota/GetCotaEntries", opts...)
	if err != nil {
		return nil, err
	}
	x := &cotaGetCotaEntriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cota_GetCotaEntriesClient interface {
	Recv() (*GetCotaEntriesResponse, error)
	grpc.ClientStream
}

type cotaGetCotaEntriesClient struct {
	grpc.ClientStream
}

func (x *cotaGetCotaEntriesClient) Recv() (*GetCotaEntriesResponse, error) {
	m := new(GetCotaEntriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cotaClient) GetClassMetadata(ctx context.Context, in *GetClassMetadataRequest, opts ...grpc.CallOption) (Cota_GetClassMetadataClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cota_ServiceDesc.Streams[1], "/cotaparser.api.v1.Cota/GetClassMetadata", opts...)
	if err != nil {
		return nil, err
	}
	x := &cotaGetClassMetadataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cota_GetClassMetadataClient interface {
	Recv() (*GetClassMetadataResponse, error)
	grpc.ClientStream
}

type cotaGetClassMetadataClient struct {
	grpc.ClientStream
}

func (x *cotaGetClassMetadataClient) Recv() (*GetClassMetadataResponse, error) {
	m := new(GetClassMetadataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cotaClient) GetIssuerMetadata(ctx context.Context, in *GetIssuerMetadataRequest, opts ...grpc.CallOption) (Cota_GetIssuerMetadataClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cota_ServiceDesc.Streams[2], "/cotaparser.api.v1.Cota/GetIssuerMetadata", opts...)
	if err != nil {
		return nil, err
	}
	x := &cotaGetIssuerMetadataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cota_GetIssuerMetadataClient interface {
	Recv() (*GetIssuerMetadataResponse, error)
	grpc.ClientStream
}

type cotaGetIssuerMetadataClient struct {
	grpc.ClientStream
}

func (x *cotaGetIssuerMetadataClient) Recv() (*GetIssuerMetadataResponse, error) {
	m := new(GetIssuerMetadataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CotaServer is the server API for Cota service.
// All implementations must embed UnimplementedCotaServer
// for forward compatibility
type CotaServer interface {
	GetCotaEntries(*GetCotaEntriesRequest, Cota_GetCotaEntriesServer) error
	GetClassMetadata(*GetClassMetadataRequest, Cota_GetClassMetadataServer) error
	GetIssuerMetadata(*GetIssuerMetadataRequest, Cota_GetIssuerMetadataServer) error
	mustEmbedUnimplementedCotaServer()
}

// UnimplementedCotaServer must be embedded to have forward compatible implementations.
type UnimplementedCotaServer struct {
}

func (UnimplementedCotaServer) GetCotaEntries(*GetCotaEntriesRequest, Cota_GetCotaEntriesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCotaEntries not implemented")
}
func (UnimplementedCotaServer) GetClassMetadata(*GetClassMetadataRequest, Cota_GetClassMetadataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetClassMetadata not implemented")
}
func (UnimplementedCotaServer) GetIssuerMetadata(*GetIssuerMetadataRequest, Cota_GetIssuerMetadataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetIssuerMetadata not implemented")
}
func (UnimplementedCotaServer) mustEmbedUnimplementedCotaServer() {}

// UnsafeCotaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CotaServer will
// result in compilation errors.
type UnsafeCotaServer interface {
	mustEmbedUnimplementedCotaServer()
}

func RegisterCotaServer(s grpc.ServiceRegistrar, srv CotaServer) {
	s.RegisterService(&Cota_ServiceDesc, srv)
}

func _Cota_GetCotaEntries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCotaEntriesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CotaServer).GetCotaEntries(m, &cotaGetCotaEntriesServer{stream})
}

type Cota_GetCotaEntriesServer interface {
	Send(*GetCotaEntriesResponse) error
	grpc.ServerStream
}

type cotaGetCotaEntriesServer struct {
	grpc.ServerStream
}

func (x *cotaGetCotaEntriesServer) Send(m *GetCotaEntriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Cota_GetClassMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetClassMetadataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CotaServer).GetClassMetadata(m, &cotaGetClassMetadataServer{stream})
}

type Cota_GetClassMetadataServer interface {
	Send(*GetClassMetadataResponse) error
	grpc.ServerStream
}

type cotaGetClassMetadataServer struct {
	grpc.ServerStream
}

func (x *cotaGetClassMetadataServer) Send(m *GetClassMetadataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Cota_GetIssuerMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetIssuerMetadataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CotaServer).GetIssuerMetadata(m, &cotaGetIssuerMetadataServer{stream})
}

type Cota_GetIssuerMetadataServer interface {
	Send(*GetIssuerMetadataResponse) error
	grpc.ServerStream
}

type cotaGetIssuerMetadataServer struct {
	grpc.ServerStream
}

func (x *cotaGetIssuerMetadataServer) Send(m *GetIssuerMetadataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Cota_ServiceDesc is the grpc.ServiceDesc for Cota service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cota_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cotaparser.api.v1.Cota",
	HandlerType: (*CotaServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCotaEntries",
			Handler:       _Cota_GetCotaEntries_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetClassMetadata",
			Handler:       _Cota_GetClassMetadata_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetIssuerMetadata",
			Handler:       _Cota_GetIssuerMetadata_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/cotaparser/v1/cota.proto",
}
