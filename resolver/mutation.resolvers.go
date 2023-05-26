package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"mock_project/ent"
	graphql1 "mock_project/graphql"
	"mock_project/internal/util"
	"mock_project/middleware/auth"

	"github.com/google/uuid"
)

// NewAccount is the resolver for the newAccount field.
func (r *accountOpsResolver) NewAccount(ctx context.Context, obj *ent.AccountOps, input ent.NewAccountInput) (*ent.Account, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }
	return r.accountService.CreateAccount(ctx, input)
}

// Login is the resolver for the login field.
func (r *accountOpsResolver) Login(ctx context.Context, obj *ent.AccountOps, input ent.AccountLogin) (*ent.AccountLoginResponse, error) {
	res, err := r.accountService.Login(ctx, input)
	return &ent.AccountLoginResponse{
		Token:  res.Token,
		Status: res.Status,
	}, err
}

// AccResetPassword is the resolver for the accResetPassword field.
func (r *accountOpsResolver) AccResetPassword(ctx context.Context, obj *ent.AccountOps, input ent.AccountResetPasswordInput) (string, error) {
	return r.accountService.AccountResetPassword(ctx, input)
}

// UpdateAccountRole is the resolver for the updateAccountRole field.
func (r *accountOpsResolver) UpdateAccountRole(ctx context.Context, obj *ent.AccountOps, input ent.UpdateAccountByIDInput) (*ent.Account, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }

	return r.accountService.UpdateAccount(ctx, input)
}

// UpdateAccountStatus is the resolver for the updateAccountStatus field.
func (r *accountOpsResolver) UpdateAccountStatus(ctx context.Context, obj *ent.AccountOps, input ent.UpdateAccountStatusRequest) (*ent.Account, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }

	return r.accountService.UpdateAccountStatus(ctx, input)
}

// NewBooking is the resolver for the newBooking field.
func (r *bookingOpsResolver) NewBooking(ctx context.Context, obj *ent.BookingOps, input ent.NewBookingInput) (*ent.Booking, error) {
	return r.bookingService.CreateBooking(ctx, input)
}

// UpdateBookingStatus is the resolver for the updateBookingStatus field.
func (r *bookingOpsResolver) UpdateBookingStatus(ctx context.Context, obj *ent.BookingOps, input ent.UpdateBookingStatusInput) (bool, error) {
	return r.bookingService.UpdateBookingStatus(ctx, input)
}

// DeleteCustomer is the resolver for the deleteCustomer field.
func (r *customerOpsResolver) DeleteCustomer(ctx context.Context, obj *ent.CustomerOps, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteCustomer - deleteCustomer"))
}

// NewFlight is the resolver for the newFlight field.
func (r *flightOpsResolver) NewFlight(ctx context.Context, obj *ent.FlightOps, input ent.NewFlightInput) (*ent.Flight, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }

	return r.flightService.CreateFlight(ctx, input)
}

// UpdateFlightInfor is the resolver for the updateFlightInfor field.
func (r *flightOpsResolver) UpdateFlightInfor(ctx context.Context, obj *ent.FlightOps, input ent.UpdateFlightByFlightCodeInput) (*ent.Flight, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }

	return r.flightService.UpdateFlightByFlightCode(ctx, input)
}

// UpdateFlightStatus is the resolver for the updateFlightStatus field.
func (r *flightOpsResolver) UpdateFlightStatus(ctx context.Context, obj *ent.FlightOps, input ent.UpdateFlightStatusInput) (*ent.Flight, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }

	return r.flightService.UpdateFlightStatus(ctx, input)
}

// Customer is the resolver for the Customer field.
func (r *mutationResolver) Customer(ctx context.Context) (*ent.CustomerOps, error) {
	return &ent.CustomerOps{}, nil
}

// Passenger is the resolver for the Passenger field.
func (r *mutationResolver) Passenger(ctx context.Context) (*ent.PassengerOps, error) {
	return &ent.PassengerOps{}, nil
}

