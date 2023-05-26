// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mock_project/ent/booking"
	"mock_project/ent/passenger"
	"mock_project/ent/predicate"
	"mock_project/ent/ticketowner"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PassengerUpdate is the builder for updating Passenger entities.
type PassengerUpdate struct {
	config
	hooks    []Hook
	mutation *PassengerMutation
}

// Where appends a list predicates to the PassengerUpdate builder.
func (pu *PassengerUpdate) Where(ps ...predicate.Passenger) *PassengerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PassengerUpdate) SetName(s string) *PassengerUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetCitizenID sets the "citizen_id" field.
func (pu *PassengerUpdate) SetCitizenID(s string) *PassengerUpdate {
	pu.mutation.SetCitizenID(s)
	return pu
}

// SetEmail sets the "email" field.
func (pu *PassengerUpdate) SetEmail(s string) *PassengerUpdate {
	pu.mutation.SetEmail(s)
	return pu
}

// SetPhone sets the "phone" field.
func (pu *PassengerUpdate) SetPhone(s string) *PassengerUpdate {
	pu.mutation.SetPhone(s)
	return pu
}

// SetAddress sets the "address" field.
func (pu *PassengerUpdate) SetAddress(s string) *PassengerUpdate {
	pu.mutation.SetAddress(s)
	return pu
}

// SetGender sets the "gender" field.
func (pu *PassengerUpdate) SetGender(pa passenger.Gender) *PassengerUpdate {
	pu.mutation.SetGender(pa)
	return pu
}

// SetDateOfBirth sets the "date_of_birth" field.
func (pu *PassengerUpdate) SetDateOfBirth(t time.Time) *PassengerUpdate {
	pu.mutation.SetDateOfBirth(t)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PassengerUpdate) SetUpdatedAt(t time.Time) *PassengerUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetTicketID sets the "ticket" edge to the TicketOwner entity by ID.
func (pu *PassengerUpdate) SetTicketID(id uuid.UUID) *PassengerUpdate {
	pu.mutation.SetTicketID(id)
	return pu
}

// SetNillableTicketID sets the "ticket" edge to the TicketOwner entity by ID if the given value is not nil.
func (pu *PassengerUpdate) SetNillableTicketID(id *uuid.UUID) *PassengerUpdate {
	if id != nil {
		pu = pu.SetTicketID(*id)
	}
	return pu
}

// SetTicket sets the "ticket" edge to the TicketOwner entity.
func (pu *PassengerUpdate) SetTicket(t *TicketOwner) *PassengerUpdate {
	return pu.SetTicketID(t.ID)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (pu *PassengerUpdate) AddBookingIDs(ids ...uuid.UUID) *PassengerUpdate {
	pu.mutation.AddBookingIDs(ids...)
	return pu
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (pu *PassengerUpdate) AddBookings(b ...*Booking) *PassengerUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddBookingIDs(ids...)
}

// Mutation returns the PassengerMutation object of the builder.
func (pu *PassengerUpdate) Mutation() *PassengerMutation {
	return pu.mutation
}

// ClearTicket clears the "ticket" edge to the TicketOwner entity.
func (pu *PassengerUpdate) ClearTicket() *PassengerUpdate {
	pu.mutation.ClearTicket()
	return pu
}

// ClearBookings clears all "bookings" edges to the Booking entity.
func (pu *PassengerUpdate) ClearBookings() *PassengerUpdate {
	pu.mutation.ClearBookings()
	return pu
}

// RemoveBookingIDs removes the "bookings" edge to Booking entities by IDs.
func (pu *PassengerUpdate) RemoveBookingIDs(ids ...uuid.UUID) *PassengerUpdate {
	pu.mutation.RemoveBookingIDs(ids...)
	return pu
}

// RemoveBookings removes "bookings" edges to Booking entities.
func (pu *PassengerUpdate) RemoveBookings(b ...*Booking) *PassengerUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveBookingIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PassengerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PassengerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PassengerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PassengerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PassengerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PassengerUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := passenger.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PassengerUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := passenger.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Passenger.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.CitizenID(); ok {
		if err := passenger.CitizenIDValidator(v); err != nil {
			return &ValidationError{Name: "citizen_id", err: fmt.Errorf(`ent: validator failed for field "Passenger.citizen_id": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Phone(); ok {
		if err := passenger.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Passenger.phone": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Gender(); ok {
		if err := passenger.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Passenger.gender": %w`, err)}
		}
	}
	return nil
}

