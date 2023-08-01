// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ordersystem/ent/order"
	"ordersystem/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeOrder = "Order"
)

// OrderMutation represents an operation that mutates the Order nodes in the graph.
type OrderMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	create_time   *time.Time
	update_time   *time.Time
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Order, error)
	predicates    []predicate.Order
}

var _ ent.Mutation = (*OrderMutation)(nil)

// orderOption allows management of the mutation configuration using functional options.
type orderOption func(*OrderMutation)

// newOrderMutation creates new mutation for the Order entity.
func newOrderMutation(c config, op Op, opts ...orderOption) *OrderMutation {
	m := &OrderMutation{
		config:        c,
		op:            op,
		typ:           TypeOrder,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOrderID sets the ID field of the mutation.
func withOrderID(id int64) orderOption {
	return func(m *OrderMutation) {
		var (
			err   error
			once  sync.Once
			value *Order
		)
		m.oldValue = func(ctx context.Context) (*Order, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Order.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOrder sets the old Order of the mutation.
func withOrder(node *Order) orderOption {
	return func(m *OrderMutation) {
		m.oldValue = func(context.Context) (*Order, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OrderMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OrderMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Order entities.
func (m *OrderMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OrderMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *OrderMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Order.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreateTime sets the "create_time" field.
func (m *OrderMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *OrderMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *OrderMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *OrderMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *OrderMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *OrderMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetName sets the "name" field.
func (m *OrderMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *OrderMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *OrderMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the OrderMutation builder.
func (m *OrderMutation) Where(ps ...predicate.Order) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the OrderMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *OrderMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Order, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *OrderMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *OrderMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Order).
func (m *OrderMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OrderMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.create_time != nil {
		fields = append(fields, order.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, order.FieldUpdateTime)
	}
	if m.name != nil {
		fields = append(fields, order.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OrderMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case order.FieldCreateTime:
		return m.CreateTime()
	case order.FieldUpdateTime:
		return m.UpdateTime()
	case order.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OrderMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case order.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case order.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case order.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Order field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) SetField(name string, value ent.Value) error {
	switch name {
	case order.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case order.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case order.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OrderMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OrderMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Order numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OrderMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OrderMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OrderMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Order nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OrderMutation) ResetField(name string) error {
	switch name {
	case order.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case order.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case order.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OrderMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OrderMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OrderMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OrderMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OrderMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OrderMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OrderMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Order unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OrderMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Order edge %s", name)
}
