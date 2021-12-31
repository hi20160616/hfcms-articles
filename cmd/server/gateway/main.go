package main

import (
	"context"
	"flag"

	"github.com/golang/glog"
	"github.com/hi20160616/hfcms-articles/internal/server/gateway"
)

var (
	endpoint   = flag.String("endpoint", "localhost:9090", "endpoint of the gRPC service")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	openAPIDir = flag.String("openapi_dir", "api/articles/v1", "path to the directory which contains OpenAPI definitions")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
		OpenAPIDir: *openAPIDir,
	}
	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}
