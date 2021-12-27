// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AttributesServiceClient is the client API for AttributesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttributesServiceClient interface {
	ListAttributes(ctx context.Context, in *ListAttributesRequest, opts ...grpc.CallOption) (*ListAttributesResponse, error)
	GetAttribute(ctx context.Context, in *GetAttributeRequest, opts ...grpc.CallOption) (*Attribute, error)
	CreateAttribute(ctx context.Context, in *CreateAttributeRequest, opts ...grpc.CallOption) (*Attribute, error)
	UpdateAttribute(ctx context.Context, in *UpdateAttributeRequest, opts ...grpc.CallOption) (*Attribute, error)
	DeleteAttribute(ctx context.Context, in *DeleteAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type attributesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttributesServiceClient(cc grpc.ClientConnInterface) AttributesServiceClient {
	return &attributesServiceClient{cc}
}

func (c *attributesServiceClient) ListAttributes(ctx context.Context, in *ListAttributesRequest, opts ...grpc.CallOption) (*ListAttributesResponse, error) {
	out := new(ListAttributesResponse)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.AttributesService/ListAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) GetAttribute(ctx context.Context, in *GetAttributeRequest, opts ...grpc.CallOption) (*Attribute, error) {
	out := new(Attribute)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.AttributesService/GetAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) CreateAttribute(ctx context.Context, in *CreateAttributeRequest, opts ...grpc.CallOption) (*Attribute, error) {
	out := new(Attribute)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.AttributesService/CreateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) UpdateAttribute(ctx context.Context, in *UpdateAttributeRequest, opts ...grpc.CallOption) (*Attribute, error) {
	out := new(Attribute)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.AttributesService/UpdateAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributesServiceClient) DeleteAttribute(ctx context.Context, in *DeleteAttributeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.AttributesService/DeleteAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttributesServiceServer is the server API for AttributesService service.
// All implementations must embed UnimplementedAttributesServiceServer
// for forward compatibility
type AttributesServiceServer interface {
	ListAttributes(context.Context, *ListAttributesRequest) (*ListAttributesResponse, error)
	GetAttribute(context.Context, *GetAttributeRequest) (*Attribute, error)
	CreateAttribute(context.Context, *CreateAttributeRequest) (*Attribute, error)
	UpdateAttribute(context.Context, *UpdateAttributeRequest) (*Attribute, error)
	DeleteAttribute(context.Context, *DeleteAttributeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAttributesServiceServer()
}

// UnimplementedAttributesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAttributesServiceServer struct {
}

func (UnimplementedAttributesServiceServer) ListAttributes(context.Context, *ListAttributesRequest) (*ListAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttributes not implemented")
}
func (UnimplementedAttributesServiceServer) GetAttribute(context.Context, *GetAttributeRequest) (*Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttribute not implemented")
}
func (UnimplementedAttributesServiceServer) CreateAttribute(context.Context, *CreateAttributeRequest) (*Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttribute not implemented")
}
func (UnimplementedAttributesServiceServer) UpdateAttribute(context.Context, *UpdateAttributeRequest) (*Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttribute not implemented")
}
func (UnimplementedAttributesServiceServer) DeleteAttribute(context.Context, *DeleteAttributeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAttribute not implemented")
}
func (UnimplementedAttributesServiceServer) mustEmbedUnimplementedAttributesServiceServer() {}

// UnsafeAttributesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttributesServiceServer will
// result in compilation errors.
type UnsafeAttributesServiceServer interface {
	mustEmbedUnimplementedAttributesServiceServer()
}

func RegisterAttributesServiceServer(s *grpc.Server, srv AttributesServiceServer) {
	s.RegisterService(&_AttributesService_serviceDesc, srv)
}

func _AttributesService_ListAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).ListAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.AttributesService/ListAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).ListAttributes(ctx, req.(*ListAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_GetAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).GetAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.AttributesService/GetAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).GetAttribute(ctx, req.(*GetAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_CreateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).CreateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.AttributesService/CreateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).CreateAttribute(ctx, req.(*CreateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_UpdateAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).UpdateAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.AttributesService/UpdateAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).UpdateAttribute(ctx, req.(*UpdateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributesService_DeleteAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributesServiceServer).DeleteAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.AttributesService/DeleteAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributesServiceServer).DeleteAttribute(ctx, req.(*DeleteAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AttributesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hfcms.articles.v1.AttributesService",
	HandlerType: (*AttributesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAttributes",
			Handler:    _AttributesService_ListAttributes_Handler,
		},
		{
			MethodName: "GetAttribute",
			Handler:    _AttributesService_GetAttribute_Handler,
		},
		{
			MethodName: "CreateAttribute",
			Handler:    _AttributesService_CreateAttribute_Handler,
		},
		{
			MethodName: "UpdateAttribute",
			Handler:    _AttributesService_UpdateAttribute_Handler,
		},
		{
			MethodName: "DeleteAttribute",
			Handler:    _AttributesService_DeleteAttribute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/articles/v1/hfcms-attributes.proto",
}
