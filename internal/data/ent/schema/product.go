package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Post.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title"),
		field.Uint64("price"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.Postgres: "timestamp with time zone",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.Postgres: "timestamp with time zone",
		}),
	}
}
