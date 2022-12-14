// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.42paris.fr/utilities/notion_service/ent/predicate"
	"gitlab.42paris.fr/utilities/notion_service/ent/status"
)

// StatusUpdate is the builder for updating Status entities.
type StatusUpdate struct {
	config
	hooks    []Hook
	mutation *StatusMutation
}

// Where appends a list predicates to the StatusUpdate builder.
func (su *StatusUpdate) Where(ps ...predicate.Status) *StatusUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetDbID sets the "db_id" field.
func (su *StatusUpdate) SetDbID(s string) *StatusUpdate {
	su.mutation.SetDbID(s)
	return su
}

// SetRowID sets the "row_id" field.
func (su *StatusUpdate) SetRowID(s string) *StatusUpdate {
	su.mutation.SetRowID(s)
	return su
}

// SetStatus sets the "status" field.
func (su *StatusUpdate) SetStatus(s string) *StatusUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetSavedAt sets the "saved_at" field.
func (su *StatusUpdate) SetSavedAt(t time.Time) *StatusUpdate {
	su.mutation.SetSavedAt(t)
	return su
}

// SetNillableSavedAt sets the "saved_at" field if the given value is not nil.
func (su *StatusUpdate) SetNillableSavedAt(t *time.Time) *StatusUpdate {
	if t != nil {
		su.SetSavedAt(*t)
	}
	return su
}

// Mutation returns the StatusMutation object of the builder.
func (su *StatusUpdate) Mutation() *StatusMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StatusUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StatusUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StatusUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StatusUpdate) check() error {
	if v, ok := su.mutation.DbID(); ok {
		if err := status.DbIDValidator(v); err != nil {
			return &ValidationError{Name: "db_id", err: fmt.Errorf(`ent: validator failed for field "Status.db_id": %w`, err)}
		}
	}
	if v, ok := su.mutation.RowID(); ok {
		if err := status.RowIDValidator(v); err != nil {
			return &ValidationError{Name: "row_id", err: fmt.Errorf(`ent: validator failed for field "Status.row_id": %w`, err)}
		}
	}
	if v, ok := su.mutation.Status(); ok {
		if err := status.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Status.status": %w`, err)}
		}
	}
	return nil
}

func (su *StatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   status.Table,
			Columns: status.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: status.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.DbID(); ok {
		_spec.SetField(status.FieldDbID, field.TypeString, value)
	}
	if value, ok := su.mutation.RowID(); ok {
		_spec.SetField(status.FieldRowID, field.TypeString, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(status.FieldStatus, field.TypeString, value)
	}
	if value, ok := su.mutation.SavedAt(); ok {
		_spec.SetField(status.FieldSavedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{status.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// StatusUpdateOne is the builder for updating a single Status entity.
type StatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StatusMutation
}

// SetDbID sets the "db_id" field.
func (suo *StatusUpdateOne) SetDbID(s string) *StatusUpdateOne {
	suo.mutation.SetDbID(s)
	return suo
}

// SetRowID sets the "row_id" field.
func (suo *StatusUpdateOne) SetRowID(s string) *StatusUpdateOne {
	suo.mutation.SetRowID(s)
	return suo
}

// SetStatus sets the "status" field.
func (suo *StatusUpdateOne) SetStatus(s string) *StatusUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetSavedAt sets the "saved_at" field.
func (suo *StatusUpdateOne) SetSavedAt(t time.Time) *StatusUpdateOne {
	suo.mutation.SetSavedAt(t)
	return suo
}

// SetNillableSavedAt sets the "saved_at" field if the given value is not nil.
func (suo *StatusUpdateOne) SetNillableSavedAt(t *time.Time) *StatusUpdateOne {
	if t != nil {
		suo.SetSavedAt(*t)
	}
	return suo
}

// Mutation returns the StatusMutation object of the builder.
func (suo *StatusUpdateOne) Mutation() *StatusMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StatusUpdateOne) Select(field string, fields ...string) *StatusUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Status entity.
func (suo *StatusUpdateOne) Save(ctx context.Context) (*Status, error) {
	var (
		err  error
		node *Status
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Status)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StatusMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StatusUpdateOne) SaveX(ctx context.Context) *Status {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StatusUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StatusUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StatusUpdateOne) check() error {
	if v, ok := suo.mutation.DbID(); ok {
		if err := status.DbIDValidator(v); err != nil {
			return &ValidationError{Name: "db_id", err: fmt.Errorf(`ent: validator failed for field "Status.db_id": %w`, err)}
		}
	}
	if v, ok := suo.mutation.RowID(); ok {
		if err := status.RowIDValidator(v); err != nil {
			return &ValidationError{Name: "row_id", err: fmt.Errorf(`ent: validator failed for field "Status.row_id": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Status(); ok {
		if err := status.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Status.status": %w`, err)}
		}
	}
	return nil
}

func (suo *StatusUpdateOne) sqlSave(ctx context.Context) (_node *Status, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   status.Table,
			Columns: status.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: status.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Status.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, status.FieldID)
		for _, f := range fields {
			if !status.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != status.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.DbID(); ok {
		_spec.SetField(status.FieldDbID, field.TypeString, value)
	}
	if value, ok := suo.mutation.RowID(); ok {
		_spec.SetField(status.FieldRowID, field.TypeString, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(status.FieldStatus, field.TypeString, value)
	}
	if value, ok := suo.mutation.SavedAt(); ok {
		_spec.SetField(status.FieldSavedAt, field.TypeTime, value)
	}
	_node = &Status{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{status.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
