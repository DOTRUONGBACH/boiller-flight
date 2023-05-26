// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mock_project/ent/booking"
	"mock_project/ent/customer"
	"mock_project/ent/flight"
	"mock_project/ent/passenger"
	"mock_project/ent/predicate"
	"mock_project/ent/ticket"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// Where appends a list predicates to the BookingUpdate builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetFlightID sets the "flight_id" field.
func (bu *BookingUpdate) SetFlightID(u uuid.UUID) *BookingUpdate {
	bu.mutation.SetFlightID(u)
	return bu
}

// SetCustomerID sets the "customer_id" field.
func (bu *BookingUpdate) SetCustomerID(u uuid.UUID) *BookingUpdate {
	bu.mutation.SetCustomerID(u)
	return bu
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableCustomerID(u *uuid.UUID) *BookingUpdate {
	if u != nil {
		bu.SetCustomerID(*u)
	}
	return bu
}

// ClearCustomerID clears the value of the "customer_id" field.
func (bu *BookingUpdate) ClearCustomerID() *BookingUpdate {
	bu.mutation.ClearCustomerID()
	return bu
}

// SetPassengerID sets the "passenger_id" field.
func (bu *BookingUpdate) SetPassengerID(u uuid.UUID) *BookingUpdate {
	bu.mutation.SetPassengerID(u)
	return bu
}

// SetNillablePassengerID sets the "passenger_id" field if the given value is not nil.
func (bu *BookingUpdate) SetNillablePassengerID(u *uuid.UUID) *BookingUpdate {
	if u != nil {
		bu.SetPassengerID(*u)
	}
	return bu
}

// ClearPassengerID clears the value of the "passenger_id" field.
func (bu *BookingUpdate) ClearPassengerID() *BookingUpdate {
	bu.mutation.ClearPassengerID()
	return bu
}

// SetEconomyTickets sets the "economy_tickets" field.
func (bu *BookingUpdate) SetEconomyTickets(i int) *BookingUpdate {
	bu.mutation.ResetEconomyTickets()
	bu.mutation.SetEconomyTickets(i)
	return bu
}

// SetNillableEconomyTickets sets the "economy_tickets" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableEconomyTickets(i *int) *BookingUpdate {
	if i != nil {
		bu.SetEconomyTickets(*i)
	}
	return bu
}

// AddEconomyTickets adds i to the "economy_tickets" field.
func (bu *BookingUpdate) AddEconomyTickets(i int) *BookingUpdate {
	bu.mutation.AddEconomyTickets(i)
	return bu
}

// SetBusinessTickets sets the "business_tickets" field.
func (bu *BookingUpdate) SetBusinessTickets(i int) *BookingUpdate {
	bu.mutation.ResetBusinessTickets()
	bu.mutation.SetBusinessTickets(i)
	return bu
}

// SetNillableBusinessTickets sets the "business_tickets" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableBusinessTickets(i *int) *BookingUpdate {
	if i != nil {
		bu.SetBusinessTickets(*i)
	}
	return bu
}

// AddBusinessTickets adds i to the "business_tickets" field.
func (bu *BookingUpdate) AddBusinessTickets(i int) *BookingUpdate {
	bu.mutation.AddBusinessTickets(i)
	return bu
}

// SetStatus sets the "status" field.
func (bu *BookingUpdate) SetStatus(b booking.Status) *BookingUpdate {
	bu.mutation.SetStatus(b)
	return bu
}

// SetType sets the "type" field.
func (bu *BookingUpdate) SetType(b booking.Type) *BookingUpdate {
	bu.mutation.SetType(b)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BookingUpdate) SetUpdatedAt(t time.Time) *BookingUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetBookingFlightID sets the "booking_flight" edge to the Flight entity by ID.
func (bu *BookingUpdate) SetBookingFlightID(id uuid.UUID) *BookingUpdate {
	bu.mutation.SetBookingFlightID(id)
	return bu
}

// SetBookingFlight sets the "booking_flight" edge to the Flight entity.
func (bu *BookingUpdate) SetBookingFlight(f *Flight) *BookingUpdate {
	return bu.SetBookingFlightID(f.ID)
}