// Account is the resolver for the Account field.
func (r *mutationResolver) Account(ctx context.Context) (*ent.AccountOps, error) {
	return &ent.AccountOps{}, nil
}

// Flight is the resolver for the Flight field.
func (r *mutationResolver) Flight(ctx context.Context) (*ent.FlightOps, error) {
	return &ent.FlightOps{}, nil
}

// Ticket is the resolver for the Ticket field.
func (r *mutationResolver) Ticket(ctx context.Context) (*ent.TicketOps, error) {
	return &ent.TicketOps{}, nil
}

// Booking is the resolver for the Booking field.
func (r *mutationResolver) Booking(ctx context.Context) (*ent.BookingOps, error) {
	return &ent.BookingOps{}, nil
}

// NewPassenger is the resolver for the newPassenger field.
func (r *passengerOpsResolver) NewPassenger(ctx context.Context, obj *ent.PassengerOps, input ent.NewPassengerInput) (*ent.Passenger, error) {
	return r.passengerService.CreatePassenger(ctx, input)
}

// UpdatePassenger is the resolver for the updatePassenger field.
func (r *passengerOpsResolver) UpdatePassenger(ctx context.Context, obj *ent.PassengerOps, input *ent.UpdatePassengerByIDInput) (*ent.Passenger, error) {
	panic(fmt.Errorf("not implemented: UpdatePassenger - updatePassenger"))
}

// DeletePassenger is the resolver for the deletePassenger field.
func (r *passengerOpsResolver) DeletePassenger(ctx context.Context, obj *ent.PassengerOps, id string) (bool, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return false, util.WrapGQLUnauthenticatedError(ctx)
	// }

	return r.passengerService.DeletePassenger(ctx, uuid.MustParse(id))
}

// NewTicket is the resolver for the newTicket field.
func (r *ticketOpsResolver) NewTicket(ctx context.Context, obj *ent.TicketOps, input ent.NewTicketInput) (*ent.Ticket, error) {
	return r.ticketService.CreateTicket(ctx, input)
}

// UpdateTicketByID is the resolver for the UpdateTicketByID field.
func (r *ticketOpsResolver) UpdateTicketByID(ctx context.Context, obj *ent.TicketOps, input ent.UpdateTicketInputByID) (*ent.Ticket, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }

	return r.ticketService.UpdateTicketById(ctx, input)
}

// AccountOps returns graphql1.AccountOpsResolver implementation.
func (r *Resolver) AccountOps() graphql1.AccountOpsResolver { return &accountOpsResolver{r} }

// BookingOps returns graphql1.BookingOpsResolver implementation.
func (r *Resolver) BookingOps() graphql1.BookingOpsResolver { return &bookingOpsResolver{r} }

// CustomerOps returns graphql1.CustomerOpsResolver implementation.
func (r *Resolver) CustomerOps() graphql1.CustomerOpsResolver { return &customerOpsResolver{r} }

// FlightOps returns graphql1.FlightOpsResolver implementation.
func (r *Resolver) FlightOps() graphql1.FlightOpsResolver { return &flightOpsResolver{r} }

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// PassengerOps returns graphql1.PassengerOpsResolver implementation.
func (r *Resolver) PassengerOps() graphql1.PassengerOpsResolver { return &passengerOpsResolver{r} }

// TicketOps returns graphql1.TicketOpsResolver implementation.
func (r *Resolver) TicketOps() graphql1.TicketOpsResolver { return &ticketOpsResolver{r} }

