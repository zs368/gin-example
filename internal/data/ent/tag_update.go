// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zs368/gin-example/internal/data/ent/predicate"
	"github.com/zs368/gin-example/internal/data/ent/tag"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TagUpdate) SetStatus(i int) *TagUpdate {
	tu.mutation.ResetStatus()
	tu.mutation.SetStatus(i)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TagUpdate) SetNillableStatus(i *int) *TagUpdate {
	if i != nil {
		tu.SetStatus(*i)
	}
	return tu
}

// AddStatus adds i to the "status" field.
func (tu *TagUpdate) AddStatus(i int) *TagUpdate {
	tu.mutation.AddStatus(i)
	return tu
}

// SetCreatedBy sets the "created_by" field.
func (tu *TagUpdate) SetCreatedBy(s string) *TagUpdate {
	tu.mutation.SetCreatedBy(s)
	return tu
}

// SetUpdatedBy sets the "updated_by" field.
func (tu *TagUpdate) SetUpdatedBy(s string) *TagUpdate {
	tu.mutation.SetUpdatedBy(s)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TagUpdate) SetCreatedAt(t time.Time) *TagUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TagUpdate) SetNillableCreatedAt(t *time.Time) *TagUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TagUpdate) SetUpdatedAt(t time.Time) *TagUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tu *TagUpdate) SetNillableUpdatedAt(t *time.Time) *TagUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TagUpdate) SetDeletedAt(t time.Time) *TagUpdate {
	tu.mutation.SetDeletedAt(t)
	return tu
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: tag.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldName,
		})
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldStatus,
		})
	}
	if value, ok := tu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldStatus,
		})
	}
	if value, ok := tu.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldCreatedBy,
		})
	}
	if value, ok := tu.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldUpdatedBy,
		})
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TagUpdateOne) SetStatus(i int) *TagUpdateOne {
	tuo.mutation.ResetStatus()
	tuo.mutation.SetStatus(i)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableStatus(i *int) *TagUpdateOne {
	if i != nil {
		tuo.SetStatus(*i)
	}
	return tuo
}

// AddStatus adds i to the "status" field.
func (tuo *TagUpdateOne) AddStatus(i int) *TagUpdateOne {
	tuo.mutation.AddStatus(i)
	return tuo
}

// SetCreatedBy sets the "created_by" field.
func (tuo *TagUpdateOne) SetCreatedBy(s string) *TagUpdateOne {
	tuo.mutation.SetCreatedBy(s)
	return tuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tuo *TagUpdateOne) SetUpdatedBy(s string) *TagUpdateOne {
	tuo.mutation.SetUpdatedBy(s)
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TagUpdateOne) SetCreatedAt(t time.Time) *TagUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableCreatedAt(t *time.Time) *TagUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TagUpdateOne) SetUpdatedAt(t time.Time) *TagUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableUpdatedAt(t *time.Time) *TagUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TagUpdateOne) SetDeletedAt(t time.Time) *TagUpdateOne {
	tuo.mutation.SetDeletedAt(t)
	return tuo
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	var (
		err  error
		node *Tag
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: tag.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Tag.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldName,
		})
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldStatus,
		})
	}
	if value, ok := tuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldStatus,
		})
	}
	if value, ok := tuo.mutation.CreatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldCreatedBy,
		})
	}
	if value, ok := tuo.mutation.UpdatedBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldUpdatedBy,
		})
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldDeletedAt,
		})
	}
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}