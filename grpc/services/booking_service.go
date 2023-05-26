package services

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/booking"
	"mock_project/internal/util"
	"mock_project/middleware/auth"
	"mock_project/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingServiceClient interface {
	GetBookingByID(ctx context.Context, id string) (*ent.Booking, error)
	CreateBooking(ctx context.Context, input ent.NewBookingInput) (*ent.Booking, error)
	UpdateBookingStatus(ctx context.Context, input ent.UpdateBookingStatusInput) (bool, error)
	DeleteBooking(ctx context.Context, id string) (bool, error)
	CustomerBookingHistory(ctx context.Context, input ent.CustomerBookingHistoryInput) ([]*ent.Booking, error)
}

type BookingHandler struct {
	BookingClient   pb.BookingServiceRPCClient
	FlightClient    pb.FlightServiceRPCClient
	CustomerClient  pb.CustomerServiceRPCClient
	PassengerClient pb.PassengerServiceRPCClient
	TicketClient    pb.TicketServiceRPCClient
}

func NewBookingHandler(BookingClient pb.BookingServiceRPCClient, FlightClient pb.FlightServiceRPCClient, CustomerClient pb.CustomerServiceRPCClient, PassengerClient pb.PassengerServiceRPCClient, TicketClient pb.TicketServiceRPCClient) BookingServiceClient {
	return &BookingHandler{
		BookingClient:   BookingClient,
		FlightClient:    FlightClient,
		CustomerClient:  CustomerClient,
		PassengerClient: PassengerClient,
		TicketClient:    TicketClient,
	}
}

