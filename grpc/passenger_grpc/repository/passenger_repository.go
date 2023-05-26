package repository

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"mock_project/ent"
	"mock_project/ent/passenger"
	"mock_project/pb"
	"os"
)

type PassengerRepository interface {
	CloseDB()
	GetPassengerByCitizenId(ctx context.Context, model *pb.GetPassengerByCitizenIdRequest) (*ent.Passenger, error)
	CreatePassenger(ctx context.Context, model *pb.CreatePassengerRequest) (*ent.Passenger, error)
	UpdatePassengerById(ctx context.Context, model *pb.UpdatePassengerByIdRequest) (*ent.Passenger, error)
	DeletePassenger(ctx context.Context, model *pb.DeletePassengerRequest) (bool, error)
}

type PostgresDB struct {
	Client *ent.Client
}

// Establish a connection to database
func NewPostgresDB(connection_str string) (PassengerRepository, error) {
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

//Get Passenger information via CitizenID
func (r *PostgresDB) GetPassengerByCitizenId(ctx context.Context, model *pb.GetPassengerByCitizenIdRequest) (*ent.Passenger, error) {
	res, err := r.Client.Passenger.Query().Where(passenger.CitizenID(model.CitizenId)).Only(ctx)
	log.Println("PASSENGER REPO",res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// Create new Passenger
func (r *PostgresDB) CreatePassenger(ctx context.Context, model *pb.CreatePassengerRequest) (*ent.Passenger, error) {
	res, err := r.Client.Passenger.Create().
		SetName(model.PassengerInput.Name).
		SetCitizenID(model.PassengerInput.CitizenId).
		SetEmail(model.PassengerInput.Email).
		SetPhone(model.PassengerInput.Phone).
		SetAddress(model.PassengerInput.Address).
		SetGender(passenger.Gender(model.PassengerInput.Gender.String())).
		SetDateOfBirth(model.PassengerInput.Dob.AsTime()).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return res, nil
}

//Update the Passenger information via CitizenID
func (r *PostgresDB) UpdatePassengerById(ctx context.Context, model *pb.UpdatePassengerByIdRequest) (*ent.Passenger, error) {
	myPassenger, err := r.Client.Passenger.Query().Where(passenger.ID(uuid.MustParse(model.Id))).Only(ctx)
	if err != nil {
		log.Println("Can not find the desired Passenger", err)
		return nil, err
	}

	myPassenger, err = myPassenger.Update().
		SetName(model.PassengerInput.Name).
		SetCitizenID(model.PassengerInput.CitizenId).
		SetEmail(model.PassengerInput.Email).
		SetPhone(model.PassengerInput.Phone).
		SetAddress(model.PassengerInput.Address).
		SetGender(passenger.Gender(model.PassengerInput.Gender.String())).
		SetDateOfBirth(model.PassengerInput.Dob.AsTime()).
		Save(ctx)

	if err != nil {
		log.Println("Update Passenger failed", err)
		return nil, err
	}
	return myPassenger, nil
}

// Delete the Passenger
func (r *PostgresDB) DeletePassenger(ctx context.Context, model *pb.DeletePassengerRequest) (bool, error) {
	err := r.Client.Passenger.DeleteOneID(uuid.MustParse(model.Id)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
