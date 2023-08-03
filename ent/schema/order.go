package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"ordersystem/common"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").DefaultFunc(func() int64 {
			snow := new(common.Snowflake)
			return snow.NextVal()
		}),
		field.String("name"),
		field.Int64("inventory_id").Optional(),
		field.Int("count").Default(0).Comment("buy count"),
		field.Int("status").Default(0).Comment("order status, 0: unprocessed, 1: processed, -1: canceled, -2: failed"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("inventory", Inventory.Type).Ref("order").Unique().Field("inventory_id"),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
