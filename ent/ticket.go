// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mock_project/ent/booking"
	"mock_project/ent/flight"
	"mock_project/ent/ticket"
	"mock_project/ent/ticketowner"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Ticket is the model entity for the Ticket schema.
type Ticket struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Class holds the value of the "class" field.
	Class ticket.Class `json:"class,omitempty"`
	// Status holds the value of the "status" field.
	Status ticket.Status `json:"status,omitempty"`
	// FlightID holds the value of the "flight_id" field.
	FlightID uuid.UUID `json:"flight_id,omitempty"`
	// BookingID holds the value of the "booking_id" field.
	BookingID uuid.UUID `json:"booking_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TicketQuery when eager-loading is set.
	Edges               TicketEdges `json:"edges"`
	ticket_owner_ticket *uuid.UUID
}

// TicketEdges holds the relations/edges for other nodes in the graph.
type TicketEdges struct {
	// FlightOwner holds the value of the flight_owner edge.
	FlightOwner *Flight `json:"flight_owner,omitempty"`
	// BookingOwner holds the value of the booking_owner edge.
	BookingOwner *Booking `json:"booking_owner,omitempty"`
	// TicketOwner holds the value of the ticket_owner edge.
	TicketOwner *TicketOwner `json:"ticket_owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// FlightOwnerOrErr returns the FlightOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEdges) FlightOwnerOrErr() (*Flight, error) {
	if e.loadedTypes[0] {
		if e.FlightOwner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: flight.Label}
		}
		return e.FlightOwner, nil
	}
	return nil, &NotLoadedError{edge: "flight_owner"}
}

// BookingOwnerOrErr returns the BookingOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEdges) BookingOwnerOrErr() (*Booking, error) {
	if e.loadedTypes[1] {
		if e.BookingOwner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: booking.Label}
		}
		return e.BookingOwner, nil
	}
	return nil, &NotLoadedError{edge: "booking_owner"}
}

// TicketOwnerOrErr returns the TicketOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEdges) TicketOwnerOrErr() (*TicketOwner, error) {
	if e.loadedTypes[2] {
		if e.TicketOwner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: ticketowner.Label}
		}
		return e.TicketOwner, nil
	}
	return nil, &NotLoadedError{edge: "ticket_owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ticket) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ticket.FieldClass, ticket.FieldStatus:
			values[i] = new(sql.NullString)
		case ticket.FieldCreatedAt, ticket.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case ticket.FieldID, ticket.FieldFlightID, ticket.FieldBookingID:
			values[i] = new(uuid.UUID)
		case ticket.ForeignKeys[0]: // ticket_owner_ticket
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ticket", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ticket fields.
func (t *Ticket) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ticket.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case ticket.FieldClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class", values[i])
			} else if value.Valid {
				t.Class = ticket.Class(value.String)
			}
		case ticket.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = ticket.Status(value.String)
			}
		case ticket.FieldFlightID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field flight_id", values[i])
			} else if value != nil {
				t.FlightID = *value
			}
		case ticket.FieldBookingID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field booking_id", values[i])
			} else if value != nil {
				t.BookingID = *value
			}
		case ticket.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case ticket.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case ticket.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field ticket_owner_ticket", values[i])
			} else if value.Valid {
				t.ticket_owner_ticket = new(uuid.UUID)
				*t.ticket_owner_ticket = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryFlightOwner queries the "flight_owner" edge of the Ticket entity.
func (t *Ticket) QueryFlightOwner() *FlightQuery {
	return (&TicketClient{config: t.config}).QueryFlightOwner(t)
}

// QueryBookingOwner queries the "booking_owner" edge of the Ticket entity.
func (t *Ticket) QueryBookingOwner() *BookingQuery {
	return (&TicketClient{config: t.config}).QueryBookingOwner(t)
}

// QueryTicketOwner queries the "ticket_owner" edge of the Ticket entity.
func (t *Ticket) QueryTicketOwner() *TicketOwnerQuery {
	return (&TicketClient{config: t.config}).QueryTicketOwner(t)
}

// Update returns a builder for updating this Ticket.
// Note that you need to call Ticket.Unwrap() before calling this method if this Ticket
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Ticket) Update() *TicketUpdateOne {
	return (&TicketClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Ticket entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Ticket) Unwrap() *Ticket {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ticket is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Ticket) String() string {
	var builder strings.Builder
	builder.WriteString("Ticket(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("class=")
	builder.WriteString(fmt.Sprintf("%v", t.Class))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("flight_id=")
	builder.WriteString(fmt.Sprintf("%v", t.FlightID))
	builder.WriteString(", ")
	builder.WriteString("booking_id=")
	builder.WriteString(fmt.Sprintf("%v", t.BookingID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Tickets is a parsable slice of Ticket.
type Tickets []*Ticket

func (t Tickets) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
