package services

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/flight"
	"mock_project/ent/ticket"

	// "mock_project/ent/ticket"
	"mock_project/pb"

	"github.com/google/uuid"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

type TicketServiceClient interface {
	GetTicketById(ctx context.Context, id uuid.UUID) (*ent.Ticket, error)
	CreateTicket(ctx context.Context, input ent.NewTicketInput) (*ent.Ticket, error)
	UpdateTicketById(ctx context.Context, input ent.UpdateTicketInputByID) (*ent.Ticket, error)
	UpdateTicketStatusByFlightId(ctx context.Context, input ent.UpdateTicketStatusInput) (bool, error)
	DeleteTicket(ctx context.Context,  id string) (bool, error)
}

type TicketHandler struct {
	TicketClient pb.TicketServiceRPCClient
}

func NewTicketHandler(TicketClient pb.TicketServiceRPCClient) TicketServiceClient {
	return &TicketHandler{
		TicketClient: TicketClient,
	}
}

func (h *TicketHandler) CreateTicket(ctx context.Context, input ent.NewTicketInput) (*ent.Ticket, error) {
	res, err := h.TicketClient.CreateTicket(ctx, &pb.CreateTicketRequest{
		TicketInput: &pb.TicketInput{
		Class: pb.TicketInput_Class(pb.TicketInput_Class_value[input.Class.String()]),
		Status: pb.TicketInput_TicketStatus(pb.TicketInput_TicketStatus_value[input.Status.String()]),
		FlightId: input.FlightID,
		BookingId: input.BookingID,
	}})

	if err != nil {
		return nil, err
	}

	return &ent.Ticket{
		ID: uuid.MustParse(res.Id),
	    Class: ticket.Class(res.Class.String()),
		Status: ticket.Status(res.Status.String()),
		CreatedAt: res.CreatedAt.AsTime(),
		UpdatedAt: res.UpdatedAt.AsTime(),
	}, nil
}

func (h *TicketHandler) GetTicketById(ctx context.Context, id uuid.UUID) (*ent.Ticket, error) {
	res, err := h.TicketClient.GetTicketById(ctx, &pb.GetTicketByIdRequest{Id: id.String()})
	if err != nil {
		log.Println("Can not find desired Ticket", err)
		return nil, err
	}
	return &ent.Ticket{
		ID: uuid.MustParse(res.Id),
		Class: ticket.Class(res.Class.String()),
		Status: ticket.Status(res.Status.String()),
		Edges: ent.TicketEdges{FlightOwner: &ent.Flight{
			ID: uuid.MustParse(res.Flight.Id), 
			FlightCode: res.Flight.FlightCode,
			From: res.Flight.From,
			To: res.Flight.To,
			DepartureDate: res.Flight.DepartureDate.AsTime(),
			DepartureTime: res.Flight.DepartureTime.AsTime(),
			Duration: int(res.Flight.Duration),
			Capacity: int(res.Flight.Capacity),
			EconomyAvailableSeat: int(res.Flight.EcomomyAvailableSeat),
			BusinessAvailableSeat: int(res.Flight.BusinessAvailableSeat),
			Status: flight.Status(res.Flight.Status.String()),
		}},
	}, nil
}

func (h *TicketHandler) UpdateTicketById(ctx context.Context, input ent.UpdateTicketInputByID) (*ent.Ticket, error) {
	res, err := h.TicketClient.UpdateTicketById(ctx, &pb.UpdateTicketByIdRequest{Id: input.ID, TicketInput: &pb.TicketInput{
		Class: pb.TicketInput_Class(pb.TicketInput_Class_value[input.Class.String()]),
		Status: pb.TicketInput_TicketStatus(pb.TicketInput_TicketStatus_value[input.Status.String()]),
		FlightId: input.FlightID,
		BookingId: input.BookingID,
	}})
	if err != nil {
		log.Println("Update Ticket failed", err)
		return nil, err
	}

	return &ent.Ticket{
		ID: uuid.MustParse(res.Id),
	    Class: ticket.Class(res.Class.String()),
		Status: ticket.Status(res.Status.String()),
		CreatedAt: res.CreatedAt.AsTime(),
		UpdatedAt: res.UpdatedAt.AsTime(),
	}, nil
}

func (h *TicketHandler) UpdateTicketStatusByFlightId(ctx context.Context, input ent.UpdateTicketStatusInput) (bool, error) {
	res, err := h.TicketClient.UpdateTicketStatusByFlightId(ctx, &pb.UpdateTicketStatusByFlightIdRequest{
		Id: input.FlightID,
		Status: pb.UpdateTicketStatusByFlightIdRequest_TicketStatus(pb.UpdateTicketStatusByFlightIdRequest_TicketStatus_value[input.Status.String()]),
	})
	if err != nil {
		return res.IsUpdated, err
	}
	return res.IsUpdated, nil
}

func (h *TicketHandler) DeleteTicket(ctx context.Context, id string) (bool, error) {
	res, err := h.TicketClient.DeleteTicket(ctx, &pb.DeleteTicketRequest{Id: id})
	if err != nil {
		return res.IsDeleted, err
	}
	return res.IsDeleted, nil
}
