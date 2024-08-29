package client

import (
	"context"
	"fmt"
	"log"

	trainbookingpb "github.com/Train-Booking/proto"
)

func DoPurchase(c trainbookingpb.TrainBookingServiceClient) {
	fmt.Println("Starting to do a Purchase RPC...")
	req := &trainbookingpb.TrainBookingRequest{
		Booking: &trainbookingpb.TrainBooking{
			FirstName: "Roshani",
			LastName:  "KK",
			ToCity:    "London",
			FromCity:  "Paris",
			SeatNo:    23,
			Amount:    20,
		},
	}

	res, err := c.Purchase(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Purchase RPC: %v", err)
	}

	fmt.Printf("Response from Purchase: %v", res)
}
