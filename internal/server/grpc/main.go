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

func NewGRPC() (*GRPC, []error) {
	ng := &GRPC{}
	var errs []error
	var err error
	ng.AS, err = service.NewArticleService()
	errs = append(errs, err)
	ng.CS, err = service.NewCategoryService()
	errs = append(errs, err)
	ng.TS, err = service.NewTagService()
	errs = append(errs, err)
	ng.ATTS, err = service.NewAttributeService()
	errs = append(errs, err)
	return ng, errs
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
	pb.RegisterArticlesAPIServer(s, gx.AS)
	pb.RegisterCategoriesAPIServer(s, gx.CS)
	pb.RegisterTagsAPIServer(s, gx.TS)
	pb.RegisterAttributesAPIServer(s, gx.ATTS)

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
