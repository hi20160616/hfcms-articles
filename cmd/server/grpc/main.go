package main

import (
	"context"
	"flag"
	"log"

	"github.com/golang/glog"
	"github.com/hi20160616/hfcms-articles/internal/server/grpc"
)

var (
	addr    = flag.String("addr", ":9090", "endpoint of the gRPC service")
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	g, err := grpc.NewGRPC()
	if err != nil {
		log.Fatal(err)
	}
	// if err := g.RunInProcessGateway(ctx, *addr); err != nil {
	//         glog.Fatal(err)
	// }
	if err := g.Run(ctx, *network, *addr); err != nil {
		glog.Fatal(err)
	}
}