// SetCustomerBookingTicketsID sets the "customer_booking_tickets" edge to the Customer entity by ID.
func (bu *BookingUpdate) SetCustomerBookingTicketsID(id uuid.UUID) *BookingUpdate {
	bu.mutation.SetCustomerBookingTicketsID(id)
	return bu
}

// SetNillableCustomerBookingTicketsID sets the "customer_booking_tickets" edge to the Customer entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillableCustomerBookingTicketsID(id *uuid.UUID) *BookingUpdate {
	if id != nil {
		bu = bu.SetCustomerBookingTicketsID(*id)
	}
	return bu
}

// SetCustomerBookingTickets sets the "customer_booking_tickets" edge to the Customer entity.
func (bu *BookingUpdate) SetCustomerBookingTickets(c *Customer) *BookingUpdate {
	return bu.SetCustomerBookingTicketsID(c.ID)
}

// SetPassengerBookingTicketsID sets the "passenger_booking_tickets" edge to the Passenger entity by ID.
func (bu *BookingUpdate) SetPassengerBookingTicketsID(id uuid.UUID) *BookingUpdate {
	bu.mutation.SetPassengerBookingTicketsID(id)
	return bu
}

// SetNillablePassengerBookingTicketsID sets the "passenger_booking_tickets" edge to the Passenger entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillablePassengerBookingTicketsID(id *uuid.UUID) *BookingUpdate {
	if id != nil {
		bu = bu.SetPassengerBookingTicketsID(*id)
	}
	return bu
}

// SetPassengerBookingTickets sets the "passenger_booking_tickets" edge to the Passenger entity.
func (bu *BookingUpdate) SetPassengerBookingTickets(p *Passenger) *BookingUpdate {
	return bu.SetPassengerBookingTicketsID(p.ID)
}

// AddBookingTicketIDs adds the "booking_tickets" edge to the Ticket entity by IDs.
func (bu *BookingUpdate) AddBookingTicketIDs(ids ...uuid.UUID) *BookingUpdate {
	bu.mutation.AddBookingTicketIDs(ids...)
	return bu
}

// AddBookingTickets adds the "booking_tickets" edges to the Ticket entity.
func (bu *BookingUpdate) AddBookingTickets(t ...*Ticket) *BookingUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddBookingTicketIDs(ids...)
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// ClearBookingFlight clears the "booking_flight" edge to the Flight entity.
func (bu *BookingUpdate) ClearBookingFlight() *BookingUpdate {
	bu.mutation.ClearBookingFlight()
	return bu
}

// ClearCustomerBookingTickets clears the "customer_booking_tickets" edge to the Customer entity.
func (bu *BookingUpdate) ClearCustomerBookingTickets() *BookingUpdate {
	bu.mutation.ClearCustomerBookingTickets()
	return bu
}

// ClearPassengerBookingTickets clears the "passenger_booking_tickets" edge to the Passenger entity.
func (bu *BookingUpdate) ClearPassengerBookingTickets() *BookingUpdate {
	bu.mutation.ClearPassengerBookingTickets()
	return bu
}

// ClearBookingTickets clears all "booking_tickets" edges to the Ticket entity.
func (bu *BookingUpdate) ClearBookingTickets() *BookingUpdate {
	bu.mutation.ClearBookingTickets()
	return bu
}

// RemoveBookingTicketIDs removes the "booking_tickets" edge to Ticket entities by IDs.
func (bu *BookingUpdate) RemoveBookingTicketIDs(ids ...uuid.UUID) *BookingUpdate {
	bu.mutation.RemoveBookingTicketIDs(ids...)
	return bu
}

