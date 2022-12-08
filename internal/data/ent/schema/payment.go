package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Post.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("order_no"),
		field.String("transaction_id"),
		field.String("payment_type"),
		field.String("trade_type"),
		field.String("trade_state"),
		field.Int64("payer_total"),
		field.Text("content"),

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
