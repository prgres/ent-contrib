// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithstrings"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithStringsCreate is the builder for creating a MessageWithStrings entity.
type MessageWithStringsCreate struct {
	config
	mutation *MessageWithStringsMutation
	hooks    []Hook
}

// SetStrings sets the "strings" field.
func (mwsc *MessageWithStringsCreate) SetStrings(s []string) *MessageWithStringsCreate {
	mwsc.mutation.SetStrings(s)
	return mwsc
}

// Mutation returns the MessageWithStringsMutation object of the builder.
func (mwsc *MessageWithStringsCreate) Mutation() *MessageWithStringsMutation {
	return mwsc.mutation
}

// Save creates the MessageWithStrings in the database.
func (mwsc *MessageWithStringsCreate) Save(ctx context.Context) (*MessageWithStrings, error) {
	return withHooks[*MessageWithStrings, MessageWithStringsMutation](ctx, mwsc.sqlSave, mwsc.mutation, mwsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mwsc *MessageWithStringsCreate) SaveX(ctx context.Context) *MessageWithStrings {
	v, err := mwsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwsc *MessageWithStringsCreate) Exec(ctx context.Context) error {
	_, err := mwsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwsc *MessageWithStringsCreate) ExecX(ctx context.Context) {
	if err := mwsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mwsc *MessageWithStringsCreate) check() error {
	if _, ok := mwsc.mutation.Strings(); !ok {
		return &ValidationError{Name: "strings", err: errors.New(`ent: missing required field "MessageWithStrings.strings"`)}
	}
	return nil
}

func (mwsc *MessageWithStringsCreate) sqlSave(ctx context.Context) (*MessageWithStrings, error) {
	if err := mwsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mwsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mwsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mwsc.mutation.id = &_node.ID
	mwsc.mutation.done = true
	return _node, nil
}

func (mwsc *MessageWithStringsCreate) createSpec() (*MessageWithStrings, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageWithStrings{config: mwsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: messagewithstrings.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithstrings.FieldID,
			},
		}
	)
	if value, ok := mwsc.mutation.Strings(); ok {
		_spec.SetField(messagewithstrings.FieldStrings, field.TypeJSON, value)
		_node.Strings = value
	}
	return _node, _spec
}

// MessageWithStringsCreateBulk is the builder for creating many MessageWithStrings entities in bulk.
type MessageWithStringsCreateBulk struct {
	config
	builders []*MessageWithStringsCreate
}

// Save creates the MessageWithStrings entities in the database.
func (mwscb *MessageWithStringsCreateBulk) Save(ctx context.Context) ([]*MessageWithStrings, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mwscb.builders))
	nodes := make([]*MessageWithStrings, len(mwscb.builders))
	mutators := make([]Mutator, len(mwscb.builders))
	for i := range mwscb.builders {
		func(i int, root context.Context) {
			builder := mwscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageWithStringsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mwscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mwscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mwscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mwscb *MessageWithStringsCreateBulk) SaveX(ctx context.Context) []*MessageWithStrings {
	v, err := mwscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwscb *MessageWithStringsCreateBulk) Exec(ctx context.Context) error {
	_, err := mwscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwscb *MessageWithStringsCreateBulk) ExecX(ctx context.Context) {
	if err := mwscb.Exec(ctx); err != nil {
		panic(err)
	}
}
