package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mock_project/ent"
	graphql1 "mock_project/graphql"
)

// ID is the resolver for the id field.
func (r *flightResolver) ID(ctx context.Context, obj *ent.Flight) (string, error) {
	return obj.ID.String(), nil
}

// Status is the resolver for the status field.
func (r *flightResolver) Status(ctx context.Context, obj *ent.Flight) (ent.FlightStatus, error) {
	return ent.FlightStatus(obj.Status), nil
}

// Flight returns graphql1.FlightResolver implementation.
func (r *Resolver) Flight() graphql1.FlightResolver { return &flightResolver{r} }

type flightResolver struct{ *Resolver }
