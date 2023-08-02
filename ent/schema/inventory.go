package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"ordersystem/common"
)

// Inventory holds the schema definition for the Inventory entity.
type Inventory struct {
	ent.Schema
}

// Fields of the Inventory.
func (Inventory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").DefaultFunc(func() int64 {
			snow := new(common.Snowflake)
			return snow.NextVal()
		}),
		field.String("name"),
		field.String("description"),
		field.Int("total"),
	}
}

// Edges of the Inventory.
func (Inventory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type),
	}
}

func (Inventory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
