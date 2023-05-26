package handler

import (
	"context"
	"log"
	_ "mock_project/ent/ticket"
	"mock_project/grpc/ticket_grpc/repository"
	"mock_project/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TicketServiceImp interface {
	GetTicketById(ctx context.Context, model *pb.GetTicketByIdRequest) (*pb.Ticket, error)
	CreateTicket(ctx context.Context, model *pb.CreateTicketRequest) (*pb.Ticket, error)
	UpdateTicketById(ctx context.Context, model *pb.UpdateTicketByIdRequest) (*pb.Ticket, error)
	UpdateTicketStatusByFlightId(ctx context.Context, model *pb.UpdateTicketStatusByFlightIdRequest) (*pb.UpdateTicketStatusByFlightIdResponse, error)
	DeleteTicket(ctx context.Context, model *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error)
}

type TicketService struct {
	TicketRepository repository.TicketRepository
	pb.UnimplementedTicketServiceRPCServer
}

func NewTicketrHandler(TicketRepository repository.TicketRepository) (*TicketService, error) {
	return &TicketService{
		TicketRepository: TicketRepository,
	}, nil
}

func (s *TicketService) GetTicketById(ctx context.Context, model *pb.GetTicketByIdRequest) (*pb.Ticket, error) {
	ticket, flight, err := s.TicketRepository.GetTicketById(ctx, model)
	if err != nil {
		return nil, err
	}

	return &pb.Ticket{
		Id:   	ticket.ID.String(),
		Class: 	pb.Ticket_Class(pb.Ticket_Class_value[ticket.Class.String()]),
		Flight: &pb.Flight{
			Id: flight.ID.String(), 
			FlightCode: flight.FlightCode, 
			From: flight.From, 
			To: flight.To, 
			Duration: int32(flight.Duration), 
			Capacity: int32(flight.Capacity), 
			EcomomyAvailableSeat: int32(flight.EconomyAvailableSeat),
			BusinessAvailableSeat: int32(flight.BusinessAvailableSeat),
			Status: pb.Flight_Status(pb.Flight_Status_value[flight.Status.String()]),},	
		CreatedAt: timestamppb.New(ticket.CreatedAt),
		UpdatedAt: timestamppb.New(ticket.UpdatedAt),		
	}, nil
}

func (s *TicketService) CreateTicket(ctx context.Context, model *pb.CreateTicketRequest) (*pb.Ticket, error) {
	res, err := s.TicketRepository.CreateTicket(ctx, model)
	if err != nil {
		return nil, err
	}

	return &pb.Ticket{
		Id:   	res.ID.String(),
		Class: 	pb.Ticket_Class(pb.Ticket_Class_value[res.Class.String()]),
		Status: pb.Ticket_Status(pb.Ticket_Status_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),		
	},nil	
}

func (s *TicketService) UpdateTicketById(ctx context.Context, model *pb.UpdateTicketByIdRequest) (*pb.Ticket, error) {
	res, err := s.TicketRepository.UpdateTicketById(ctx, model)
	if err != nil {
		log.Println("Updated failed:", err)
	}
	return &pb.Ticket{
		Id:   	res.ID.String(),
		Class: 	pb.Ticket_Class(pb.Ticket_Class_value[res.Class.String()]),
		Status: pb.Ticket_Status(pb.Ticket_Status_value[res.Status.String()]),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	},nil	
}

func (s *TicketService) UpdateTicketStatusByFlightId(ctx context.Context, model *pb.UpdateTicketStatusByFlightIdRequest) (*pb.UpdateTicketStatusByFlightIdResponse, error) {
	is_updated, err := s.TicketRepository.UpdateTicketStatusByFlightId(ctx, model)
	if err != nil {
		return &pb.UpdateTicketStatusByFlightIdResponse{
			IsUpdated: is_updated,
		},err	
	}
	return &pb.UpdateTicketStatusByFlightIdResponse{
			IsUpdated: is_updated,
	},nil	
}

func (s *TicketService) DeleteTicket(ctx context.Context, model *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	res, err := s.TicketRepository.DeleteTicket(ctx, model)
	if err != nil {
		log.Println("Delete error: ", err)
		return &pb.DeleteTicketResponse{
			IsDeleted: res,
		}, err
	}
	return &pb.DeleteTicketResponse{
		IsDeleted: res,
	}, nil
}
