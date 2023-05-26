package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"mock_project/ent"
	graphql1 "mock_project/graphql"
)

// ID is the resolver for the id field.
func (r *bookingResolver) ID(ctx context.Context, obj *ent.Booking) (string, error) {
	return obj.ID.String(), nil
}

// TotalEconomyTicket is the resolver for the totalEconomyTicket field.
func (r *bookingResolver) TotalEconomyTicket(ctx context.Context, obj *ent.Booking) (int, error) {
	return obj.EconomyTickets, nil
}

// TotalBusinessTicket is the resolver for the totalBusinessTicket field.
func (r *bookingResolver) TotalBusinessTicket(ctx context.Context, obj *ent.Booking) (int, error) {
	return obj.BusinessTickets, nil
}

// Status is the resolver for the status field.
func (r *bookingResolver) Status(ctx context.Context, obj *ent.Booking) (ent.BookingStatus, error) {
	return ent.BookingStatus(obj.Status), nil
}

// TotalEconomyTicket is the resolver for the totalEconomyTicket field.
func (r *updateBookingInputResolver) TotalEconomyTicket(ctx context.Context, obj *ent.UpdateBookingInput, data int) error {
	panic(fmt.Errorf("not implemented: TotalEconomyTicket - totalEconomyTicket"))
}

// TotalBusinessTicket is the resolver for the totalBusinessTicket field.
func (r *updateBookingInputResolver) TotalBusinessTicket(ctx context.Context, obj *ent.UpdateBookingInput, data int) error {
	panic(fmt.Errorf("not implemented: TotalBusinessTicket - totalBusinessTicket"))
}

// Status is the resolver for the status field.
func (r *updateBookingInputResolver) Status(ctx context.Context, obj *ent.UpdateBookingInput, data ent.BookingStatus) error {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Booking returns graphql1.BookingResolver implementation.
func (r *Resolver) Booking() graphql1.BookingResolver { return &bookingResolver{r} }

// UpdateBookingInput returns graphql1.UpdateBookingInputResolver implementation.
func (r *Resolver) UpdateBookingInput() graphql1.UpdateBookingInputResolver {
	return &updateBookingInputResolver{r}
}

type bookingResolver struct{ *Resolver }
type updateBookingInputResolver struct{ *Resolver }
