package handler

import (
	"context"
	"log"
	_ "mock_project/ent/passenger"
	"mock_project/grpc/passenger_grpc/repository"
	"mock_project/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PassengerServiceImp interface {
	GetPassengerByCitizenId(ctx context.Context, model *pb.GetPassengerByCitizenIdRequest) (*pb.Passenger, error)
	CreatePassenger(ctx context.Context, model *pb.CreatePassengerRequest) (*pb.Passenger, error)
	UpdatePassengerById(ctx context.Context, model *pb.UpdatePassengerByIdRequest) (*pb.Passenger, error)
	DeletePassenger(ctx context.Context, model *pb.DeletePassengerRequest) (bool, error)
}

type PassengerService struct {
	PassengerRepository repository.PassengerRepository
	pb.UnimplementedPassengerServiceRPCServer
}

func NewPassengerrHandler(PassengerRepository repository.PassengerRepository) (*PassengerService, error) {
	return &PassengerService{
		PassengerRepository: PassengerRepository,
	}, nil
}

func (s *PassengerService) GetPassengerByCitizenId(ctx context.Context, model *pb.GetPassengerByCitizenIdRequest) (*pb.Passenger, error) {
	res, err := s.PassengerRepository.GetPassengerByCitizenId(ctx, model)
	if err != nil {
		return nil, err
	}
	log.Println("PASSENGER HANDLER",res)
	return &pb.Passenger{
		Id: res.ID.String(),
		Name: res.Name,
		CitizenId: res.CitizenID,
		Phone: res.Phone,
		Email: res.Email,
		Address: res.Address,
		Gender: pb.Passenger_Gender(pb.Passenger_Gender_value[res.Gender.String()]),
		Dob: timestamppb.New(res.DateOfBirth),
	}, nil
}

func (s *PassengerService) CreatePassenger(ctx context.Context, model *pb.CreatePassengerRequest) (*pb.Passenger, error) {
	res, err := s.PassengerRepository.CreatePassenger(ctx, model)
	if err != nil {
		log.Println("Created failed:", err)
		return nil, err
	}

	return &pb.Passenger{
		Id: res.ID.String(),
		Name: res.Name,
		CitizenId: res.CitizenID,
		Phone: res.Phone,
		Email: res.Email,
		Address: res.Address,
		Gender: pb.Passenger_Gender(pb.Passenger_Gender_value[res.Gender.String()]),
		Dob: timestamppb.New(res.DateOfBirth),
	}, nil
}


func (s *PassengerService) UpdatePassengerById(ctx context.Context, model *pb.UpdatePassengerByIdRequest) (*pb.Passenger, error) {
	res, err := s.PassengerRepository.UpdatePassengerById(ctx, model)
	if err != nil {
		log.Println("Updated failed:",err)
	}
	return &pb.Passenger{
		Id: res.ID.String(),
		Name: res.Name,
		CitizenId: res.CitizenID,
		Phone: res.Phone,
		Email: res.Email,
		Address: res.Address,
		Gender: pb.Passenger_Gender(pb.Passenger_Gender_value[res.Gender.String()]),
		Dob: timestamppb.New(res.DateOfBirth),
	}, nil
}


func (s *PassengerService) DeletePassenger(ctx context.Context, model *pb.DeletePassengerRequest) (*pb.DeletePassengerResponse, error) {
	res, err := s.PassengerRepository.DeletePassenger(ctx, model)
	if err != nil {
		log.Println("Delete error: ", err)
		return &pb.DeletePassengerResponse{
			IsDeleted: res,
		}, err
	}
	return &pb.DeletePassengerResponse{
		IsDeleted: res,
	}, nil
}