func (h *BookingHandler) CreateBooking(ctx context.Context, input ent.NewBookingInput) (*ent.Booking, error) {

	myFlight, err1 := h.FlightClient.GetFlightByFlightCode(ctx, &pb.GetFlightByFlightCodeRequest{FlightCode: input.FlightCode})

	if err1 != nil {
		return nil, util.WrapGQLInternalError(ctx)
	}

	//Check if user is an already existed customer in database
	myCustomer, err2 := h.CustomerClient.GetCustomerByCitizenId(ctx, &pb.GetCustomerByCitizenIdRequest{CitizenId: input.CitizenID})
	if err2 != nil {

		//Check if user is an already existed passenger in database
		myPassenger, err3 := h.PassengerClient.GetPassengerByCitizenId(ctx, &pb.GetPassengerByCitizenIdRequest{CitizenId: input.CitizenID})
		if err3 != nil {

			//Create new Passenger record
			newPassenger, _ := h.PassengerClient.CreatePassenger(ctx, &pb.CreatePassengerRequest{PassengerInput: &pb.PassengerInput{
				Name:      input.Name,
				CitizenId: input.CitizenID,
				Phone:     input.Phone,
				Email:     input.Email,
				Address:   input.Address,
				Gender:    pb.PassengerInput_Gender(pb.Passenger_Gender_value[string(input.Gender)]),
				Dob:       timestamppb.New(input.DateOfBirth),
			}})

			if myFlight.EcomomyAvailableSeat >= int32(input.TotalEconomyTicket) && myFlight.BusinessAvailableSeat >= int32(input.TotalBusinessTicket) && myFlight.Status.String() == "Scheduled" {
				//Create Booking
				myBooking, errBooking := h.BookingClient.CreateBooking(ctx, &pb.CreateBookingRequest{
					BookingInput: &pb.BookingInput{
						TotalEconomyTicket:  int32(input.TotalEconomyTicket),
						TotalBusinessTicket: int32(input.TotalBusinessTicket),
						Status:              pb.BookingInput_Success,
						Type:                pb.BookingInput_Type(pb.BookingInput_Type_value[string(input.Type)]),
						FlightId:            myFlight.Id,
						PassengerId:         newPassenger.Id,
					}})
				log.Println("ERROR BOOKING", errBooking)
				if errBooking != nil {
					return nil, util.WrapGQLInvalidActionError(ctx)
				}

				//Create Economy Ticket
				for i := 0; i < input.TotalEconomyTicket; i++ {
					_, err1 := h.TicketClient.CreateTicket(ctx, &pb.CreateTicketRequest{
						TicketInput: &pb.TicketInput{
							Class:     pb.TicketInput_Economy,
							Status:    pb.TicketInput_Paid,
							FlightId:  myFlight.Id,
							BookingId: myBooking.Id,
						}})
					if err1 != nil {
						log.Println("Can not create ticket", err1)
					}

				}

				//Create Business Ticket
				for j := 0; j < input.TotalBusinessTicket; j++ {
					_, err := h.TicketClient.CreateTicket(ctx, &pb.CreateTicketRequest{
						TicketInput: &pb.TicketInput{
							Class:     pb.TicketInput_Business,
							Status:    pb.TicketInput_Paid,
							FlightId:  myFlight.Id,
							BookingId: myBooking.Id,
						}})
					if err != nil {
						log.Println("Can not create ticket", err)
					}
				}

				_, err := h.FlightClient.UpdateAvailableSeat(ctx, &pb.UpdateFlightAvailableSeatRequest{
					FlightCode:            myFlight.FlightCode,
					EcomomyAvailableSeat:  myFlight.EcomomyAvailableSeat - int32(input.TotalEconomyTicket),
					BusinessAvailableSeat: myFlight.BusinessAvailableSeat - int32(input.TotalBusinessTicket),
				})
	
				if err != nil {
					log.Println("CAN NOT UPDATE FLIGHT AVAILABLE SEAT", err)
				}
	
				// if(myFlight.BusinessAvailableSeat == 0 && myFlight.EcomomyAvailableSeat==0){
				// 	log.Println("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")
				// 	_, err := h.FlightClient.UpdateFlightStatus(ctx, &pb.UpdateFlightStatusRequest{
				// 		Id: myFlight.Id,
				// 		Status: pb.UpdateFlightStatusRequest_Full,
				// 	})
	
				// 	if err != nil {
				// 		log.Println("CAN NOT UPDATE FLIGHT AVAILABLE STATUS", err)
				// 	}
				// }

				return &ent.Booking{
					ID:              uuid.MustParse(myBooking.Id),
					EconomyTickets:  int(myBooking.TotalEconomyTicket),
					BusinessTickets: int(myBooking.TotalBusinessTicket),
					Status:          booking.Status(myBooking.Status.String()),
					Type:            booking.Type(myBooking.Type.String()),
				}, nil
			} else {
				return nil, util.WrapGQLInvalidActionError(ctx)
			}
		} else {
			//Create Booking
			if myFlight.EcomomyAvailableSeat >= int32(input.TotalEconomyTicket) && myFlight.BusinessAvailableSeat >= int32(input.TotalBusinessTicket) && myFlight.Status.String() == "Scheduled" {
				//Create Booking

				myBooking, errBooking := h.BookingClient.CreateBooking(ctx, &pb.CreateBookingRequest{
					BookingInput: &pb.BookingInput{
						TotalEconomyTicket:  int32(input.TotalEconomyTicket),
						TotalBusinessTicket: int32(input.TotalBusinessTicket),
						Status:              pb.BookingInput_Success,
						Type:                pb.BookingInput_Type(pb.BookingInput_Type_value[string(input.Type)]),
						FlightId:            myFlight.Id,
						PassengerId:         myPassenger.Id,
					}})

				if errBooking != nil {
					return nil, util.WrapGQLInvalidActionError(ctx)

				}

				//Create Economy Ticket
				for i := 0; i < input.TotalEconomyTicket; i++ {
					_, err := h.TicketClient.CreateTicket(ctx, &pb.CreateTicketRequest{
						TicketInput: &pb.TicketInput{
							Class:     pb.TicketInput_Economy,
							Status:    pb.TicketInput_Paid,
							FlightId:  myFlight.Id,
							BookingId: myBooking.Id,
						}})
					if err != nil {
						log.Println("Can not create ticket", err)
					}
				}

				//Create Business Ticket
				for j := 0; j < input.TotalBusinessTicket; j++ {

					_, err := h.TicketClient.CreateTicket(ctx, &pb.CreateTicketRequest{
						TicketInput: &pb.TicketInput{
							Class:     pb.TicketInput_Business,
							Status:    pb.TicketInput_Paid,
							FlightId:  myFlight.Id,
							BookingId: myBooking.Id,
						}})

					if err != nil {
						log.Println("Can not create ticket", err)
					}
				}

				_, err := h.FlightClient.UpdateAvailableSeat(ctx, &pb.UpdateFlightAvailableSeatRequest{
					FlightCode:            myFlight.FlightCode,
					EcomomyAvailableSeat:  myFlight.EcomomyAvailableSeat - int32(input.TotalEconomyTicket),
					BusinessAvailableSeat: myFlight.BusinessAvailableSeat - int32(input.TotalBusinessTicket),
				})
	
				if err != nil {
					log.Println("CAN NOT UPDATE FLIGHT AVAILABLE SEAT", err)
				}
	
				// if(myFlight.BusinessAvailableSeat == 0 && myFlight.EcomomyAvailableSeat == 0){
				// 	log.Println("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")
				// 	_, err := h.FlightClient.UpdateFlightStatus(ctx, &pb.UpdateFlightStatusRequest{
				// 		Id: myFlight.Id,
				// 		Status: pb.UpdateFlightStatusRequest_Full,
				// 	})
	
				// 	if err != nil {
				// 		log.Println("CAN NOT UPDATE FLIGHT AVAILABLE STATUS", err)
				// 	}
				// }

				return &ent.Booking{
					ID:              uuid.MustParse(myBooking.Id),
					EconomyTickets:  int(myBooking.TotalEconomyTicket),
					BusinessTickets: int(myBooking.TotalBusinessTicket),
					Status:          booking.Status(myBooking.Status.String()),
					Type:            booking.Type(myBooking.Type.String()),
				}, nil
			} else {
				return nil, util.WrapGQLInvalidActionError(ctx)
			}
		}
	}

	is_authenticated := auth.ForContext(ctx)

	if is_authenticated != nil {
		//Create Booking
		if myFlight.EcomomyAvailableSeat >= int32(input.TotalEconomyTicket) && myFlight.BusinessAvailableSeat >= int32(input.TotalBusinessTicket) && myFlight.Status.String() == "Scheduled" {
			//Create Booking
			myBooking, errBooking := h.BookingClient.CreateBooking(ctx, &pb.CreateBookingRequest{
				BookingInput: &pb.BookingInput{
					TotalEconomyTicket:  int32(input.TotalEconomyTicket),
					TotalBusinessTicket: int32(input.TotalBusinessTicket),
					Status:              pb.BookingInput_Success,
					Type:                pb.BookingInput_Type(pb.BookingInput_Type_value[string(input.Type)]),
					FlightId:            myFlight.Id,
					CustomerId:          myCustomer.Id,
				}})

			if errBooking != nil {
				return nil, util.WrapGQLInvalidActionError(ctx)
			}

			//Create Economy Ticket
			for i := 0; i < input.TotalEconomyTicket; i++ {
				_, err := h.TicketClient.CreateTicket(ctx, &pb.CreateTicketRequest{
					TicketInput: &pb.TicketInput{
						Class:     pb.TicketInput_Economy,
						Status:    pb.TicketInput_Paid,
						FlightId:  myFlight.Id,
						BookingId: myBooking.Id,
					}})
				if err != nil {
					log.Println("Can not create ticket", err)
				}
			}

			//Create Business Ticket
			for j := 0; j < input.TotalBusinessTicket; j++ {

				_, err := h.TicketClient.CreateTicket(ctx, &pb.CreateTicketRequest{
					TicketInput: &pb.TicketInput{
						Class:     pb.TicketInput_Business,
						Status:    pb.TicketInput_Paid,
						FlightId:  myFlight.Id,
						BookingId: myBooking.Id,
					}})

				if err != nil {
					log.Println("Can not create ticket", err)
				}
			}

			_, err := h.FlightClient.UpdateAvailableSeat(ctx, &pb.UpdateFlightAvailableSeatRequest{
				FlightCode:            myFlight.FlightCode,
				EcomomyAvailableSeat:  myFlight.EcomomyAvailableSeat - int32(input.TotalEconomyTicket),
				BusinessAvailableSeat: myFlight.BusinessAvailableSeat - int32(input.TotalBusinessTicket),
			})

			if err != nil {
				log.Println("CAN NOT UPDATE FLIGHT AVAILABLE SEAT", err)
			}

			// if(myFlight.BusinessAvailableSeat == 0 && myFlight.EcomomyAvailableSeat==0){
			// 	log.Println("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")
			// 	_, err := h.FlightClient.UpdateFlightStatus(ctx, &pb.UpdateFlightStatusRequest{
			// 		Id: myFlight.Id,
			// 		Status: pb.UpdateFlightStatusRequest_Full,
			// 	})

			// 	if err != nil {
			// 		log.Println("CAN NOT UPDATE FLIGHT AVAILABLE STATUS", err)
			// 	}
			// }
			
		
			return &ent.Booking{
				ID:              uuid.MustParse(myBooking.Id),
				EconomyTickets:  int(myBooking.TotalEconomyTicket),
				BusinessTickets: int(myBooking.TotalBusinessTicket),
				Status:          booking.Status(myBooking.Status.String()),
				Type:            booking.Type(myBooking.Type.String()),
			}, nil
		} else {
			return nil, util.WrapGQLInvalidActionError(ctx)
		}
	}
	return nil, util.WrapGQLUnauthenticatedError(ctx)
}

