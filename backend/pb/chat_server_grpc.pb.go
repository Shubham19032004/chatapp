// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: chat_server.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChatApp_HelloServer_FullMethodName = "/pb.ChatApp/HelloServer"
)

// ChatAppClient is the client API for ChatApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatAppClient interface {
	HelloServer(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Send, Receive], error)
}

type chatAppClient struct {
	cc grpc.ClientConnInterface
}

func NewChatAppClient(cc grpc.ClientConnInterface) ChatAppClient {
	return &chatAppClient{cc}
}

func (c *chatAppClient) HelloServer(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Send, Receive], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatApp_ServiceDesc.Streams[0], ChatApp_HelloServer_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Send, Receive]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatApp_HelloServerClient = grpc.BidiStreamingClient[Send, Receive]

// ChatAppServer is the server API for ChatApp service.
// All implementations must embed UnimplementedChatAppServer
// for forward compatibility.
type ChatAppServer interface {
	HelloServer(grpc.BidiStreamingServer[Send, Receive]) error
	mustEmbedUnimplementedChatAppServer()
}

// UnimplementedChatAppServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatAppServer struct{}

func (UnimplementedChatAppServer) HelloServer(grpc.BidiStreamingServer[Send, Receive]) error {
	return status.Errorf(codes.Unimplemented, "method HelloServer not implemented")
}
func (UnimplementedChatAppServer) mustEmbedUnimplementedChatAppServer() {}
func (UnimplementedChatAppServer) testEmbeddedByValue()                 {}

// UnsafeChatAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatAppServer will
// result in compilation errors.
type UnsafeChatAppServer interface {
	mustEmbedUnimplementedChatAppServer()
}

func RegisterChatAppServer(s grpc.ServiceRegistrar, srv ChatAppServer) {
	// If the following call pancis, it indicates UnimplementedChatAppServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatApp_ServiceDesc, srv)
}

func _ChatApp_HelloServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatAppServer).HelloServer(&grpc.GenericServerStream[Send, Receive]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatApp_HelloServerServer = grpc.BidiStreamingServer[Send, Receive]

// ChatApp_ServiceDesc is the grpc.ServiceDesc for ChatApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ChatApp",
	HandlerType: (*ChatAppServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloServer",
			Handler:       _ChatApp_HelloServer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat_server.proto",
}
