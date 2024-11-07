package main

import (
	"context"
	"log"
	"net/http"

	trainbookingpb "github.com/Train-Booking/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := trainbookingpb.RegisterTrainTicketServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Error in registering the service:%v", err)
	}

	err = http.ListenAndServe(":8010", mux)
	if err != nil {
		log.Fatalf("Error in listening the service:%v", err)
	}
}
