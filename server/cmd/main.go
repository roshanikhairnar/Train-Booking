package main

import (
	"log"
	"net"
	"sync"

	trainbookingpb "github.com/Train-Booking/proto"
	"github.com/Train-Booking/server"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	trainbookingpb.RegisterTrainTicketServiceServer(s, &server.Server{
		UnimplementedTrainTicketServiceServer: trainbookingpb.UnimplementedTrainTicketServiceServer{},
		Mu:                                    sync.Mutex{},
		BookingMap:                            make(map[string]*trainbookingpb.PurchaseResponse),
		Users:                                 make(map[string]*trainbookingpb.UserWithSeat),
		SeatCount:                             0,
	})

	log.Printf("Server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
