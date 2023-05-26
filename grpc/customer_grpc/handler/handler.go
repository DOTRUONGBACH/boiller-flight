package handler

import (
	"context"
	"log"
	_ "mock_project/ent/customer"
	"mock_project/grpc/customer_grpc/repository"
	"mock_project/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerServiceImp interface {
	GetCustomerByCitizenId(ctx context.Context, model *pb.GetCustomerByCitizenIdRequest) (*pb.Customer, error)
	CreateCustomer(ctx context.Context, model *pb.CreateCustomerRequest) (*pb.Customer, error)
	UpdateCustomer(ctx context.Context, model *pb.UpdateCustomerRequest) (*pb.Customer, error)
	DeleteCustomer(ctx context.Context, model *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error)
}

type CustomerService struct {
	customerRepository repository.CustomerRepository
	pb.UnimplementedCustomerServiceRPCServer
}

func NewCustomerrHandler(customerRepository repository.CustomerRepository) (*CustomerService, error) {
	return &CustomerService{
		customerRepository: customerRepository,
	}, nil
}

func (s *CustomerService) GetCustomerByCitizenId(ctx context.Context, model *pb.GetCustomerByCitizenIdRequest) (*pb.Customer, error) {
	res, err := s.customerRepository.GetCustomerByCitizenId(ctx, model)
	if err != nil {
		return nil, err
	}
	return &pb.Customer{
		Id:      res.ID.String(),
		Name:    res.Name,
		Phone:   res.Phone,
		Address: res.Address,
		Gender: pb.Customer_Gender(pb.CreateCustomerRequest_Gender_value[res.Gender.String()]),
		Dob:       timestamppb.New(res.DateOfBirth),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}

func (s *CustomerService) CreateCustomer(ctx context.Context, model *pb.CreateCustomerRequest) (*pb.Customer, error) {
	res, err := s.customerRepository.CreateCustomer(ctx, model)
	if err != nil {
		log.Println("Created failed:", err)
		return nil, err
	}


	return &pb.Customer{
		Id:        res.ID.String(),
		Name:      res.Name,
		Phone:     res.Phone,
		Address:   res.Address,
		Gender:    pb.Customer_Gender(pb.CreateCustomerRequest_Gender_value[res.Gender.String()]),
		Dob:       timestamppb.New(res.DateOfBirth),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, model *pb.UpdateCustomerRequest) (*pb.Customer, error) {
	res, err := s.customerRepository.UpdateCustomer(ctx, model)
	if err != nil {
		log.Fatalf("Updated failed: %s", err)
	}
	return &pb.Customer{
		Id:      res.ID.String(),
		Name:    res.Name,
		Phone:   res.Phone,
		Address: res.Address,
		// Gender: pb.Customer_Gender(res.Gender),
		Dob:       timestamppb.New(res.DateOfBirth),
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, model *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error) {
	res, err := s.customerRepository.DeleteCustomer(ctx, model)
	log.Println("--------DELETE RESULT----------:", res)
	if err != nil {
		log.Println("Delete error: ", err)
		return &pb.DeleteCustomerResponse{
			IsDeleted: res,
		}, err
	}
	return &pb.DeleteCustomerResponse{
		IsDeleted: res,
	}, nil
}
