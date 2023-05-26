// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mock_project/ent/customer"
	"mock_project/ent/ticketowner"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CitizenID holds the value of the "citizen_id" field.
	CitizenID string `json:"citizen_id,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender customer.Gender `json:"gender,omitempty"`
	// DateOfBirth holds the value of the "date_of_birth" field.
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges CustomerEdges `json:"edges"`
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Accounts holds the value of the accounts edge.
	Accounts []*Account `json:"accounts,omitempty"`
	// Bookings holds the value of the bookings edge.
	Bookings []*Booking `json:"bookings,omitempty"`
	// Ticket holds the value of the ticket edge.
	Ticket *TicketOwner `json:"ticket,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedAccounts map[string][]*Account
	namedBookings map[string][]*Booking
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) AccountsOrErr() ([]*Account, error) {
	if e.loadedTypes[0] {
		return e.Accounts, nil
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// BookingsOrErr returns the Bookings value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) BookingsOrErr() ([]*Booking, error) {
	if e.loadedTypes[1] {
		return e.Bookings, nil
	}
	return nil, &NotLoadedError{edge: "bookings"}
}

// TicketOrErr returns the Ticket value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) TicketOrErr() (*TicketOwner, error) {
	if e.loadedTypes[2] {
		if e.Ticket == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: ticketowner.Label}
		}
		return e.Ticket, nil
	}
	return nil, &NotLoadedError{edge: "ticket"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldName, customer.FieldCitizenID, customer.FieldPhone, customer.FieldAddress, customer.FieldGender:
			values[i] = new(sql.NullString)
		case customer.FieldDateOfBirth, customer.FieldCreatedAt, customer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case customer.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Customer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case customer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case customer.FieldCitizenID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field citizen_id", values[i])
			} else if value.Valid {
				c.CitizenID = value.String
			}
		case customer.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case customer.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = value.String
			}
		case customer.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				c.Gender = customer.Gender(value.String)
			}
		case customer.FieldDateOfBirth:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_of_birth", values[i])
			} else if value.Valid {
				c.DateOfBirth = value.Time
			}
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryAccounts queries the "accounts" edge of the Customer entity.
func (c *Customer) QueryAccounts() *AccountQuery {
	return (&CustomerClient{config: c.config}).QueryAccounts(c)
}

// QueryBookings queries the "bookings" edge of the Customer entity.
func (c *Customer) QueryBookings() *BookingQuery {
	return (&CustomerClient{config: c.config}).QueryBookings(c)
}

// QueryTicket queries the "ticket" edge of the Customer entity.
func (c *Customer) QueryTicket() *TicketOwnerQuery {
	return (&CustomerClient{config: c.config}).QueryTicket(c)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("citizen_id=")
	builder.WriteString(c.CitizenID)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(c.Address)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", c.Gender))
	builder.WriteString(", ")
	builder.WriteString("date_of_birth=")
	builder.WriteString(c.DateOfBirth.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedAccounts returns the Accounts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Customer) NamedAccounts(name string) ([]*Account, error) {
	if c.Edges.namedAccounts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedAccounts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Customer) appendNamedAccounts(name string, edges ...*Account) {
	if c.Edges.namedAccounts == nil {
		c.Edges.namedAccounts = make(map[string][]*Account)
	}
	if len(edges) == 0 {
		c.Edges.namedAccounts[name] = []*Account{}
	} else {
		c.Edges.namedAccounts[name] = append(c.Edges.namedAccounts[name], edges...)
	}
}

// NamedBookings returns the Bookings named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Customer) NamedBookings(name string) ([]*Booking, error) {
	if c.Edges.namedBookings == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedBookings[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Customer) appendNamedBookings(name string, edges ...*Booking) {
	if c.Edges.namedBookings == nil {
		c.Edges.namedBookings = make(map[string][]*Booking)
	}
	if len(edges) == 0 {
		c.Edges.namedBookings[name] = []*Booking{}
	} else {
		c.Edges.namedBookings[name] = append(c.Edges.namedBookings[name], edges...)
	}
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
