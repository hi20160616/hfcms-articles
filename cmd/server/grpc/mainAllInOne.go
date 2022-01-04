package main

//
// import (
//         "context"
//         "flag"
//         "fmt"
//         "net"
//
//         pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
//         v1 "github.com/hi20160616/hfcms-articles/api/articles/v1"
//         "github.com/hi20160616/hfcms-articles/internal/service"
//         "google.golang.org/grpc"
// )
//
// var (
//         addr    = flag.String("addr", ":9090", "endpoint of the gRPC service")
//         network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
// )
//
// type server struct {
//         pb.UnimplementedArticlesAPIServer
// }
//
// func (s *server) ListArticles(ctx context.Context, in *pb.ListArticlesRequest) (*pb.ListArticlesResponse, error) {
//         us, err := service.NewArticleService()
//         if err != nil {
//                 return nil, err
//         }
//         as, err := us.ListArticles(ctx, &v1.ListArticlesRequest{})
//         if err != nil {
//                 return nil, err
//         }
//         resp := []*pb.Article{}
//         for _, a := range as.Articles {
//                 resp = append(resp, &pb.Article{
//                         ArticleId:  a.ArticleId,
//                         Title:      a.Title,
//                         Content:    a.Content,
//                         CategoryId: int32(a.CategoryId),
//                         UserId:     int32(a.UserId),
//                         UpdateTime: a.UpdateTime,
//                 })
//         }
//         return &pb.ListArticlesResponse{Articles: resp}, nil
// }
//
// func main() {
//         flag.Parse()
//         lis, err := net.Listen("tcp", *addr)
//         if err != nil {
//                 fmt.Println(err)
//         }
//         s := grpc.NewServer()
//         pb.RegisterArticlesAPIServer(s, &server{})
//         if err := s.Serve(lis); err != nil {
//                 fmt.Println(err)
//         }
// }