type accountOpsResolver struct{ *Resolver }
type bookingOpsResolver struct{ *Resolver }
type customerOpsResolver struct{ *Resolver }
type flightOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type passengerOpsResolver struct{ *Resolver }
type ticketOpsResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) NewPassenger(ctx context.Context, input ent.NewPassengerInput) (*ent.Passenger, error) {
	return r.passengerService.CreatePassenger(ctx, input)
}
func (r *mutationResolver) DeletePassenger(ctx context.Context, id string) (bool, error) {
	is_authenticated := auth.ForContext(ctx)

	if is_authenticated.Role != "Administrator" {
		return false, util.WrapGQLUnauthenticatedError(ctx)
	}

	return r.passengerService.DeletePassenger(ctx, uuid.MustParse(id))
}
func (r *mutationResolver) NewFlight(ctx context.Context, input ent.NewFlightInput) (*ent.Flight, error) {
	// is_authenticated := auth.ForContext(ctx)

	// if is_authenticated.Role != "Administrator" {
	// 	return nil, util.WrapGQLUnauthenticatedError(ctx)
	// }

	return r.flightService.CreateFlight(ctx, input)
}
func (r *mutationResolver) UpdateFlightInfor(ctx context.Context, input ent.UpdateFlightByFlightCodeInput) (*ent.Flight, error) {
	is_authenticated := auth.ForContext(ctx)

	if is_authenticated.Role != "Administrator" {
		return nil, util.WrapGQLUnauthenticatedError(ctx)
	}

	return r.flightService.UpdateFlightByFlightCode(ctx, input)
}
func (r *mutationResolver) UpdateFlightStatus(ctx context.Context, input ent.UpdateFlightStatusInput) (*ent.Flight, error) {
	is_authenticated := auth.ForContext(ctx)

	if is_authenticated.Role != "Administrator" {
		return nil, util.WrapGQLUnauthenticatedError(ctx)
	}

	return r.flightService.UpdateFlightStatus(ctx, input)
}
func (r *mutationResolver) NewTicket(ctx context.Context, input ent.NewTicketInput) (*ent.Ticket, error) {
	return r.ticketService.CreateTicket(ctx, input)
}
func (r *mutationResolver) UpdateTicketByID(ctx context.Context, input ent.UpdateTicketInputByID) (*ent.Ticket, error) {
	is_authenticated := auth.ForContext(ctx)

	if is_authenticated.Role != "Administrator" {
		return nil, util.WrapGQLUnauthenticatedError(ctx)
	}

	return r.ticketService.UpdateTicketById(ctx, input)
}
func (r *mutationResolver) NewBooking(ctx context.Context, input ent.NewBookingInput) (*ent.Booking, error) {
	return r.bookingService.CreateBooking(ctx, input)
}
func (r *mutationResolver) UpdateBookingStatus(ctx context.Context, input ent.UpdateBookingStatusInput) (*ent.Booking, error) {
	panic(fmt.Errorf("not implemented: UpdateBookingStatus - updateBookingStatus"))
}
func (r *mutationResolver) NewAccount(ctx context.Context, input ent.NewAccountInput) (*ent.Account, error) {

	return r.accountService.CreateAccount(ctx, input)
}
func (r *mutationResolver) Login(ctx context.Context, input ent.AccountLogin) (*ent.AccountLoginResponse, error) {
	res, err := r.accountService.Login(ctx, input)
	return &ent.AccountLoginResponse{
		Token:  res.Token,
		Status: res.Status,
	}, err
}
func (r *mutationResolver) AccResetPassword(ctx context.Context, input ent.AccountResetPasswordInput) (string, error) {
	return r.accountService.AccountResetPassword(ctx, input)
}
func (r *mutationResolver) UpdateAccountRole(ctx context.Context, input ent.UpdateAccountByIDInput) (*ent.Account, error) {

	is_authenticated := auth.ForContext(ctx)

	if is_authenticated.Role != "Administrator" {
		return nil, util.WrapGQLUnauthenticatedError(ctx)
	}

	return r.accountService.UpdateAccount(ctx, input)
}
func (r *mutationResolver) UpdateAccountStatus(ctx context.Context, input ent.UpdateAccountStatusRequest) (*ent.Account, error) {

	is_authenticated := auth.ForContext(ctx)

	if is_authenticated.Role != "Administrator" {
		return nil, util.WrapGQLUnauthenticatedError(ctx)
	}

	return r.accountService.UpdateAccountStatus(ctx, input)
}