// RemoveBookingTickets removes "booking_tickets" edges to Ticket entities.
func (bu *BookingUpdate) RemoveBookingTickets(t ...*Ticket) *BookingUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveBookingTicketIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bu.defaults()
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BookingUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := booking.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookingUpdate) check() error {
	if v, ok := bu.mutation.EconomyTickets(); ok {
		if err := booking.EconomyTicketsValidator(v); err != nil {
			return &ValidationError{Name: "economy_tickets", err: fmt.Errorf(`ent: validator failed for field "Booking.economy_tickets": %w`, err)}
		}
	}
	if v, ok := bu.mutation.BusinessTickets(); ok {
		if err := booking.BusinessTicketsValidator(v); err != nil {
			return &ValidationError{Name: "business_tickets", err: fmt.Errorf(`ent: validator failed for field "Booking.business_tickets": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	if v, ok := bu.mutation.GetType(); ok {
		if err := booking.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Booking.type": %w`, err)}
		}
	}
	if _, ok := bu.mutation.BookingFlightID(); bu.mutation.BookingFlightCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Booking.booking_flight"`)
	}
	return nil
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: booking.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.EconomyTickets(); ok {
		_spec.SetField(booking.FieldEconomyTickets, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedEconomyTickets(); ok {
		_spec.AddField(booking.FieldEconomyTickets, field.TypeInt, value)
	}
	if value, ok := bu.mutation.BusinessTickets(); ok {
		_spec.SetField(booking.FieldBusinessTickets, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedBusinessTickets(); ok {
		_spec.AddField(booking.FieldBusinessTickets, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := bu.mutation.GetType(); ok {
		_spec.SetField(booking.FieldType, field.TypeEnum, value)
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
	}
	if bu.mutation.BookingFlightCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.BookingFlightTable,
			Columns: []string{booking.BookingFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flight.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BookingFlightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.BookingFlightTable,
			Columns: []string{booking.BookingFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flight.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.CustomerBookingTicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CustomerBookingTicketsTable,
			Columns: []string{booking.CustomerBookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.CustomerBookingTicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CustomerBookingTicketsTable,
			Columns: []string{booking.CustomerBookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.PassengerBookingTicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.PassengerBookingTicketsTable,
			Columns: []string{booking.PassengerBookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passenger.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.PassengerBookingTicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.PassengerBookingTicketsTable,
			Columns: []string{booking.PassengerBookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passenger.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.BookingTicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.BookingTicketsTable,
			Columns: []string{booking.BookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticket.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedBookingTicketsIDs(); len(nodes) > 0 && !bu.mutation.BookingTicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.BookingTicketsTable,
			Columns: []string{booking.BookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BookingTicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.BookingTicketsTable,
			Columns: []string{booking.BookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookingMutation
}

// SetFlightID sets the "flight_id" field.
func (buo *BookingUpdateOne) SetFlightID(u uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetFlightID(u)
	return buo
}

// SetCustomerID sets the "customer_id" field.
func (buo *BookingUpdateOne) SetCustomerID(u uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetCustomerID(u)
	return buo
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableCustomerID(u *uuid.UUID) *BookingUpdateOne {
	if u != nil {
		buo.SetCustomerID(*u)
	}
	return buo
}

// ClearCustomerID clears the value of the "customer_id" field.
func (buo *BookingUpdateOne) ClearCustomerID() *BookingUpdateOne {
	buo.mutation.ClearCustomerID()
	return buo
}

// SetPassengerID sets the "passenger_id" field.
func (buo *BookingUpdateOne) SetPassengerID(u uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetPassengerID(u)
	return buo
}

// SetNillablePassengerID sets the "passenger_id" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillablePassengerID(u *uuid.UUID) *BookingUpdateOne {
	if u != nil {
		buo.SetPassengerID(*u)
	}
	return buo
}

// ClearPassengerID clears the value of the "passenger_id" field.
func (buo *BookingUpdateOne) ClearPassengerID() *BookingUpdateOne {
	buo.mutation.ClearPassengerID()
	return buo
}

// SetEconomyTickets sets the "economy_tickets" field.
func (buo *BookingUpdateOne) SetEconomyTickets(i int) *BookingUpdateOne {
	buo.mutation.ResetEconomyTickets()
	buo.mutation.SetEconomyTickets(i)
	return buo
}

// SetNillableEconomyTickets sets the "economy_tickets" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableEconomyTickets(i *int) *BookingUpdateOne {
	if i != nil {
		buo.SetEconomyTickets(*i)
	}
	return buo
}

// AddEconomyTickets adds i to the "economy_tickets" field.
func (buo *BookingUpdateOne) AddEconomyTickets(i int) *BookingUpdateOne {
	buo.mutation.AddEconomyTickets(i)
	return buo
}

// SetBusinessTickets sets the "business_tickets" field.
func (buo *BookingUpdateOne) SetBusinessTickets(i int) *BookingUpdateOne {
	buo.mutation.ResetBusinessTickets()
	buo.mutation.SetBusinessTickets(i)
	return buo
}

// SetNillableBusinessTickets sets the "business_tickets" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableBusinessTickets(i *int) *BookingUpdateOne {
	if i != nil {
		buo.SetBusinessTickets(*i)
	}
	return buo
}

// AddBusinessTickets adds i to the "business_tickets" field.
func (buo *BookingUpdateOne) AddBusinessTickets(i int) *BookingUpdateOne {
	buo.mutation.AddBusinessTickets(i)
	return buo
}

// SetStatus sets the "status" field.
func (buo *BookingUpdateOne) SetStatus(b booking.Status) *BookingUpdateOne {
	buo.mutation.SetStatus(b)
	return buo
}

// SetType sets the "type" field.
func (buo *BookingUpdateOne) SetType(b booking.Type) *BookingUpdateOne {
	buo.mutation.SetType(b)
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BookingUpdateOne) SetUpdatedAt(t time.Time) *BookingUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetBookingFlightID sets the "booking_flight" edge to the Flight entity by ID.
func (buo *BookingUpdateOne) SetBookingFlightID(id uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetBookingFlightID(id)
	return buo
}

// SetBookingFlight sets the "booking_flight" edge to the Flight entity.
func (buo *BookingUpdateOne) SetBookingFlight(f *Flight) *BookingUpdateOne {
	return buo.SetBookingFlightID(f.ID)
}

// SetCustomerBookingTicketsID sets the "customer_booking_tickets" edge to the Customer entity by ID.
func (buo *BookingUpdateOne) SetCustomerBookingTicketsID(id uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetCustomerBookingTicketsID(id)
	return buo
}

// SetNillableCustomerBookingTicketsID sets the "customer_booking_tickets" edge to the Customer entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableCustomerBookingTicketsID(id *uuid.UUID) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetCustomerBookingTicketsID(*id)
	}
	return buo
}

// SetCustomerBookingTickets sets the "customer_booking_tickets" edge to the Customer entity.
func (buo *BookingUpdateOne) SetCustomerBookingTickets(c *Customer) *BookingUpdateOne {
	return buo.SetCustomerBookingTicketsID(c.ID)
}

// SetPassengerBookingTicketsID sets the "passenger_booking_tickets" edge to the Passenger entity by ID.
func (buo *BookingUpdateOne) SetPassengerBookingTicketsID(id uuid.UUID) *BookingUpdateOne {
	buo.mutation.SetPassengerBookingTicketsID(id)
	return buo
}

// SetNillablePassengerBookingTicketsID sets the "passenger_booking_tickets" edge to the Passenger entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillablePassengerBookingTicketsID(id *uuid.UUID) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetPassengerBookingTicketsID(*id)
	}
	return buo
}

// SetPassengerBookingTickets sets the "passenger_booking_tickets" edge to the Passenger entity.
func (buo *BookingUpdateOne) SetPassengerBookingTickets(p *Passenger) *BookingUpdateOne {
	return buo.SetPassengerBookingTicketsID(p.ID)
}

// AddBookingTicketIDs adds the "booking_tickets" edge to the Ticket entity by IDs.
func (buo *BookingUpdateOne) AddBookingTicketIDs(ids ...uuid.UUID) *BookingUpdateOne {
	buo.mutation.AddBookingTicketIDs(ids...)
	return buo
}

// AddBookingTickets adds the "booking_tickets" edges to the Ticket entity.
func (buo *BookingUpdateOne) AddBookingTickets(t ...*Ticket) *BookingUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddBookingTicketIDs(ids...)
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// ClearBookingFlight clears the "booking_flight" edge to the Flight entity.
func (buo *BookingUpdateOne) ClearBookingFlight() *BookingUpdateOne {
	buo.mutation.ClearBookingFlight()
	return buo
}

// ClearCustomerBookingTickets clears the "customer_booking_tickets" edge to the Customer entity.
func (buo *BookingUpdateOne) ClearCustomerBookingTickets() *BookingUpdateOne {
	buo.mutation.ClearCustomerBookingTickets()
	return buo
}

// ClearPassengerBookingTickets clears the "passenger_booking_tickets" edge to the Passenger entity.
func (buo *BookingUpdateOne) ClearPassengerBookingTickets() *BookingUpdateOne {
	buo.mutation.ClearPassengerBookingTickets()
	return buo
}

// ClearBookingTickets clears all "booking_tickets" edges to the Ticket entity.
func (buo *BookingUpdateOne) ClearBookingTickets() *BookingUpdateOne {
	buo.mutation.ClearBookingTickets()
	return buo
}

// RemoveBookingTicketIDs removes the "booking_tickets" edge to Ticket entities by IDs.
func (buo *BookingUpdateOne) RemoveBookingTicketIDs(ids ...uuid.UUID) *BookingUpdateOne {
	buo.mutation.RemoveBookingTicketIDs(ids...)
	return buo
}

// RemoveBookingTickets removes "booking_tickets" edges to Ticket entities.
func (buo *BookingUpdateOne) RemoveBookingTickets(t ...*Ticket) *BookingUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveBookingTicketIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookingUpdateOne) Select(field string, fields ...string) *BookingUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Booking entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	var (
		err  error
		node *Booking
	)
	buo.defaults()
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Booking)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BookingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BookingUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := booking.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookingUpdateOne) check() error {
	if v, ok := buo.mutation.EconomyTickets(); ok {
		if err := booking.EconomyTicketsValidator(v); err != nil {
			return &ValidationError{Name: "economy_tickets", err: fmt.Errorf(`ent: validator failed for field "Booking.economy_tickets": %w`, err)}
		}
	}
	if v, ok := buo.mutation.BusinessTickets(); ok {
		if err := booking.BusinessTicketsValidator(v); err != nil {
			return &ValidationError{Name: "business_tickets", err: fmt.Errorf(`ent: validator failed for field "Booking.business_tickets": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	if v, ok := buo.mutation.GetType(); ok {
		if err := booking.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Booking.type": %w`, err)}
		}
	}
	if _, ok := buo.mutation.BookingFlightID(); buo.mutation.BookingFlightCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Booking.booking_flight"`)
	}
	return nil
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (_node *Booking, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: booking.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Booking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for _, f := range fields {
			if !booking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.EconomyTickets(); ok {
		_spec.SetField(booking.FieldEconomyTickets, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedEconomyTickets(); ok {
		_spec.AddField(booking.FieldEconomyTickets, field.TypeInt, value)
	}
	if value, ok := buo.mutation.BusinessTickets(); ok {
		_spec.SetField(booking.FieldBusinessTickets, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedBusinessTickets(); ok {
		_spec.AddField(booking.FieldBusinessTickets, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := buo.mutation.GetType(); ok {
		_spec.SetField(booking.FieldType, field.TypeEnum, value)
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
	}
	if buo.mutation.BookingFlightCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.BookingFlightTable,
			Columns: []string{booking.BookingFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flight.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BookingFlightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.BookingFlightTable,
			Columns: []string{booking.BookingFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flight.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.CustomerBookingTicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CustomerBookingTicketsTable,
			Columns: []string{booking.CustomerBookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.CustomerBookingTicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.CustomerBookingTicketsTable,
			Columns: []string{booking.CustomerBookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.PassengerBookingTicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.PassengerBookingTicketsTable,
			Columns: []string{booking.PassengerBookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passenger.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.PassengerBookingTicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.PassengerBookingTicketsTable,
			Columns: []string{booking.PassengerBookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: passenger.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.BookingTicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.BookingTicketsTable,
			Columns: []string{booking.BookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticket.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedBookingTicketsIDs(); len(nodes) > 0 && !buo.mutation.BookingTicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.BookingTicketsTable,
			Columns: []string{booking.BookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BookingTicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.BookingTicketsTable,
			Columns: []string{booking.BookingTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Booking{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
