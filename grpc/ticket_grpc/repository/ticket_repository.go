package repository

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"mock_project/ent"
	"mock_project/ent/ticket"
	"mock_project/pb"
	"os"
)

type TicketRepository interface {
	CloseDB()
	GetTicketById(ctx context.Context, model *pb.GetTicketByIdRequest) (*ent.Ticket, *ent.Flight, error)
	CreateTicket(ctx context.Context, model *pb.CreateTicketRequest) (*ent.Ticket, error)
	UpdateTicketById(ctx context.Context, model *pb.UpdateTicketByIdRequest) (*ent.Ticket, error)
	UpdateTicketStatusByFlightId(ctx context.Context, model *pb.UpdateTicketStatusByFlightIdRequest) (bool, error)
	DeleteTicket(ctx context.Context, model *pb.DeleteTicketRequest) (bool, error)
}

type PostgresDB struct {
	Client *ent.Client
}

func NewPostgresDB(connection_str string) (TicketRepository, error) {
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

func (r *PostgresDB) GetTicketById(ctx context.Context, model *pb.GetTicketByIdRequest) (*ent.Ticket, *ent.Flight, error) {
	res1, err1 := r.Client.Ticket.Query().Where(ticket.ID(uuid.MustParse(model.Id))).Only(ctx)
	if err1 != nil {
		log.Println(err1)
		return nil, nil, err1
	}

	res2, err2 := res1.QueryFlightOwner().Only(ctx)

	if err2 != nil {
		log.Println("Ticket flight does not found", err2)
		return nil, nil, err1
	}

	return res1, res2, nil
}

func (r *PostgresDB) CreateTicket(ctx context.Context, model *pb.CreateTicketRequest) (*ent.Ticket, error) {
	res, err := r.Client.Ticket.Create().
		SetClass(ticket.Class(model.TicketInput.Class.String())).
		SetFlightOwnerID(uuid.MustParse(model.TicketInput.FlightId)).
		SetBookingOwnerID(uuid.MustParse(model.TicketInput.BookingId)).
		SetStatus(ticket.Status(model.TicketInput.Status.String())).
		Save(ctx)
	if err != nil {
		log.Println("CREATE TICKET FAILED REPOSITORY", err)
		return nil, err
	}
	log.Println("CREATE TICKET SUCCESSED", res)
	return res, nil
}

func (r *PostgresDB) UpdateTicketById(ctx context.Context, model *pb.UpdateTicketByIdRequest) (*ent.Ticket, error) {
	res, err := r.Client.Ticket.Query().Where(ticket.ID(uuid.MustParse(model.Id))).Only(ctx)
	if err != nil {
		log.Println("Can not find the desired Ticket", err)
		return nil, err
	}
	res, err = res.Update().
		SetClass(ticket.Class(model.TicketInput.Class.String())).
		SetFlightOwnerID(uuid.MustParse(model.TicketInput.FlightId)).
		SetBookingOwnerID(uuid.MustParse(model.TicketInput.BookingId)).
		SetStatus(ticket.Status(model.TicketInput.Status.String())).
		Save(ctx)
	if err != nil {
		log.Println("Update Ticket failed", err)
		return nil, err
	}
	return res, nil
}

func (r *PostgresDB) UpdateTicketStatusByFlightId(ctx context.Context, model *pb.UpdateTicketStatusByFlightIdRequest) (bool, error) {
	// res, err := r.Client.Ticket.Query().Where(ticket.FlightID(uuid.MustParse(model.Id))).All(ctx)
	// if err != nil {
	// 	log.Println("Can not find the desired Ticket", err)
	// 	return nil, err
	// }
	err := r.Client.Ticket.Update().
								Where(ticket.FlightID(uuid.MustParse(model.Id))).
								SetStatus(ticket.Status(model.Status.String())).
								Exec(ctx)

	if err != nil {
		log.Println("Update Ticket failed", err)
		return false, err
	}
	return true, nil
}

func (r *PostgresDB) DeleteTicket(ctx context.Context, model *pb.DeleteTicketRequest) (bool, error) {
	err := r.Client.Ticket.DeleteOneID(uuid.MustParse(model.Id)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
