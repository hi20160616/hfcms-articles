package grpc

import (
	"context"
	"net"

	"github.com/golang/glog"
	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/internal/service"
	"google.golang.org/grpc"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	as, err := service.NewArticleService()
	if err != nil {
		return err
	}
	atts, err := service.NewAttributeService()
	if err != nil {
		return err
	}
	cs, err := service.NewCategoryService()
	if err != nil {
		return err
	}
	ts, err := service.NewTagService()
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterArticlesAPIServer(s, as)
	pb.RegisterAttributesAPIServer(s, atts)
	pb.RegisterCategoriesAPIServer(s, cs)
	pb.RegisterTagsAPIServer(s, ts)
	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	glog.Infof("gRPC starting listening at %s", address)
	return s.Serve(l)
}
