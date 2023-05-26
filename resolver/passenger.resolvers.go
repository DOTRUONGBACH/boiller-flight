package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mock_project/ent"
	graphql1 "mock_project/graphql"
	"time"
)

// ID is the resolver for the id field.
func (r *passengerResolver) ID(ctx context.Context, obj *ent.Passenger) (string, error) {
	return obj.ID.String(), nil
}

// Gender is the resolver for the gender field.
func (r *passengerResolver) Gender(ctx context.Context, obj *ent.Passenger) (ent.PassengerGender, error) {
	return ent.PassengerGender(obj.Gender), nil
}

// Dob is the resolver for the dob field.
func (r *passengerResolver) Dob(ctx context.Context, obj *ent.Passenger) (*time.Time, error) {
	return &obj.DateOfBirth, nil
}

// Passenger returns graphql1.PassengerResolver implementation.
func (r *Resolver) Passenger() graphql1.PassengerResolver { return &passengerResolver{r} }

type passengerResolver struct{ *Resolver }
