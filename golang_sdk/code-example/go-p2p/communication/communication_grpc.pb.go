// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: communication.proto

package communication

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

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	InitNode(ctx context.Context, opts ...grpc.CallOption) (Node_InitNodeClient, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) InitNode(ctx context.Context, opts ...grpc.CallOption) (Node_InitNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[0], "/Node/InitNode", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeInitNodeClient{stream}
	return x, nil
}

type Node_InitNodeClient interface {
	Send(*NodeReq) error
	Recv() (*NodeRes, error)
	grpc.ClientStream
}

type nodeInitNodeClient struct {
	grpc.ClientStream
}

func (x *nodeInitNodeClient) Send(m *NodeReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeInitNodeClient) Recv() (*NodeRes, error) {
	m := new(NodeRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	InitNode(Node_InitNodeServer) error
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) InitNode(Node_InitNodeServer) error {
	return status.Errorf(codes.Unimplemented, "method InitNode not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_InitNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).InitNode(&nodeInitNodeServer{stream})
}

type Node_InitNodeServer interface {
	Send(*NodeRes) error
	Recv() (*NodeReq, error)
	grpc.ServerStream
}

type nodeInitNodeServer struct {
	grpc.ServerStream
}

func (x *nodeInitNodeServer) Send(m *NodeRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeInitNodeServer) Recv() (*NodeReq, error) {
	m := new(NodeReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Node",
	HandlerType: (*NodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InitNode",
			Handler:       _Node_InitNode_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "communication.proto",
}
