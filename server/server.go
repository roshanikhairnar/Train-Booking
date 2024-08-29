package server

import (
	"context"
	"fmt"
	"sync"

	trainbookingpb "github.com/Train-Booking/proto"
	"golang.org/x/exp/rand"
)

type Server struct {
	trainbookingpb.UnimplementedTrainTicketServiceServer
	Mu         sync.Mutex
	BookingMap map[string]*trainbookingpb.PurchaseResponse
	Users      map[string]*trainbookingpb.UserWithSeat
	SeatCount  int
}

func (s *Server) SubmitPurchase(ticketStream trainbookingpb.TrainTicketService_SubmitPurchaseServer) error {
	for {
		req, err := ticketStream.Recv()
		if err != nil {
			return err
		}
		s.Mu.Lock()
		seat := fmt.Sprintf("%d", s.SeatCount+1)
		s.SeatCount++

		bookingReceipt := &trainbookingpb.PurchaseResponse{
			TicketId:   fmt.Sprintf("TNO%d", rand.Intn(10000)),
			From:       req.GetFrom(),
			To:         req.GetTo(),
			User:       req.GetUser(),
			Price:      20.0,
			SeatNumber: seat,
			Section:    "A",
		}

		s.BookingMap[req.GetUser().GetId()] = bookingReceipt
		s.Users[req.GetUser().GetId()] = &trainbookingpb.UserWithSeat{
			User:       req.GetUser(),
			SeatNumber: seat,
		}
		s.Mu.Unlock()

		if err := ticketStream.Send(bookingReceipt); err != nil {
			return err
		}
	}
}

func (s *Server) GetTicketDetails(ctx context.Context, req *trainbookingpb.GetTicketRequest) (*trainbookingpb.GetTicketResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	ticket, ok := s.BookingMap[req.UserId]
	if !ok {
		return nil, fmt.Errorf("ticket not found for the userID %v", req.UserId)
	}
	resp := trainbookingpb.GetTicketResponse{
		TicketId: ticket.TicketId,
		From:     ticket.From,
		To:       ticket.To,
		User: &trainbookingpb.User{
			Id:        ticket.User.Id,
			FirstName: ticket.User.FirstName,
			LastName:  ticket.User.LastName,
			Email:     ticket.User.Email,
		},
		Price:      ticket.Price,
		SeatNumber: ticket.SeatNumber,
		Section:    ticket.Section,
	}
	return &resp, nil
}

func (s *Server) GetUsersBySection(ctx context.Context, req *trainbookingpb.GetUsersBySectionRequest) (*trainbookingpb.GetUsersBySectionResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	section := req.Section
	var users []*trainbookingpb.UserWithSeat
	for _, val := range s.BookingMap {
		if val.Section == section {
			users = append(users, &trainbookingpb.UserWithSeat{
				User:       val.User,
				SeatNumber: val.SeatNumber,
			})
		}
	}
	resp := trainbookingpb.GetUsersBySectionResponse{
		Users: users,
	}
	return &resp, nil
}

func (s *Server) RemoveUser(ctx context.Context, req *trainbookingpb.RemoveUserRequest) (*trainbookingpb.RemoveUserResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	delete(s.BookingMap, req.UserId)

	delete(s.Users, req.UserId)
	return &trainbookingpb.RemoveUserResponse{
		Success: true,
	}, nil
}

func (s *Server) ModifySeat(ctx context.Context, req *trainbookingpb.ModifySeatRequest) (*trainbookingpb.ModifySeatResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	for _, val := range s.BookingMap {
		if req.NewSeatNumber == val.SeatNumber && req.UserId != val.User.Id {
			return &trainbookingpb.ModifySeatResponse{
				Success: false,
			}, nil
		}
	}
	s.BookingMap[req.UserId].SeatNumber = req.NewSeatNumber
	s.Users[req.UserId].SeatNumber = req.NewSeatNumber
	return &trainbookingpb.ModifySeatResponse{
		Success: true,
	}, nil
}
