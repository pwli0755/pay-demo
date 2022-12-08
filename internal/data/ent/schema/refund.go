package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Refund holds the schema definition for the Refund entity.
type Refund struct {
	ent.Schema
}

// Fields of the Post.
func (Refund) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("order_no"),
		field.String("refund_no"),
		field.String("refund_id"),
		field.Int64("total_fee"),
		field.Int64("refund"),
		field.String("reason"),
		field.String("refund_status"),

		field.Text("content_return"),
		field.Text("content_notify"),

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
