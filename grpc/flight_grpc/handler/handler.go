package handler

import (
	"context"
	"log"
	_ "mock_project/ent/flight"
	"mock_project/grpc/flight_grpc/repository"
	"mock_project/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightServiceImp interface {
	GetFlightByFlightCode(ctx context.Context, model *pb.GetFlightByFlightCodeRequest) (*pb.Flight, error)
	CreateFlight(ctx context.Context, model *pb.CreateFlightRequest) (*pb.Flight, error)
	UpdateFlightByFlightCode(ctx context.Context, model *pb.UpdateFlightByFlightCodeRequest) (*pb.Flight, error)
	UpdateAvailableSeat(ctx context.Context, model *pb.UpdateFlightAvailableSeatRequest) (*pb.Flight, error)
	UpdateFlightStatus(ctx context.Context, model *pb.UpdateFlightStatusRequest) (*pb.Flight, error)
	DeleteFlight(ctx context.Context, model *pb.DeleteFlightRequest) (bool, error)
	FindFlightByDateRange(ctx context.Context, model *pb.FindFlightByDateRangeRequest) (*pb.Flights, error)
}

type FlightService struct {
	FlightRepository repository.FlightRepository
	pb.UnimplementedFlightServiceRPCServer
}

func NewFlightrHandler(FlightRepository repository.FlightRepository) (*FlightService, error) {
	return &FlightService{
		FlightRepository: FlightRepository,
	}, nil
}

func (s *FlightService) GetFlightByFlightCode(ctx context.Context, model *pb.GetFlightByFlightCodeRequest) (*pb.Flight, error) {
	
	res, err := s.FlightRepository.GetFlightByFlightCode(ctx, model)
	if err != nil {
		return nil, err
	}

	return &pb.Flight{
		Id:                    res.ID.String(),
		FlightCode:            res.FlightCode,
		From:                  res.From,
		To:                    res.To,
		DepartureDate:         timestamppb.New(res.DepartureDate),
		DepartureTime:         timestamppb.New(res.DepartureTime),
		Duration:              int32(res.Duration),
		Capacity:              int32(res.Capacity),
		EcomomyAvailableSeat:  int32(res.EconomyAvailableSeat),
		BusinessAvailableSeat: int32(res.BusinessAvailableSeat),
		Status:                pb.Flight_Status(pb.FlightInput_Status_value[res.Status.String()]),
	}, nil
}

func (s *FlightService) CreateFlight(ctx context.Context, model *pb.CreateFlightRequest) (*pb.Flight, error) {
	res, err := s.FlightRepository.CreateFlight(ctx, model)
	if err != nil {
		log.Println("Created failed:", err)
		return nil, err
	}

	return &pb.Flight{
		Id:                    res.ID.String(),
		FlightCode:            res.FlightCode,
		From:                  res.From,
		To:                    res.To,
		DepartureDate:         timestamppb.New(res.DepartureDate),
		DepartureTime:         timestamppb.New(res.DepartureTime),
		Duration:              int32(res.Duration),
		Capacity:              int32(res.Capacity),
		EcomomyAvailableSeat:  int32(res.EconomyAvailableSeat),
		BusinessAvailableSeat: int32(res.BusinessAvailableSeat),
		Status:                pb.Flight_Status(pb.FlightInput_Status_value[res.Status.String()]),
	}, nil
}

func (s *FlightService) UpdateFlightByFlightCode(ctx context.Context, model *pb.UpdateFlightByFlightCodeRequest) (*pb.Flight, error) {
	
	res, err := s.FlightRepository.UpdateFlightByFlightCode(ctx, model)
	if err != nil {
		log.Fatalf("Updated failed: %s", err)
	}

	return &pb.Flight{
		Id:                    res.ID.String(),
		FlightCode:            res.FlightCode,
		From:                  res.From,
		To:                    res.To,
		DepartureDate:         timestamppb.New(res.DepartureDate),
		DepartureTime:         timestamppb.New(res.DepartureTime),
		Duration:              int32(res.Duration),
		Capacity:              int32(res.Capacity),
		EcomomyAvailableSeat:  int32(res.EconomyAvailableSeat),
		BusinessAvailableSeat: int32(res.BusinessAvailableSeat),
		Status:                pb.Flight_Status(pb.FlightInput_Status_value[res.Status.String()]),
	}, nil
}

