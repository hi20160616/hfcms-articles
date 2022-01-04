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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArticlesAPIClient is the client API for ArticlesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesAPIClient interface {
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*Article, error)
	SearchArticles(ctx context.Context, in *SearchArticlesRequest, opts ...grpc.CallOption) (*SearchArticlesResponse, error)
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*Article, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*Article, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type articlesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesAPIClient(cc grpc.ClientConnInterface) ArticlesAPIClient {
	return &articlesAPIClient{cc}
}

func (c *articlesAPIClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error) {
	out := new(ListArticlesResponse)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.ArticlesAPI/ListArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesAPIClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.ArticlesAPI/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesAPIClient) SearchArticles(ctx context.Context, in *SearchArticlesRequest, opts ...grpc.CallOption) (*SearchArticlesResponse, error) {
	out := new(SearchArticlesResponse)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.ArticlesAPI/SearchArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesAPIClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.ArticlesAPI/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesAPIClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.ArticlesAPI/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesAPIClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hfcms.articles.v1.ArticlesAPI/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesAPIServer is the server API for ArticlesAPI service.
// All implementations must embed UnimplementedArticlesAPIServer
// for forward compatibility
type ArticlesAPIServer interface {
	ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error)
	GetArticle(context.Context, *GetArticleRequest) (*Article, error)
	SearchArticles(context.Context, *SearchArticlesRequest) (*SearchArticlesResponse, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*Article, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*Article, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedArticlesAPIServer()
}

// UnimplementedArticlesAPIServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesAPIServer struct {
}

func (UnimplementedArticlesAPIServer) ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedArticlesAPIServer) GetArticle(context.Context, *GetArticleRequest) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticlesAPIServer) SearchArticles(context.Context, *SearchArticlesRequest) (*SearchArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArticles not implemented")
}
func (UnimplementedArticlesAPIServer) CreateArticle(context.Context, *CreateArticleRequest) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticlesAPIServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticlesAPIServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticlesAPIServer) mustEmbedUnimplementedArticlesAPIServer() {}

// UnsafeArticlesAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticlesAPIServer will
// result in compilation errors.
type UnsafeArticlesAPIServer interface {
	mustEmbedUnimplementedArticlesAPIServer()
}

func RegisterArticlesAPIServer(s grpc.ServiceRegistrar, srv ArticlesAPIServer) {
	s.RegisterService(&ArticlesAPI_ServiceDesc, srv)
}

func _ArticlesAPI_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesAPIServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.ArticlesAPI/ListArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesAPIServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticlesAPI_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesAPIServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.ArticlesAPI/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesAPIServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticlesAPI_SearchArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesAPIServer).SearchArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.ArticlesAPI/SearchArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesAPIServer).SearchArticles(ctx, req.(*SearchArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticlesAPI_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesAPIServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.ArticlesAPI/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesAPIServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticlesAPI_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesAPIServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.ArticlesAPI/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesAPIServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticlesAPI_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesAPIServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hfcms.articles.v1.ArticlesAPI/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesAPIServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticlesAPI_ServiceDesc is the grpc.ServiceDesc for ArticlesAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticlesAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hfcms.articles.v1.ArticlesAPI",
	HandlerType: (*ArticlesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListArticles",
			Handler:    _ArticlesAPI_ListArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ArticlesAPI_GetArticle_Handler,
		},
		{
			MethodName: "SearchArticles",
			Handler:    _ArticlesAPI_SearchArticles_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _ArticlesAPI_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _ArticlesAPI_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ArticlesAPI_DeleteArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/articles/v1/hfcms-articles.proto",
}
