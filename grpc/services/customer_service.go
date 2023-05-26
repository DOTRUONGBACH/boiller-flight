package services

import (
	"context"
	"mock_project/ent"
	_"mock_project/ent/customer"
	"mock_project/pb"

	"github.com/google/uuid"
	_"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerServiceClient interface {
	GetCustomerByCitizenId(ctx context.Context, citizen_id string) (*ent.Customer, error)
	// CreateCustomer(ctx context.Context, input ent.NewCustomerInput) (*ent.Customer, error)
	UpdateCustomer(ctx context.Context, id uuid.UUID) (*ent.Customer, error)
	DeleteCustomer(ctx context.Context, id uuid.UUID) (bool, error)
}

type CustomerHandler struct {
	CustomerClient pb.CustomerServiceRPCClient
}

func NewCustomerHandler(CustomerClient pb.CustomerServiceRPCClient) CustomerServiceClient {
	return &CustomerHandler{
		CustomerClient: CustomerClient,
	}
}

func (h *CustomerHandler) GetCustomerByCitizenId(ctx context.Context, citizen_id string) (*ent.Customer, error){
	res, err := h.CustomerClient.GetCustomerByCitizenId(ctx, &pb.GetCustomerByCitizenIdRequest{CitizenId: citizen_id})
	if err != nil {
		return nil, err
	}
	
	return &ent.Customer{
		ID: uuid.MustParse(res.Id),
		Name: res.Name,
		CitizenID: res.CitizenId,
		Phone: res.Phone,
		//Gender: customer.Gender(res.Gender),
		Address: res.Address,
		DateOfBirth: res.Dob.AsTime(),
	}, nil
}

// func (h *CustomerHandler) CreateCustomer(ctx context.Context, input ent.NewCustomerInput) (*ent.Customer, error){
// 	res, err := h.CustomerClient.CreateCustomer(ctx, &pb.CreateCustomerRequest{Name: input.Name, Phone: input.Phone, Address: input.Address, Gender: pb.CreateCustomerRequest_Gender(pb.CreateCustomerRequest_Gender_value[input.Gender.String()]) ,Dob: timestamppb.New(input.Dob)})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &ent.Customer{
// 		ID: uuid.MustParse(res.Id),
// 		Name: res.Name,
// 		Phone: res.Phone,
// 		// Gender: customer.Gender(res.Gender),
// 		Address: res.Address,
// 		// Role: customer.Role(res.Role),
// 		DateOfBirth: res.Dob.AsTime(),
// 	}, nil
// }

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, id uuid.UUID) (*ent.Customer, error){
	res, err := h.CustomerClient.UpdateCustomer(ctx , &pb.UpdateCustomerRequest{Id: id.String()})
	if err != nil {
		return nil, err
	}
	return &ent.Customer{
		ID: uuid.MustParse(res.Id),
		Name: res.Name,
		Phone: res.Phone,
		Address: res.Address,
		DateOfBirth: res.Dob.AsTime(),
	}, nil
}

func (h *CustomerHandler) DeleteCustomer(ctx context.Context, id uuid.UUID) (bool, error){
	res, err := h.CustomerClient.DeleteCustomer(ctx, &pb.DeleteCustomerRequest{Id: id.String()})
	if err != nil {
		return res.IsDeleted, err
	}
	return res.IsDeleted, nil
}