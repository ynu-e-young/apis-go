// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: apis/htpp/machine/robot/v1/robot.proto

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

const (
	Robot_AppendCoordinate_FullMethodName = "/htpp.machine.robot.v1.Robot/AppendCoordinate"
	Robot_Zero_FullMethodName             = "/htpp.machine.robot.v1.Robot/Zero"
	Robot_GetMotorStatus_FullMethodName   = "/htpp.machine.robot.v1.Robot/GetMotorStatus"
)

// RobotClient is the client API for Robot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotClient interface {
	AppendCoordinate(ctx context.Context, in *CoordinateRequest, opts ...grpc.CallOption) (*CoordinateReply, error)
	Zero(ctx context.Context, in *ZeroRequest, opts ...grpc.CallOption) (*ZeroReply, error)
	GetMotorStatus(ctx context.Context, in *MotorInfoRequest, opts ...grpc.CallOption) (*MotorInfoReply, error)
}

type robotClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotClient(cc grpc.ClientConnInterface) RobotClient {
	return &robotClient{cc}
}

func (c *robotClient) AppendCoordinate(ctx context.Context, in *CoordinateRequest, opts ...grpc.CallOption) (*CoordinateReply, error) {
	out := new(CoordinateReply)
	err := c.cc.Invoke(ctx, Robot_AppendCoordinate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotClient) Zero(ctx context.Context, in *ZeroRequest, opts ...grpc.CallOption) (*ZeroReply, error) {
	out := new(ZeroReply)
	err := c.cc.Invoke(ctx, Robot_Zero_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotClient) GetMotorStatus(ctx context.Context, in *MotorInfoRequest, opts ...grpc.CallOption) (*MotorInfoReply, error) {
	out := new(MotorInfoReply)
	err := c.cc.Invoke(ctx, Robot_GetMotorStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotServer is the server API for Robot service.
// All implementations must embed UnimplementedRobotServer
// for forward compatibility
type RobotServer interface {
	AppendCoordinate(context.Context, *CoordinateRequest) (*CoordinateReply, error)
	Zero(context.Context, *ZeroRequest) (*ZeroReply, error)
	GetMotorStatus(context.Context, *MotorInfoRequest) (*MotorInfoReply, error)
	mustEmbedUnimplementedRobotServer()
}

// UnimplementedRobotServer must be embedded to have forward compatible implementations.
type UnimplementedRobotServer struct {
}

func (UnimplementedRobotServer) AppendCoordinate(context.Context, *CoordinateRequest) (*CoordinateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendCoordinate not implemented")
}
func (UnimplementedRobotServer) Zero(context.Context, *ZeroRequest) (*ZeroReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Zero not implemented")
}
func (UnimplementedRobotServer) GetMotorStatus(context.Context, *MotorInfoRequest) (*MotorInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMotorStatus not implemented")
}
func (UnimplementedRobotServer) mustEmbedUnimplementedRobotServer() {}

// UnsafeRobotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotServer will
// result in compilation errors.
type UnsafeRobotServer interface {
	mustEmbedUnimplementedRobotServer()
}

func RegisterRobotServer(s grpc.ServiceRegistrar, srv RobotServer) {
	s.RegisterService(&Robot_ServiceDesc, srv)
}

func _Robot_AppendCoordinate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoordinateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServer).AppendCoordinate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Robot_AppendCoordinate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServer).AppendCoordinate(ctx, req.(*CoordinateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robot_Zero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServer).Zero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Robot_Zero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServer).Zero(ctx, req.(*ZeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robot_GetMotorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MotorInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServer).GetMotorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Robot_GetMotorStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServer).GetMotorStatus(ctx, req.(*MotorInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Robot_ServiceDesc is the grpc.ServiceDesc for Robot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Robot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "htpp.machine.robot.v1.Robot",
	HandlerType: (*RobotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendCoordinate",
			Handler:    _Robot_AppendCoordinate_Handler,
		},
		{
			MethodName: "Zero",
			Handler:    _Robot_Zero_Handler,
		},
		{
			MethodName: "GetMotorStatus",
			Handler:    _Robot_GetMotorStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/htpp/machine/robot/v1/robot.proto",
}
