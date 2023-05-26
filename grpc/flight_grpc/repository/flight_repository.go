package repository

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"mock_project/ent"
	"mock_project/ent/flight"
	"mock_project/pb"
	"os"
)

type FlightRepository interface {
	CloseDB()
	GetFlightByFlightCode(ctx context.Context, model *pb.GetFlightByFlightCodeRequest) (*ent.Flight, error)
	CreateFlight(ctx context.Context, model *pb.CreateFlightRequest) (*ent.Flight, error)
	UpdateFlightByFlightCode(ctx context.Context, model *pb.UpdateFlightByFlightCodeRequest) (*ent.Flight, error)
	UpdateAvailableSeat(ctx context.Context, model *pb.UpdateFlightAvailableSeatRequest) (*ent.Flight, error)
	UpdateFlightStatus(ctx context.Context, model *pb.UpdateFlightStatusRequest) (*ent.Flight, error)
	DeleteFlight(ctx context.Context, model *pb.DeleteFlightRequest) (bool, error)
	FindFlightByDateRange(ctx context.Context, model *pb.FindFlightByDateRangeRequest) ([]*ent.Flight, error)
}

type PostgresDB struct {
	Client *ent.Client
}

// Establish a connection to database
func NewPostgresDB(connection_str string) (FlightRepository, error) {
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

//Get flight information via its ID
func (r *PostgresDB) GetFlightByFlightCode(ctx context.Context, model *pb.GetFlightByFlightCodeRequest) (*ent.Flight, error) {
	res, err := r.Client.Flight.Query().Where(flight.FlightCode(model.FlightCode)).Only(ctx)
	log.Println(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// Create new flight
func (r *PostgresDB) CreateFlight(ctx context.Context, model *pb.CreateFlightRequest) (*ent.Flight, error) {
	res, err := r.Client.Flight.Create().
			SetFlightCode(model.FlightInput.FlightCode).
			SetFrom(model.FlightInput.From).
			SetTo(model.FlightInput.To).
			SetDepartureDate(model.FlightInput.DepartureDate.AsTime()).
			SetDepartureTime(model.FlightInput.DepartureTime.AsTime()).
			SetDuration(int(model.FlightInput.Duration)).
			SetCapacity(int(model.FlightInput.Capacity)).
			SetEconomyAvailableSeat(int(model.FlightInput.EcomomyAvailableSeat)).
			SetBusinessAvailableSeat(int(model.FlightInput.BusinessAvailableSeat)).
			SetStatus(flight.Status(model.FlightInput.Status.String())).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//Update the flight information via its flight code
func (r *PostgresDB) UpdateFlightByFlightCode(ctx context.Context, model *pb.UpdateFlightByFlightCodeRequest) (*ent.Flight, error) {
	myFlight, err := r.Client.Flight.Query().Where(flight.FlightCode(model.FlightInput.FlightCode)).Only(ctx)
	if err != nil {
		log.Println("Can not find the desired flight", err)
		return nil, err
	}
	myFlight, err = myFlight.Update().
			SetFrom(model.FlightInput.From).
			SetTo(model.FlightInput.To).
			SetDepartureDate(model.FlightInput.DepartureDate.AsTime()).
			SetDepartureTime(model.FlightInput.DepartureTime.AsTime()).
			SetDuration(int(model.FlightInput.Duration)).
			SetCapacity(int(model.FlightInput.Capacity)).
			SetEconomyAvailableSeat(int(model.FlightInput.EcomomyAvailableSeat)).
			SetBusinessAvailableSeat(int(model.FlightInput.BusinessAvailableSeat)).
			SetStatus(flight.Status(model.FlightInput.Status.String())).Save(ctx)
	if err != nil {
		log.Println("Update flight failed", err)
		return nil, err
	}
	log.Println("UPDATE REPOOOOOOOOOOOOOOOOOOOOOOOOOO111",myFlight)
	return myFlight, nil
}

//Update the flight information via its flight code
func (r *PostgresDB) UpdateAvailableSeat(ctx context.Context, model *pb.UpdateFlightAvailableSeatRequest) (*ent.Flight, error) {
	myFlight, err := r.Client.Flight.Query().Where(flight.FlightCode(model.FlightCode)).Only(ctx)
	if err != nil {
		log.Println("Can not find the desired flight", err)
		return nil, err
	}

	myFlight, err = myFlight.Update().
			SetEconomyAvailableSeat(int(model.EcomomyAvailableSeat)).
			SetBusinessAvailableSeat(int(model.BusinessAvailableSeat)).Save(ctx)

	if err != nil {
		log.Println("Update flight failed", err)
		return nil, err
	}
	log.Println("UPDATE REPOOOOOOOOOOOOOOOOOOOOOOOOOO222",myFlight)
	return myFlight, nil
}

//Update the flight information via its flight code
func (r *PostgresDB) UpdateFlightStatus(ctx context.Context, model *pb.UpdateFlightStatusRequest) (*ent.Flight, error) {


	myFlight, err := r.Client.Flight.Query().Where(flight.ID(uuid.MustParse(model.Id))).Only(ctx)
	if err != nil {
		log.Println("Can not find the desired flight", err)
		return nil, err
	}

	myFlight, err = myFlight.Update().SetStatus(flight.Status(model.Status.String())).Save(ctx)
	
	if err != nil {
		log.Println("Update flight failed", err)
		return nil, err
	}
	return myFlight, nil
}

// Delete the flight
func (r *PostgresDB) DeleteFlight(ctx context.Context, model *pb.DeleteFlightRequest) (bool, error) {
	err := r.Client.Flight.DeleteOneID(uuid.MustParse(model.Id)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}


// Find the flight information by id
func (r *PostgresDB) FindFlightByDateRange(ctx context.Context, model *pb.FindFlightByDateRangeRequest) ([]*ent.Flight, error) {
	log.Println("REPOREPOREPOREPOREPO")
	res, err := r.Client.Flight.Query().Where(flight.DepartureDateGTE(model.FirstDate.AsTime()), flight.DepartureDateLTE(model.SecondDate.AsTime())).All(ctx)
	
	log.Println("REPOREPOREPOREPOREPO",res)
	
	if err != nil {
		return nil, err
	}
	return res, nil
}
