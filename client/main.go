package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	blogpb "jasper.com/blog/proto/blog/gen"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to blog service: %v", err)
		return
	}
	defer conn.Close()

	client := blogpb.NewBlogServiceClient(conn)
	resp, err := client.GetArticles(context.Background(), &blogpb.GetArticlesRequest{
		Page: 1,
		Size: 10,
	})
	if err != nil {
		log.Fatalf("cannot get articles: %v", err)
		return
	}
	log.Printf("Response: %v", resp.Items)
}
