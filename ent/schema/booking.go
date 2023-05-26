package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID"), entproto.Field(1)),
		field.UUID("flight_id", uuid.UUID{}).Annotations(entgql.OrderField("FLIGHT_ID"), entproto.Field(2)),
		field.UUID("customer_id", uuid.UUID{}).Optional().Annotations(entgql.OrderField("CUSTOMER_ID"), entproto.Field(3)),
		field.UUID("passenger_id", uuid.UUID{}).Optional().Annotations(entgql.OrderField("PASSENGER_ID"), entproto.Field(4)),
		field.Int("economy_tickets").Positive().Default(0).Annotations(entgql.OrderField("ECONOMY_TICKET"), entproto.Field(5)),
		field.Int("business_tickets").Positive().Default(0).Annotations(entgql.OrderField("BUSINESS_TICKET"), entproto.Field(6)),
		field.Enum("status").Values("Success", "Fail", "Canceled").Annotations(entgql.OrderField("STATUS"), entproto.Field(7)),
		field.Enum("type").Values("OneWay", "RoundTrip").Annotations(entgql.OrderField("TYPE"), entproto.Field(8)),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT"), entproto.Field(9)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entproto.Field(10)),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("booking_flight", Flight.Type).Field("flight_id").
			 Ref("belongs_to").Unique().Required(),
		edge.From("customer_booking_tickets", Customer.Type).Field("customer_id").
			 Ref("bookings").Unique(),
		edge.From("passenger_booking_tickets", Passenger.Type).Field("passenger_id").
			 Ref("bookings").Unique(),
		edge.To("booking_tickets", Ticket.Type),
	}
}

func (Booking) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
    }
}
