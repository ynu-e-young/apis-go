// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: apis/htpp/machine/service/v1/machine.proto

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
	Machine_FindByUserId_FullMethodName   = "/htpp.machine.service.v1.Machine/FindByUserId"
	Machine_Create_FullMethodName         = "/htpp.machine.service.v1.Machine/Create"
	Machine_Update_FullMethodName         = "/htpp.machine.service.v1.Machine/Update"
	Machine_Get_FullMethodName            = "/htpp.machine.service.v1.Machine/Get"
	Machine_Move_FullMethodName           = "/htpp.machine.service.v1.Machine/Move"
	Machine_Zero_FullMethodName           = "/htpp.machine.service.v1.Machine/Zero"
	Machine_GetMotorStatus_FullMethodName = "/htpp.machine.service.v1.Machine/GetMotorStatus"
	Machine_MoveDone_FullMethodName       = "/htpp.machine.service.v1.Machine/MoveDone"
	Machine_CreateCronJob_FullMethodName  = "/htpp.machine.service.v1.Machine/CreateCronJob"
	Machine_DeleteCronJob_FullMethodName  = "/htpp.machine.service.v1.Machine/DeleteCronJob"
	Machine_ListCronJob_FullMethodName    = "/htpp.machine.service.v1.Machine/ListCronJob"
)

// MachineClient is the client API for Machine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineClient interface {
	FindByUserId(ctx context.Context, in *FindByUserIdRequest, opts ...grpc.CallOption) (*MachinesReply, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*MachineReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MachineReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*MachineReply, error)
	Move(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveReply, error)
	Zero(ctx context.Context, in *ZeroRequest, opts ...grpc.CallOption) (*ZeroReply, error)
	GetMotorStatus(ctx context.Context, in *GetMotorStatusRequest, opts ...grpc.CallOption) (*GetMotorStatusReply, error)
	MoveDone(ctx context.Context, in *MoveDoneRequest, opts ...grpc.CallOption) (*MoveDoneReply, error)
	CreateCronJob(ctx context.Context, in *CreateCronJobRequest, opts ...grpc.CallOption) (*CronJobReply, error)
	DeleteCronJob(ctx context.Context, in *DeleteCronJobRequest, opts ...grpc.CallOption) (*DeleteCronJobReply, error)
	ListCronJob(ctx context.Context, in *ListCronJobRequest, opts ...grpc.CallOption) (*CronJobsReply, error)
}

type machineClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineClient(cc grpc.ClientConnInterface) MachineClient {
	return &machineClient{cc}
}

