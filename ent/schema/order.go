package schema

import (
	"entgo.io/ent"
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
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
