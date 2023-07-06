// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: service.proto

package proto

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

// AureliusAgentManagerClient is the client API for AureliusAgentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AureliusAgentManagerClient interface {
	AgentChat(ctx context.Context, opts ...grpc.CallOption) (AureliusAgentManager_AgentChatClient, error)
}

type aureliusAgentManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewAureliusAgentManagerClient(cc grpc.ClientConnInterface) AureliusAgentManagerClient {
	return &aureliusAgentManagerClient{cc}
}

func (c *aureliusAgentManagerClient) AgentChat(ctx context.Context, opts ...grpc.CallOption) (AureliusAgentManager_AgentChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &AureliusAgentManager_ServiceDesc.Streams[0], "/AureliusAgentManager/AgentChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &aureliusAgentManagerAgentChatClient{stream}
	return x, nil
}

type AureliusAgentManager_AgentChatClient interface {
	Send(*ChatRequest) error
	Recv() (*AgentChatResponse, error)
	grpc.ClientStream
}

type aureliusAgentManagerAgentChatClient struct {
	grpc.ClientStream
}

func (x *aureliusAgentManagerAgentChatClient) Send(m *ChatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aureliusAgentManagerAgentChatClient) Recv() (*AgentChatResponse, error) {
	m := new(AgentChatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AureliusAgentManagerServer is the server API for AureliusAgentManager service.
// All implementations must embed UnimplementedAureliusAgentManagerServer
// for forward compatibility
type AureliusAgentManagerServer interface {
	AgentChat(AureliusAgentManager_AgentChatServer) error
	mustEmbedUnimplementedAureliusAgentManagerServer()
}

// UnimplementedAureliusAgentManagerServer must be embedded to have forward compatible implementations.
type UnimplementedAureliusAgentManagerServer struct {
}

func (UnimplementedAureliusAgentManagerServer) AgentChat(AureliusAgentManager_AgentChatServer) error {
	return status.Errorf(codes.Unimplemented, "method AgentChat not implemented")
}
func (UnimplementedAureliusAgentManagerServer) mustEmbedUnimplementedAureliusAgentManagerServer() {}

// UnsafeAureliusAgentManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AureliusAgentManagerServer will
// result in compilation errors.
type UnsafeAureliusAgentManagerServer interface {
	mustEmbedUnimplementedAureliusAgentManagerServer()
}

func RegisterAureliusAgentManagerServer(s grpc.ServiceRegistrar, srv AureliusAgentManagerServer) {
	s.RegisterService(&AureliusAgentManager_ServiceDesc, srv)
}

func _AureliusAgentManager_AgentChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AureliusAgentManagerServer).AgentChat(&aureliusAgentManagerAgentChatServer{stream})
}

type AureliusAgentManager_AgentChatServer interface {
	Send(*AgentChatResponse) error
	Recv() (*ChatRequest, error)
	grpc.ServerStream
}

type aureliusAgentManagerAgentChatServer struct {
	grpc.ServerStream
}

func (x *aureliusAgentManagerAgentChatServer) Send(m *AgentChatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aureliusAgentManagerAgentChatServer) Recv() (*ChatRequest, error) {
	m := new(ChatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AureliusAgentManager_ServiceDesc is the grpc.ServiceDesc for AureliusAgentManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AureliusAgentManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AureliusAgentManager",
	HandlerType: (*AureliusAgentManagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AgentChat",
			Handler:       _AureliusAgentManager_AgentChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
