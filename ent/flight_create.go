// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mock_project/ent/booking"
	"mock_project/ent/flight"
	"mock_project/ent/ticket"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FlightCreate is the builder for creating a Flight entity.
type FlightCreate struct {
	config
	mutation *FlightMutation
	hooks    []Hook
}

// SetFlightCode sets the "flight_code" field.
func (fc *FlightCreate) SetFlightCode(s string) *FlightCreate {
	fc.mutation.SetFlightCode(s)
	return fc
}

// SetFrom sets the "from" field.
func (fc *FlightCreate) SetFrom(s string) *FlightCreate {
	fc.mutation.SetFrom(s)
	return fc
}

// SetTo sets the "to" field.
func (fc *FlightCreate) SetTo(s string) *FlightCreate {
	fc.mutation.SetTo(s)
	return fc
}

// SetDepartureDate sets the "departure_date" field.
func (fc *FlightCreate) SetDepartureDate(t time.Time) *FlightCreate {
	fc.mutation.SetDepartureDate(t)
	return fc
}

// SetDepartureTime sets the "departure_time" field.
func (fc *FlightCreate) SetDepartureTime(t time.Time) *FlightCreate {
	fc.mutation.SetDepartureTime(t)
	return fc
}

// SetDuration sets the "duration" field.
func (fc *FlightCreate) SetDuration(i int) *FlightCreate {
	fc.mutation.SetDuration(i)
	return fc
}

// SetCapacity sets the "capacity" field.
func (fc *FlightCreate) SetCapacity(i int) *FlightCreate {
	fc.mutation.SetCapacity(i)
	return fc
}

// SetEconomyAvailableSeat sets the "economy_available_seat" field.
func (fc *FlightCreate) SetEconomyAvailableSeat(i int) *FlightCreate {
	fc.mutation.SetEconomyAvailableSeat(i)
	return fc
}

// SetBusinessAvailableSeat sets the "business_available_seat" field.
func (fc *FlightCreate) SetBusinessAvailableSeat(i int) *FlightCreate {
	fc.mutation.SetBusinessAvailableSeat(i)
	return fc
}

// SetStatus sets the "status" field.
func (fc *FlightCreate) SetStatus(f flight.Status) *FlightCreate {
	fc.mutation.SetStatus(f)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FlightCreate) SetCreatedAt(t time.Time) *FlightCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FlightCreate) SetNillableCreatedAt(t *time.Time) *FlightCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FlightCreate) SetUpdatedAt(t time.Time) *FlightCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FlightCreate) SetNillableUpdatedAt(t *time.Time) *FlightCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FlightCreate) SetID(u uuid.UUID) *FlightCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FlightCreate) SetNillableID(u *uuid.UUID) *FlightCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// AddBelongsToIDs adds the "belongs_to" edge to the Booking entity by IDs.
func (fc *FlightCreate) AddBelongsToIDs(ids ...uuid.UUID) *FlightCreate {
	fc.mutation.AddBelongsToIDs(ids...)
	return fc
}

// AddBelongsTo adds the "belongs_to" edges to the Booking entity.
func (fc *FlightCreate) AddBelongsTo(b ...*Booking) *FlightCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fc.AddBelongsToIDs(ids...)
}

// AddFlightTicketIDs adds the "flight_tickets" edge to the Ticket entity by IDs.
func (fc *FlightCreate) AddFlightTicketIDs(ids ...uuid.UUID) *FlightCreate {
	fc.mutation.AddFlightTicketIDs(ids...)
	return fc
}

// AddFlightTickets adds the "flight_tickets" edges to the Ticket entity.
func (fc *FlightCreate) AddFlightTickets(t ...*Ticket) *FlightCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fc.AddFlightTicketIDs(ids...)
}

// Mutation returns the FlightMutation object of the builder.
func (fc *FlightCreate) Mutation() *FlightMutation {
	return fc.mutation
}

