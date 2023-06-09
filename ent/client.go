// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"mock_project/ent/migrate"

	"mock_project/ent/account"
	"mock_project/ent/booking"
	"mock_project/ent/customer"
	"mock_project/ent/flight"
	"mock_project/ent/passenger"
	"mock_project/ent/ticket"
	"mock_project/ent/ticketowner"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Booking is the client for interacting with the Booking builders.
	Booking *BookingClient
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Flight is the client for interacting with the Flight builders.
	Flight *FlightClient
	// Passenger is the client for interacting with the Passenger builders.
	Passenger *PassengerClient
	// Ticket is the client for interacting with the Ticket builders.
	Ticket *TicketClient
	// TicketOwner is the client for interacting with the TicketOwner builders.
	TicketOwner *TicketOwnerClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Booking = NewBookingClient(c.config)
	c.Customer = NewCustomerClient(c.config)
	c.Flight = NewFlightClient(c.config)
	c.Passenger = NewPassengerClient(c.config)
	c.Ticket = NewTicketClient(c.config)
	c.TicketOwner = NewTicketOwnerClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Account:     NewAccountClient(cfg),
		Booking:     NewBookingClient(cfg),
		Customer:    NewCustomerClient(cfg),
		Flight:      NewFlightClient(cfg),
		Passenger:   NewPassengerClient(cfg),
		Ticket:      NewTicketClient(cfg),
		TicketOwner: NewTicketOwnerClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Account:     NewAccountClient(cfg),
		Booking:     NewBookingClient(cfg),
		Customer:    NewCustomerClient(cfg),
		Flight:      NewFlightClient(cfg),
		Passenger:   NewPassengerClient(cfg),
		Ticket:      NewTicketClient(cfg),
		TicketOwner: NewTicketOwnerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Booking.Use(hooks...)
	c.Customer.Use(hooks...)
	c.Flight.Use(hooks...)
	c.Passenger.Use(hooks...)
	c.Ticket.Use(hooks...)
	c.TicketOwner.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uuid.UUID) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id uuid.UUID) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uuid.UUID) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uuid.UUID) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccOwner queries the acc_owner edge of a Account.
func (c *AccountClient) QueryAccOwner(a *Account) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, account.AccOwnerTable, account.AccOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// BookingClient is a client for the Booking schema.
type BookingClient struct {
	config
}

