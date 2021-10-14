// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shiny_disco_api

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

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (Core_ConnectClient, error)
}

type coreClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreClient(cc grpc.ClientConnInterface) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (Core_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Core_ServiceDesc.Streams[0], "/shinydisco.api.v1.Core/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Core_ConnectClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type coreConnectClient struct {
	grpc.ClientStream
}

func (x *coreConnectClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CoreServer is the server API for Core service.
// All implementations should embed UnimplementedCoreServer
// for forward compatibility
type CoreServer interface {
	Connect(*ConnectRequest, Core_ConnectServer) error
}

// UnimplementedCoreServer should be embedded to have forward compatible implementations.
type UnimplementedCoreServer struct {
}

func (UnimplementedCoreServer) Connect(*ConnectRequest, Core_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

// UnsafeCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreServer will
// result in compilation errors.
type UnsafeCoreServer interface {
	mustEmbedUnimplementedCoreServer()
}

func RegisterCoreServer(s grpc.ServiceRegistrar, srv CoreServer) {
	s.RegisterService(&Core_ServiceDesc, srv)
}

func _Core_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreServer).Connect(m, &coreConnectServer{stream})
}

type Core_ConnectServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type coreConnectServer struct {
	grpc.ServerStream
}

func (x *coreConnectServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// Core_ServiceDesc is the grpc.ServiceDesc for Core service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Core_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shinydisco.api.v1.Core",
	HandlerType: (*CoreServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Core_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/na4ma4/shiny-disco-api/discovery.proto",
}

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (Server_DiscoverClient, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (Server_DiscoverClient, error) {
	stream, err := c.cc.NewStream(ctx, &Server_ServiceDesc.Streams[0], "/shinydisco.api.v1.Server/Discover", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverDiscoverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Server_DiscoverClient interface {
	Recv() (*Service, error)
	grpc.ClientStream
}

type serverDiscoverClient struct {
	grpc.ClientStream
}

func (x *serverDiscoverClient) Recv() (*Service, error) {
	m := new(Service)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerServer is the server API for Server service.
// All implementations should embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	Discover(*DiscoverRequest, Server_DiscoverServer) error
}

// UnimplementedServerServer should be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) Discover(*DiscoverRequest, Server_DiscoverServer) error {
	return status.Errorf(codes.Unimplemented, "method Discover not implemented")
}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DiscoverRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerServer).Discover(m, &serverDiscoverServer{stream})
}

type Server_DiscoverServer interface {
	Send(*Service) error
	grpc.ServerStream
}

type serverDiscoverServer struct {
	grpc.ServerStream
}

func (x *serverDiscoverServer) Send(m *Service) error {
	return x.ServerStream.SendMsg(m)
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shinydisco.api.v1.Server",
	HandlerType: (*ServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _Server_Discover_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/na4ma4/shiny-disco-api/discovery.proto",
}
