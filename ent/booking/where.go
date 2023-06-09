// Code generated by ent, DO NOT EDIT.

package booking

import (
	"mock_project/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FlightID applies equality check predicate on the "flight_id" field. It's identical to FlightIDEQ.
func FlightID(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlightID), v))
	})
}

// CustomerID applies equality check predicate on the "customer_id" field. It's identical to CustomerIDEQ.
func CustomerID(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomerID), v))
	})
}

// PassengerID applies equality check predicate on the "passenger_id" field. It's identical to PassengerIDEQ.
func PassengerID(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassengerID), v))
	})
}

// EconomyTickets applies equality check predicate on the "economy_tickets" field. It's identical to EconomyTicketsEQ.
func EconomyTickets(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEconomyTickets), v))
	})
}

// BusinessTickets applies equality check predicate on the "business_tickets" field. It's identical to BusinessTicketsEQ.
func BusinessTickets(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBusinessTickets), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// FlightIDEQ applies the EQ predicate on the "flight_id" field.
func FlightIDEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFlightID), v))
	})
}

// FlightIDNEQ applies the NEQ predicate on the "flight_id" field.
func FlightIDNEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFlightID), v))
	})
}

// FlightIDIn applies the In predicate on the "flight_id" field.
func FlightIDIn(vs ...uuid.UUID) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFlightID), v...))
	})
}

// FlightIDNotIn applies the NotIn predicate on the "flight_id" field.
func FlightIDNotIn(vs ...uuid.UUID) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFlightID), v...))
	})
}

// CustomerIDEQ applies the EQ predicate on the "customer_id" field.
func CustomerIDEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomerID), v))
	})
}

// CustomerIDNEQ applies the NEQ predicate on the "customer_id" field.
func CustomerIDNEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCustomerID), v))
	})
}

// CustomerIDIn applies the In predicate on the "customer_id" field.
func CustomerIDIn(vs ...uuid.UUID) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCustomerID), v...))
	})
}

// CustomerIDNotIn applies the NotIn predicate on the "customer_id" field.
func CustomerIDNotIn(vs ...uuid.UUID) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCustomerID), v...))
	})
}

// CustomerIDIsNil applies the IsNil predicate on the "customer_id" field.
func CustomerIDIsNil() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCustomerID)))
	})
}

// CustomerIDNotNil applies the NotNil predicate on the "customer_id" field.
func CustomerIDNotNil() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCustomerID)))
	})
}

// PassengerIDEQ applies the EQ predicate on the "passenger_id" field.
func PassengerIDEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassengerID), v))
	})
}

// PassengerIDNEQ applies the NEQ predicate on the "passenger_id" field.
func PassengerIDNEQ(v uuid.UUID) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassengerID), v))
	})
}

// PassengerIDIn applies the In predicate on the "passenger_id" field.
func PassengerIDIn(vs ...uuid.UUID) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPassengerID), v...))
	})
}

// PassengerIDNotIn applies the NotIn predicate on the "passenger_id" field.
func PassengerIDNotIn(vs ...uuid.UUID) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPassengerID), v...))
	})
}

// PassengerIDIsNil applies the IsNil predicate on the "passenger_id" field.
func PassengerIDIsNil() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPassengerID)))
	})
}

// PassengerIDNotNil applies the NotNil predicate on the "passenger_id" field.
func PassengerIDNotNil() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPassengerID)))
	})
}

// EconomyTicketsEQ applies the EQ predicate on the "economy_tickets" field.
func EconomyTicketsEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEconomyTickets), v))
	})
}

// EconomyTicketsNEQ applies the NEQ predicate on the "economy_tickets" field.
func EconomyTicketsNEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEconomyTickets), v))
	})
}

// EconomyTicketsIn applies the In predicate on the "economy_tickets" field.
func EconomyTicketsIn(vs ...int) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEconomyTickets), v...))
	})
}

