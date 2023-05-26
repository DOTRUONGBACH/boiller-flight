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
func (r *ticketResolver) ID(ctx context.Context, obj *ent.Ticket) (string, error) {
	return obj.ID.String(), nil
}

// Class is the resolver for the class field.
func (r *ticketResolver) Class(ctx context.Context, obj *ent.Ticket) (ent.Class, error) {
	return ent.Class(obj.Class), nil
}

// Status is the resolver for the status field.
func (r *ticketResolver) Status(ctx context.Context, obj *ent.Ticket) (ent.TicketStatus, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Ticket returns graphql1.TicketResolver implementation.
func (r *Resolver) Ticket() graphql1.TicketResolver { return &ticketResolver{r} }

type ticketResolver struct{ *Resolver }
