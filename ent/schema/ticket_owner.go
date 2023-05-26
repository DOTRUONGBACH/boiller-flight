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

// TicketOwner holds the schema definition for the TicketOwner entity.
type TicketOwner struct {
	ent.Schema
}

// Fields of the TicketOwner.
func (TicketOwner) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID"), entproto.Field(1)),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT"), entproto.Field(2)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entproto.Field(3)),
	}
}

// Edges of the TicketOwner.
func (TicketOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ticket", Ticket.Type).Unique(),
		edge.From("customer_owner", Customer.Type).
			 Ref("ticket").Unique(),
		edge.From("passenger_owner", Passenger.Type).
			 Ref("ticket").Unique(),
	}
}

func (TicketOwner) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
    }
}
