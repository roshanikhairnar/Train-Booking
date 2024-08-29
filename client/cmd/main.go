package main

import (
	"log"

	"github.com/Train-Booking/client"
	trainbookingpb "github.com/Train-Booking/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := trainbookingpb.NewTrainBookingServiceClient(conn)

	client.DoPurchase(c)
}
