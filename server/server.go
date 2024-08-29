package server

import (
	"context"
	"fmt"

	trainbookingpb "github.com/Train-Booking/proto"
)

type Server struct {
	trainbookingpb.UnimplementedTrainBookingServiceServer
}

var (
	TicketId   = 1
	bookingMap = make(map[int32]trainbookingpb.TrainBooking, 0)
)

func (*Server) Purchase(ctx context.Context, req *trainbookingpb.TrainBookingRequest) (*trainbookingpb.TrainBookingResponse, error) {
	fmt.Printf("Ticket Purchase function was invoked with %v\n", req)
	booking := req.GetBooking()

	res := &trainbookingpb.TrainBookingResponse{
		FirstName: booking.FirstName,
		LastName:  booking.LastName,
		ToCity:    booking.ToCity,
		FromCity:  booking.FromCity,
		TicketId:  int32(TicketId) + 1,
		SeatNo:    booking.SeatNo,
		Amount:    booking.Amount,
	}
	bookingMap[res.TicketId] = trainbookingpb.TrainBooking{
		FirstName: res.FirstName,
		LastName:  res.LastName,
		ToCity:    res.ToCity,
		FromCity:  res.FromCity,
		TicketId:  res.TicketId,
		SeatNo:    res.SeatNo,
		Amount:    res.Amount,
	}
	return res, nil
}