// NewBookingClient returns a client for the Booking from the given config.
func NewBookingClient(c config) *BookingClient {
	return &BookingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `booking.Hooks(f(g(h())))`.
func (c *BookingClient) Use(hooks ...Hook) {
	c.hooks.Booking = append(c.hooks.Booking, hooks...)
}

// Create returns a builder for creating a Booking entity.
func (c *BookingClient) Create() *BookingCreate {
	mutation := newBookingMutation(c.config, OpCreate)
	return &BookingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Booking entities.
func (c *BookingClient) CreateBulk(builders ...*BookingCreate) *BookingCreateBulk {
	return &BookingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Booking.
func (c *BookingClient) Update() *BookingUpdate {
	mutation := newBookingMutation(c.config, OpUpdate)
	return &BookingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookingClient) UpdateOne(b *Booking) *BookingUpdateOne {
	mutation := newBookingMutation(c.config, OpUpdateOne, withBooking(b))
	return &BookingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookingClient) UpdateOneID(id uuid.UUID) *BookingUpdateOne {
	mutation := newBookingMutation(c.config, OpUpdateOne, withBookingID(id))
	return &BookingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Booking.
func (c *BookingClient) Delete() *BookingDelete {
	mutation := newBookingMutation(c.config, OpDelete)
	return &BookingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BookingClient) DeleteOne(b *Booking) *BookingDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BookingClient) DeleteOneID(id uuid.UUID) *BookingDeleteOne {
	builder := c.Delete().Where(booking.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookingDeleteOne{builder}
}

// Query returns a query builder for Booking.
func (c *BookingClient) Query() *BookingQuery {
	return &BookingQuery{
		config: c.config,
	}
}

// Get returns a Booking entity by its id.
func (c *BookingClient) Get(ctx context.Context, id uuid.UUID) (*Booking, error) {
	return c.Query().Where(booking.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookingClient) GetX(ctx context.Context, id uuid.UUID) *Booking {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBookingFlight queries the booking_flight edge of a Booking.
func (c *BookingClient) QueryBookingFlight(b *Booking) *FlightQuery {
	query := &FlightQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.FieldID, id),
			sqlgraph.To(flight.Table, flight.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, booking.BookingFlightTable, booking.BookingFlightColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCustomerBookingTickets queries the customer_booking_tickets edge of a Booking.
func (c *BookingClient) QueryCustomerBookingTickets(b *Booking) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, booking.CustomerBookingTicketsTable, booking.CustomerBookingTicketsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPassengerBookingTickets queries the passenger_booking_tickets edge of a Booking.
func (c *BookingClient) QueryPassengerBookingTickets(b *Booking) *PassengerQuery {
	query := &PassengerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.FieldID, id),
			sqlgraph.To(passenger.Table, passenger.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, booking.PassengerBookingTicketsTable, booking.PassengerBookingTicketsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBookingTickets queries the booking_tickets edge of a Booking.
func (c *BookingClient) QueryBookingTickets(b *Booking) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, booking.BookingTicketsTable, booking.BookingTicketsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookingClient) Hooks() []Hook {
	return c.hooks.Booking
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Create returns a builder for creating a Customer entity.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Customer entities.
func (c *CustomerClient) CreateBulk(builders ...*CustomerCreate) *CustomerCreateBulk {
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id uuid.UUID) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerClient) DeleteOneID(id uuid.UUID) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Query returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{
		config: c.config,
	}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id uuid.UUID) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id uuid.UUID) *Customer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccounts queries the accounts edge of a Customer.
func (c *CustomerClient) QueryAccounts(cu *Customer) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.AccountsTable, customer.AccountsColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBookings queries the bookings edge of a Customer.
func (c *CustomerClient) QueryBookings(cu *Customer) *BookingQuery {
	query := &BookingQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.BookingsTable, customer.BookingsColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTicket queries the ticket edge of a Customer.
func (c *CustomerClient) QueryTicket(cu *Customer) *TicketOwnerQuery {
	query := &TicketOwnerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(ticketowner.Table, ticketowner.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, customer.TicketTable, customer.TicketColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// FlightClient is a client for the Flight schema.
type FlightClient struct {
	config
}

// NewFlightClient returns a client for the Flight from the given config.
func NewFlightClient(c config) *FlightClient {
	return &FlightClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `flight.Hooks(f(g(h())))`.
func (c *FlightClient) Use(hooks ...Hook) {
	c.hooks.Flight = append(c.hooks.Flight, hooks...)
}

// Create returns a builder for creating a Flight entity.
func (c *FlightClient) Create() *FlightCreate {
	mutation := newFlightMutation(c.config, OpCreate)
	return &FlightCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Flight entities.
func (c *FlightClient) CreateBulk(builders ...*FlightCreate) *FlightCreateBulk {
	return &FlightCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Flight.
func (c *FlightClient) Update() *FlightUpdate {
	mutation := newFlightMutation(c.config, OpUpdate)
	return &FlightUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FlightClient) UpdateOne(f *Flight) *FlightUpdateOne {
	mutation := newFlightMutation(c.config, OpUpdateOne, withFlight(f))
	return &FlightUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FlightClient) UpdateOneID(id uuid.UUID) *FlightUpdateOne {
	mutation := newFlightMutation(c.config, OpUpdateOne, withFlightID(id))
	return &FlightUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Flight.
func (c *FlightClient) Delete() *FlightDelete {
	mutation := newFlightMutation(c.config, OpDelete)
	return &FlightDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FlightClient) DeleteOne(f *Flight) *FlightDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FlightClient) DeleteOneID(id uuid.UUID) *FlightDeleteOne {
	builder := c.Delete().Where(flight.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FlightDeleteOne{builder}
}

// Query returns a query builder for Flight.
func (c *FlightClient) Query() *FlightQuery {
	return &FlightQuery{
		config: c.config,
	}
}

// Get returns a Flight entity by its id.
func (c *FlightClient) Get(ctx context.Context, id uuid.UUID) (*Flight, error) {
	return c.Query().Where(flight.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FlightClient) GetX(ctx context.Context, id uuid.UUID) *Flight {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBelongsTo queries the belongs_to edge of a Flight.
func (c *FlightClient) QueryBelongsTo(f *Flight) *BookingQuery {
	query := &BookingQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flight.Table, flight.FieldID, id),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flight.BelongsToTable, flight.BelongsToColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFlightTickets queries the flight_tickets edge of a Flight.
func (c *FlightClient) QueryFlightTickets(f *Flight) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(flight.Table, flight.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, flight.FlightTicketsTable, flight.FlightTicketsColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FlightClient) Hooks() []Hook {
	return c.hooks.Flight
}

// PassengerClient is a client for the Passenger schema.
type PassengerClient struct {
	config
}

// NewPassengerClient returns a client for the Passenger from the given config.
func NewPassengerClient(c config) *PassengerClient {
	return &PassengerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `passenger.Hooks(f(g(h())))`.
func (c *PassengerClient) Use(hooks ...Hook) {
	c.hooks.Passenger = append(c.hooks.Passenger, hooks...)
}

// Create returns a builder for creating a Passenger entity.
func (c *PassengerClient) Create() *PassengerCreate {
	mutation := newPassengerMutation(c.config, OpCreate)
	return &PassengerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Passenger entities.
func (c *PassengerClient) CreateBulk(builders ...*PassengerCreate) *PassengerCreateBulk {
	return &PassengerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Passenger.
func (c *PassengerClient) Update() *PassengerUpdate {
	mutation := newPassengerMutation(c.config, OpUpdate)
	return &PassengerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PassengerClient) UpdateOne(pa *Passenger) *PassengerUpdateOne {
	mutation := newPassengerMutation(c.config, OpUpdateOne, withPassenger(pa))
	return &PassengerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PassengerClient) UpdateOneID(id uuid.UUID) *PassengerUpdateOne {
	mutation := newPassengerMutation(c.config, OpUpdateOne, withPassengerID(id))
	return &PassengerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Passenger.
func (c *PassengerClient) Delete() *PassengerDelete {
	mutation := newPassengerMutation(c.config, OpDelete)
	return &PassengerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PassengerClient) DeleteOne(pa *Passenger) *PassengerDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PassengerClient) DeleteOneID(id uuid.UUID) *PassengerDeleteOne {
	builder := c.Delete().Where(passenger.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PassengerDeleteOne{builder}
}

// Query returns a query builder for Passenger.
func (c *PassengerClient) Query() *PassengerQuery {
	return &PassengerQuery{
		config: c.config,
	}
}

// Get returns a Passenger entity by its id.
func (c *PassengerClient) Get(ctx context.Context, id uuid.UUID) (*Passenger, error) {
	return c.Query().Where(passenger.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PassengerClient) GetX(ctx context.Context, id uuid.UUID) *Passenger {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTicket queries the ticket edge of a Passenger.
func (c *PassengerClient) QueryTicket(pa *Passenger) *TicketOwnerQuery {
	query := &TicketOwnerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(passenger.Table, passenger.FieldID, id),
			sqlgraph.To(ticketowner.Table, ticketowner.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, passenger.TicketTable, passenger.TicketColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBookings queries the bookings edge of a Passenger.
func (c *PassengerClient) QueryBookings(pa *Passenger) *BookingQuery {
	query := &BookingQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(passenger.Table, passenger.FieldID, id),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, passenger.BookingsTable, passenger.BookingsColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PassengerClient) Hooks() []Hook {
	return c.hooks.Passenger
}

// TicketClient is a client for the Ticket schema.
type TicketClient struct {
	config
}

// NewTicketClient returns a client for the Ticket from the given config.
func NewTicketClient(c config) *TicketClient {
	return &TicketClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ticket.Hooks(f(g(h())))`.
func (c *TicketClient) Use(hooks ...Hook) {
	c.hooks.Ticket = append(c.hooks.Ticket, hooks...)
}

// Create returns a builder for creating a Ticket entity.
func (c *TicketClient) Create() *TicketCreate {
	mutation := newTicketMutation(c.config, OpCreate)
	return &TicketCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ticket entities.
func (c *TicketClient) CreateBulk(builders ...*TicketCreate) *TicketCreateBulk {
	return &TicketCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ticket.
func (c *TicketClient) Update() *TicketUpdate {
	mutation := newTicketMutation(c.config, OpUpdate)
	return &TicketUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TicketClient) UpdateOne(t *Ticket) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicket(t))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TicketClient) UpdateOneID(id uuid.UUID) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicketID(id))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ticket.
func (c *TicketClient) Delete() *TicketDelete {
	mutation := newTicketMutation(c.config, OpDelete)
	return &TicketDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TicketClient) DeleteOne(t *Ticket) *TicketDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TicketClient) DeleteOneID(id uuid.UUID) *TicketDeleteOne {
	builder := c.Delete().Where(ticket.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TicketDeleteOne{builder}
}

// Query returns a query builder for Ticket.
func (c *TicketClient) Query() *TicketQuery {
	return &TicketQuery{
		config: c.config,
	}
}

// Get returns a Ticket entity by its id.
func (c *TicketClient) Get(ctx context.Context, id uuid.UUID) (*Ticket, error) {
	return c.Query().Where(ticket.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TicketClient) GetX(ctx context.Context, id uuid.UUID) *Ticket {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFlightOwner queries the flight_owner edge of a Ticket.
func (c *TicketClient) QueryFlightOwner(t *Ticket) *FlightQuery {
	query := &FlightQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(flight.Table, flight.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticket.FlightOwnerTable, ticket.FlightOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBookingOwner queries the booking_owner edge of a Ticket.
func (c *TicketClient) QueryBookingOwner(t *Ticket) *BookingQuery {
	query := &BookingQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticket.BookingOwnerTable, ticket.BookingOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTicketOwner queries the ticket_owner edge of a Ticket.
func (c *TicketClient) QueryTicketOwner(t *Ticket) *TicketOwnerQuery {
	query := &TicketOwnerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(ticketowner.Table, ticketowner.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ticket.TicketOwnerTable, ticket.TicketOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TicketClient) Hooks() []Hook {
	return c.hooks.Ticket
}

// TicketOwnerClient is a client for the TicketOwner schema.
type TicketOwnerClient struct {
	config
}

// NewTicketOwnerClient returns a client for the TicketOwner from the given config.
func NewTicketOwnerClient(c config) *TicketOwnerClient {
	return &TicketOwnerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ticketowner.Hooks(f(g(h())))`.
func (c *TicketOwnerClient) Use(hooks ...Hook) {
	c.hooks.TicketOwner = append(c.hooks.TicketOwner, hooks...)
}

// Create returns a builder for creating a TicketOwner entity.
func (c *TicketOwnerClient) Create() *TicketOwnerCreate {
	mutation := newTicketOwnerMutation(c.config, OpCreate)
	return &TicketOwnerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TicketOwner entities.
func (c *TicketOwnerClient) CreateBulk(builders ...*TicketOwnerCreate) *TicketOwnerCreateBulk {
	return &TicketOwnerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TicketOwner.
func (c *TicketOwnerClient) Update() *TicketOwnerUpdate {
	mutation := newTicketOwnerMutation(c.config, OpUpdate)
	return &TicketOwnerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TicketOwnerClient) UpdateOne(to *TicketOwner) *TicketOwnerUpdateOne {
	mutation := newTicketOwnerMutation(c.config, OpUpdateOne, withTicketOwner(to))
	return &TicketOwnerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TicketOwnerClient) UpdateOneID(id uuid.UUID) *TicketOwnerUpdateOne {
	mutation := newTicketOwnerMutation(c.config, OpUpdateOne, withTicketOwnerID(id))
	return &TicketOwnerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TicketOwner.
func (c *TicketOwnerClient) Delete() *TicketOwnerDelete {
	mutation := newTicketOwnerMutation(c.config, OpDelete)
	return &TicketOwnerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TicketOwnerClient) DeleteOne(to *TicketOwner) *TicketOwnerDeleteOne {
	return c.DeleteOneID(to.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TicketOwnerClient) DeleteOneID(id uuid.UUID) *TicketOwnerDeleteOne {
	builder := c.Delete().Where(ticketowner.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TicketOwnerDeleteOne{builder}
}

// Query returns a query builder for TicketOwner.
func (c *TicketOwnerClient) Query() *TicketOwnerQuery {
	return &TicketOwnerQuery{
		config: c.config,
	}
}

// Get returns a TicketOwner entity by its id.
func (c *TicketOwnerClient) Get(ctx context.Context, id uuid.UUID) (*TicketOwner, error) {
	return c.Query().Where(ticketowner.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TicketOwnerClient) GetX(ctx context.Context, id uuid.UUID) *TicketOwner {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTicket queries the ticket edge of a TicketOwner.
func (c *TicketOwnerClient) QueryTicket(to *TicketOwner) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := to.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketowner.Table, ticketowner.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ticketowner.TicketTable, ticketowner.TicketColumn),
		)
		fromV = sqlgraph.Neighbors(to.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCustomerOwner queries the customer_owner edge of a TicketOwner.
func (c *TicketOwnerClient) QueryCustomerOwner(to *TicketOwner) *CustomerQuery {
	query := &CustomerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := to.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketowner.Table, ticketowner.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ticketowner.CustomerOwnerTable, ticketowner.CustomerOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(to.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPassengerOwner queries the passenger_owner edge of a TicketOwner.
func (c *TicketOwnerClient) QueryPassengerOwner(to *TicketOwner) *PassengerQuery {
	query := &PassengerQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := to.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketowner.Table, ticketowner.FieldID, id),
			sqlgraph.To(passenger.Table, passenger.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ticketowner.PassengerOwnerTable, ticketowner.PassengerOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(to.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TicketOwnerClient) Hooks() []Hook {
	return c.hooks.TicketOwner
}
