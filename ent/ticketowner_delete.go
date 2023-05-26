// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"mock_project/ent/predicate"
	"mock_project/ent/ticketowner"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TicketOwnerDelete is the builder for deleting a TicketOwner entity.
type TicketOwnerDelete struct {
	config
	hooks    []Hook
	mutation *TicketOwnerMutation
}

// Where appends a list predicates to the TicketOwnerDelete builder.
func (tod *TicketOwnerDelete) Where(ps ...predicate.TicketOwner) *TicketOwnerDelete {
	tod.mutation.Where(ps...)
	return tod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tod *TicketOwnerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tod.hooks) == 0 {
		affected, err = tod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TicketOwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tod.mutation = mutation
			affected, err = tod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tod.hooks) - 1; i >= 0; i-- {
			if tod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tod *TicketOwnerDelete) ExecX(ctx context.Context) int {
	n, err := tod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tod *TicketOwnerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: ticketowner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ticketowner.FieldID,
			},
		},
	}
	if ps := tod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TicketOwnerDeleteOne is the builder for deleting a single TicketOwner entity.
type TicketOwnerDeleteOne struct {
	tod *TicketOwnerDelete
}

// Exec executes the deletion query.
func (todo *TicketOwnerDeleteOne) Exec(ctx context.Context) error {
	n, err := todo.tod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ticketowner.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (todo *TicketOwnerDeleteOne) ExecX(ctx context.Context) {
	todo.tod.ExecX(ctx)
}