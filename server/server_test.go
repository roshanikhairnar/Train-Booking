package server

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"

	trainbookingpb "github.com/Train-Booking/proto"
	"github.com/stretchr/testify/assert"
)

var (
	modifyError = fmt.Errorf("seat is already taken by other user")
)

func TestServer_ModifySeat(t *testing.T) {
	type fields struct {
		UnimplementedTrainTicketServiceServer trainbookingpb.UnimplementedTrainTicketServiceServer
		Mu                                    sync.Mutex
		BookingMap                            map[string]*trainbookingpb.PurchaseResponse
		Users                                 map[string]*trainbookingpb.UserWithSeat
		SeatCount                             int
	}
	type args struct {
		ctx context.Context
		req *trainbookingpb.ModifySeatRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trainbookingpb.ModifySeatResponse
		wantErr error
	}{
		{
			name: "Seat is already taken",
			fields: fields{
				UnimplementedTrainTicketServiceServer: trainbookingpb.UnimplementedTrainTicketServiceServer{},
				Mu:                                    sync.Mutex{},
				BookingMap: map[string]*trainbookingpb.PurchaseResponse{
					"user-111": {
						TicketId: "TNO1121",
						From:     "London",
						To:       "France",
						User: &trainbookingpb.User{
							Id:        "user-111",
							FirstName: "kk",
							LastName:  "kk",
							Email:     "kk",
						},
						Price:      20,
						SeatNumber: "2",
						Section:    "A",
					},
					"user-121": {
						TicketId: "TNO5121",
						From:     "London",
						To:       "France",
						User: &trainbookingpb.User{
							Id:        "user-121",
							FirstName: "kyk",
							LastName:  "kyk",
							Email:     "kyk",
						},
						Price:      20,
						SeatNumber: "1",
						Section:    "A",
					}},
				Users:     map[string]*trainbookingpb.UserWithSeat{},
				SeatCount: 2,
			},
			args: args{
				ctx: context.Background(),
				req: &trainbookingpb.ModifySeatRequest{
					UserId:        "user-111",
					NewSeatNumber: "1",
				},
			},
			want: &trainbookingpb.ModifySeatResponse{
				Success: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedTrainTicketServiceServer: tt.fields.UnimplementedTrainTicketServiceServer,
				Mu:                                    tt.fields.Mu,
				BookingMap:                            tt.fields.BookingMap,
				Users:                                 tt.fields.Users,
				SeatCount:                             tt.fields.SeatCount,
			}
			got, err := s.ModifySeat(tt.args.ctx, tt.args.req)
			assert.ErrorContains(t, err, "seat is already taken by other user")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ModifySeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_RemoveUser(t *testing.T) {
	type fields struct {
		UnimplementedTrainTicketServiceServer trainbookingpb.UnimplementedTrainTicketServiceServer
		Mu                                    sync.Mutex
		BookingMap                            map[string]*trainbookingpb.PurchaseResponse
		Users                                 map[string]*trainbookingpb.UserWithSeat
		SeatCount                             int
	}
	type args struct {
		ctx context.Context
		req *trainbookingpb.RemoveUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trainbookingpb.RemoveUserResponse
		wantErr bool
	}{
		{
			name: "user not found",
			fields: fields{
				UnimplementedTrainTicketServiceServer: trainbookingpb.UnimplementedTrainTicketServiceServer{},
				Mu:                                    sync.Mutex{},
				BookingMap: map[string]*trainbookingpb.PurchaseResponse{
					"user-111": {
						TicketId: "TNO1121",
						From:     "London",
						To:       "France",
						User: &trainbookingpb.User{
							Id:        "user-111",
							FirstName: "kk",
							LastName:  "kk",
							Email:     "kk",
						},
						Price:      20,
						SeatNumber: "2",
						Section:    "A",
					},
					"user-121": {
						TicketId: "TNO5121",
						From:     "London",
						To:       "France",
						User: &trainbookingpb.User{
							Id:        "user-121",
							FirstName: "kyk",
							LastName:  "kyk",
							Email:     "kyk",
						},
						Price:      20,
						SeatNumber: "1",
						Section:    "A",
					}},
				Users:     map[string]*trainbookingpb.UserWithSeat{},
				SeatCount: 2,
			},
			args: args{
				ctx: nil,
				req: &trainbookingpb.RemoveUserRequest{
					UserId: "user-999",
				},
			},
			want: &trainbookingpb.RemoveUserResponse{
				Success: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedTrainTicketServiceServer: tt.fields.UnimplementedTrainTicketServiceServer,
				Mu:                                    tt.fields.Mu,
				BookingMap:                            tt.fields.BookingMap,
				Users:                                 tt.fields.Users,
				SeatCount:                             tt.fields.SeatCount,
			}
			got, err := s.RemoveUser(tt.args.ctx, tt.args.req)
			assert.ErrorContains(t, err, "no user found with id")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.RemoveUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetTicketDetails(t *testing.T) {
	type fields struct {
		UnimplementedTrainTicketServiceServer trainbookingpb.UnimplementedTrainTicketServiceServer
		Mu                                    sync.Mutex
		BookingMap                            map[string]*trainbookingpb.PurchaseResponse
		Users                                 map[string]*trainbookingpb.UserWithSeat
		SeatCount                             int
	}
	type args struct {
		ctx context.Context
		req *trainbookingpb.GetTicketRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trainbookingpb.GetTicketResponse
		wantErr bool
	}{

		{
			name: "user not found to get ticket details",
			fields: fields{
				UnimplementedTrainTicketServiceServer: trainbookingpb.UnimplementedTrainTicketServiceServer{},
				Mu:                                    sync.Mutex{},
				BookingMap: map[string]*trainbookingpb.PurchaseResponse{
					"user-111": {
						TicketId: "TNO1121",
						From:     "London",
						To:       "France",
						User: &trainbookingpb.User{
							Id:        "user-111",
							FirstName: "kk",
							LastName:  "kk",
							Email:     "kk",
						},
						Price:      20,
						SeatNumber: "2",
						Section:    "A",
					},
					"user-121": {
						TicketId: "TNO5121",
						From:     "London",
						To:       "France",
						User: &trainbookingpb.User{
							Id:        "user-121",
							FirstName: "kyk",
							LastName:  "kyk",
							Email:     "kyk",
						},
						Price:      20,
						SeatNumber: "1",
						Section:    "A",
					}},
				Users:     map[string]*trainbookingpb.UserWithSeat{},
				SeatCount: 2,
			},
			args: args{
				ctx: nil,
				req: &trainbookingpb.GetTicketRequest{
					UserId: "user-456874",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedTrainTicketServiceServer: tt.fields.UnimplementedTrainTicketServiceServer,
				Mu:                                    tt.fields.Mu,
				BookingMap:                            tt.fields.BookingMap,
				Users:                                 tt.fields.Users,
				SeatCount:                             tt.fields.SeatCount,
			}
			got, err := s.GetTicketDetails(tt.args.ctx, tt.args.req)
			assert.ErrorContains(t, err, "ticket not found for the userID,user has not booked ticket")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetTicketDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
