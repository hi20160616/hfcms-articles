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

// TagsAPIClient is the client API for TagsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagsAPIClient interface {
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error)
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*Tag, error)
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error)
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*Tag, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tagsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTagsAPIClient(cc grpc.ClientConnInterface) TagsAPIClient {
	return &tagsAPIClient{cc}
}

func (c *tagsAPIClient) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagsAPI/ListTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsAPIClient) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagsAPI/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsAPIClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagsAPI/CreateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsAPIClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagsAPI/UpdateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsAPIClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.TagsAPI/DeleteTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsAPIServer is the server API for TagsAPI service.
// All implementations must embed UnimplementedTagsAPIServer
// for forward compatibility
type TagsAPIServer interface {
	ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
	GetTag(context.Context, *GetTagRequest) (*Tag, error)
	CreateTag(context.Context, *CreateTagRequest) (*Tag, error)
	UpdateTag(context.Context, *UpdateTagRequest) (*Tag, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTagsAPIServer()
}

// UnimplementedTagsAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTagsAPIServer struct {
}

func (UnimplementedTagsAPIServer) ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTags not implemented")
}
func (UnimplementedTagsAPIServer) GetTag(context.Context, *GetTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedTagsAPIServer) CreateTag(context.Context, *CreateTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedTagsAPIServer) UpdateTag(context.Context, *UpdateTagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedTagsAPIServer) DeleteTag(context.Context, *DeleteTagRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedTagsAPIServer) mustEmbedUnimplementedTagsAPIServer() {}

// UnsafeTagsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagsAPIServer will
// result in compilation errors.
type UnsafeTagsAPIServer interface {
	mustEmbedUnimplementedTagsAPIServer()
}

func RegisterTagsAPIServer(s *grpc.Server, srv TagsAPIServer) {
	s.RegisterService(&_TagsAPI_serviceDesc, srv)
}

func _TagsAPI_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsAPIServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagsAPI/ListTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsAPIServer).ListTags(ctx, req.(*ListTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsAPI_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsAPIServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagsAPI/GetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsAPIServer).GetTag(ctx, req.(*GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsAPI_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsAPIServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagsAPI/CreateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsAPIServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsAPI_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsAPIServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagsAPI/UpdateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsAPIServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagsAPI_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsAPIServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.TagsAPI/DeleteTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsAPIServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagsAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hfcms.articles.v1.TagsAPI",
	HandlerType: (*TagsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTags",
			Handler:    _TagsAPI_ListTags_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _TagsAPI_GetTag_Handler,
		},
		{
			MethodName: "CreateTag",
			Handler:    _TagsAPI_CreateTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _TagsAPI_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _TagsAPI_DeleteTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/articles/v1/hfcms-tags.proto",
}
