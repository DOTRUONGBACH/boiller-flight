// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mock_project/ent/flight"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Flight is the model entity for the Flight schema.
type Flight struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// FlightCode holds the value of the "flight_code" field.
	FlightCode string `json:"flight_code,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// DepartureDate holds the value of the "departure_date" field.
	DepartureDate time.Time `json:"departure_date,omitempty"`
	// DepartureTime holds the value of the "departure_time" field.
	DepartureTime time.Time `json:"departure_time,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// Capacity holds the value of the "capacity" field.
	Capacity int `json:"capacity,omitempty"`
	// EconomyAvailableSeat holds the value of the "economy_available_seat" field.
	EconomyAvailableSeat int `json:"economy_available_seat,omitempty"`
	// BusinessAvailableSeat holds the value of the "business_available_seat" field.
	BusinessAvailableSeat int `json:"business_available_seat,omitempty"`
	// Status holds the value of the "status" field.
	Status flight.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FlightQuery when eager-loading is set.
	Edges FlightEdges `json:"edges"`
}

// FlightEdges holds the relations/edges for other nodes in the graph.
type FlightEdges struct {
	// BelongsTo holds the value of the belongs_to edge.
	BelongsTo []*Booking `json:"belongs_to,omitempty"`
	// FlightTickets holds the value of the flight_tickets edge.
	FlightTickets []*Ticket `json:"flight_tickets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedBelongsTo     map[string][]*Booking
	namedFlightTickets map[string][]*Ticket
}

// BelongsToOrErr returns the BelongsTo value or an error if the edge
// was not loaded in eager-loading.
func (e FlightEdges) BelongsToOrErr() ([]*Booking, error) {
	if e.loadedTypes[0] {
		return e.BelongsTo, nil
	}
	return nil, &NotLoadedError{edge: "belongs_to"}
}

// FlightTicketsOrErr returns the FlightTickets value or an error if the edge
// was not loaded in eager-loading.
func (e FlightEdges) FlightTicketsOrErr() ([]*Ticket, error) {
	if e.loadedTypes[1] {
		return e.FlightTickets, nil
	}
	return nil, &NotLoadedError{edge: "flight_tickets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Flight) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case flight.FieldDuration, flight.FieldCapacity, flight.FieldEconomyAvailableSeat, flight.FieldBusinessAvailableSeat:
			values[i] = new(sql.NullInt64)
		case flight.FieldFlightCode, flight.FieldFrom, flight.FieldTo, flight.FieldStatus:
			values[i] = new(sql.NullString)
		case flight.FieldDepartureDate, flight.FieldDepartureTime, flight.FieldCreatedAt, flight.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case flight.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Flight", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Flight fields.
func (f *Flight) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case flight.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case flight.FieldFlightCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field flight_code", values[i])
			} else if value.Valid {
				f.FlightCode = value.String
			}
		case flight.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				f.From = value.String
			}
		case flight.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				f.To = value.String
			}
		case flight.FieldDepartureDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field departure_date", values[i])
			} else if value.Valid {
				f.DepartureDate = value.Time
			}
		case flight.FieldDepartureTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field departure_time", values[i])
			} else if value.Valid {
				f.DepartureTime = value.Time
			}
		case flight.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				f.Duration = int(value.Int64)
			}
		case flight.FieldCapacity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field capacity", values[i])
			} else if value.Valid {
				f.Capacity = int(value.Int64)
			}
		case flight.FieldEconomyAvailableSeat:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field economy_available_seat", values[i])
			} else if value.Valid {
				f.EconomyAvailableSeat = int(value.Int64)
			}
		case flight.FieldBusinessAvailableSeat:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field business_available_seat", values[i])
			} else if value.Valid {
				f.BusinessAvailableSeat = int(value.Int64)
			}
		case flight.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = flight.Status(value.String)
			}
		case flight.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case flight.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryBelongsTo queries the "belongs_to" edge of the Flight entity.
func (f *Flight) QueryBelongsTo() *BookingQuery {
	return (&FlightClient{config: f.config}).QueryBelongsTo(f)
}

// QueryFlightTickets queries the "flight_tickets" edge of the Flight entity.
func (f *Flight) QueryFlightTickets() *TicketQuery {
	return (&FlightClient{config: f.config}).QueryFlightTickets(f)
}

// Update returns a builder for updating this Flight.
// Note that you need to call Flight.Unwrap() before calling this method if this Flight
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Flight) Update() *FlightUpdateOne {
	return (&FlightClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Flight entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Flight) Unwrap() *Flight {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Flight is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Flight) String() string {
	var builder strings.Builder
	builder.WriteString("Flight(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("flight_code=")
	builder.WriteString(f.FlightCode)
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(f.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(f.To)
	builder.WriteString(", ")
	builder.WriteString("departure_date=")
	builder.WriteString(f.DepartureDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("departure_time=")
	builder.WriteString(f.DepartureTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", f.Duration))
	builder.WriteString(", ")
	builder.WriteString("capacity=")
	builder.WriteString(fmt.Sprintf("%v", f.Capacity))
	builder.WriteString(", ")
	builder.WriteString("economy_available_seat=")
	builder.WriteString(fmt.Sprintf("%v", f.EconomyAvailableSeat))
	builder.WriteString(", ")
	builder.WriteString("business_available_seat=")
	builder.WriteString(fmt.Sprintf("%v", f.BusinessAvailableSeat))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", f.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedBelongsTo returns the BelongsTo named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *Flight) NamedBelongsTo(name string) ([]*Booking, error) {
	if f.Edges.namedBelongsTo == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedBelongsTo[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *Flight) appendNamedBelongsTo(name string, edges ...*Booking) {
	if f.Edges.namedBelongsTo == nil {
		f.Edges.namedBelongsTo = make(map[string][]*Booking)
	}
	if len(edges) == 0 {
		f.Edges.namedBelongsTo[name] = []*Booking{}
	} else {
		f.Edges.namedBelongsTo[name] = append(f.Edges.namedBelongsTo[name], edges...)
	}
}

// NamedFlightTickets returns the FlightTickets named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *Flight) NamedFlightTickets(name string) ([]*Ticket, error) {
	if f.Edges.namedFlightTickets == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedFlightTickets[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *Flight) appendNamedFlightTickets(name string, edges ...*Ticket) {
	if f.Edges.namedFlightTickets == nil {
		f.Edges.namedFlightTickets = make(map[string][]*Ticket)
	}
	if len(edges) == 0 {
		f.Edges.namedFlightTickets[name] = []*Ticket{}
	} else {
		f.Edges.namedFlightTickets[name] = append(f.Edges.namedFlightTickets[name], edges...)
	}
}

// Flights is a parsable slice of Flight.
type Flights []*Flight

func (f Flights) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}