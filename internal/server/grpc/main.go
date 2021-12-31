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
	AS service.ArticleService
}

func NewGRPC() *GRPC {
	return &GRPC{
		AS: *service.NewArticleService(),
	}
}

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func (g *GRPC) Run(ctx context.Context, gx *GRPC, network, address string) error {
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
	// pb.RegisterCategoriesAPIServer(s, newCagegoriesAPIServer())
	// pb.RegisterTagsAPIServer(s, newTagsAPIServer())
	// pb.RegisterAttributesAPIServer(s, newAttributesAPIServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

// RunInProcessGateway starts the invoke in process http gateway.
func RunInProcessGateway(ctx context.Context, gx *GRPC, addr string, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)

	pb.RegisterArticlesAPIHandlerServer(ctx, mux, gx.AS.UnimplementedArticlesAPIServer)
	// pb.RegisterCategoriesAPIHandlerServer(ctx, mux, newArticlesAPIServer())
	// pb.RegisterTagsAPIHandlerServer(ctx, mux, newArticlesAPIServer())
	// pb.RegisterAttributesAPIHandlerServer(ctx, mux, newArticlesAPIServer())

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
