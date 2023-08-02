// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"ordersystem/ent/inventory"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Inventory is the model entity for the Inventory schema.
type Inventory struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Total holds the value of the "total" field.
	Total int `json:"total,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InventoryQuery when eager-loading is set.
	Edges        InventoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// InventoryEdges holds the relations/edges for other nodes in the graph.
type InventoryEdges struct {
	// Order holds the value of the order edge.
	Order []*Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading.
func (e InventoryEdges) OrderOrErr() ([]*Order, error) {
	if e.loadedTypes[0] {
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Inventory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case inventory.FieldID, inventory.FieldTotal:
			values[i] = new(sql.NullInt64)
		case inventory.FieldName, inventory.FieldDescription:
			values[i] = new(sql.NullString)
		case inventory.FieldCreateTime, inventory.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Inventory fields.
func (i *Inventory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case inventory.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int64(value.Int64)
		case inventory.FieldCreateTime:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[j])
			} else if value.Valid {
				i.CreateTime = value.Time
			}
		case inventory.FieldUpdateTime:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[j])
			} else if value.Valid {
				i.UpdateTime = value.Time
			}
		case inventory.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case inventory.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		case inventory.FieldTotal:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[j])
			} else if value.Valid {
				i.Total = int(value.Int64)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Inventory.
// This includes values selected through modifiers, order, etc.
func (i *Inventory) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the Inventory entity.
func (i *Inventory) QueryOrder() *OrderQuery {
	return NewInventoryClient(i.config).QueryOrder(i)
}

// Update returns a builder for updating this Inventory.
// Note that you need to call Inventory.Unwrap() before calling this method if this Inventory
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Inventory) Update() *InventoryUpdateOne {
	return NewInventoryClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Inventory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Inventory) Unwrap() *Inventory {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Inventory is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Inventory) String() string {
	var builder strings.Builder
	builder.WriteString("Inventory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("create_time=")
	builder.WriteString(i.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(i.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteString(", ")
	builder.WriteString("total=")
	builder.WriteString(fmt.Sprintf("%v", i.Total))
	builder.WriteByte(')')
	return builder.String()
}

// Inventories is a parsable slice of Inventory.
type Inventories []*Inventory
