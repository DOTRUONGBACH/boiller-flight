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

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID"), entproto.Field(1)),
		field.Enum("class").Values("Economy", "Business").Annotations(entgql.OrderField("CLASS"), entproto.Field(2)),
		field.Enum("status").Values("Paid", "Canceled").Annotations(entgql.OrderField("STATUS"), entproto.Field(3)),
		field.UUID("flight_id", uuid.UUID{}).Annotations(entgql.OrderField("FLIGHT_ID"), entproto.Field(4)),
		field.UUID("booking_id", uuid.UUID{}).Annotations(entgql.OrderField("BOOKING_ID"), entproto.Field(5)),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT"), entproto.Field(6)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entproto.Field(7)),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("flight_owner", Flight.Type).Field("flight_id").
			Ref("flight_tickets").Unique().Required(),
		edge.From("booking_owner", Booking.Type).Field("booking_id").
			Ref("booking_tickets").Unique().Required(),
		edge.From("ticket_owner", TicketOwner.Type).
			Ref("ticket").Unique(),
	}
}

func (Ticket) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