// Save creates the Flight in the database.
func (fc *FlightCreate) Save(ctx context.Context) (*Flight, error) {
	var (
		err  error
		node *Flight
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FlightMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Flight)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FlightMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FlightCreate) SaveX(ctx context.Context) *Flight {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FlightCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FlightCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FlightCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := flight.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := flight.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := flight.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FlightCreate) check() error {
	if _, ok := fc.mutation.FlightCode(); !ok {
		return &ValidationError{Name: "flight_code", err: errors.New(`ent: missing required field "Flight.flight_code"`)}
	}
	if v, ok := fc.mutation.FlightCode(); ok {
		if err := flight.FlightCodeValidator(v); err != nil {
			return &ValidationError{Name: "flight_code", err: fmt.Errorf(`ent: validator failed for field "Flight.flight_code": %w`, err)}
		}
	}
	if _, ok := fc.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "Flight.from"`)}
	}
	if v, ok := fc.mutation.From(); ok {
		if err := flight.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`ent: validator failed for field "Flight.from": %w`, err)}
		}
	}
	if _, ok := fc.mutation.To(); !ok {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required field "Flight.to"`)}
	}
	if v, ok := fc.mutation.To(); ok {
		if err := flight.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "Flight.to": %w`, err)}
		}
	}
	if _, ok := fc.mutation.DepartureDate(); !ok {
		return &ValidationError{Name: "departure_date", err: errors.New(`ent: missing required field "Flight.departure_date"`)}
	}
	if _, ok := fc.mutation.DepartureTime(); !ok {
		return &ValidationError{Name: "departure_time", err: errors.New(`ent: missing required field "Flight.departure_time"`)}
	}
	if _, ok := fc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Flight.duration"`)}
	}
	if v, ok := fc.mutation.Duration(); ok {
		if err := flight.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Flight.duration": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Capacity(); !ok {
		return &ValidationError{Name: "capacity", err: errors.New(`ent: missing required field "Flight.capacity"`)}
	}
	if v, ok := fc.mutation.Capacity(); ok {
		if err := flight.CapacityValidator(v); err != nil {
			return &ValidationError{Name: "capacity", err: fmt.Errorf(`ent: validator failed for field "Flight.capacity": %w`, err)}
		}
	}
	if _, ok := fc.mutation.EconomyAvailableSeat(); !ok {
		return &ValidationError{Name: "economy_available_seat", err: errors.New(`ent: missing required field "Flight.economy_available_seat"`)}
	}
	if v, ok := fc.mutation.EconomyAvailableSeat(); ok {
		if err := flight.EconomyAvailableSeatValidator(v); err != nil {
			return &ValidationError{Name: "economy_available_seat", err: fmt.Errorf(`ent: validator failed for field "Flight.economy_available_seat": %w`, err)}
		}
	}
	if _, ok := fc.mutation.BusinessAvailableSeat(); !ok {
		return &ValidationError{Name: "business_available_seat", err: errors.New(`ent: missing required field "Flight.business_available_seat"`)}
	}
	if v, ok := fc.mutation.BusinessAvailableSeat(); ok {
		if err := flight.BusinessAvailableSeatValidator(v); err != nil {
			return &ValidationError{Name: "business_available_seat", err: fmt.Errorf(`ent: validator failed for field "Flight.business_available_seat": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Flight.status"`)}
	}
	if v, ok := fc.mutation.Status(); ok {
		if err := flight.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Flight.status": %w`, err)}
		}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Flight.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Flight.updated_at"`)}
	}
	return nil
}

func (fc *FlightCreate) sqlSave(ctx context.Context) (*Flight, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (fc *FlightCreate) createSpec() (*Flight, *sqlgraph.CreateSpec) {
	var (
		_node = &Flight{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flight.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: flight.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.FlightCode(); ok {
		_spec.SetField(flight.FieldFlightCode, field.TypeString, value)
		_node.FlightCode = value
	}
	if value, ok := fc.mutation.From(); ok {
		_spec.SetField(flight.FieldFrom, field.TypeString, value)
		_node.From = value
	}
	if value, ok := fc.mutation.To(); ok {
		_spec.SetField(flight.FieldTo, field.TypeString, value)
		_node.To = value
	}
	if value, ok := fc.mutation.DepartureDate(); ok {
		_spec.SetField(flight.FieldDepartureDate, field.TypeTime, value)
		_node.DepartureDate = value
	}
	if value, ok := fc.mutation.DepartureTime(); ok {
		_spec.SetField(flight.FieldDepartureTime, field.TypeTime, value)
		_node.DepartureTime = value
	}
	if value, ok := fc.mutation.Duration(); ok {
		_spec.SetField(flight.FieldDuration, field.TypeInt, value)
		_node.Duration = value
	}
	if value, ok := fc.mutation.Capacity(); ok {
		_spec.SetField(flight.FieldCapacity, field.TypeInt, value)
		_node.Capacity = value
	}
	if value, ok := fc.mutation.EconomyAvailableSeat(); ok {
		_spec.SetField(flight.FieldEconomyAvailableSeat, field.TypeInt, value)
		_node.EconomyAvailableSeat = value
	}
	if value, ok := fc.mutation.BusinessAvailableSeat(); ok {
		_spec.SetField(flight.FieldBusinessAvailableSeat, field.TypeInt, value)
		_node.BusinessAvailableSeat = value
	}
	if value, ok := fc.mutation.Status(); ok {
		_spec.SetField(flight.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(flight.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(flight.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := fc.mutation.BelongsToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flight.BelongsToTable,
			Columns: []string{flight.BelongsToColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FlightTicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flight.FlightTicketsTable,
			Columns: []string{flight.FlightTicketsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FlightCreateBulk is the builder for creating many Flight entities in bulk.
type FlightCreateBulk struct {
	config
	builders []*FlightCreate
}

// Save creates the Flight entities in the database.
func (fcb *FlightCreateBulk) Save(ctx context.Context) ([]*Flight, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Flight, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FlightMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FlightCreateBulk) SaveX(ctx context.Context) []*Flight {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FlightCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FlightCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
