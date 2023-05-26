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

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID"), entproto.Field(1)),
		field.String("email").Unique().Annotations(entgql.OrderField("EMAIL"), entproto.Field(2)),
		field.String("password").Sensitive().Unique().Annotations(entgql.OrderField("PASSWORD"), entproto.Field(3)),
		field.Enum("role").Values("Administrator", "Subscriber").Annotations(entgql.OrderField("ROLE"), entproto.Field(4)),
		field.Enum("status").Values("Inactive", "Active").Annotations(entgql.OrderField("STATUS"), entproto.Field(5)),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT"), entproto.Field(6)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entproto.Field(7)),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("acc_owner", Customer.Type).
			 Ref("accounts").
			 Unique(),
	}
}

func (Account) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
    }
}
