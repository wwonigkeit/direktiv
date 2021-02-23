// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vorteil/direktiv/pkg/secrets/ent/bucketsecret"
	"github.com/vorteil/direktiv/pkg/secrets/ent/predicate"
)

// BucketSecretUpdate is the builder for updating BucketSecret entities.
type BucketSecretUpdate struct {
	config
	hooks    []Hook
	mutation *BucketSecretMutation
}

// Where adds a new predicate for the BucketSecretUpdate builder.
func (bsu *BucketSecretUpdate) Where(ps ...predicate.BucketSecret) *BucketSecretUpdate {
	bsu.mutation.predicates = append(bsu.mutation.predicates, ps...)
	return bsu
}

// SetNs sets the "ns" field.
func (bsu *BucketSecretUpdate) SetNs(s string) *BucketSecretUpdate {
	bsu.mutation.SetNs(s)
	return bsu
}

// SetName sets the "name" field.
func (bsu *BucketSecretUpdate) SetName(s string) *BucketSecretUpdate {
	bsu.mutation.SetName(s)
	return bsu
}

// SetSecret sets the "secret" field.
func (bsu *BucketSecretUpdate) SetSecret(b []byte) *BucketSecretUpdate {
	bsu.mutation.SetSecret(b)
	return bsu
}

// SetType sets the "type" field.
func (bsu *BucketSecretUpdate) SetType(i int) *BucketSecretUpdate {
	bsu.mutation.ResetType()
	bsu.mutation.SetType(i)
	return bsu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bsu *BucketSecretUpdate) SetNillableType(i *int) *BucketSecretUpdate {
	if i != nil {
		bsu.SetType(*i)
	}
	return bsu
}

// AddType adds i to the "type" field.
func (bsu *BucketSecretUpdate) AddType(i int) *BucketSecretUpdate {
	bsu.mutation.AddType(i)
	return bsu
}

// Mutation returns the BucketSecretMutation object of the builder.
func (bsu *BucketSecretUpdate) Mutation() *BucketSecretMutation {
	return bsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bsu *BucketSecretUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bsu.hooks) == 0 {
		affected, err = bsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BucketSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bsu.mutation = mutation
			affected, err = bsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bsu.hooks) - 1; i >= 0; i-- {
			mut = bsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bsu *BucketSecretUpdate) SaveX(ctx context.Context) int {
	affected, err := bsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bsu *BucketSecretUpdate) Exec(ctx context.Context) error {
	_, err := bsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsu *BucketSecretUpdate) ExecX(ctx context.Context) {
	if err := bsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bsu *BucketSecretUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bucketsecret.Table,
			Columns: bucketsecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bucketsecret.FieldID,
			},
		},
	}
	if ps := bsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bsu.mutation.Ns(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bucketsecret.FieldNs,
		})
	}
	if value, ok := bsu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bucketsecret.FieldName,
		})
	}
	if value, ok := bsu.mutation.Secret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: bucketsecret.FieldSecret,
		})
	}
	if value, ok := bsu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bucketsecret.FieldType,
		})
	}
	if value, ok := bsu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bucketsecret.FieldType,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bucketsecret.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BucketSecretUpdateOne is the builder for updating a single BucketSecret entity.
type BucketSecretUpdateOne struct {
	config
	hooks    []Hook
	mutation *BucketSecretMutation
}

// SetNs sets the "ns" field.
func (bsuo *BucketSecretUpdateOne) SetNs(s string) *BucketSecretUpdateOne {
	bsuo.mutation.SetNs(s)
	return bsuo
}

// SetName sets the "name" field.
func (bsuo *BucketSecretUpdateOne) SetName(s string) *BucketSecretUpdateOne {
	bsuo.mutation.SetName(s)
	return bsuo
}

// SetSecret sets the "secret" field.
func (bsuo *BucketSecretUpdateOne) SetSecret(b []byte) *BucketSecretUpdateOne {
	bsuo.mutation.SetSecret(b)
	return bsuo
}

// SetType sets the "type" field.
func (bsuo *BucketSecretUpdateOne) SetType(i int) *BucketSecretUpdateOne {
	bsuo.mutation.ResetType()
	bsuo.mutation.SetType(i)
	return bsuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bsuo *BucketSecretUpdateOne) SetNillableType(i *int) *BucketSecretUpdateOne {
	if i != nil {
		bsuo.SetType(*i)
	}
	return bsuo
}

// AddType adds i to the "type" field.
func (bsuo *BucketSecretUpdateOne) AddType(i int) *BucketSecretUpdateOne {
	bsuo.mutation.AddType(i)
	return bsuo
}

// Mutation returns the BucketSecretMutation object of the builder.
func (bsuo *BucketSecretUpdateOne) Mutation() *BucketSecretMutation {
	return bsuo.mutation
}

// Save executes the query and returns the updated BucketSecret entity.
func (bsuo *BucketSecretUpdateOne) Save(ctx context.Context) (*BucketSecret, error) {
	var (
		err  error
		node *BucketSecret
	)
	if len(bsuo.hooks) == 0 {
		node, err = bsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BucketSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bsuo.mutation = mutation
			node, err = bsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bsuo.hooks) - 1; i >= 0; i-- {
			mut = bsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (bsuo *BucketSecretUpdateOne) SaveX(ctx context.Context) *BucketSecret {
	node, err := bsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bsuo *BucketSecretUpdateOne) Exec(ctx context.Context) error {
	_, err := bsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsuo *BucketSecretUpdateOne) ExecX(ctx context.Context) {
	if err := bsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bsuo *BucketSecretUpdateOne) sqlSave(ctx context.Context) (_node *BucketSecret, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bucketsecret.Table,
			Columns: bucketsecret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bucketsecret.FieldID,
			},
		},
	}
	id, ok := bsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing BucketSecret.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := bsuo.mutation.Ns(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bucketsecret.FieldNs,
		})
	}
	if value, ok := bsuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bucketsecret.FieldName,
		})
	}
	if value, ok := bsuo.mutation.Secret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: bucketsecret.FieldSecret,
		})
	}
	if value, ok := bsuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bucketsecret.FieldType,
		})
	}
	if value, ok := bsuo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bucketsecret.FieldType,
		})
	}
	_node = &BucketSecret{config: bsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bucketsecret.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}