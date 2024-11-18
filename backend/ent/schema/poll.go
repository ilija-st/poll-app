package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Poll holds the schema definition for the Poll entity.
type Poll struct {
	ent.Schema
}

// Fields of the Poll.
func (Poll) Fields() []ent.Field {
	return []ent.Field{
		field.String("question").NotEmpty().MaxLen(1024),
		field.String("status").NotEmpty().Default("open"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (Poll) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("polls").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
		edge.To("poll_options", PollOption.Type),
	}
}