func (c *machineClient) FindByUserId(ctx context.Context, in *FindByUserIdRequest, opts ...grpc.CallOption) (*MachinesReply, error) {
	out := new(MachinesReply)
	err := c.cc.Invoke(ctx, Machine_FindByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*MachineReply, error) {
	out := new(MachineReply)
	err := c.cc.Invoke(ctx, Machine_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*MachineReply, error) {
	out := new(MachineReply)
	err := c.cc.Invoke(ctx, Machine_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*MachineReply, error) {
	out := new(MachineReply)
	err := c.cc.Invoke(ctx, Machine_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) Move(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveReply, error) {
	out := new(MoveReply)
	err := c.cc.Invoke(ctx, Machine_Move_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) Zero(ctx context.Context, in *ZeroRequest, opts ...grpc.CallOption) (*ZeroReply, error) {
	out := new(ZeroReply)
	err := c.cc.Invoke(ctx, Machine_Zero_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) GetMotorStatus(ctx context.Context, in *GetMotorStatusRequest, opts ...grpc.CallOption) (*GetMotorStatusReply, error) {
	out := new(GetMotorStatusReply)
	err := c.cc.Invoke(ctx, Machine_GetMotorStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) MoveDone(ctx context.Context, in *MoveDoneRequest, opts ...grpc.CallOption) (*MoveDoneReply, error) {
	out := new(MoveDoneReply)
	err := c.cc.Invoke(ctx, Machine_MoveDone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) CreateCronJob(ctx context.Context, in *CreateCronJobRequest, opts ...grpc.CallOption) (*CronJobReply, error) {
	out := new(CronJobReply)
	err := c.cc.Invoke(ctx, Machine_CreateCronJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) DeleteCronJob(ctx context.Context, in *DeleteCronJobRequest, opts ...grpc.CallOption) (*DeleteCronJobReply, error) {
	out := new(DeleteCronJobReply)
	err := c.cc.Invoke(ctx, Machine_DeleteCronJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineClient) ListCronJob(ctx context.Context, in *ListCronJobRequest, opts ...grpc.CallOption) (*CronJobsReply, error) {
	out := new(CronJobsReply)
	err := c.cc.Invoke(ctx, Machine_ListCronJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServer is the server API for Machine service.
// All implementations must embed UnimplementedMachineServer
// for forward compatibility
type MachineServer interface {
	FindByUserId(context.Context, *FindByUserIdRequest) (*MachinesReply, error)
	Create(context.Context, *CreateRequest) (*MachineReply, error)
	Update(context.Context, *UpdateRequest) (*MachineReply, error)
	Get(context.Context, *GetRequest) (*MachineReply, error)
	Move(context.Context, *MoveRequest) (*MoveReply, error)
	Zero(context.Context, *ZeroRequest) (*ZeroReply, error)
	GetMotorStatus(context.Context, *GetMotorStatusRequest) (*GetMotorStatusReply, error)
	MoveDone(context.Context, *MoveDoneRequest) (*MoveDoneReply, error)
	CreateCronJob(context.Context, *CreateCronJobRequest) (*CronJobReply, error)
	DeleteCronJob(context.Context, *DeleteCronJobRequest) (*DeleteCronJobReply, error)
	ListCronJob(context.Context, *ListCronJobRequest) (*CronJobsReply, error)
	mustEmbedUnimplementedMachineServer()
}

// UnimplementedMachineServer must be embedded to have forward compatible implementations.
type UnimplementedMachineServer struct {
}

func (UnimplementedMachineServer) FindByUserId(context.Context, *FindByUserIdRequest) (*MachinesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserId not implemented")
}
func (UnimplementedMachineServer) Create(context.Context, *CreateRequest) (*MachineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMachineServer) Update(context.Context, *UpdateRequest) (*MachineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMachineServer) Get(context.Context, *GetRequest) (*MachineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMachineServer) Move(context.Context, *MoveRequest) (*MoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedMachineServer) Zero(context.Context, *ZeroRequest) (*ZeroReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Zero not implemented")
}
func (UnimplementedMachineServer) GetMotorStatus(context.Context, *GetMotorStatusRequest) (*GetMotorStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMotorStatus not implemented")
}
func (UnimplementedMachineServer) MoveDone(context.Context, *MoveDoneRequest) (*MoveDoneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveDone not implemented")
}
func (UnimplementedMachineServer) CreateCronJob(context.Context, *CreateCronJobRequest) (*CronJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCronJob not implemented")
}
func (UnimplementedMachineServer) DeleteCronJob(context.Context, *DeleteCronJobRequest) (*DeleteCronJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCronJob not implemented")
}
func (UnimplementedMachineServer) ListCronJob(context.Context, *ListCronJobRequest) (*CronJobsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCronJob not implemented")
}
func (UnimplementedMachineServer) mustEmbedUnimplementedMachineServer() {}

// UnsafeMachineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineServer will
// result in compilation errors.
type UnsafeMachineServer interface {
	mustEmbedUnimplementedMachineServer()
}

func RegisterMachineServer(s grpc.ServiceRegistrar, srv MachineServer) {
	s.RegisterService(&Machine_ServiceDesc, srv)
}

func _Machine_FindByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).FindByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_FindByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).FindByUserId(ctx, req.(*FindByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_Move_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).Move(ctx, req.(*MoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_Zero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).Zero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_Zero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).Zero(ctx, req.(*ZeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_GetMotorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMotorStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).GetMotorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_GetMotorStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).GetMotorStatus(ctx, req.(*GetMotorStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_MoveDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).MoveDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_MoveDone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).MoveDone(ctx, req.(*MoveDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_CreateCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).CreateCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_CreateCronJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).CreateCronJob(ctx, req.(*CreateCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_DeleteCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).DeleteCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_DeleteCronJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).DeleteCronJob(ctx, req.(*DeleteCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Machine_ListCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).ListCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Machine_ListCronJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).ListCronJob(ctx, req.(*ListCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Machine_ServiceDesc is the grpc.ServiceDesc for Machine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Machine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "htpp.machine.service.v1.Machine",
	HandlerType: (*MachineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByUserId",
			Handler:    _Machine_FindByUserId_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Machine_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Machine_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Machine_Get_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _Machine_Move_Handler,
		},
		{
			MethodName: "Zero",
			Handler:    _Machine_Zero_Handler,
		},
		{
			MethodName: "GetMotorStatus",
			Handler:    _Machine_GetMotorStatus_Handler,
		},
		{
			MethodName: "MoveDone",
			Handler:    _Machine_MoveDone_Handler,
		},
		{
			MethodName: "CreateCronJob",
			Handler:    _Machine_CreateCronJob_Handler,
		},
		{
			MethodName: "DeleteCronJob",
			Handler:    _Machine_DeleteCronJob_Handler,
		},
		{
			MethodName: "ListCronJob",
			Handler:    _Machine_ListCronJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/htpp/machine/service/v1/machine.proto",
}