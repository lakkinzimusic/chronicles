package main

import (
	"fmt"
	"github.com/lakkinzimusic/chronicles/config"
	pb "github.com/lakkinzimusic/chronicles/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"github.com/lakkinzimusic/chronicles/service"
	"log"
	"net"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.GetConfig("chronicles.yaml")
	listener, err := net.Listen(cfg.Server.Protocol, cfg.Server.Port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	fmt.Println(fmt.Sprintf("mongodb://%s:%s/", cfg.DB.Host, cfg.DB.Port))
	client := CreateMongoClient(context.Background(), fmt.Sprintf("mongodb://%s:%s/", cfg.DB.Host, cfg.DB.Port))
	collection := client.Database(cfg.DB.Database).Collection(cfg.DB.Collection)
	pb.RegisterChronicleServiceServer(grpcServer, service.NewChronicleService(service.NewStore(collection)))
	grpcServer.Serve(listener)
}

func CreateMongoClient(ctx context.Context, uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}