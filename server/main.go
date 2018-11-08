package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/rugwirobaker/quavo/models/cache"

	"google.golang.org/grpc"
)

var port = 4000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	srv := NewCacheService()
	cache.RegisterCacheServer(grpcServer, srv)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	fmt.Println("starting quavo server...")
	go func() {
		fmt.Printf("quavo service listening on port \"%d\"\n", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("shutting down with error: %s", err)
		}
	}()
	<-stop
	grpcServer.GracefulStop()
	log.Println("quavo service shutting down...")
}
