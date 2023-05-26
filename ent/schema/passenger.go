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

// Passenger holds the schema definition for the Passenger entity.
type Passenger struct {
	ent.Schema
}

// Fields of the Passenger.
func (Passenger) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID"), entproto.Field(1)),
		field.String("name").MaxLen(255).Annotations(entgql.OrderField("NAME"), entproto.Field(2)),
		field.String("citizen_id").MinLen(9).MaxLen(12).Unique().Annotations(entgql.OrderField("CITIZEN_ID"), entproto.Field(3)),
		field.String("email").Annotations(entgql.OrderField("EMAIL"), entproto.Field(4)),
		field.String("phone").MaxLen(10).Annotations(entgql.OrderField("PHONE"), entproto.Field(5)),
		field.String("address").Annotations(entgql.OrderField("ADDRESS"), entproto.Field(6)),
		field.Enum("gender").Values("Male", "Female", "Other").Annotations(entgql.OrderField("GENDER"), entproto.Field(7)),
		field.Time("date_of_birth").Annotations(entgql.OrderField("DOB"), entproto.Field(8)),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT"), entproto.Field(9)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entproto.Field(10)),
	}
}

// Edges of the Passenger.
func (Passenger) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ticket", TicketOwner.Type).Unique(),
		edge.To("bookings", Booking.Type),
	}
}

func (Passenger) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
