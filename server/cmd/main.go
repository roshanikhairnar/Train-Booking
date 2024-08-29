package main

import (
	"fmt"
	"log"
	"net"

	trainbookingpb "github.com/Train-Booking/proto"
	"github.com/Train-Booking/server"
	"google.golang.org/grpc"
	//"github.com/Train-Booking/proto/train_booking/trainbookingpb"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Print("Server started")
	s := grpc.NewServer()
	trainbookingpb.RegisterTrainBookingServiceServer(s, &server.Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