// EconomyTicketsNotIn applies the NotIn predicate on the "economy_tickets" field.
func EconomyTicketsNotIn(vs ...int) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEconomyTickets), v...))
	})
}

// EconomyTicketsGT applies the GT predicate on the "economy_tickets" field.
func EconomyTicketsGT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEconomyTickets), v))
	})
}

// EconomyTicketsGTE applies the GTE predicate on the "economy_tickets" field.
func EconomyTicketsGTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEconomyTickets), v))
	})
}

// EconomyTicketsLT applies the LT predicate on the "economy_tickets" field.
func EconomyTicketsLT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEconomyTickets), v))
	})
}

// EconomyTicketsLTE applies the LTE predicate on the "economy_tickets" field.
func EconomyTicketsLTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEconomyTickets), v))
	})
}

// BusinessTicketsEQ applies the EQ predicate on the "business_tickets" field.
func BusinessTicketsEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBusinessTickets), v))
	})
}

// BusinessTicketsNEQ applies the NEQ predicate on the "business_tickets" field.
func BusinessTicketsNEQ(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBusinessTickets), v))
	})
}

// BusinessTicketsIn applies the In predicate on the "business_tickets" field.
func BusinessTicketsIn(vs ...int) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBusinessTickets), v...))
	})
}

// BusinessTicketsNotIn applies the NotIn predicate on the "business_tickets" field.
func BusinessTicketsNotIn(vs ...int) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBusinessTickets), v...))
	})
}

// BusinessTicketsGT applies the GT predicate on the "business_tickets" field.
func BusinessTicketsGT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBusinessTickets), v))
	})
}

// BusinessTicketsGTE applies the GTE predicate on the "business_tickets" field.
func BusinessTicketsGTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBusinessTickets), v))
	})
}

// BusinessTicketsLT applies the LT predicate on the "business_tickets" field.
func BusinessTicketsLT(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBusinessTickets), v))
	})
}

// BusinessTicketsLTE applies the LTE predicate on the "business_tickets" field.
func BusinessTicketsLTE(v int) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBusinessTickets), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Booking {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasBookingFlight applies the HasEdge predicate on the "booking_flight" edge.
func HasBookingFlight() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookingFlightTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BookingFlightTable, BookingFlightColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookingFlightWith applies the HasEdge predicate on the "booking_flight" edge with a given conditions (other predicates).
func HasBookingFlightWith(preds ...predicate.Flight) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookingFlightInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BookingFlightTable, BookingFlightColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomerBookingTickets applies the HasEdge predicate on the "customer_booking_tickets" edge.
func HasCustomerBookingTickets() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerBookingTicketsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerBookingTicketsTable, CustomerBookingTicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerBookingTicketsWith applies the HasEdge predicate on the "customer_booking_tickets" edge with a given conditions (other predicates).
func HasCustomerBookingTicketsWith(preds ...predicate.Customer) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerBookingTicketsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerBookingTicketsTable, CustomerBookingTicketsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPassengerBookingTickets applies the HasEdge predicate on the "passenger_booking_tickets" edge.
func HasPassengerBookingTickets() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PassengerBookingTicketsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PassengerBookingTicketsTable, PassengerBookingTicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPassengerBookingTicketsWith applies the HasEdge predicate on the "passenger_booking_tickets" edge with a given conditions (other predicates).
func HasPassengerBookingTicketsWith(preds ...predicate.Passenger) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PassengerBookingTicketsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PassengerBookingTicketsTable, PassengerBookingTicketsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBookingTickets applies the HasEdge predicate on the "booking_tickets" edge.
func HasBookingTickets() predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookingTicketsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookingTicketsTable, BookingTicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookingTicketsWith applies the HasEdge predicate on the "booking_tickets" edge with a given conditions (other predicates).
func HasBookingTicketsWith(preds ...predicate.Ticket) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BookingTicketsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BookingTicketsTable, BookingTicketsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Booking) predicate.Booking {
	return predicate.Booking(func(s *sql.Selector) {
		p(s.Not())
	})
}
