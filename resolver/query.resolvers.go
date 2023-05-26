package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mock_project/ent"
	graphql1 "mock_project/graphql"

	"github.com/google/uuid"
)

// Accounts is the resolver for the Accounts field.
func (r *accountQueryResolver) Accounts(ctx context.Context, obj *ent.AccountQuery, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.AccountOrder) (*ent.AccountConnection, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }
	return r.client.Account.Query().Paginate(ctx, after, first, before, last, ent.WithAccountOrder(orderBy))
}

// GetAccountByID is the resolver for the GetAccountByID field.
func (r *accountQueryResolver) GetAccountByID(ctx context.Context, obj *ent.AccountQuery, input string) (*ent.AccountByIDResponse, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }
	res, err := r.accountService.GetAccountByID(ctx, uuid.MustParse(input))
	return &ent.AccountByIDResponse{
		Email:    res.Email,
		Role:     ent.Role(res.Role),
		AccOwner: &ent.Customer{ID: res.Edges.AccOwner.ID, Name: res.Edges.AccOwner.Name, CitizenID: res.Edges.AccOwner.CitizenID, Phone: res.Edges.AccOwner.Phone, Address: res.Edges.AccOwner.Address, Gender: res.Edges.AccOwner.Gender, DateOfBirth: res.Edges.AccOwner.DateOfBirth},
	}, err
}

// Bookings is the resolver for the Bookings field.
func (r *bookingQueryResolver) Bookings(ctx context.Context, obj *ent.BookingQuery, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.BookingOrder) (*ent.BookingConnection, error) {
	return r.client.Booking.Query().Paginate(ctx, after, first, before, last, ent.WithBookingOrder(orderBy))
}

// GetBookingByID is the resolver for the GetBookingByID field.
func (r *bookingQueryResolver) GetBookingByID(ctx context.Context, obj *ent.BookingQuery, input string) (*ent.Booking, error) {
	return r.bookingService.GetBookingByID(ctx, input)
}

// ViewBookingHistory is the resolver for the viewBookingHistory field.
func (r *bookingQueryResolver) ViewBookingHistory(ctx context.Context, obj *ent.BookingQuery, input ent.CustomerBookingHistoryInput) ([]*ent.Booking, error) {
	return r.bookingService.CustomerBookingHistory(ctx, input)
}

// Customers is the resolver for the Customers field.
func (r *customerQueryResolver) Customers(ctx context.Context, obj *ent.CustomerQuery, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CustomerOrder) (*ent.CustomerConnection, error) {
	// is_authenticated := auth.ForContext(ctx)
	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }
	return r.client.Customer.Query().Paginate(ctx, after, first, before, last, ent.WithCustomerOrder(orderBy))
}

// Flights is the resolver for the Flights field.
func (r *flightQueryResolver) Flights(ctx context.Context, obj *ent.FlightQuery, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.FlightOrder) (*ent.FlightConnection, error) {
	return r.client.Flight.Query().Paginate(ctx, after, first, before, last, ent.WithFlightOrder(orderBy))
}

// FlightsByDateRange is the resolver for the FlightsByDateRange field.
func (r *flightQueryResolver) FlightsByDateRange(ctx context.Context, obj *ent.FlightQuery, input ent.GetFlightByDateRangeInput) ([]*ent.Flight, error) {
	return r.flightService.FindFlightByDateRange(ctx, input)
}

// GetFlightByFlightCode is the resolver for the GetFlightByFlightCode field.
func (r *flightQueryResolver) GetFlightByFlightCode(ctx context.Context, obj *ent.FlightQuery, input string) (*ent.Flight, error) {
	return r.flightService.GetFlightByFlightCode(ctx, input)
}

// Passengers is the resolver for the Passengers field.
func (r *passengerQueryResolver) Passengers(ctx context.Context, obj *ent.PassengerQuery, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PassengerOrder) (*ent.PassengerConnection, error) {
	// is_authenticated := auth.ForContext(ctx)
	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }
	return r.client.Passenger.Query().Paginate(ctx, after, first, before, last, ent.WithPassengerOrder(orderBy))
}

// GetPassengerByCitizenID is the resolver for the GetPassengerByCitizenID field.
func (r *passengerQueryResolver) GetPassengerByCitizenID(ctx context.Context, obj *ent.PassengerQuery, input string) (*ent.Passenger, error) {
	// is_authenticated := auth.ForContext(ctx)
	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }
	return r.passengerService.GetPassengerByCitizenId(ctx, input)
}

