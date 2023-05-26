package repository

import (
	"context"
	"log"
	"mock_project/ent"
	"mock_project/ent/booking"
	"mock_project/pb"
	"os"
	_ "github.com/lib/pq"
	"github.com/google/uuid"
)

type BookingRepository interface {
	GetBookingByID(ctx context.Context, model *pb.GetBookingByIdRequest) (*ent.Booking, error)
	CreateBooking(ctx context.Context,modelBooking *pb.CreateBookingRequest) (*ent.Booking, error)
	UpdateBookingStatus(ctx context.Context, model *pb.UpdateBookingStatusRequest) (bool, error)	
	DeleteBooking(ctx context.Context, model *pb.DeleteBookingRequest) (bool ,error)
    CustomerBookingHistory(ctx context.Context, model *pb.CustomerBookingHistoryRequest) ([]*ent.Booking, error)
	CloseDB()
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (BookingRepository, error) {
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

func (r *PostgresDB) GetBookingByID(ctx context.Context, model *pb.GetBookingByIdRequest) (*ent.Booking, error){
	res, err := r.Client.Booking.Query().Where(booking.ID(uuid.MustParse(model.Id))).Only(ctx)
	if err != nil {
		log.Println("Can not find the desired boooking",err)
		return nil, err
	}

	return res, nil
}

func (r *PostgresDB) CustomerBookingHistory(ctx context.Context, model *pb.CustomerBookingHistoryRequest) ([]*ent.Booking, error){
	res, err := r.Client.Booking.Query().Where(booking.CustomerID(uuid.MustParse(model.CustomerId))).All(ctx)


	if err != nil {
		log.Println("Can not find the desired boookings",err)
		return nil, err
	}
	return res, nil
}

func (r *PostgresDB) CreateBooking(ctx context.Context, modelBooking *pb.CreateBookingRequest) (*ent.Booking, error){

	var customerUUID *uuid.UUID
	if modelBooking.BookingInput.CustomerId == ""{
		customerUUID = nil
	}else{
		convert := uuid.MustParse(modelBooking.BookingInput.CustomerId)
		customerUUID = &convert
	}

	var passengerUUID *uuid.UUID
	if modelBooking.BookingInput.PassengerId == ""{
		passengerUUID = nil
	}else{
		convert := uuid.MustParse(modelBooking.BookingInput.PassengerId)
		passengerUUID = &convert
	}

	res, err := r.Client.Booking.Create().
		SetEconomyTickets(int(modelBooking.BookingInput.TotalEconomyTicket)).
		SetBusinessTickets(int(modelBooking.BookingInput.TotalBusinessTicket)).
		SetFlightID(uuid.MustParse(modelBooking.BookingInput.FlightId)).
		SetType(booking.Type(modelBooking.BookingInput.Type.String())).
		SetStatus(booking.Status(modelBooking.BookingInput.Status.String())).
		SetBookingFlightID(uuid.MustParse(modelBooking.BookingInput.FlightId)).
		// SetCustomerID(uuid.MustParse(modelBooking.BookingInput.PassengerId)).
		// SetPassengerID(uuid.MustParse(modelBooking.BookingInput.PassengerId)).
		SetNillableCustomerBookingTicketsID(customerUUID).
		SetNillablePassengerBookingTicketsID(passengerUUID).
		Save(ctx)
	
	if err != nil {
		log.Println("Create booking failed",err)
		return nil, err
	}

	return res, nil
}


func (r *PostgresDB) UpdateBookingStatus(ctx context.Context, model *pb.UpdateBookingStatusRequest) (bool, error) {


	err := r.Client.Booking.Update().
								Where(booking.FlightID(uuid.MustParse(model.Id))).
								SetStatus(booking.Status(model.Status.String())).
								Exec(ctx)


	if err != nil {
		log.Println("Update Booking failed", err)
		return false, err
	}
	return true, nil
}

func (r *PostgresDB) DeleteBooking(ctx context.Context, model *pb.DeleteBookingRequest) (bool, error) {
	err := r.Client.Booking.DeleteOneID(uuid.MustParse(model.Id)).Exec(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

