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
func (r *customerResolver) ID(ctx context.Context, obj *ent.Customer) (string, error) {
	return obj.ID.String(), nil
}

// Gender is the resolver for the gender field.
func (r *customerResolver) Gender(ctx context.Context, obj *ent.Customer) (ent.CustomerGender, error) {
	return ent.CustomerGender(obj.Gender), nil
}

// Dob is the resolver for the dob field.
func (r *customerResolver) Dob(ctx context.Context, obj *ent.Customer) (*time.Time, error) {
	return &obj.DateOfBirth, nil
}

// Customer returns graphql1.CustomerResolver implementation.
func (r *Resolver) Customer() graphql1.CustomerResolver { return &customerResolver{r} }

type customerResolver struct{ *Resolver }
