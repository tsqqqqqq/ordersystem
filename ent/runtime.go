// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ordersystem/ent/inventory"
	"ordersystem/ent/order"
	"ordersystem/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	inventoryMixin := schema.Inventory{}.Mixin()
	inventoryMixinFields0 := inventoryMixin[0].Fields()
	_ = inventoryMixinFields0
	inventoryFields := schema.Inventory{}.Fields()
	_ = inventoryFields
	// inventoryDescCreateTime is the schema descriptor for create_time field.
	inventoryDescCreateTime := inventoryMixinFields0[0].Descriptor()
	// inventory.DefaultCreateTime holds the default value on creation for the create_time field.
	inventory.DefaultCreateTime = inventoryDescCreateTime.Default.(func() time.Time)
	// inventoryDescUpdateTime is the schema descriptor for update_time field.
	inventoryDescUpdateTime := inventoryMixinFields0[1].Descriptor()
	// inventory.DefaultUpdateTime holds the default value on creation for the update_time field.
	inventory.DefaultUpdateTime = inventoryDescUpdateTime.Default.(func() time.Time)
	// inventory.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	inventory.UpdateDefaultUpdateTime = inventoryDescUpdateTime.UpdateDefault.(func() time.Time)
	// inventoryDescID is the schema descriptor for id field.
	inventoryDescID := inventoryFields[0].Descriptor()
	// inventory.DefaultID holds the default value on creation for the id field.
	inventory.DefaultID = inventoryDescID.Default.(func() int64)
	orderMixin := schema.Order{}.Mixin()
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreateTime is the schema descriptor for create_time field.
	orderDescCreateTime := orderMixinFields0[0].Descriptor()
	// order.DefaultCreateTime holds the default value on creation for the create_time field.
	order.DefaultCreateTime = orderDescCreateTime.Default.(func() time.Time)
	// orderDescUpdateTime is the schema descriptor for update_time field.
	orderDescUpdateTime := orderMixinFields0[1].Descriptor()
	// order.DefaultUpdateTime holds the default value on creation for the update_time field.
	order.DefaultUpdateTime = orderDescUpdateTime.Default.(func() time.Time)
	// order.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	order.UpdateDefaultUpdateTime = orderDescUpdateTime.UpdateDefault.(func() time.Time)
	// orderDescCount is the schema descriptor for count field.
	orderDescCount := orderFields[3].Descriptor()
	// order.DefaultCount holds the default value on creation for the count field.
	order.DefaultCount = orderDescCount.Default.(int)
	// orderDescStatus is the schema descriptor for status field.
	orderDescStatus := orderFields[4].Descriptor()
	// order.DefaultStatus holds the default value on creation for the status field.
	order.DefaultStatus = orderDescStatus.Default.(int)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderFields[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() int64)
}
