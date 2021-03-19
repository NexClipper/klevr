// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agent

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

// TaskSendClient is the client API for TaskSend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskSendClient interface {
	SendTask(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	StatusCheck(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Status, error)
}

type taskSendClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskSendClient(cc grpc.ClientConnInterface) TaskSendClient {
	return &taskSendClient{cc}
}

func (c *taskSendClient) SendTask(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/agent.TaskSend/SendTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskSendClient) StatusCheck(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/agent.TaskSend/StatusCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskSendServer is the server API for TaskSend service.
// All implementations must embed UnimplementedTaskSendServer
// for forward compatibility
type TaskSendServer interface {
	SendTask(context.Context, *Message) (*Message, error)
	StatusCheck(context.Context, *Status) (*Status, error)
	//mustEmbedUnimplementedTaskSendServer()
}

// UnimplementedTaskSendServer must be embedded to have forward compatible implementations.
type UnimplementedTaskSendServer struct {
}

func (UnimplementedTaskSendServer) SendTask(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTask not implemented")
}
func (UnimplementedTaskSendServer) StatusCheck(context.Context, *Status) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusCheck not implemented")
}

//func (UnimplementedTaskSendServer) mustEmbedUnimplementedTaskSendServer() {}

// UnsafeTaskSendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskSendServer will
// result in compilation errors.
type UnsafeTaskSendServer interface {
	mustEmbedUnimplementedTaskSendServer()
}

func RegisterTaskSendServer(s grpc.ServiceRegistrar, srv TaskSendServer) {
	s.RegisterService(&TaskSend_ServiceDesc, srv)
}

func _TaskSend_SendTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskSendServer).SendTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.TaskSend/SendTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskSendServer).SendTask(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskSend_StatusCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskSendServer).StatusCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.TaskSend/StatusCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskSendServer).StatusCheck(ctx, req.(*Status))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskSend_ServiceDesc is the grpc.ServiceDesc for TaskSend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskSend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.TaskSend",
	HandlerType: (*TaskSendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTask",
			Handler:    _TaskSend_SendTask_Handler,
		},
		{
			MethodName: "StatusCheck",
			Handler:    _TaskSend_StatusCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/task.proto",
}
