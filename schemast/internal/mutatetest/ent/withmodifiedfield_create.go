// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/user"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withmodifiedfield"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithModifiedFieldCreate is the builder for creating a WithModifiedField entity.
type WithModifiedFieldCreate struct {
	config
	mutation *WithModifiedFieldMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wmfc *WithModifiedFieldCreate) SetName(s string) *WithModifiedFieldCreate {
	wmfc.mutation.SetName(s)
	return wmfc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (wmfc *WithModifiedFieldCreate) SetOwnerID(id int) *WithModifiedFieldCreate {
	wmfc.mutation.SetOwnerID(id)
	return wmfc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (wmfc *WithModifiedFieldCreate) SetNillableOwnerID(id *int) *WithModifiedFieldCreate {
	if id != nil {
		wmfc = wmfc.SetOwnerID(*id)
	}
	return wmfc
}

// SetOwner sets the "owner" edge to the User entity.
func (wmfc *WithModifiedFieldCreate) SetOwner(u *User) *WithModifiedFieldCreate {
	return wmfc.SetOwnerID(u.ID)
}

// Mutation returns the WithModifiedFieldMutation object of the builder.
func (wmfc *WithModifiedFieldCreate) Mutation() *WithModifiedFieldMutation {
	return wmfc.mutation
}

// Save creates the WithModifiedField in the database.
func (wmfc *WithModifiedFieldCreate) Save(ctx context.Context) (*WithModifiedField, error) {
	return withHooks[*WithModifiedField, WithModifiedFieldMutation](ctx, wmfc.sqlSave, wmfc.mutation, wmfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wmfc *WithModifiedFieldCreate) SaveX(ctx context.Context) *WithModifiedField {
	v, err := wmfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wmfc *WithModifiedFieldCreate) Exec(ctx context.Context) error {
	_, err := wmfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wmfc *WithModifiedFieldCreate) ExecX(ctx context.Context) {
	if err := wmfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wmfc *WithModifiedFieldCreate) check() error {
	if _, ok := wmfc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WithModifiedField.name"`)}
	}
	if v, ok := wmfc.mutation.Name(); ok {
		if err := withmodifiedfield.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WithModifiedField.name": %w`, err)}
		}
	}
	return nil
}

func (wmfc *WithModifiedFieldCreate) sqlSave(ctx context.Context) (*WithModifiedField, error) {
	if err := wmfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wmfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wmfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wmfc.mutation.id = &_node.ID
	wmfc.mutation.done = true
	return _node, nil
}

func (wmfc *WithModifiedFieldCreate) createSpec() (*WithModifiedField, *sqlgraph.CreateSpec) {
	var (
		_node = &WithModifiedField{config: wmfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: withmodifiedfield.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withmodifiedfield.FieldID,
			},
		}
	)
	if value, ok := wmfc.mutation.Name(); ok {
		_spec.SetField(withmodifiedfield.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := wmfc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   withmodifiedfield.OwnerTable,
			Columns: []string{withmodifiedfield.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.with_modified_field_owner = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WithModifiedFieldCreateBulk is the builder for creating many WithModifiedField entities in bulk.
type WithModifiedFieldCreateBulk struct {
	config
	builders []*WithModifiedFieldCreate
}

// Save creates the WithModifiedField entities in the database.
func (wmfcb *WithModifiedFieldCreateBulk) Save(ctx context.Context) ([]*WithModifiedField, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wmfcb.builders))
	nodes := make([]*WithModifiedField, len(wmfcb.builders))
	mutators := make([]Mutator, len(wmfcb.builders))
	for i := range wmfcb.builders {
		func(i int, root context.Context) {
			builder := wmfcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WithModifiedFieldMutation)
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
					_, err = mutators[i+1].Mutate(root, wmfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wmfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wmfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wmfcb *WithModifiedFieldCreateBulk) SaveX(ctx context.Context) []*WithModifiedField {
	v, err := wmfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wmfcb *WithModifiedFieldCreateBulk) Exec(ctx context.Context) error {
	_, err := wmfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wmfcb *WithModifiedFieldCreateBulk) ExecX(ctx context.Context) {
	if err := wmfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
