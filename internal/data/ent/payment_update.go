// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pay/internal/data/ent/payment"
	"pay/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetOrderNo sets the "order_no" field.
func (pu *PaymentUpdate) SetOrderNo(s string) *PaymentUpdate {
	pu.mutation.SetOrderNo(s)
	return pu
}

// SetTransactionID sets the "transaction_id" field.
func (pu *PaymentUpdate) SetTransactionID(s string) *PaymentUpdate {
	pu.mutation.SetTransactionID(s)
	return pu
}

// SetPaymentType sets the "payment_type" field.
func (pu *PaymentUpdate) SetPaymentType(s string) *PaymentUpdate {
	pu.mutation.SetPaymentType(s)
	return pu
}

// SetTradeType sets the "trade_type" field.
func (pu *PaymentUpdate) SetTradeType(s string) *PaymentUpdate {
	pu.mutation.SetTradeType(s)
	return pu
}

// SetTradeState sets the "trade_state" field.
func (pu *PaymentUpdate) SetTradeState(s string) *PaymentUpdate {
	pu.mutation.SetTradeState(s)
	return pu
}

// SetPayerTotal sets the "payer_total" field.
func (pu *PaymentUpdate) SetPayerTotal(i int64) *PaymentUpdate {
	pu.mutation.ResetPayerTotal()
	pu.mutation.SetPayerTotal(i)
	return pu
}

// AddPayerTotal adds i to the "payer_total" field.
func (pu *PaymentUpdate) AddPayerTotal(i int64) *PaymentUpdate {
	pu.mutation.AddPayerTotal(i)
	return pu
}

// SetContent sets the "content" field.
func (pu *PaymentUpdate) SetContent(s string) *PaymentUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PaymentUpdate) SetCreatedAt(t time.Time) *PaymentUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableCreatedAt(t *time.Time) *PaymentUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PaymentUpdate) SetUpdatedAt(t time.Time) *PaymentUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableUpdatedAt(t *time.Time) *PaymentUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: payment.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.OrderNo(); ok {
		_spec.SetField(payment.FieldOrderNo, field.TypeString, value)
	}
	if value, ok := pu.mutation.TransactionID(); ok {
		_spec.SetField(payment.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := pu.mutation.PaymentType(); ok {
		_spec.SetField(payment.FieldPaymentType, field.TypeString, value)
	}
	if value, ok := pu.mutation.TradeType(); ok {
		_spec.SetField(payment.FieldTradeType, field.TypeString, value)
	}
	if value, ok := pu.mutation.TradeState(); ok {
		_spec.SetField(payment.FieldTradeState, field.TypeString, value)
	}
	if value, ok := pu.mutation.PayerTotal(); ok {
		_spec.SetField(payment.FieldPayerTotal, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedPayerTotal(); ok {
		_spec.AddField(payment.FieldPayerTotal, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(payment.FieldContent, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(payment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(payment.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetOrderNo sets the "order_no" field.
func (puo *PaymentUpdateOne) SetOrderNo(s string) *PaymentUpdateOne {
	puo.mutation.SetOrderNo(s)
	return puo
}

// SetTransactionID sets the "transaction_id" field.
func (puo *PaymentUpdateOne) SetTransactionID(s string) *PaymentUpdateOne {
	puo.mutation.SetTransactionID(s)
	return puo
}

// SetPaymentType sets the "payment_type" field.
func (puo *PaymentUpdateOne) SetPaymentType(s string) *PaymentUpdateOne {
	puo.mutation.SetPaymentType(s)
	return puo
}

// SetTradeType sets the "trade_type" field.
func (puo *PaymentUpdateOne) SetTradeType(s string) *PaymentUpdateOne {
	puo.mutation.SetTradeType(s)
	return puo
}

// SetTradeState sets the "trade_state" field.
func (puo *PaymentUpdateOne) SetTradeState(s string) *PaymentUpdateOne {
	puo.mutation.SetTradeState(s)
	return puo
}

// SetPayerTotal sets the "payer_total" field.
func (puo *PaymentUpdateOne) SetPayerTotal(i int64) *PaymentUpdateOne {
	puo.mutation.ResetPayerTotal()
	puo.mutation.SetPayerTotal(i)
	return puo
}

// AddPayerTotal adds i to the "payer_total" field.
func (puo *PaymentUpdateOne) AddPayerTotal(i int64) *PaymentUpdateOne {
	puo.mutation.AddPayerTotal(i)
	return puo
}

// SetContent sets the "content" field.
func (puo *PaymentUpdateOne) SetContent(s string) *PaymentUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PaymentUpdateOne) SetCreatedAt(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableCreatedAt(t *time.Time) *PaymentUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PaymentUpdateOne) SetUpdatedAt(t time.Time) *PaymentUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableUpdatedAt(t *time.Time) *PaymentUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	var (
		err  error
		node *Payment
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Payment)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PaymentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payment.Table,
			Columns: payment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: payment.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.OrderNo(); ok {
		_spec.SetField(payment.FieldOrderNo, field.TypeString, value)
	}
	if value, ok := puo.mutation.TransactionID(); ok {
		_spec.SetField(payment.FieldTransactionID, field.TypeString, value)
	}
	if value, ok := puo.mutation.PaymentType(); ok {
		_spec.SetField(payment.FieldPaymentType, field.TypeString, value)
	}
	if value, ok := puo.mutation.TradeType(); ok {
		_spec.SetField(payment.FieldTradeType, field.TypeString, value)
	}
	if value, ok := puo.mutation.TradeState(); ok {
		_spec.SetField(payment.FieldTradeState, field.TypeString, value)
	}
	if value, ok := puo.mutation.PayerTotal(); ok {
		_spec.SetField(payment.FieldPayerTotal, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedPayerTotal(); ok {
		_spec.AddField(payment.FieldPayerTotal, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(payment.FieldContent, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(payment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(payment.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
