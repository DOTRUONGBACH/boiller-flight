// Code generated by ent, DO NOT EDIT.

package ent

import (
	"mock_project/ent/account"
	"mock_project/ent/booking"
	"mock_project/ent/customer"
	"mock_project/ent/flight"
	"mock_project/ent/passenger"
	"mock_project/ent/ticket"
	"time"

	"github.com/google/uuid"
)

// CreateAccountInput represents a mutation input for creating accounts.
type CreateAccountInput struct {
	Email      string
	Password   string
	Role       account.Role
	Status     account.Status
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	AccOwnerID *uuid.UUID
}

// Mutate applies the CreateAccountInput on the AccountMutation builder.
func (i *CreateAccountInput) Mutate(m *AccountMutation) {
	m.SetEmail(i.Email)
	m.SetPassword(i.Password)
	m.SetRole(i.Role)
	m.SetStatus(i.Status)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.AccOwnerID; v != nil {
		m.SetAccOwnerID(*v)
	}
}

// SetInput applies the change-set in the CreateAccountInput on the AccountCreate builder.
func (c *AccountCreate) SetInput(i CreateAccountInput) *AccountCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAccountInput represents a mutation input for updating accounts.
type UpdateAccountInput struct {
	Email         *string
	Password      *string
	Role          *account.Role
	Status        *account.Status
	UpdatedAt     *time.Time
	ClearAccOwner bool
	AccOwnerID    *uuid.UUID
}