func (s *FlightService) UpdateAvailableSeat(ctx context.Context, model *pb.UpdateFlightAvailableSeatRequest) (*pb.Flight, error) {
	
	res, err := s.FlightRepository.UpdateAvailableSeat(ctx, model)
	if err != nil {
		log.Fatalf("Updated failed: %s", err)
	}

	return &pb.Flight{
		Id:                    res.ID.String(),
		FlightCode:            res.FlightCode,
		From:                  res.From,
		To:                    res.To,
		DepartureDate:         timestamppb.New(res.DepartureDate),
		DepartureTime:         timestamppb.New(res.DepartureTime),
		Duration:              int32(res.Duration),
		Capacity:              int32(res.Capacity),
		EcomomyAvailableSeat:  int32(res.EconomyAvailableSeat),
		BusinessAvailableSeat: int32(res.BusinessAvailableSeat),
		Status:                pb.Flight_Status(pb.FlightInput_Status_value[res.Status.String()]),
	}, nil
}

func (s *FlightService) UpdateFlightStatus(ctx context.Context, model *pb.UpdateFlightStatusRequest) (*pb.Flight, error) {

	res, err := s.FlightRepository.UpdateFlightStatus(ctx, model)
	if err != nil {
		log.Fatalf("Updated failed: %s", err)
	}

	return &pb.Flight{
		Id:                    res.ID.String(),
		FlightCode:            res.FlightCode,
		From:                  res.From,
		To:                    res.To,
		DepartureDate:         timestamppb.New(res.DepartureDate),
		DepartureTime:         timestamppb.New(res.DepartureTime),
		Duration:              int32(res.Duration),
		Capacity:              int32(res.Capacity),
		EcomomyAvailableSeat:  int32(res.EconomyAvailableSeat),
		BusinessAvailableSeat: int32(res.BusinessAvailableSeat),
		Status:                pb.Flight_Status(pb.FlightInput_Status_value[res.Status.String()]),
	}, nil
}


func (s *FlightService) DeleteFlight(ctx context.Context, model *pb.DeleteFlightRequest) (*pb.DeleteFlightResponse, error) {
	res, err := s.FlightRepository.DeleteFlight(ctx, model)
	if err != nil {
		log.Println("Delete error: ", err)
		return &pb.DeleteFlightResponse{
			IsDeleted: res,
		}, err
	}
	return &pb.DeleteFlightResponse{
		IsDeleted: res,
	}, nil
}

func (s *FlightService) FindFlightByDateRange(ctx context.Context, model *pb.FindFlightByDateRangeRequest) (*pb.Flights, error) {

	res, err := s.FlightRepository.FindFlightByDateRange(ctx, model)
	if err != nil {
		
		return nil, err
	}


	var pRes []*pb.Flight 

	for _, e := range res {
		temp := &pb.Flight{
				Id:            e.ID.String(),
				FlightCode:    e.FlightCode,
				From:          e.From,
				To:            e.To,
				DepartureDate: timestamppb.New(e.DepartureDate),
				DepartureTime: timestamppb.New(e.DepartureTime),
				Duration:      int32(e.Duration),
				Capacity:      int32(e.Capacity),
				EcomomyAvailableSeat: int32(e.EconomyAvailableSeat),
				BusinessAvailableSeat: int32(e.BusinessAvailableSeat),
				Status:        pb.Flight_Status(pb.FlightInput_Status_value[e.Status.String()]),
		}
		// pRes = append(pRes.Flight, temp)
		pRes = append(pRes, temp)
	}

	return &pb.Flights{Flights: pRes}, nil
}