// Account is the resolver for the Account field.
func (r *queryResolver) Account(ctx context.Context) (*ent.AccountQuery, error) {
	return &ent.AccountQuery{}, nil
}

// Customer is the resolver for the Customer field.
func (r *queryResolver) Customer(ctx context.Context) (*ent.CustomerQuery, error) {
	return &ent.CustomerQuery{}, nil
}

// Passenger is the resolver for the Passenger field.
func (r *queryResolver) Passenger(ctx context.Context) (*ent.PassengerQuery, error) {
	return &ent.PassengerQuery{}, nil
}

// Ticket is the resolver for the Ticket field.
func (r *queryResolver) Ticket(ctx context.Context) (*ent.TicketQuery, error) {
	return &ent.TicketQuery{}, nil
}

// Flight is the resolver for the Flight field.
func (r *queryResolver) Flight(ctx context.Context) (*ent.FlightQuery, error) {
	return &ent.FlightQuery{}, nil
}

// Booking is the resolver for the Booking field.
func (r *queryResolver) Booking(ctx context.Context) (*ent.BookingQuery, error) {
	return &ent.BookingQuery{}, nil
}

// Tickets is the resolver for the Tickets field.
func (r *ticketQueryResolver) Tickets(ctx context.Context, obj *ent.TicketQuery, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TicketOrder) (*ent.TicketConnection, error) {
	// is_authenticated := auth.ForContext(ctx)
	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }
	return r.client.Ticket.Query().Paginate(ctx, after, first, before, last, ent.WithTicketOrder(orderBy))
}

// GetTicketByID is the resolver for the GetTicketByID field.
func (r *ticketQueryResolver) GetTicketByID(ctx context.Context, obj *ent.TicketQuery, input string) (*ent.TicketByIDResponse, error) {
	res, err := r.ticketService.GetTicketById(ctx, uuid.MustParse(input))
	return &ent.TicketByIDResponse{
		Class: ent.Class(res.Class),
		Flight: &ent.Flight{
			ID:                    res.Edges.FlightOwner.ID,
			FlightCode:            res.Edges.FlightOwner.FlightCode,
			From:                  res.Edges.FlightOwner.From,
			To:                    res.Edges.FlightOwner.To,
			DepartureDate:         res.Edges.FlightOwner.DepartureDate,
			DepartureTime:         res.Edges.FlightOwner.DepartureTime,
			Duration:              res.Edges.FlightOwner.Duration,
			Capacity:              res.Edges.FlightOwner.Capacity,
			EconomyAvailableSeat:  res.Edges.FlightOwner.EconomyAvailableSeat,
			BusinessAvailableSeat: res.Edges.FlightOwner.BusinessAvailableSeat,
			Status:                res.Edges.FlightOwner.Status,
		},
	}, err
}

// AccountQuery returns graphql1.AccountQueryResolver implementation.
func (r *Resolver) AccountQuery() graphql1.AccountQueryResolver { return &accountQueryResolver{r} }

// BookingQuery returns graphql1.BookingQueryResolver implementation.
func (r *Resolver) BookingQuery() graphql1.BookingQueryResolver { return &bookingQueryResolver{r} }

// CustomerQuery returns graphql1.CustomerQueryResolver implementation.
func (r *Resolver) CustomerQuery() graphql1.CustomerQueryResolver { return &customerQueryResolver{r} }

// FlightQuery returns graphql1.FlightQueryResolver implementation.
func (r *Resolver) FlightQuery() graphql1.FlightQueryResolver { return &flightQueryResolver{r} }

// PassengerQuery returns graphql1.PassengerQueryResolver implementation.
func (r *Resolver) PassengerQuery() graphql1.PassengerQueryResolver {
	return &passengerQueryResolver{r}
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

// TicketQuery returns graphql1.TicketQueryResolver implementation.
func (r *Resolver) TicketQuery() graphql1.TicketQueryResolver { return &ticketQueryResolver{r} }

type accountQueryResolver struct{ *Resolver }
type bookingQueryResolver struct{ *Resolver }
type customerQueryResolver struct{ *Resolver }
type flightQueryResolver struct{ *Resolver }
type passengerQueryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type ticketQueryResolver struct{ *Resolver }