func (pu *PassengerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passenger.Table,
			Columns: passenger.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: passenger.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(passenger.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.CitizenID(); ok {
		_spec.SetField(passenger.FieldCitizenID, field.TypeString, value)
	}
	if value, ok := pu.mutation.Email(); ok {
		_spec.SetField(passenger.FieldEmail, field.TypeString, value)
	}
	if value, ok := pu.mutation.Phone(); ok {
		_spec.SetField(passenger.FieldPhone, field.TypeString, value)
	}
	if value, ok := pu.mutation.Address(); ok {
		_spec.SetField(passenger.FieldAddress, field.TypeString, value)
	}
	if value, ok := pu.mutation.Gender(); ok {
		_spec.SetField(passenger.FieldGender, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.DateOfBirth(); ok {
		_spec.SetField(passenger.FieldDateOfBirth, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(passenger.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.TicketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   passenger.TicketTable,
			Columns: []string{passenger.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticketowner.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TicketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   passenger.TicketTable,
			Columns: []string{passenger.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticketowner.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   passenger.BookingsTable,
			Columns: []string{passenger.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: booking.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedBookingsIDs(); len(nodes) > 0 && !pu.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   passenger.BookingsTable,
			Columns: []string{passenger.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   passenger.BookingsTable,
			Columns: []string{passenger.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passenger.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PassengerUpdateOne is the builder for updating a single Passenger entity.
type PassengerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PassengerMutation
}

// SetName sets the "name" field.
func (puo *PassengerUpdateOne) SetName(s string) *PassengerUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetCitizenID sets the "citizen_id" field.
func (puo *PassengerUpdateOne) SetCitizenID(s string) *PassengerUpdateOne {
	puo.mutation.SetCitizenID(s)
	return puo
}

// SetEmail sets the "email" field.
func (puo *PassengerUpdateOne) SetEmail(s string) *PassengerUpdateOne {
	puo.mutation.SetEmail(s)
	return puo
}

// SetPhone sets the "phone" field.
func (puo *PassengerUpdateOne) SetPhone(s string) *PassengerUpdateOne {
	puo.mutation.SetPhone(s)
	return puo
}

// SetAddress sets the "address" field.
func (puo *PassengerUpdateOne) SetAddress(s string) *PassengerUpdateOne {
	puo.mutation.SetAddress(s)
	return puo
}

// SetGender sets the "gender" field.
func (puo *PassengerUpdateOne) SetGender(pa passenger.Gender) *PassengerUpdateOne {
	puo.mutation.SetGender(pa)
	return puo
}

// SetDateOfBirth sets the "date_of_birth" field.
func (puo *PassengerUpdateOne) SetDateOfBirth(t time.Time) *PassengerUpdateOne {
	puo.mutation.SetDateOfBirth(t)
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PassengerUpdateOne) SetUpdatedAt(t time.Time) *PassengerUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetTicketID sets the "ticket" edge to the TicketOwner entity by ID.
func (puo *PassengerUpdateOne) SetTicketID(id uuid.UUID) *PassengerUpdateOne {
	puo.mutation.SetTicketID(id)
	return puo
}

// SetNillableTicketID sets the "ticket" edge to the TicketOwner entity by ID if the given value is not nil.
func (puo *PassengerUpdateOne) SetNillableTicketID(id *uuid.UUID) *PassengerUpdateOne {
	if id != nil {
		puo = puo.SetTicketID(*id)
	}
	return puo
}

// SetTicket sets the "ticket" edge to the TicketOwner entity.
func (puo *PassengerUpdateOne) SetTicket(t *TicketOwner) *PassengerUpdateOne {
	return puo.SetTicketID(t.ID)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (puo *PassengerUpdateOne) AddBookingIDs(ids ...uuid.UUID) *PassengerUpdateOne {
	puo.mutation.AddBookingIDs(ids...)
	return puo
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (puo *PassengerUpdateOne) AddBookings(b ...*Booking) *PassengerUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddBookingIDs(ids...)
}

// Mutation returns the PassengerMutation object of the builder.
func (puo *PassengerUpdateOne) Mutation() *PassengerMutation {
	return puo.mutation
}

// ClearTicket clears the "ticket" edge to the TicketOwner entity.
func (puo *PassengerUpdateOne) ClearTicket() *PassengerUpdateOne {
	puo.mutation.ClearTicket()
	return puo
}

// ClearBookings clears all "bookings" edges to the Booking entity.
func (puo *PassengerUpdateOne) ClearBookings() *PassengerUpdateOne {
	puo.mutation.ClearBookings()
	return puo
}

// RemoveBookingIDs removes the "bookings" edge to Booking entities by IDs.
func (puo *PassengerUpdateOne) RemoveBookingIDs(ids ...uuid.UUID) *PassengerUpdateOne {
	puo.mutation.RemoveBookingIDs(ids...)
	return puo
}

// RemoveBookings removes "bookings" edges to Booking entities.
func (puo *PassengerUpdateOne) RemoveBookings(b ...*Booking) *PassengerUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveBookingIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PassengerUpdateOne) Select(field string, fields ...string) *PassengerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Passenger entity.
func (puo *PassengerUpdateOne) Save(ctx context.Context) (*Passenger, error) {
	var (
		err  error
		node *Passenger
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PassengerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Passenger)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PassengerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PassengerUpdateOne) SaveX(ctx context.Context) *Passenger {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PassengerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PassengerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PassengerUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := passenger.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PassengerUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := passenger.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Passenger.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.CitizenID(); ok {
		if err := passenger.CitizenIDValidator(v); err != nil {
			return &ValidationError{Name: "citizen_id", err: fmt.Errorf(`ent: validator failed for field "Passenger.citizen_id": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Phone(); ok {
		if err := passenger.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Passenger.phone": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Gender(); ok {
		if err := passenger.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Passenger.gender": %w`, err)}
		}
	}
	return nil
}

func (puo *PassengerUpdateOne) sqlSave(ctx context.Context) (_node *Passenger, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passenger.Table,
			Columns: passenger.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: passenger.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Passenger.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passenger.FieldID)
		for _, f := range fields {
			if !passenger.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passenger.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(passenger.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.CitizenID(); ok {
		_spec.SetField(passenger.FieldCitizenID, field.TypeString, value)
	}
	if value, ok := puo.mutation.Email(); ok {
		_spec.SetField(passenger.FieldEmail, field.TypeString, value)
	}
	if value, ok := puo.mutation.Phone(); ok {
		_spec.SetField(passenger.FieldPhone, field.TypeString, value)
	}
	if value, ok := puo.mutation.Address(); ok {
		_spec.SetField(passenger.FieldAddress, field.TypeString, value)
	}
	if value, ok := puo.mutation.Gender(); ok {
		_spec.SetField(passenger.FieldGender, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.DateOfBirth(); ok {
		_spec.SetField(passenger.FieldDateOfBirth, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(passenger.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.TicketCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   passenger.TicketTable,
			Columns: []string{passenger.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticketowner.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TicketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   passenger.TicketTable,
			Columns: []string{passenger.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticketowner.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   passenger.BookingsTable,
			Columns: []string{passenger.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: booking.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedBookingsIDs(); len(nodes) > 0 && !puo.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   passenger.BookingsTable,
			Columns: []string{passenger.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   passenger.BookingsTable,
			Columns: []string{passenger.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Passenger{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passenger.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
