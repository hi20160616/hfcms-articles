package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/configs"
	myerr "github.com/hi20160616/hfcms-articles/errors"
	"github.com/hi20160616/hfcms-articles/internal/service"
	"google.golang.org/grpc"
)

type GRPC struct {
	s   *grpc.Server
	l   net.Listener
	cfg *configs.Config
}

func NewGRPCServer(cfg *configs.Config) (*GRPC, error) {
	// s := grpc.NewServer(
	//         grpc.KeepaliveParams(keepalive.ServerParameters{
	//                 MaxConnectionIdle: 5 * time.Minute,
	//         }),
	// )
	s := grpc.NewServer()
	as, err := service.NewArticleService()
	if err != nil {
		return nil, err
	}

	pb.RegisterArticlesAPIServer(s, as)
	l, err := net.Listen("tcp", cfg.API.GRPC.Addr)
	if err != nil {
		return nil, err
	}
	return &GRPC{s, l, cfg}, nil
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
	log.Println("Server gRPC on ", g.cfg.API.GRPC.Addr)
	return g.s.Serve(g.l)
}

func (g *GRPC) Stop(ctx context.Context) error {
	defer g.s.GracefulStop()
	<-ctx.Done()
	return nil
}

func (g *GRPC) StartRESTFul(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e)
			myerr.PanicLog(e)
		}
	}()

	conn, err := grpc.DialContext(ctx,
		g.cfg.API.GRPC.Addr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}
	gwmux := runtime.NewServeMux()
	if err = pb.RegisterArticlesAPIHandler(ctx, gwmux, conn); err != nil {
		return err
	}
	gwServer := &http.Server{Addr: g.cfg.API.HTTP.Addr, Handler: gwmux}
	go func() {
		<-ctx.Done()
		log.Println("Server gRPC-Gateway stop!")
		if err := gwServer.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server gRPC-Gateway on ", g.cfg.API.HTTP.Addr)
	return gwServer.ListenAndServe()
}
