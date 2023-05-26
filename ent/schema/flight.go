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

// Flight holds the schema definition for the Flight entity.
type Flight struct {
	ent.Schema
}

// Fields of the Flight.
func (Flight) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID"), entproto.Field(1)),
		field.String("flight_code").MinLen(6).Annotations(entgql.OrderField("FLIGHT_CODE"), entproto.Field(2)),
		field.String("from").MaxLen(255).Annotations(entgql.OrderField("FROM"), entproto.Field(3)),
		field.String("to").MaxLen(255).Annotations(entgql.OrderField("TO"), entproto.Field(4)),
		field.Time("departure_date").Annotations(entgql.OrderField("DEPARTURE_DATE"), entproto.Field(5)),
		field.Time("departure_time").Annotations(entgql.OrderField("DEPARTURE_TIME"), entproto.Field(6)),
		field.Int("duration").Positive().Annotations(entgql.OrderField("DURATION"), entproto.Field(7)),
		field.Int("capacity").Positive().Annotations(entgql.OrderField("CAPACITY"), entproto.Field(8)),
		field.Int("economy_available_seat").Min(0).Annotations(entgql.OrderField("ECONOMY_AVAILABLE_SEAT"), entproto.Field(9)),
		field.Int("business_available_seat").Min(0).Annotations(entgql.OrderField("BUSINESS_AVAILABLE_SEAT"), entproto.Field(10)),
		field.Enum("status").Values("Scheduled", "Delayed", "Departed", "Arrived", "Canceled", "Full" ).Annotations(entgql.OrderField("STATUS"), entproto.Field(11)),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT"), entproto.Field(12)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entproto.Field(13)),
	}
}

// Edges of the Flight.
func (Flight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("belongs_to", Booking.Type),
		edge.To("flight_tickets", Ticket.Type),
	}
}

func (Flight) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
    }
}
