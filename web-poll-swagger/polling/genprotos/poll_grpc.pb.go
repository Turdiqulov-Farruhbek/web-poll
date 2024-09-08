// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: protos/poll.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PollService_Create_FullMethodName  = "/polling.PollService/Create"
	PollService_Update_FullMethodName  = "/polling.PollService/Update"
	PollService_Delete_FullMethodName  = "/polling.PollService/Delete"
	PollService_GetByID_FullMethodName = "/polling.PollService/GetByID"
	PollService_GetAll_FullMethodName  = "/polling.PollService/GetAll"
)

// PollServiceClient is the client API for PollService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PollServiceClient interface {
	Create(ctx context.Context, in *PollCreateReq, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *PollUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error)
	GetByID(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*PollGetByIDRes, error)
	GetAll(ctx context.Context, in *PollGetAllReq, opts ...grpc.CallOption) (*PollGetAllRes, error)
}

type pollServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPollServiceClient(cc grpc.ClientConnInterface) PollServiceClient {
	return &pollServiceClient{cc}
}

func (c *pollServiceClient) Create(ctx context.Context, in *PollCreateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, PollService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollServiceClient) Update(ctx context.Context, in *PollUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, PollService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollServiceClient) Delete(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, PollService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollServiceClient) GetByID(ctx context.Context, in *ByID, opts ...grpc.CallOption) (*PollGetByIDRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PollGetByIDRes)
	err := c.cc.Invoke(ctx, PollService_GetByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pollServiceClient) GetAll(ctx context.Context, in *PollGetAllReq, opts ...grpc.CallOption) (*PollGetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PollGetAllRes)
	err := c.cc.Invoke(ctx, PollService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PollServiceServer is the server API for PollService service.
// All implementations must embed UnimplementedPollServiceServer
// for forward compatibility
type PollServiceServer interface {
	Create(context.Context, *PollCreateReq) (*Void, error)
	Update(context.Context, *PollUpdateReq) (*Void, error)
	Delete(context.Context, *ByID) (*Void, error)
	GetByID(context.Context, *ByID) (*PollGetByIDRes, error)
	GetAll(context.Context, *PollGetAllReq) (*PollGetAllRes, error)
	mustEmbedUnimplementedPollServiceServer()
}

// UnimplementedPollServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPollServiceServer struct {
}

func (UnimplementedPollServiceServer) Create(context.Context, *PollCreateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPollServiceServer) Update(context.Context, *PollUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPollServiceServer) Delete(context.Context, *ByID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPollServiceServer) GetByID(context.Context, *ByID) (*PollGetByIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedPollServiceServer) GetAll(context.Context, *PollGetAllReq) (*PollGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPollServiceServer) mustEmbedUnimplementedPollServiceServer() {}

// UnsafePollServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PollServiceServer will
// result in compilation errors.
type UnsafePollServiceServer interface {
	mustEmbedUnimplementedPollServiceServer()
}

func RegisterPollServiceServer(s grpc.ServiceRegistrar, srv PollServiceServer) {
	s.RegisterService(&PollService_ServiceDesc, srv)
}

func _PollService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PollService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServiceServer).Create(ctx, req.(*PollCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PollService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PollService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServiceServer).Update(ctx, req.(*PollUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PollService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PollService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServiceServer).Delete(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PollService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PollService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServiceServer).GetByID(ctx, req.(*ByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PollService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PollServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PollService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PollServiceServer).GetAll(ctx, req.(*PollGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PollService_ServiceDesc is the grpc.ServiceDesc for PollService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PollService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "polling.PollService",
	HandlerType: (*PollServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PollService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PollService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PollService_Delete_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _PollService_GetByID_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PollService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/poll.proto",
}