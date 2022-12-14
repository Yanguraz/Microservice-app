// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: predict/predict.proto

package predict

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

// PredictServiceClient is the client API for PredictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PredictServiceClient interface {
	CreatePredict(ctx context.Context, in *CreatePredictRequest, opts ...grpc.CallOption) (*CreatePredictResponse, error)
	GetPredicts(ctx context.Context, in *GetPredictsRequest, opts ...grpc.CallOption) (*GetPredictsResponse, error)
	GetPredict(ctx context.Context, in *GetPredictRequest, opts ...grpc.CallOption) (*GetPredictResponse, error)
	UpdatePredict(ctx context.Context, in *UpdatePredictRequest, opts ...grpc.CallOption) (*UpdatePredictResponse, error)
	DeletePredict(ctx context.Context, in *DeletePredictRequest, opts ...grpc.CallOption) (*DeletePredictResponse, error)
}

type predictServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictServiceClient(cc grpc.ClientConnInterface) PredictServiceClient {
	return &predictServiceClient{cc}
}

func (c *predictServiceClient) CreatePredict(ctx context.Context, in *CreatePredictRequest, opts ...grpc.CallOption) (*CreatePredictResponse, error) {
	out := new(CreatePredictResponse)
	err := c.cc.Invoke(ctx, "/predict.PredictService/CreatePredict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictServiceClient) GetPredicts(ctx context.Context, in *GetPredictsRequest, opts ...grpc.CallOption) (*GetPredictsResponse, error) {
	out := new(GetPredictsResponse)
	err := c.cc.Invoke(ctx, "/predict.PredictService/GetPredicts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictServiceClient) GetPredict(ctx context.Context, in *GetPredictRequest, opts ...grpc.CallOption) (*GetPredictResponse, error) {
	out := new(GetPredictResponse)
	err := c.cc.Invoke(ctx, "/predict.PredictService/GetPredict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictServiceClient) UpdatePredict(ctx context.Context, in *UpdatePredictRequest, opts ...grpc.CallOption) (*UpdatePredictResponse, error) {
	out := new(UpdatePredictResponse)
	err := c.cc.Invoke(ctx, "/predict.PredictService/UpdatePredict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictServiceClient) DeletePredict(ctx context.Context, in *DeletePredictRequest, opts ...grpc.CallOption) (*DeletePredictResponse, error) {
	out := new(DeletePredictResponse)
	err := c.cc.Invoke(ctx, "/predict.PredictService/DeletePredict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictServiceServer is the server API for PredictService service.
// All implementations must embed UnimplementedPredictServiceServer
// for forward compatibility
type PredictServiceServer interface {
	CreatePredict(context.Context, *CreatePredictRequest) (*CreatePredictResponse, error)
	GetPredicts(context.Context, *GetPredictsRequest) (*GetPredictsResponse, error)
	GetPredict(context.Context, *GetPredictRequest) (*GetPredictResponse, error)
	UpdatePredict(context.Context, *UpdatePredictRequest) (*UpdatePredictResponse, error)
	DeletePredict(context.Context, *DeletePredictRequest) (*DeletePredictResponse, error)
	mustEmbedUnimplementedPredictServiceServer()
}

// UnimplementedPredictServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPredictServiceServer struct {
}

func (UnimplementedPredictServiceServer) CreatePredict(context.Context, *CreatePredictRequest) (*CreatePredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePredict not implemented")
}
func (UnimplementedPredictServiceServer) GetPredicts(context.Context, *GetPredictsRequest) (*GetPredictsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredicts not implemented")
}
func (UnimplementedPredictServiceServer) GetPredict(context.Context, *GetPredictRequest) (*GetPredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredict not implemented")
}
func (UnimplementedPredictServiceServer) UpdatePredict(context.Context, *UpdatePredictRequest) (*UpdatePredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePredict not implemented")
}
func (UnimplementedPredictServiceServer) DeletePredict(context.Context, *DeletePredictRequest) (*DeletePredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePredict not implemented")
}
func (UnimplementedPredictServiceServer) mustEmbedUnimplementedPredictServiceServer() {}

// UnsafePredictServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PredictServiceServer will
// result in compilation errors.
type UnsafePredictServiceServer interface {
	mustEmbedUnimplementedPredictServiceServer()
}

func RegisterPredictServiceServer(s grpc.ServiceRegistrar, srv PredictServiceServer) {
	s.RegisterService(&PredictService_ServiceDesc, srv)
}

func _PredictService_CreatePredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServiceServer).CreatePredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predict.PredictService/CreatePredict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServiceServer).CreatePredict(ctx, req.(*CreatePredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictService_GetPredicts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPredictsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServiceServer).GetPredicts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predict.PredictService/GetPredicts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServiceServer).GetPredicts(ctx, req.(*GetPredictsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictService_GetPredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServiceServer).GetPredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predict.PredictService/GetPredict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServiceServer).GetPredict(ctx, req.(*GetPredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictService_UpdatePredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServiceServer).UpdatePredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predict.PredictService/UpdatePredict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServiceServer).UpdatePredict(ctx, req.(*UpdatePredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictService_DeletePredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServiceServer).DeletePredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predict.PredictService/DeletePredict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServiceServer).DeletePredict(ctx, req.(*DeletePredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PredictService_ServiceDesc is the grpc.ServiceDesc for PredictService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PredictService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "predict.PredictService",
	HandlerType: (*PredictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePredict",
			Handler:    _PredictService_CreatePredict_Handler,
		},
		{
			MethodName: "GetPredicts",
			Handler:    _PredictService_GetPredicts_Handler,
		},
		{
			MethodName: "GetPredict",
			Handler:    _PredictService_GetPredict_Handler,
		},
		{
			MethodName: "UpdatePredict",
			Handler:    _PredictService_UpdatePredict_Handler,
		},
		{
			MethodName: "DeletePredict",
			Handler:    _PredictService_DeletePredict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "predict/predict.proto",
}
