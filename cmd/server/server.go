package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hi20160616/hfcms-articles/configs"
	theGateway "github.com/hi20160616/hfcms-articles/internal/server/gateway"
	theGRPC "github.com/hi20160616/hfcms-articles/internal/server/grpc"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	cfg := configs.NewConfig("hfcms-articles")

	// gRPC
	g.Go(func() error {
		return theGRPC.Run(ctx, "tcp", cfg.API.GRPC.Addr)
	})

	// gRPC-gateway
	opts := theGateway.Options{
		Addr: cfg.API.HTTP.Addr,
		GRPCServer: theGateway.Endpoint{
			Network: cfg.API.GRPC.Network,
			Addr:    cfg.API.GRPC.Addr,
		},
	}
	g.Go(func() error {
		return theGateway.Run(ctx, opts)
	})

	// Graceful stop
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	g.Go(func() error {
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("not canceled by context: %s", err)
		} else {
			log.Println(err)
		}
	}
}
