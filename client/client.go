package client

// func SubmitPurchase(client trainbookingpb.TrainTicketServiceClient, user *trainbookingpb.User, to, from string) {
// 	ticketStream, err := client.SubmitPurchase(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error in streaming %v", err)
// 	}
// 	err = ticketStream.Send(&trainbookingpb.PurchaseRequest{
// 		User: &trainbookingpb.User{
// 			Id:        user.Id,
// 			FirstName: user.FirstName,
// 			LastName:  user.LastName,
// 			Email:     user.Email,
// 		},
// 		From: from,
// 		To:   to,
// 	})
// 	if err != nil {
// 		log.Fatalf("Error in sending request %v", err)
// 	}
// 	response, err := ticketStream.Recv()
// 	if err != nil {
// 		log.Fatalf("could not receive purchase ticket response: %v", err)
// 	}
// 	log.Printf("Response Recv %v", response)
// }

// func GetTicketDetails(client trainbookingpb.TrainTicketServiceClient, userId string) {
// 	ticket, err := client.GetTicketDetails(context.Background(), &trainbookingpb.GetTicketRequest{
// 		UserId: userId,
// 	})
// 	if err != nil {
// 		log.Fatalf("Error in Get Ticket Details%v", err)
// 	}
// 	log.Printf("Ticket details for userid %v is %v", userId, ticket)
// }

// func GetUsersBySection(client trainbookingpb.TrainTicketServiceClient, section string) {
// 	response, err := client.GetUsersBySection(context.Background(), &trainbookingpb.GetUsersBySectionRequest{
// 		Section: section,
// 	})
// 	if err != nil {
// 		log.Fatalf("could not get users by section: %v", err)
// 	}

// 	for _, userWithSeat := range response.GetUsers() {
// 		user := userWithSeat.GetUser()
// 		log.Printf("User: %s %s, Email: %s, Seat: %s\n", user.GetFirstName(), user.GetLastName(), user.GetEmail(), userWithSeat.GetSeatNumber())
// 	}
// }

// func RemoveUser(client trainbookingpb.TrainTicketServiceClient, userId string) {
// 	resp, err := client.RemoveUser(context.Background(), &trainbookingpb.RemoveUserRequest{
// 		UserId: userId,
// 	})
// 	if err != nil {
// 		log.Fatalf("could not remove users: %v", err)
// 	}
// 	log.Printf("User removed successfully %v", resp)
// }

// func ModifySeat(client trainbookingpb.TrainTicketServiceClient, userId, seatNo string) {
// 	resp, err := client.ModifySeat(context.Background(), &trainbookingpb.ModifySeatRequest{
// 		UserId:        userId,
// 		NewSeatNumber: seatNo,
// 	})
// 	if err != nil {
// 		log.Fatalf("could not modify seat: %v", err)
// 	}
// 	log.Printf("Successfully Modified seat for user %v %v", userId, resp)
// }
