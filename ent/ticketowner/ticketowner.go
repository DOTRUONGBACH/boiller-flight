// Code generated by ent, DO NOT EDIT.

package ticketowner

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ticketowner type in the database.
	Label = "ticket_owner"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTicket holds the string denoting the ticket edge name in mutations.
	EdgeTicket = "ticket"
	// EdgeCustomerOwner holds the string denoting the customer_owner edge name in mutations.
	EdgeCustomerOwner = "customer_owner"
	// EdgePassengerOwner holds the string denoting the passenger_owner edge name in mutations.
	EdgePassengerOwner = "passenger_owner"
	// Table holds the table name of the ticketowner in the database.
	Table = "ticket_owners"
	// TicketTable is the table that holds the ticket relation/edge.
	TicketTable = "tickets"
	// TicketInverseTable is the table name for the Ticket entity.
	// It exists in this package in order to avoid circular dependency with the "ticket" package.
	TicketInverseTable = "tickets"
	// TicketColumn is the table column denoting the ticket relation/edge.
	TicketColumn = "ticket_owner_ticket"
	// CustomerOwnerTable is the table that holds the customer_owner relation/edge.
	CustomerOwnerTable = "ticket_owners"
	// CustomerOwnerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerOwnerInverseTable = "customers"
	// CustomerOwnerColumn is the table column denoting the customer_owner relation/edge.
	CustomerOwnerColumn = "customer_ticket"
	// PassengerOwnerTable is the table that holds the passenger_owner relation/edge.
	PassengerOwnerTable = "ticket_owners"
	// PassengerOwnerInverseTable is the table name for the Passenger entity.
	// It exists in this package in order to avoid circular dependency with the "passenger" package.
	PassengerOwnerInverseTable = "passengers"
	// PassengerOwnerColumn is the table column denoting the passenger_owner relation/edge.
	PassengerOwnerColumn = "passenger_ticket"
)

// Columns holds all SQL columns for ticketowner fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ticket_owners"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"customer_ticket",
	"passenger_ticket",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
