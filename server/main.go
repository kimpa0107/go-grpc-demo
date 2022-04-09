package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	blogpb "jasper.com/blog/proto/blog/gen"
)

type Service struct{}

func (s *Service) GetArticles(ctx context.Context, req *blogpb.GetArticlesRequest) (*blogpb.GetArticlesResponse, error) {
	articles := []*blogpb.Article{
		{
			Id:         "article-1",
			Title:      "Article 1",
			CoverImage: "",
			Category: &blogpb.Category{
				Id:   1,
				Name: "Go",
			},
			Tags: []string{"Go", "GoLang"},
		},
		{
			Id:         "article-2",
			Title:      "Article 2",
			CoverImage: "",
			Category: &blogpb.Category{
				Id:   2,
				Name: "gRPC",
			},
			Tags: []string{"grpc"},
		},
	}

	return &blogpb.GetArticlesResponse{
		Items: articles,
		Extra: &blogpb.Extra{
			Page:  req.Page,
			Size:  req.Size,
			Total: 100,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &Service{})
	log.Fatal(s.Serve(lis))
}
