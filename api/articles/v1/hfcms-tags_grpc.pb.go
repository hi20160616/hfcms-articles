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

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error)
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*Tag, error)
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error)
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*Tag, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagService/ListTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagService/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagService/CreateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagService/UpdateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagService/DeleteTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
	GetTag(context.Context, *GetTagRequest) (*Tag, error)
	CreateTag(context.Context, *CreateTagRequest) (*Tag, error)
	UpdateTag(context.Context, *UpdateTagRequest) (*Tag, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTags not implemented")
}
func (UnimplementedTagServiceServer) GetTag(context.Context, *GetTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedTagServiceServer) CreateTag(context.Context, *CreateTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedTagServiceServer) UpdateTag(context.Context, *UpdateTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedTagServiceServer) DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s *grpc.Server, srv TagServiceServer) {
	s.RegisterService(&_TagService_serviceDesc, srv)
}

func _TagService_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagService/ListTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).ListTags(ctx, req.(*ListTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagService/GetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTag(ctx, req.(*GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagService/CreateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagService/UpdateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagService/DeleteTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hfcms.articles.v1.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTags",
			Handler:    _TagService_ListTags_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _TagService_GetTag_Handler,
		},
		{
			MethodName: "CreateTag",
			Handler:    _TagService_CreateTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _TagService_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _TagService_DeleteTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/articles/v1/hfcms-tags.proto",
}