func (h *BookingHandler) GetBookingByID(ctx context.Context, id string) (*ent.Booking, error) {
	res, err := h.BookingClient.GetBookingByID(ctx, &pb.GetBookingByIdRequest{Id: id})
	if err != nil {
		log.Println("Can not find desired Booking", err)
		return nil, err
	}
	return &ent.Booking{
		ID:              uuid.MustParse(res.Id),
		EconomyTickets:  int(res.TotalEconomyTicket),
		BusinessTickets: int(res.TotalBusinessTicket),
		Status:          booking.Status(res.Status.String()),
		Type:            booking.Type(res.Type.String()),
	}, nil
}

func (h *BookingHandler) CustomerBookingHistory(ctx context.Context, input ent.CustomerBookingHistoryInput) ([]*ent.Booking, error) {
	pRes, err := h.BookingClient.CustomerBookingHistory(ctx, &pb.CustomerBookingHistoryRequest{
		CustomerId: input.CustomerID,
	})

	log.Println("HANDLERHANDLER PRES",pRes)

	if err != nil {
		log.Println("Can not find desired Bookings", err)
		return nil, err
	}

	var res []*ent.Booking

	log.Println("HANDLERHANDLER PRES", res)


	for _, e := range pRes.Bookings {
		temp := &ent.Booking{
			ID:              uuid.MustParse(e.Id),
			EconomyTickets:  int(e.TotalEconomyTicket),
			BusinessTickets: int(e.TotalBusinessTicket),
			Status:          booking.Status(e.Status.String()),
			Type:            booking.Type(e.Type.String()),
		}
		res = append(res, temp)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *BookingHandler) UpdateBookingStatus(ctx context.Context, input ent.UpdateBookingStatusInput) (bool, error) {

	log.Println("SSSSSSSSSSSSSSSSSSSSSSSSSSSSSS", input.Status.String())
	log.Println("SSSSSSSSSSSSSSSSSSSSSSSSSSSSSS", pb.UpdateBookingStatusRequest_BookingStatus(pb.UpdateBookingStatusRequest_BookingStatus_value[input.Status.String()]))

	res, err := h.BookingClient.UpdateBookingStatus(ctx, &pb.UpdateBookingStatusRequest{
		Id:     input.ID,
		Status: pb.UpdateBookingStatusRequest_BookingStatus(pb.UpdateBookingStatusRequest_BookingStatus_value[input.Status.String()]),
	})
	if err != nil {
		return res.IsUpdated, err
	}
	return res.IsUpdated, nil
}

func (h *BookingHandler) DeleteBooking(ctx context.Context, id string) (bool, error) {
	res, err := h.BookingClient.DeleteBooking(ctx, &pb.DeleteBookingRequest{Id: id})
	if err != nil {
		return res.IsDeleted, err
	}
	return res.IsDeleted, nil
}
