// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: cmd/004-grpc/model/messages.proto

package model

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

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MathServiceClient interface {
	Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Sub(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Mul(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Div(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Mod(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
}

type mathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMathServiceClient(cc grpc.ClientConnInterface) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MathService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Sub(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MathService/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Mul(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MathService/Mul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Div(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MathService/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Mod(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MathService/Mod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
// All implementations must embed UnimplementedMathServiceServer
// for forward compatibility
type MathServiceServer interface {
	Add(context.Context, *MathRequest) (*MathResponse, error)
	Sub(context.Context, *MathRequest) (*MathResponse, error)
	Mul(context.Context, *MathRequest) (*MathResponse, error)
	Div(context.Context, *MathRequest) (*MathResponse, error)
	Mod(context.Context, *MathRequest) (*MathResponse, error)
	mustEmbedUnimplementedMathServiceServer()
}

// UnimplementedMathServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMathServiceServer struct {
}

func (UnimplementedMathServiceServer) Add(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedMathServiceServer) Sub(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (UnimplementedMathServiceServer) Mul(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul not implemented")
}
func (UnimplementedMathServiceServer) Div(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}
func (UnimplementedMathServiceServer) Mod(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mod not implemented")
}
func (UnimplementedMathServiceServer) mustEmbedUnimplementedMathServiceServer() {}

// UnsafeMathServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MathServiceServer will
// result in compilation errors.
type UnsafeMathServiceServer interface {
	mustEmbedUnimplementedMathServiceServer()
}

func RegisterMathServiceServer(s grpc.ServiceRegistrar, srv MathServiceServer) {
	s.RegisterService(&MathService_ServiceDesc, srv)
}

func _MathService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MathService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Add(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MathService/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Sub(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MathService/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Mul(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MathService/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Div(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Mod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Mod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MathService/Mod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Mod(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MathService_ServiceDesc is the grpc.ServiceDesc for MathService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MathService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MathService_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _MathService_Sub_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _MathService_Mul_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _MathService_Div_Handler,
		},
		{
			MethodName: "Mod",
			Handler:    _MathService_Mod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/004-grpc/model/messages.proto",
}
