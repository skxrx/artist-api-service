package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/skxrx/artist-api-service/internal/api"
	pb "github.com/skxrx/artist-api-service/proto/parser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Can't connect parser: %v", err)
	}
	defer conn.Close()

	client := pb.NewParserServiceClient(conn)
	handler := api.NewHandler(client)

	r := gin.Default()
	r.GET("/artist/:artistName/events", handler.GetEventsHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
