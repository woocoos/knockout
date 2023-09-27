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
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent/appdict"
	"github.com/woocoos/knockout/ent/appdictitem"
)

// AppDictItemCreate is the builder for creating a AppDictItem entity.
type AppDictItemCreate struct {
	config
	mutation *AppDictItemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (adic *AppDictItemCreate) SetCreatedBy(i int) *AppDictItemCreate {
	adic.mutation.SetCreatedBy(i)
	return adic
}

// SetCreatedAt sets the "created_at" field.
func (adic *AppDictItemCreate) SetCreatedAt(t time.Time) *AppDictItemCreate {
	adic.mutation.SetCreatedAt(t)
	return adic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (adic *AppDictItemCreate) SetNillableCreatedAt(t *time.Time) *AppDictItemCreate {
	if t != nil {
		adic.SetCreatedAt(*t)
	}
	return adic
}

// SetUpdatedBy sets the "updated_by" field.
func (adic *AppDictItemCreate) SetUpdatedBy(i int) *AppDictItemCreate {
	adic.mutation.SetUpdatedBy(i)
	return adic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (adic *AppDictItemCreate) SetNillableUpdatedBy(i *int) *AppDictItemCreate {
	if i != nil {
		adic.SetUpdatedBy(*i)
	}
	return adic
}

// SetUpdatedAt sets the "updated_at" field.
func (adic *AppDictItemCreate) SetUpdatedAt(t time.Time) *AppDictItemCreate {
	adic.mutation.SetUpdatedAt(t)
	return adic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (adic *AppDictItemCreate) SetNillableUpdatedAt(t *time.Time) *AppDictItemCreate {
	if t != nil {
		adic.SetUpdatedAt(*t)
	}
	return adic
}

// SetOrgID sets the "org_id" field.
func (adic *AppDictItemCreate) SetOrgID(i int) *AppDictItemCreate {
	adic.mutation.SetOrgID(i)
	return adic
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (adic *AppDictItemCreate) SetNillableOrgID(i *int) *AppDictItemCreate {
	if i != nil {
		adic.SetOrgID(*i)
	}
	return adic
}

// SetDictID sets the "dict_id" field.
func (adic *AppDictItemCreate) SetDictID(i int) *AppDictItemCreate {
	adic.mutation.SetDictID(i)
	return adic
}

// SetNillableDictID sets the "dict_id" field if the given value is not nil.
func (adic *AppDictItemCreate) SetNillableDictID(i *int) *AppDictItemCreate {
	if i != nil {
		adic.SetDictID(*i)
	}
	return adic
}

// SetRefCode sets the "ref_code" field.
func (adic *AppDictItemCreate) SetRefCode(s string) *AppDictItemCreate {
	adic.mutation.SetRefCode(s)
	return adic
}

// SetCode sets the "code" field.
func (adic *AppDictItemCreate) SetCode(s string) *AppDictItemCreate {
	adic.mutation.SetCode(s)
	return adic
}

// SetName sets the "name" field.
func (adic *AppDictItemCreate) SetName(s string) *AppDictItemCreate {
	adic.mutation.SetName(s)
	return adic
}

// SetComments sets the "comments" field.
func (adic *AppDictItemCreate) SetComments(s string) *AppDictItemCreate {
	adic.mutation.SetComments(s)
	return adic
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (adic *AppDictItemCreate) SetNillableComments(s *string) *AppDictItemCreate {
	if s != nil {
		adic.SetComments(*s)
	}
	return adic
}

// SetDisplaySort sets the "display_sort" field.
func (adic *AppDictItemCreate) SetDisplaySort(i int32) *AppDictItemCreate {
	adic.mutation.SetDisplaySort(i)
	return adic
}

// SetNillableDisplaySort sets the "display_sort" field if the given value is not nil.
func (adic *AppDictItemCreate) SetNillableDisplaySort(i *int32) *AppDictItemCreate {
	if i != nil {
		adic.SetDisplaySort(*i)
	}
	return adic
}

// SetStatus sets the "status" field.
func (adic *AppDictItemCreate) SetStatus(ts typex.SimpleStatus) *AppDictItemCreate {
	adic.mutation.SetStatus(ts)
	return adic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (adic *AppDictItemCreate) SetNillableStatus(ts *typex.SimpleStatus) *AppDictItemCreate {
	if ts != nil {
		adic.SetStatus(*ts)
	}
	return adic
}

// SetID sets the "id" field.
func (adic *AppDictItemCreate) SetID(i int) *AppDictItemCreate {
	adic.mutation.SetID(i)
	return adic
}

// SetDict sets the "dict" edge to the AppDict entity.
func (adic *AppDictItemCreate) SetDict(a *AppDict) *AppDictItemCreate {
	return adic.SetDictID(a.ID)
}

// Mutation returns the AppDictItemMutation object of the builder.
func (adic *AppDictItemCreate) Mutation() *AppDictItemMutation {
	return adic.mutation
}

// Save creates the AppDictItem in the database.
func (adic *AppDictItemCreate) Save(ctx context.Context) (*AppDictItem, error) {
	if err := adic.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, adic.sqlSave, adic.mutation, adic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (adic *AppDictItemCreate) SaveX(ctx context.Context) *AppDictItem {
	v, err := adic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (adic *AppDictItemCreate) Exec(ctx context.Context) error {
	_, err := adic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adic *AppDictItemCreate) ExecX(ctx context.Context) {
	if err := adic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (adic *AppDictItemCreate) defaults() error {
	if _, ok := adic.mutation.CreatedAt(); !ok {
		if appdictitem.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appdictitem.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appdictitem.DefaultCreatedAt()
		adic.mutation.SetCreatedAt(v)
	}
	if _, ok := adic.mutation.Status(); !ok {
		v := appdictitem.DefaultStatus
		adic.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (adic *AppDictItemCreate) check() error {
	if _, ok := adic.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "AppDictItem.created_by"`)}
	}
	if _, ok := adic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppDictItem.created_at"`)}
	}
	if _, ok := adic.mutation.RefCode(); !ok {
		return &ValidationError{Name: "ref_code", err: errors.New(`ent: missing required field "AppDictItem.ref_code"`)}
	}
	if _, ok := adic.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "AppDictItem.code"`)}
	}
	if v, ok := adic.mutation.Code(); ok {
		if err := appdictitem.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "AppDictItem.code": %w`, err)}
		}
	}
	if _, ok := adic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AppDictItem.name"`)}
	}
	if v, ok := adic.mutation.Name(); ok {
		if err := appdictitem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AppDictItem.name": %w`, err)}
		}
	}
	if v, ok := adic.mutation.Status(); ok {
		if err := appdictitem.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "AppDictItem.status": %w`, err)}
		}
	}
	return nil
}

func (adic *AppDictItemCreate) sqlSave(ctx context.Context) (*AppDictItem, error) {
	if err := adic.check(); err != nil {
		return nil, err
	}
	_node, _spec := adic.createSpec()
	if err := sqlgraph.CreateNode(ctx, adic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	adic.mutation.id = &_node.ID
	adic.mutation.done = true
	return _node, nil
}

func (adic *AppDictItemCreate) createSpec() (*AppDictItem, *sqlgraph.CreateSpec) {
	var (
		_node = &AppDictItem{config: adic.config}
		_spec = sqlgraph.NewCreateSpec(appdictitem.Table, sqlgraph.NewFieldSpec(appdictitem.FieldID, field.TypeInt))
	)
	_spec.OnConflict = adic.conflict
	if id, ok := adic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := adic.mutation.CreatedBy(); ok {
		_spec.SetField(appdictitem.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := adic.mutation.CreatedAt(); ok {
		_spec.SetField(appdictitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := adic.mutation.UpdatedBy(); ok {
		_spec.SetField(appdictitem.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := adic.mutation.UpdatedAt(); ok {
		_spec.SetField(appdictitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := adic.mutation.OrgID(); ok {
		_spec.SetField(appdictitem.FieldOrgID, field.TypeInt, value)
		_node.OrgID = value
	}
	if value, ok := adic.mutation.RefCode(); ok {
		_spec.SetField(appdictitem.FieldRefCode, field.TypeString, value)
		_node.RefCode = value
	}
	if value, ok := adic.mutation.Code(); ok {
		_spec.SetField(appdictitem.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := adic.mutation.Name(); ok {
		_spec.SetField(appdictitem.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := adic.mutation.Comments(); ok {
		_spec.SetField(appdictitem.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if value, ok := adic.mutation.DisplaySort(); ok {
		_spec.SetField(appdictitem.FieldDisplaySort, field.TypeInt32, value)
		_node.DisplaySort = value
	}
	if value, ok := adic.mutation.Status(); ok {
		_spec.SetField(appdictitem.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := adic.mutation.DictIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appdictitem.DictTable,
			Columns: []string{appdictitem.DictColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appdict.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DictID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppDictItem.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppDictItemUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (adic *AppDictItemCreate) OnConflict(opts ...sql.ConflictOption) *AppDictItemUpsertOne {
	adic.conflict = opts
	return &AppDictItemUpsertOne{
		create: adic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppDictItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (adic *AppDictItemCreate) OnConflictColumns(columns ...string) *AppDictItemUpsertOne {
	adic.conflict = append(adic.conflict, sql.ConflictColumns(columns...))
	return &AppDictItemUpsertOne{
		create: adic,
	}
}

type (
	// AppDictItemUpsertOne is the builder for "upsert"-ing
	//  one AppDictItem node.
	AppDictItemUpsertOne struct {
		create *AppDictItemCreate
	}

	// AppDictItemUpsert is the "OnConflict" setter.
	AppDictItemUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *AppDictItemUpsert) SetUpdatedBy(v int) *AppDictItemUpsert {
	u.Set(appdictitem.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppDictItemUpsert) UpdateUpdatedBy() *AppDictItemUpsert {
	u.SetExcluded(appdictitem.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppDictItemUpsert) AddUpdatedBy(v int) *AppDictItemUpsert {
	u.Add(appdictitem.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppDictItemUpsert) ClearUpdatedBy() *AppDictItemUpsert {
	u.SetNull(appdictitem.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppDictItemUpsert) SetUpdatedAt(v time.Time) *AppDictItemUpsert {
	u.Set(appdictitem.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppDictItemUpsert) UpdateUpdatedAt() *AppDictItemUpsert {
	u.SetExcluded(appdictitem.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppDictItemUpsert) ClearUpdatedAt() *AppDictItemUpsert {
	u.SetNull(appdictitem.FieldUpdatedAt)
	return u
}

// SetRefCode sets the "ref_code" field.
func (u *AppDictItemUpsert) SetRefCode(v string) *AppDictItemUpsert {
	u.Set(appdictitem.FieldRefCode, v)
	return u
}

// UpdateRefCode sets the "ref_code" field to the value that was provided on create.
func (u *AppDictItemUpsert) UpdateRefCode() *AppDictItemUpsert {
	u.SetExcluded(appdictitem.FieldRefCode)
	return u
}

// SetName sets the "name" field.
func (u *AppDictItemUpsert) SetName(v string) *AppDictItemUpsert {
	u.Set(appdictitem.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppDictItemUpsert) UpdateName() *AppDictItemUpsert {
	u.SetExcluded(appdictitem.FieldName)
	return u
}

// SetComments sets the "comments" field.
func (u *AppDictItemUpsert) SetComments(v string) *AppDictItemUpsert {
	u.Set(appdictitem.FieldComments, v)
	return u
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppDictItemUpsert) UpdateComments() *AppDictItemUpsert {
	u.SetExcluded(appdictitem.FieldComments)
	return u
}

// ClearComments clears the value of the "comments" field.
func (u *AppDictItemUpsert) ClearComments() *AppDictItemUpsert {
	u.SetNull(appdictitem.FieldComments)
	return u
}

// SetDisplaySort sets the "display_sort" field.
func (u *AppDictItemUpsert) SetDisplaySort(v int32) *AppDictItemUpsert {
	u.Set(appdictitem.FieldDisplaySort, v)
	return u
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *AppDictItemUpsert) UpdateDisplaySort() *AppDictItemUpsert {
	u.SetExcluded(appdictitem.FieldDisplaySort)
	return u
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *AppDictItemUpsert) AddDisplaySort(v int32) *AppDictItemUpsert {
	u.Add(appdictitem.FieldDisplaySort, v)
	return u
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *AppDictItemUpsert) ClearDisplaySort() *AppDictItemUpsert {
	u.SetNull(appdictitem.FieldDisplaySort)
	return u
}

// SetStatus sets the "status" field.
func (u *AppDictItemUpsert) SetStatus(v typex.SimpleStatus) *AppDictItemUpsert {
	u.Set(appdictitem.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppDictItemUpsert) UpdateStatus() *AppDictItemUpsert {
	u.SetExcluded(appdictitem.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *AppDictItemUpsert) ClearStatus() *AppDictItemUpsert {
	u.SetNull(appdictitem.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppDictItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appdictitem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppDictItemUpsertOne) UpdateNewValues() *AppDictItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appdictitem.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(appdictitem.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(appdictitem.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.OrgID(); exists {
			s.SetIgnore(appdictitem.FieldOrgID)
		}
		if _, exists := u.create.mutation.DictID(); exists {
			s.SetIgnore(appdictitem.FieldDictID)
		}
		if _, exists := u.create.mutation.Code(); exists {
			s.SetIgnore(appdictitem.FieldCode)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppDictItem.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppDictItemUpsertOne) Ignore() *AppDictItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppDictItemUpsertOne) DoNothing() *AppDictItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppDictItemCreate.OnConflict
// documentation for more info.
func (u *AppDictItemUpsertOne) Update(set func(*AppDictItemUpsert)) *AppDictItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppDictItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AppDictItemUpsertOne) SetUpdatedBy(v int) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppDictItemUpsertOne) AddUpdatedBy(v int) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppDictItemUpsertOne) UpdateUpdatedBy() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppDictItemUpsertOne) ClearUpdatedBy() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppDictItemUpsertOne) SetUpdatedAt(v time.Time) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppDictItemUpsertOne) UpdateUpdatedAt() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppDictItemUpsertOne) ClearUpdatedAt() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetRefCode sets the "ref_code" field.
func (u *AppDictItemUpsertOne) SetRefCode(v string) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetRefCode(v)
	})
}

// UpdateRefCode sets the "ref_code" field to the value that was provided on create.
func (u *AppDictItemUpsertOne) UpdateRefCode() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateRefCode()
	})
}

// SetName sets the "name" field.
func (u *AppDictItemUpsertOne) SetName(v string) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppDictItemUpsertOne) UpdateName() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateName()
	})
}

// SetComments sets the "comments" field.
func (u *AppDictItemUpsertOne) SetComments(v string) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppDictItemUpsertOne) UpdateComments() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *AppDictItemUpsertOne) ClearComments() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearComments()
	})
}

// SetDisplaySort sets the "display_sort" field.
func (u *AppDictItemUpsertOne) SetDisplaySort(v int32) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetDisplaySort(v)
	})
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *AppDictItemUpsertOne) AddDisplaySort(v int32) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.AddDisplaySort(v)
	})
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *AppDictItemUpsertOne) UpdateDisplaySort() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateDisplaySort()
	})
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *AppDictItemUpsertOne) ClearDisplaySort() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearDisplaySort()
	})
}

// SetStatus sets the "status" field.
func (u *AppDictItemUpsertOne) SetStatus(v typex.SimpleStatus) *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppDictItemUpsertOne) UpdateStatus() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *AppDictItemUpsertOne) ClearStatus() *AppDictItemUpsertOne {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *AppDictItemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppDictItemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppDictItemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppDictItemUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppDictItemUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppDictItemCreateBulk is the builder for creating many AppDictItem entities in bulk.
type AppDictItemCreateBulk struct {
	config
	err      error
	builders []*AppDictItemCreate
	conflict []sql.ConflictOption
}

// Save creates the AppDictItem entities in the database.
func (adicb *AppDictItemCreateBulk) Save(ctx context.Context) ([]*AppDictItem, error) {
	if adicb.err != nil {
		return nil, adicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(adicb.builders))
	nodes := make([]*AppDictItem, len(adicb.builders))
	mutators := make([]Mutator, len(adicb.builders))
	for i := range adicb.builders {
		func(i int, root context.Context) {
			builder := adicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppDictItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, adicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = adicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, adicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, adicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (adicb *AppDictItemCreateBulk) SaveX(ctx context.Context) []*AppDictItem {
	v, err := adicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (adicb *AppDictItemCreateBulk) Exec(ctx context.Context) error {
	_, err := adicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adicb *AppDictItemCreateBulk) ExecX(ctx context.Context) {
	if err := adicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppDictItem.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppDictItemUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (adicb *AppDictItemCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppDictItemUpsertBulk {
	adicb.conflict = opts
	return &AppDictItemUpsertBulk{
		create: adicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppDictItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (adicb *AppDictItemCreateBulk) OnConflictColumns(columns ...string) *AppDictItemUpsertBulk {
	adicb.conflict = append(adicb.conflict, sql.ConflictColumns(columns...))
	return &AppDictItemUpsertBulk{
		create: adicb,
	}
}

// AppDictItemUpsertBulk is the builder for "upsert"-ing
// a bulk of AppDictItem nodes.
type AppDictItemUpsertBulk struct {
	create *AppDictItemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppDictItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appdictitem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppDictItemUpsertBulk) UpdateNewValues() *AppDictItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appdictitem.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(appdictitem.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(appdictitem.FieldCreatedAt)
			}
			if _, exists := b.mutation.OrgID(); exists {
				s.SetIgnore(appdictitem.FieldOrgID)
			}
			if _, exists := b.mutation.DictID(); exists {
				s.SetIgnore(appdictitem.FieldDictID)
			}
			if _, exists := b.mutation.Code(); exists {
				s.SetIgnore(appdictitem.FieldCode)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppDictItem.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppDictItemUpsertBulk) Ignore() *AppDictItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppDictItemUpsertBulk) DoNothing() *AppDictItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppDictItemCreateBulk.OnConflict
// documentation for more info.
func (u *AppDictItemUpsertBulk) Update(set func(*AppDictItemUpsert)) *AppDictItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppDictItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AppDictItemUpsertBulk) SetUpdatedBy(v int) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppDictItemUpsertBulk) AddUpdatedBy(v int) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppDictItemUpsertBulk) UpdateUpdatedBy() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppDictItemUpsertBulk) ClearUpdatedBy() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppDictItemUpsertBulk) SetUpdatedAt(v time.Time) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppDictItemUpsertBulk) UpdateUpdatedAt() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppDictItemUpsertBulk) ClearUpdatedAt() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetRefCode sets the "ref_code" field.
func (u *AppDictItemUpsertBulk) SetRefCode(v string) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetRefCode(v)
	})
}

// UpdateRefCode sets the "ref_code" field to the value that was provided on create.
func (u *AppDictItemUpsertBulk) UpdateRefCode() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateRefCode()
	})
}

// SetName sets the "name" field.
func (u *AppDictItemUpsertBulk) SetName(v string) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppDictItemUpsertBulk) UpdateName() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateName()
	})
}

// SetComments sets the "comments" field.
func (u *AppDictItemUpsertBulk) SetComments(v string) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppDictItemUpsertBulk) UpdateComments() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *AppDictItemUpsertBulk) ClearComments() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearComments()
	})
}

// SetDisplaySort sets the "display_sort" field.
func (u *AppDictItemUpsertBulk) SetDisplaySort(v int32) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetDisplaySort(v)
	})
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *AppDictItemUpsertBulk) AddDisplaySort(v int32) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.AddDisplaySort(v)
	})
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *AppDictItemUpsertBulk) UpdateDisplaySort() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateDisplaySort()
	})
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *AppDictItemUpsertBulk) ClearDisplaySort() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearDisplaySort()
	})
}

// SetStatus sets the "status" field.
func (u *AppDictItemUpsertBulk) SetStatus(v typex.SimpleStatus) *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppDictItemUpsertBulk) UpdateStatus() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *AppDictItemUpsertBulk) ClearStatus() *AppDictItemUpsertBulk {
	return u.Update(func(s *AppDictItemUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *AppDictItemUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppDictItemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppDictItemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppDictItemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}