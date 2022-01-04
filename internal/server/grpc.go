package server

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	myerr "github.com/hi20160616/hfcms-articles/errors"
	"github.com/hi20160616/hfcms-articles/internal/service"
	"google.golang.org/grpc"
)

var (
	addr    = flag.String("addr", ":9090", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

type GRPC struct {
	s *grpc.Server
	l net.Listener
}

func NewGRPCServer() (*GRPC, error) {
	// t, err := time.ParseDuration("1s")
	// if err != nil {
	//         return nil, err
	// }
	// opts := []grpc.ServerOption{
	//         grpc.ConnectionTimeout(t),
	// }
	// s := grpc.NewServer(opts...)
	s := grpc.NewServer()
	as, err := service.NewArticleService()
	if err != nil {
		return nil, err
	}

	pb.RegisterArticlesAPIServer(s, as)
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		return nil, err
	}
	return &GRPC{s, l}, nil
}

func (g *GRPC) Start(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e)
			myerr.PanicLog(e)
		}
	}()
	defer g.l.Close()
	return g.s.Serve(g.l)
}

func (g *GRPC) Stop(ctx context.Context) error {
	g.s.GracefulStop()
	return nil
}
