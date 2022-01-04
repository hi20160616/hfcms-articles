package grpc

import (
	"context"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/internal/service"
	"google.golang.org/grpc"
)

type GRPC struct {
	AS   *service.ArticleService
	CS   *service.CategoryService
	TS   *service.TagService
	ATTS *service.AttributeService
}

func NewGRPC() (*GRPC, error) {
	as, err := service.NewArticleService()
	if err != nil {
		return nil, err
	}

	cs, err := service.NewCategoryService()
	if err != nil {
		return nil, err
	}

	ts, err := service.NewTagService()
	if err != nil {
		return nil, err
	}

	atts, err := service.NewAttributeService()
	if err != nil {
		return nil, err
	}

	return &GRPC{AS: as, CS: cs, TS: ts, ATTS: atts}, nil
}

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func (gx *GRPC) Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	pb.RegisterArticlesAPIServer(s, gx.AS.UnimplementedArticlesAPIServer)
	pb.RegisterCategoriesAPIServer(s, gx.CS.UnimplementedCategoriesAPIServer)
	pb.RegisterTagsAPIServer(s, gx.TS.UnimplementedTagsAPIServer)
	pb.RegisterAttributesAPIServer(s, gx.ATTS.UnimplementedAttributesAPIServer)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

// RunInProcessGateway starts the invoke in process http gateway.
func (gx *GRPC) RunInProcessGateway(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)

	pb.RegisterArticlesAPIHandlerServer(ctx, mux, gx.AS.UnimplementedArticlesAPIServer)
	pb.RegisterCategoriesAPIHandlerServer(ctx, mux, gx.CS.UnimplementedCategoriesAPIServer)
	pb.RegisterTagsAPIHandlerServer(ctx, mux, gx.TS.UnimplementedTagsAPIServer)
	pb.RegisterAttributesAPIHandlerServer(ctx, mux, gx.ATTS.UnimplementedAttributesAPIServer)

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http gateway server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http gateway server: %v", err)
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil

}
