package services

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/passenger"
	"mock_project/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PassengerServiceClient interface {
	GetPassengerByCitizenId(ctx context.Context, citizen_id string) (*ent.Passenger, error)
	CreatePassenger(ctx context.Context, input ent.NewPassengerInput) (*ent.Passenger, error)
	UpdatePassengerById(ctx context.Context, input *ent.UpdatePassengerByIDInput) (*ent.Passenger, error)
	DeletePassenger(ctx context.Context, id uuid.UUID) (bool, error)
}

type PassengerHandler struct {
	PassengerClient pb.PassengerServiceRPCClient
}

func NewPassengerHandler(PassengerClient pb.PassengerServiceRPCClient) PassengerServiceClient {
	return &PassengerHandler{
		PassengerClient: PassengerClient,
	}
}

func (h *PassengerHandler) GetPassengerByCitizenId(ctx context.Context, citizen_id string) (*ent.Passenger, error){
	res, err := h.PassengerClient.GetPassengerByCitizenId(ctx, &pb.GetPassengerByCitizenIdRequest{CitizenId: citizen_id})
	if err != nil {
		return nil, err
	}
	log.Println("PASSENGER SERVICE",res)	
	return &ent.Passenger{
		ID: uuid.MustParse(res.Id),
		Name: res.Name,
		CitizenID: res.CitizenId,
		Phone: res.Phone,
		Email: res.Email,
		Gender: passenger.Gender(res.Gender.String()),
		Address: res.Address,
		DateOfBirth: res.Dob.AsTime(),
	}, nil
}

func (h *PassengerHandler) CreatePassenger(ctx context.Context, input ent.NewPassengerInput) (*ent.Passenger, error){
	res, err := h.PassengerClient.CreatePassenger(ctx, &pb.CreatePassengerRequest{
		PassengerInput: &pb.PassengerInput{
			Name: input.Passenger.Name,
			CitizenId: input.Passenger.CitizenID,
			Phone: input.Passenger.Phone,
			Email: input.Passenger.Email,
			Address: input.Passenger.Address,
			Gender: pb.PassengerInput_Gender(pb.Passenger_Gender_value[string(input.Passenger.Gender)]),
			Dob: timestamppb.New(input.Passenger.Dob),
		},
	})

	if err != nil {
		return nil, err
	}
	return &ent.Passenger{
		ID: uuid.MustParse(res.Id),
		Name: res.Name,
		CitizenID: res.CitizenId,
		Phone: res.Phone,
		Email: res.Email,
		Gender: passenger.Gender(res.Gender.String()),
		Address: res.Address,
		DateOfBirth: res.Dob.AsTime(),
	}, nil
}

func (h *PassengerHandler) UpdatePassengerById(ctx context.Context, input *ent.UpdatePassengerByIDInput) (*ent.Passenger, error){
	res, err := h.PassengerClient.UpdatePassengerById(ctx , &pb.UpdatePassengerByIdRequest{
			Id: input.ID,
			PassengerInput: &pb.PassengerInput{
			Name: input.Passenger.Name,
			CitizenId: input.Passenger.CitizenID,
			Phone: input.Passenger.Phone,
			Email: input.Passenger.Email,
			Address: input.Passenger.Address,
			Gender: pb.PassengerInput_Gender(pb.Passenger_Gender_value[string(input.Passenger.Gender)]),
			Dob: timestamppb.New(input.Passenger.Dob),
		},
	})
	if err != nil {
		return nil, err
	}
	return &ent.Passenger{
		ID: uuid.MustParse(res.Id),
		Name: res.Name,
		CitizenID: res.CitizenId,
		Phone: res.Phone,
		Email: res.Email,
		Gender: passenger.Gender(res.Gender.String()),
		Address: res.Address,
		DateOfBirth: res.Dob.AsTime(),
	}, nil
}

func (h *PassengerHandler) DeletePassenger(ctx context.Context, id uuid.UUID) (bool, error){
	res, err := h.PassengerClient.DeletePassenger(ctx, &pb.DeletePassengerRequest{Id: id.String()})
	if err != nil {
		return res.IsDeleted, err
	}
	return res.IsDeleted, nil
}