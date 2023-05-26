package repository

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/customer"
	"mock_project/pb"
	"os"
	_ "github.com/lib/pq"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	CloseDB()
	GetCustomerByCitizenId(ctx context.Context, model *pb.GetCustomerByCitizenIdRequest) (*ent.Customer, error)
	CreateCustomer(ctx context.Context, model *pb.CreateCustomerRequest) (*ent.Customer, error)
	UpdateCustomer(ctx context.Context, model *pb.UpdateCustomerRequest) (*ent.Customer, error)
	DeleteCustomer(ctx context.Context, model *pb.DeleteCustomerRequest) (bool ,error)
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (CustomerRepository, error) {
	client, err := ent.Open("postgres", connection_str)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to creating db schema from the automation migration tool")
		os.Exit(1)
	}

	return &PostgresDB{Client: client}, nil
}

func (r *PostgresDB) CloseDB() {
	r.Client.Close()
}

func (r *PostgresDB) GetCustomerByCitizenId(ctx context.Context, model *pb.GetCustomerByCitizenIdRequest) (*ent.Customer, error) {
	res, err := r.Client.Customer.Query().Where(customer.CitizenID(model.CitizenId)).Only(ctx)
	log.Println("REPO:",res)
	if err != nil {
		log.Println("ERROR REPO", err)
		return nil, err
	}
	return res, nil
}

func (r *PostgresDB) CreateCustomer(ctx context.Context, model *pb.CreateCustomerRequest) (*ent.Customer, error) {
	res, err := r.Client.Customer.Create().SetName(model.Name).SetCitizenID(model.CitizenId).SetPhone(model.Phone).SetAddress(model.Address).SetGender(customer.Gender(model.Gender.String())).SetDateOfBirth(model.Dob.AsTime()).Save(ctx)
	if err != nil{
		return nil, err
	}
	return res, nil
}

func (r *PostgresDB) UpdateCustomer(ctx context.Context, model *pb.UpdateCustomerRequest) (*ent.Customer, error) {
	res, err := r.Client.Customer.UpdateOneID(uuid.MustParse(model.Id)).SetName(model.Name).SetCitizenID(model.CitizenId).SetPhone(model.Phone).SetAddress(model.Address).SetGender(customer.Gender(model.Gender.String())).SetDateOfBirth(model.Dob.AsTime()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *PostgresDB) DeleteCustomer(ctx context.Context, model *pb.DeleteCustomerRequest) (bool, error) {
	err := r.Client.Customer.DeleteOneID(uuid.MustParse(model.Id)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
