package main

import (
	"context"
	"flag"
	"log"
	"net"

	"cloud.google.com/go/pubsub"
	"github.com/emicklei/parcello/v1"
	"google.golang.org/grpc"
)

var verbose = flag.Bool("v", false, "verbose logging")
var version = "0.1"

func main() {
	flag.Parse()
	log.Println("parcello, the pubsub delivery service -- version", version)
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// create pubsub client
	if len(config.Project) == 0 {
		log.Fatal("missing project in config")
	}
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.Project)
	if err != nil {
		log.Fatalf("failed to create PubSub client: %v", err)
	}

	config.checkTopicsAndSubscriptions(client)

	// schedule parcel listeners
	service := &deliveryServiceImpl{client: client, config: config}
	for _, each := range config.Queues {
		go loopReceiveParcels(client, each, service)
	}

	// start server
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterDeliveryServiceServer(grpcServer, service)
	log.Fatal(grpcServer.Serve(lis))
}
