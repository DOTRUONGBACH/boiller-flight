package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mock_project/ent"
	graphql1 "mock_project/graphql"
)

// ID is the resolver for the id field.
func (r *accountResolver) ID(ctx context.Context, obj *ent.Account) (string, error) {
	return obj.ID.String(), nil
}

// OldEmail is the resolver for the oldEmail field.
func (r *accountResolver) OldEmail(ctx context.Context, obj *ent.Account) (string, error) {
	return "AAAAAAa", nil
}

// Role is the resolver for the role field.
func (r *accountResolver) Role(ctx context.Context, obj *ent.Account) (ent.Role, error) {
	return ent.Role(obj.Role), nil
}

// Account returns graphql1.AccountResolver implementation.
func (r *Resolver) Account() graphql1.AccountResolver { return &accountResolver{r} }

type accountResolver struct{ *Resolver }
