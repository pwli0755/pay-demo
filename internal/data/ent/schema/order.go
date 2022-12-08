package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Post.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title"),
		field.String("order_no"),
		field.String("payment_type"),
		field.String("user_id"),
		field.Int64("product_id"),
		field.Int64("total_fee"),
		field.String("order_status"),

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