// Mutate applies the UpdateAccountInput on the AccountMutation builder.
func (i *UpdateAccountInput) Mutate(m *AccountMutation) {
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.Role; v != nil {
		m.SetRole(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearAccOwner {
		m.ClearAccOwner()
	}
	if v := i.AccOwnerID; v != nil {
		m.SetAccOwnerID(*v)
	}
}

// SetInput applies the change-set in the UpdateAccountInput on the AccountUpdate builder.
func (c *AccountUpdate) SetInput(i UpdateAccountInput) *AccountUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAccountInput on the AccountUpdateOne builder.
func (c *AccountUpdateOne) SetInput(i UpdateAccountInput) *AccountUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateBookingInput represents a mutation input for creating bookings.
type CreateBookingInput struct {
	EconomyTickets            *int
	BusinessTickets           *int
	Status                    booking.Status
	Type                      booking.Type
	CreatedAt                 *time.Time
	UpdatedAt                 *time.Time
	BookingFlightID           uuid.UUID
	CustomerBookingTicketsID  *uuid.UUID
	PassengerBookingTicketsID *uuid.UUID
	BookingTicketIDs          []uuid.UUID
}

// Mutate applies the CreateBookingInput on the BookingMutation builder.
func (i *CreateBookingInput) Mutate(m *BookingMutation) {
	if v := i.EconomyTickets; v != nil {
		m.SetEconomyTickets(*v)
	}
	if v := i.BusinessTickets; v != nil {
		m.SetBusinessTickets(*v)
	}
	m.SetStatus(i.Status)
	m.SetType(i.Type)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetBookingFlightID(i.BookingFlightID)
	if v := i.CustomerBookingTicketsID; v != nil {
		m.SetCustomerBookingTicketsID(*v)
	}
	if v := i.PassengerBookingTicketsID; v != nil {
		m.SetPassengerBookingTicketsID(*v)
	}
	if v := i.BookingTicketIDs; len(v) > 0 {
		m.AddBookingTicketIDs(v...)
	}
}

// SetInput applies the change-set in the CreateBookingInput on the BookingCreate builder.
func (c *BookingCreate) SetInput(i CreateBookingInput) *BookingCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateBookingInput represents a mutation input for updating bookings.
type UpdateBookingInput struct {
	EconomyTickets               *int
	BusinessTickets              *int
	Status                       *booking.Status
	Type                         *booking.Type
	UpdatedAt                    *time.Time
	ClearBookingFlight           bool
	BookingFlightID              *uuid.UUID
	ClearCustomerBookingTickets  bool
	CustomerBookingTicketsID     *uuid.UUID
	ClearPassengerBookingTickets bool
	PassengerBookingTicketsID    *uuid.UUID
	AddBookingTicketIDs          []uuid.UUID
	RemoveBookingTicketIDs       []uuid.UUID
}

// Mutate applies the UpdateBookingInput on the BookingMutation builder.
func (i *UpdateBookingInput) Mutate(m *BookingMutation) {
	if v := i.EconomyTickets; v != nil {
		m.SetEconomyTickets(*v)
	}
	if v := i.BusinessTickets; v != nil {
		m.SetBusinessTickets(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearBookingFlight {
		m.ClearBookingFlight()
	}
	if v := i.BookingFlightID; v != nil {
		m.SetBookingFlightID(*v)
	}
	if i.ClearCustomerBookingTickets {
		m.ClearCustomerBookingTickets()
	}
	if v := i.CustomerBookingTicketsID; v != nil {
		m.SetCustomerBookingTicketsID(*v)
	}
	if i.ClearPassengerBookingTickets {
		m.ClearPassengerBookingTickets()
	}
	if v := i.PassengerBookingTicketsID; v != nil {
		m.SetPassengerBookingTicketsID(*v)
	}
	if v := i.AddBookingTicketIDs; len(v) > 0 {
		m.AddBookingTicketIDs(v...)
	}
	if v := i.RemoveBookingTicketIDs; len(v) > 0 {
		m.RemoveBookingTicketIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateBookingInput on the BookingUpdate builder.
func (c *BookingUpdate) SetInput(i UpdateBookingInput) *BookingUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateBookingInput on the BookingUpdateOne builder.
func (c *BookingUpdateOne) SetInput(i UpdateBookingInput) *BookingUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateCustomerInput represents a mutation input for creating customers.
type CreateCustomerInput struct {
	Name        string
	CitizenID   string
	Phone       string
	Address     string
	Gender      customer.Gender
	DateOfBirth time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	AccountIDs  []uuid.UUID
	BookingIDs  []uuid.UUID
	TicketID    *uuid.UUID
}

// Mutate applies the CreateCustomerInput on the CustomerMutation builder.
func (i *CreateCustomerInput) Mutate(m *CustomerMutation) {
	m.SetName(i.Name)
	m.SetCitizenID(i.CitizenID)
	m.SetPhone(i.Phone)
	m.SetAddress(i.Address)
	m.SetGender(i.Gender)
	m.SetDateOfBirth(i.DateOfBirth)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.AccountIDs; len(v) > 0 {
		m.AddAccountIDs(v...)
	}
	if v := i.BookingIDs; len(v) > 0 {
		m.AddBookingIDs(v...)
	}
	if v := i.TicketID; v != nil {
		m.SetTicketID(*v)
	}
}

// SetInput applies the change-set in the CreateCustomerInput on the CustomerCreate builder.
func (c *CustomerCreate) SetInput(i CreateCustomerInput) *CustomerCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCustomerInput represents a mutation input for updating customers.
type UpdateCustomerInput struct {
	Name             *string
	CitizenID        *string
	Phone            *string
	Address          *string
	Gender           *customer.Gender
	DateOfBirth      *time.Time
	UpdatedAt        *time.Time
	AddAccountIDs    []uuid.UUID
	RemoveAccountIDs []uuid.UUID
	AddBookingIDs    []uuid.UUID
	RemoveBookingIDs []uuid.UUID
	ClearTicket      bool
	TicketID         *uuid.UUID
}

// Mutate applies the UpdateCustomerInput on the CustomerMutation builder.
func (i *UpdateCustomerInput) Mutate(m *CustomerMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.CitizenID; v != nil {
		m.SetCitizenID(*v)
	}
	if v := i.Phone; v != nil {
		m.SetPhone(*v)
	}
	if v := i.Address; v != nil {
		m.SetAddress(*v)
	}
	if v := i.Gender; v != nil {
		m.SetGender(*v)
	}
	if v := i.DateOfBirth; v != nil {
		m.SetDateOfBirth(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.AddAccountIDs; len(v) > 0 {
		m.AddAccountIDs(v...)
	}
	if v := i.RemoveAccountIDs; len(v) > 0 {
		m.RemoveAccountIDs(v...)
	}
	if v := i.AddBookingIDs; len(v) > 0 {
		m.AddBookingIDs(v...)
	}
	if v := i.RemoveBookingIDs; len(v) > 0 {
		m.RemoveBookingIDs(v...)
	}
	if i.ClearTicket {
		m.ClearTicket()
	}
	if v := i.TicketID; v != nil {
		m.SetTicketID(*v)
	}
}

// SetInput applies the change-set in the UpdateCustomerInput on the CustomerUpdate builder.
func (c *CustomerUpdate) SetInput(i UpdateCustomerInput) *CustomerUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCustomerInput on the CustomerUpdateOne builder.
func (c *CustomerUpdateOne) SetInput(i UpdateCustomerInput) *CustomerUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateFlightInput represents a mutation input for creating flights.
type CreateFlightInput struct {
	FlightCode            string
	From                  string
	To                    string
	DepartureDate         time.Time
	DepartureTime         time.Time
	Duration              int
	Capacity              int
	EconomyAvailableSeat  int
	BusinessAvailableSeat int
	Status                flight.Status
	CreatedAt             *time.Time
	UpdatedAt             *time.Time
	BelongsToIDs          []uuid.UUID
	FlightTicketIDs       []uuid.UUID
}

// Mutate applies the CreateFlightInput on the FlightMutation builder.
func (i *CreateFlightInput) Mutate(m *FlightMutation) {
	m.SetFlightCode(i.FlightCode)
	m.SetFrom(i.From)
	m.SetTo(i.To)
	m.SetDepartureDate(i.DepartureDate)
	m.SetDepartureTime(i.DepartureTime)
	m.SetDuration(i.Duration)
	m.SetCapacity(i.Capacity)
	m.SetEconomyAvailableSeat(i.EconomyAvailableSeat)
	m.SetBusinessAvailableSeat(i.BusinessAvailableSeat)
	m.SetStatus(i.Status)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.BelongsToIDs; len(v) > 0 {
		m.AddBelongsToIDs(v...)
	}
	if v := i.FlightTicketIDs; len(v) > 0 {
		m.AddFlightTicketIDs(v...)
	}
}

// SetInput applies the change-set in the CreateFlightInput on the FlightCreate builder.
func (c *FlightCreate) SetInput(i CreateFlightInput) *FlightCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateFlightInput represents a mutation input for updating flights.
type UpdateFlightInput struct {
	FlightCode            *string
	From                  *string
	To                    *string
	DepartureDate         *time.Time
	DepartureTime         *time.Time
	Duration              *int
	Capacity              *int
	EconomyAvailableSeat  *int
	BusinessAvailableSeat *int
	Status                *flight.Status
	UpdatedAt             *time.Time
	AddBelongsToIDs       []uuid.UUID
	RemoveBelongsToIDs    []uuid.UUID
	AddFlightTicketIDs    []uuid.UUID
	RemoveFlightTicketIDs []uuid.UUID
}

// Mutate applies the UpdateFlightInput on the FlightMutation builder.
func (i *UpdateFlightInput) Mutate(m *FlightMutation) {
	if v := i.FlightCode; v != nil {
		m.SetFlightCode(*v)
	}
	if v := i.From; v != nil {
		m.SetFrom(*v)
	}
	if v := i.To; v != nil {
		m.SetTo(*v)
	}
	if v := i.DepartureDate; v != nil {
		m.SetDepartureDate(*v)
	}
	if v := i.DepartureTime; v != nil {
		m.SetDepartureTime(*v)
	}
	if v := i.Duration; v != nil {
		m.SetDuration(*v)
	}
	if v := i.Capacity; v != nil {
		m.SetCapacity(*v)
	}
	if v := i.EconomyAvailableSeat; v != nil {
		m.SetEconomyAvailableSeat(*v)
	}
	if v := i.BusinessAvailableSeat; v != nil {
		m.SetBusinessAvailableSeat(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.AddBelongsToIDs; len(v) > 0 {
		m.AddBelongsToIDs(v...)
	}
	if v := i.RemoveBelongsToIDs; len(v) > 0 {
		m.RemoveBelongsToIDs(v...)
	}
	if v := i.AddFlightTicketIDs; len(v) > 0 {
		m.AddFlightTicketIDs(v...)
	}
	if v := i.RemoveFlightTicketIDs; len(v) > 0 {
		m.RemoveFlightTicketIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateFlightInput on the FlightUpdate builder.
func (c *FlightUpdate) SetInput(i UpdateFlightInput) *FlightUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateFlightInput on the FlightUpdateOne builder.
func (c *FlightUpdateOne) SetInput(i UpdateFlightInput) *FlightUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreatePassengerInput represents a mutation input for creating passengers.
type CreatePassengerInput struct {
	Name        string
	CitizenID   string
	Email       string
	Phone       string
	Address     string
	Gender      passenger.Gender
	DateOfBirth time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	TicketID    *uuid.UUID
	BookingIDs  []uuid.UUID
}

// Mutate applies the CreatePassengerInput on the PassengerMutation builder.
func (i *CreatePassengerInput) Mutate(m *PassengerMutation) {
	m.SetName(i.Name)
	m.SetCitizenID(i.CitizenID)
	m.SetEmail(i.Email)
	m.SetPhone(i.Phone)
	m.SetAddress(i.Address)
	m.SetGender(i.Gender)
	m.SetDateOfBirth(i.DateOfBirth)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.TicketID; v != nil {
		m.SetTicketID(*v)
	}
	if v := i.BookingIDs; len(v) > 0 {
		m.AddBookingIDs(v...)
	}
}

// SetInput applies the change-set in the CreatePassengerInput on the PassengerCreate builder.
func (c *PassengerCreate) SetInput(i CreatePassengerInput) *PassengerCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePassengerInput represents a mutation input for updating passengers.
type UpdatePassengerInput struct {
	Name             *string
	CitizenID        *string
	Email            *string
	Phone            *string
	Address          *string
	Gender           *passenger.Gender
	DateOfBirth      *time.Time
	UpdatedAt        *time.Time
	ClearTicket      bool
	TicketID         *uuid.UUID
	AddBookingIDs    []uuid.UUID
	RemoveBookingIDs []uuid.UUID
}

// Mutate applies the UpdatePassengerInput on the PassengerMutation builder.
func (i *UpdatePassengerInput) Mutate(m *PassengerMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.CitizenID; v != nil {
		m.SetCitizenID(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Phone; v != nil {
		m.SetPhone(*v)
	}
	if v := i.Address; v != nil {
		m.SetAddress(*v)
	}
	if v := i.Gender; v != nil {
		m.SetGender(*v)
	}
	if v := i.DateOfBirth; v != nil {
		m.SetDateOfBirth(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearTicket {
		m.ClearTicket()
	}
	if v := i.TicketID; v != nil {
		m.SetTicketID(*v)
	}
	if v := i.AddBookingIDs; len(v) > 0 {
		m.AddBookingIDs(v...)
	}
	if v := i.RemoveBookingIDs; len(v) > 0 {
		m.RemoveBookingIDs(v...)
	}
}

// SetInput applies the change-set in the UpdatePassengerInput on the PassengerUpdate builder.
func (c *PassengerUpdate) SetInput(i UpdatePassengerInput) *PassengerUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePassengerInput on the PassengerUpdateOne builder.
func (c *PassengerUpdateOne) SetInput(i UpdatePassengerInput) *PassengerUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTicketInput represents a mutation input for creating tickets.
type CreateTicketInput struct {
	Class          ticket.Class
	Status         ticket.Status
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	FlightOwnerID  uuid.UUID
	BookingOwnerID uuid.UUID
	TicketOwnerID  *uuid.UUID
}

// Mutate applies the CreateTicketInput on the TicketMutation builder.
func (i *CreateTicketInput) Mutate(m *TicketMutation) {
	m.SetClass(i.Class)
	m.SetStatus(i.Status)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetFlightOwnerID(i.FlightOwnerID)
	m.SetBookingOwnerID(i.BookingOwnerID)
	if v := i.TicketOwnerID; v != nil {
		m.SetTicketOwnerID(*v)
	}
}

// SetInput applies the change-set in the CreateTicketInput on the TicketCreate builder.
func (c *TicketCreate) SetInput(i CreateTicketInput) *TicketCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTicketInput represents a mutation input for updating tickets.
type UpdateTicketInput struct {
	Class             *ticket.Class
	Status            *ticket.Status
	UpdatedAt         *time.Time
	ClearFlightOwner  bool
	FlightOwnerID     *uuid.UUID
	ClearBookingOwner bool
	BookingOwnerID    *uuid.UUID
	ClearTicketOwner  bool
	TicketOwnerID     *uuid.UUID
}

// Mutate applies the UpdateTicketInput on the TicketMutation builder.
func (i *UpdateTicketInput) Mutate(m *TicketMutation) {
	if v := i.Class; v != nil {
		m.SetClass(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearFlightOwner {
		m.ClearFlightOwner()
	}
	if v := i.FlightOwnerID; v != nil {
		m.SetFlightOwnerID(*v)
	}
	if i.ClearBookingOwner {
		m.ClearBookingOwner()
	}
	if v := i.BookingOwnerID; v != nil {
		m.SetBookingOwnerID(*v)
	}
	if i.ClearTicketOwner {
		m.ClearTicketOwner()
	}
	if v := i.TicketOwnerID; v != nil {
		m.SetTicketOwnerID(*v)
	}
}

// SetInput applies the change-set in the UpdateTicketInput on the TicketUpdate builder.
func (c *TicketUpdate) SetInput(i UpdateTicketInput) *TicketUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTicketInput on the TicketUpdateOne builder.
func (c *TicketUpdateOne) SetInput(i UpdateTicketInput) *TicketUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTicketOwnerInput represents a mutation input for creating ticketowners.
type CreateTicketOwnerInput struct {
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	TicketID         *uuid.UUID
	CustomerOwnerID  *uuid.UUID
	PassengerOwnerID *uuid.UUID
}

// Mutate applies the CreateTicketOwnerInput on the TicketOwnerMutation builder.
func (i *CreateTicketOwnerInput) Mutate(m *TicketOwnerMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.TicketID; v != nil {
		m.SetTicketID(*v)
	}
	if v := i.CustomerOwnerID; v != nil {
		m.SetCustomerOwnerID(*v)
	}
	if v := i.PassengerOwnerID; v != nil {
		m.SetPassengerOwnerID(*v)
	}
}

// SetInput applies the change-set in the CreateTicketOwnerInput on the TicketOwnerCreate builder.
func (c *TicketOwnerCreate) SetInput(i CreateTicketOwnerInput) *TicketOwnerCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTicketOwnerInput represents a mutation input for updating ticketowners.
type UpdateTicketOwnerInput struct {
	UpdatedAt           *time.Time
	ClearTicket         bool
	TicketID            *uuid.UUID
	ClearCustomerOwner  bool
	CustomerOwnerID     *uuid.UUID
	ClearPassengerOwner bool
	PassengerOwnerID    *uuid.UUID
}

// Mutate applies the UpdateTicketOwnerInput on the TicketOwnerMutation builder.
func (i *UpdateTicketOwnerInput) Mutate(m *TicketOwnerMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearTicket {
		m.ClearTicket()
	}
	if v := i.TicketID; v != nil {
		m.SetTicketID(*v)
	}
	if i.ClearCustomerOwner {
		m.ClearCustomerOwner()
	}
	if v := i.CustomerOwnerID; v != nil {
		m.SetCustomerOwnerID(*v)
	}
	if i.ClearPassengerOwner {
		m.ClearPassengerOwner()
	}
	if v := i.PassengerOwnerID; v != nil {
		m.SetPassengerOwnerID(*v)
	}
}

// SetInput applies the change-set in the UpdateTicketOwnerInput on the TicketOwnerUpdate builder.
func (c *TicketOwnerUpdate) SetInput(i UpdateTicketOwnerInput) *TicketOwnerUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTicketOwnerInput on the TicketOwnerUpdateOne builder.
func (c *TicketOwnerUpdateOne) SetInput(i UpdateTicketOwnerInput) *TicketOwnerUpdateOne {
	i.Mutate(c.Mutation())
	return c
}