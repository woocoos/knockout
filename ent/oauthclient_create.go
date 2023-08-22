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
	"github.com/woocoos/knockout/ent/oauthclient"
	"github.com/woocoos/knockout/ent/user"
)

// OauthClientCreate is the builder for creating a OauthClient entity.
type OauthClientCreate struct {
	config
	mutation *OauthClientMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (occ *OauthClientCreate) SetCreatedBy(i int) *OauthClientCreate {
	occ.mutation.SetCreatedBy(i)
	return occ
}

// SetCreatedAt sets the "created_at" field.
func (occ *OauthClientCreate) SetCreatedAt(t time.Time) *OauthClientCreate {
	occ.mutation.SetCreatedAt(t)
	return occ
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (occ *OauthClientCreate) SetNillableCreatedAt(t *time.Time) *OauthClientCreate {
	if t != nil {
		occ.SetCreatedAt(*t)
	}
	return occ
}

// SetUpdatedBy sets the "updated_by" field.
func (occ *OauthClientCreate) SetUpdatedBy(i int) *OauthClientCreate {
	occ.mutation.SetUpdatedBy(i)
	return occ
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (occ *OauthClientCreate) SetNillableUpdatedBy(i *int) *OauthClientCreate {
	if i != nil {
		occ.SetUpdatedBy(*i)
	}
	return occ
}

// SetUpdatedAt sets the "updated_at" field.
func (occ *OauthClientCreate) SetUpdatedAt(t time.Time) *OauthClientCreate {
	occ.mutation.SetUpdatedAt(t)
	return occ
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (occ *OauthClientCreate) SetNillableUpdatedAt(t *time.Time) *OauthClientCreate {
	if t != nil {
		occ.SetUpdatedAt(*t)
	}
	return occ
}

// SetName sets the "name" field.
func (occ *OauthClientCreate) SetName(s string) *OauthClientCreate {
	occ.mutation.SetName(s)
	return occ
}

// SetClientID sets the "client_id" field.
func (occ *OauthClientCreate) SetClientID(s string) *OauthClientCreate {
	occ.mutation.SetClientID(s)
	return occ
}

// SetClientSecret sets the "client_secret" field.
func (occ *OauthClientCreate) SetClientSecret(s string) *OauthClientCreate {
	occ.mutation.SetClientSecret(s)
	return occ
}

// SetGrantTypes sets the "grant_types" field.
func (occ *OauthClientCreate) SetGrantTypes(ot oauthclient.GrantTypes) *OauthClientCreate {
	occ.mutation.SetGrantTypes(ot)
	return occ
}

// SetUserID sets the "user_id" field.
func (occ *OauthClientCreate) SetUserID(i int) *OauthClientCreate {
	occ.mutation.SetUserID(i)
	return occ
}

// SetLastAuthAt sets the "last_auth_at" field.
func (occ *OauthClientCreate) SetLastAuthAt(t time.Time) *OauthClientCreate {
	occ.mutation.SetLastAuthAt(t)
	return occ
}

// SetNillableLastAuthAt sets the "last_auth_at" field if the given value is not nil.
func (occ *OauthClientCreate) SetNillableLastAuthAt(t *time.Time) *OauthClientCreate {
	if t != nil {
		occ.SetLastAuthAt(*t)
	}
	return occ
}

// SetStatus sets the "status" field.
func (occ *OauthClientCreate) SetStatus(ts typex.SimpleStatus) *OauthClientCreate {
	occ.mutation.SetStatus(ts)
	return occ
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (occ *OauthClientCreate) SetNillableStatus(ts *typex.SimpleStatus) *OauthClientCreate {
	if ts != nil {
		occ.SetStatus(*ts)
	}
	return occ
}

// SetID sets the "id" field.
func (occ *OauthClientCreate) SetID(i int) *OauthClientCreate {
	occ.mutation.SetID(i)
	return occ
}

// SetNillableID sets the "id" field if the given value is not nil.
func (occ *OauthClientCreate) SetNillableID(i *int) *OauthClientCreate {
	if i != nil {
		occ.SetID(*i)
	}
	return occ
}

// SetUser sets the "user" edge to the User entity.
func (occ *OauthClientCreate) SetUser(u *User) *OauthClientCreate {
	return occ.SetUserID(u.ID)
}

// Mutation returns the OauthClientMutation object of the builder.
func (occ *OauthClientCreate) Mutation() *OauthClientMutation {
	return occ.mutation
}

// Save creates the OauthClient in the database.
func (occ *OauthClientCreate) Save(ctx context.Context) (*OauthClient, error) {
	if err := occ.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, occ.sqlSave, occ.mutation, occ.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (occ *OauthClientCreate) SaveX(ctx context.Context) *OauthClient {
	v, err := occ.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occ *OauthClientCreate) Exec(ctx context.Context) error {
	_, err := occ.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occ *OauthClientCreate) ExecX(ctx context.Context) {
	if err := occ.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (occ *OauthClientCreate) defaults() error {
	if _, ok := occ.mutation.CreatedAt(); !ok {
		if oauthclient.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized oauthclient.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := oauthclient.DefaultCreatedAt()
		occ.mutation.SetCreatedAt(v)
	}
	if _, ok := occ.mutation.Status(); !ok {
		v := oauthclient.DefaultStatus
		occ.mutation.SetStatus(v)
	}
	if _, ok := occ.mutation.ID(); !ok {
		if oauthclient.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized oauthclient.DefaultID (forgotten import ent/runtime?)")
		}
		v := oauthclient.DefaultID()
		occ.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (occ *OauthClientCreate) check() error {
	if _, ok := occ.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "OauthClient.created_by"`)}
	}
	if _, ok := occ.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OauthClient.created_at"`)}
	}
	if _, ok := occ.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OauthClient.name"`)}
	}
	if v, ok := occ.mutation.Name(); ok {
		if err := oauthclient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OauthClient.name": %w`, err)}
		}
	}
	if _, ok := occ.mutation.ClientID(); !ok {
		return &ValidationError{Name: "client_id", err: errors.New(`ent: missing required field "OauthClient.client_id"`)}
	}
	if v, ok := occ.mutation.ClientID(); ok {
		if err := oauthclient.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`ent: validator failed for field "OauthClient.client_id": %w`, err)}
		}
	}
	if _, ok := occ.mutation.ClientSecret(); !ok {
		return &ValidationError{Name: "client_secret", err: errors.New(`ent: missing required field "OauthClient.client_secret"`)}
	}
	if v, ok := occ.mutation.ClientSecret(); ok {
		if err := oauthclient.ClientSecretValidator(v); err != nil {
			return &ValidationError{Name: "client_secret", err: fmt.Errorf(`ent: validator failed for field "OauthClient.client_secret": %w`, err)}
		}
	}
	if _, ok := occ.mutation.GrantTypes(); !ok {
		return &ValidationError{Name: "grant_types", err: errors.New(`ent: missing required field "OauthClient.grant_types"`)}
	}
	if v, ok := occ.mutation.GrantTypes(); ok {
		if err := oauthclient.GrantTypesValidator(v); err != nil {
			return &ValidationError{Name: "grant_types", err: fmt.Errorf(`ent: validator failed for field "OauthClient.grant_types": %w`, err)}
		}
	}
	if _, ok := occ.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "OauthClient.user_id"`)}
	}
	if _, ok := occ.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "OauthClient.status"`)}
	}
	if v, ok := occ.mutation.Status(); ok {
		if err := oauthclient.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OauthClient.status": %w`, err)}
		}
	}
	if _, ok := occ.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "OauthClient.user"`)}
	}
	return nil
}

func (occ *OauthClientCreate) sqlSave(ctx context.Context) (*OauthClient, error) {
	if err := occ.check(); err != nil {
		return nil, err
	}
	_node, _spec := occ.createSpec()
	if err := sqlgraph.CreateNode(ctx, occ.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	occ.mutation.id = &_node.ID
	occ.mutation.done = true
	return _node, nil
}

func (occ *OauthClientCreate) createSpec() (*OauthClient, *sqlgraph.CreateSpec) {
	var (
		_node = &OauthClient{config: occ.config}
		_spec = sqlgraph.NewCreateSpec(oauthclient.Table, sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeInt))
	)
	_spec.OnConflict = occ.conflict
	if id, ok := occ.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := occ.mutation.CreatedBy(); ok {
		_spec.SetField(oauthclient.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := occ.mutation.CreatedAt(); ok {
		_spec.SetField(oauthclient.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := occ.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthclient.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := occ.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthclient.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := occ.mutation.Name(); ok {
		_spec.SetField(oauthclient.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := occ.mutation.ClientID(); ok {
		_spec.SetField(oauthclient.FieldClientID, field.TypeString, value)
		_node.ClientID = value
	}
	if value, ok := occ.mutation.ClientSecret(); ok {
		_spec.SetField(oauthclient.FieldClientSecret, field.TypeString, value)
		_node.ClientSecret = value
	}
	if value, ok := occ.mutation.GrantTypes(); ok {
		_spec.SetField(oauthclient.FieldGrantTypes, field.TypeEnum, value)
		_node.GrantTypes = value
	}
	if value, ok := occ.mutation.LastAuthAt(); ok {
		_spec.SetField(oauthclient.FieldLastAuthAt, field.TypeTime, value)
		_node.LastAuthAt = value
	}
	if value, ok := occ.mutation.Status(); ok {
		_spec.SetField(oauthclient.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := occ.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthclient.UserTable,
			Columns: []string{oauthclient.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OauthClient.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OauthClientUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (occ *OauthClientCreate) OnConflict(opts ...sql.ConflictOption) *OauthClientUpsertOne {
	occ.conflict = opts
	return &OauthClientUpsertOne{
		create: occ,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OauthClient.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (occ *OauthClientCreate) OnConflictColumns(columns ...string) *OauthClientUpsertOne {
	occ.conflict = append(occ.conflict, sql.ConflictColumns(columns...))
	return &OauthClientUpsertOne{
		create: occ,
	}
}

type (
	// OauthClientUpsertOne is the builder for "upsert"-ing
	//  one OauthClient node.
	OauthClientUpsertOne struct {
		create *OauthClientCreate
	}

	// OauthClientUpsert is the "OnConflict" setter.
	OauthClientUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *OauthClientUpsert) SetUpdatedBy(v int) *OauthClientUpsert {
	u.Set(oauthclient.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OauthClientUpsert) UpdateUpdatedBy() *OauthClientUpsert {
	u.SetExcluded(oauthclient.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OauthClientUpsert) AddUpdatedBy(v int) *OauthClientUpsert {
	u.Add(oauthclient.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OauthClientUpsert) ClearUpdatedBy() *OauthClientUpsert {
	u.SetNull(oauthclient.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OauthClientUpsert) SetUpdatedAt(v time.Time) *OauthClientUpsert {
	u.Set(oauthclient.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OauthClientUpsert) UpdateUpdatedAt() *OauthClientUpsert {
	u.SetExcluded(oauthclient.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OauthClientUpsert) ClearUpdatedAt() *OauthClientUpsert {
	u.SetNull(oauthclient.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *OauthClientUpsert) SetName(v string) *OauthClientUpsert {
	u.Set(oauthclient.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OauthClientUpsert) UpdateName() *OauthClientUpsert {
	u.SetExcluded(oauthclient.FieldName)
	return u
}

// SetGrantTypes sets the "grant_types" field.
func (u *OauthClientUpsert) SetGrantTypes(v oauthclient.GrantTypes) *OauthClientUpsert {
	u.Set(oauthclient.FieldGrantTypes, v)
	return u
}

// UpdateGrantTypes sets the "grant_types" field to the value that was provided on create.
func (u *OauthClientUpsert) UpdateGrantTypes() *OauthClientUpsert {
	u.SetExcluded(oauthclient.FieldGrantTypes)
	return u
}

// SetUserID sets the "user_id" field.
func (u *OauthClientUpsert) SetUserID(v int) *OauthClientUpsert {
	u.Set(oauthclient.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OauthClientUpsert) UpdateUserID() *OauthClientUpsert {
	u.SetExcluded(oauthclient.FieldUserID)
	return u
}

// SetLastAuthAt sets the "last_auth_at" field.
func (u *OauthClientUpsert) SetLastAuthAt(v time.Time) *OauthClientUpsert {
	u.Set(oauthclient.FieldLastAuthAt, v)
	return u
}

// UpdateLastAuthAt sets the "last_auth_at" field to the value that was provided on create.
func (u *OauthClientUpsert) UpdateLastAuthAt() *OauthClientUpsert {
	u.SetExcluded(oauthclient.FieldLastAuthAt)
	return u
}

// ClearLastAuthAt clears the value of the "last_auth_at" field.
func (u *OauthClientUpsert) ClearLastAuthAt() *OauthClientUpsert {
	u.SetNull(oauthclient.FieldLastAuthAt)
	return u
}

// SetStatus sets the "status" field.
func (u *OauthClientUpsert) SetStatus(v typex.SimpleStatus) *OauthClientUpsert {
	u.Set(oauthclient.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OauthClientUpsert) UpdateStatus() *OauthClientUpsert {
	u.SetExcluded(oauthclient.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OauthClient.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(oauthclient.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OauthClientUpsertOne) UpdateNewValues() *OauthClientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(oauthclient.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(oauthclient.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(oauthclient.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.ClientID(); exists {
			s.SetIgnore(oauthclient.FieldClientID)
		}
		if _, exists := u.create.mutation.ClientSecret(); exists {
			s.SetIgnore(oauthclient.FieldClientSecret)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OauthClient.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OauthClientUpsertOne) Ignore() *OauthClientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OauthClientUpsertOne) DoNothing() *OauthClientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OauthClientCreate.OnConflict
// documentation for more info.
func (u *OauthClientUpsertOne) Update(set func(*OauthClientUpsert)) *OauthClientUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OauthClientUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OauthClientUpsertOne) SetUpdatedBy(v int) *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OauthClientUpsertOne) AddUpdatedBy(v int) *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OauthClientUpsertOne) UpdateUpdatedBy() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OauthClientUpsertOne) ClearUpdatedBy() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OauthClientUpsertOne) SetUpdatedAt(v time.Time) *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OauthClientUpsertOne) UpdateUpdatedAt() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OauthClientUpsertOne) ClearUpdatedAt() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *OauthClientUpsertOne) SetName(v string) *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OauthClientUpsertOne) UpdateName() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateName()
	})
}

// SetGrantTypes sets the "grant_types" field.
func (u *OauthClientUpsertOne) SetGrantTypes(v oauthclient.GrantTypes) *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetGrantTypes(v)
	})
}

// UpdateGrantTypes sets the "grant_types" field to the value that was provided on create.
func (u *OauthClientUpsertOne) UpdateGrantTypes() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateGrantTypes()
	})
}

// SetUserID sets the "user_id" field.
func (u *OauthClientUpsertOne) SetUserID(v int) *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OauthClientUpsertOne) UpdateUserID() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateUserID()
	})
}

// SetLastAuthAt sets the "last_auth_at" field.
func (u *OauthClientUpsertOne) SetLastAuthAt(v time.Time) *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetLastAuthAt(v)
	})
}

// UpdateLastAuthAt sets the "last_auth_at" field to the value that was provided on create.
func (u *OauthClientUpsertOne) UpdateLastAuthAt() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateLastAuthAt()
	})
}

// ClearLastAuthAt clears the value of the "last_auth_at" field.
func (u *OauthClientUpsertOne) ClearLastAuthAt() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.ClearLastAuthAt()
	})
}

// SetStatus sets the "status" field.
func (u *OauthClientUpsertOne) SetStatus(v typex.SimpleStatus) *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OauthClientUpsertOne) UpdateStatus() *OauthClientUpsertOne {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *OauthClientUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OauthClientCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OauthClientUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OauthClientUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OauthClientUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OauthClientCreateBulk is the builder for creating many OauthClient entities in bulk.
type OauthClientCreateBulk struct {
	config
	builders []*OauthClientCreate
	conflict []sql.ConflictOption
}

// Save creates the OauthClient entities in the database.
func (occb *OauthClientCreateBulk) Save(ctx context.Context) ([]*OauthClient, error) {
	specs := make([]*sqlgraph.CreateSpec, len(occb.builders))
	nodes := make([]*OauthClient, len(occb.builders))
	mutators := make([]Mutator, len(occb.builders))
	for i := range occb.builders {
		func(i int, root context.Context) {
			builder := occb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OauthClientMutation)
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
					_, err = mutators[i+1].Mutate(root, occb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = occb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, occb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, occb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (occb *OauthClientCreateBulk) SaveX(ctx context.Context) []*OauthClient {
	v, err := occb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occb *OauthClientCreateBulk) Exec(ctx context.Context) error {
	_, err := occb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occb *OauthClientCreateBulk) ExecX(ctx context.Context) {
	if err := occb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OauthClient.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OauthClientUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (occb *OauthClientCreateBulk) OnConflict(opts ...sql.ConflictOption) *OauthClientUpsertBulk {
	occb.conflict = opts
	return &OauthClientUpsertBulk{
		create: occb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OauthClient.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (occb *OauthClientCreateBulk) OnConflictColumns(columns ...string) *OauthClientUpsertBulk {
	occb.conflict = append(occb.conflict, sql.ConflictColumns(columns...))
	return &OauthClientUpsertBulk{
		create: occb,
	}
}

// OauthClientUpsertBulk is the builder for "upsert"-ing
// a bulk of OauthClient nodes.
type OauthClientUpsertBulk struct {
	create *OauthClientCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OauthClient.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(oauthclient.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OauthClientUpsertBulk) UpdateNewValues() *OauthClientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(oauthclient.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(oauthclient.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(oauthclient.FieldCreatedAt)
			}
			if _, exists := b.mutation.ClientID(); exists {
				s.SetIgnore(oauthclient.FieldClientID)
			}
			if _, exists := b.mutation.ClientSecret(); exists {
				s.SetIgnore(oauthclient.FieldClientSecret)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OauthClient.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OauthClientUpsertBulk) Ignore() *OauthClientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OauthClientUpsertBulk) DoNothing() *OauthClientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OauthClientCreateBulk.OnConflict
// documentation for more info.
func (u *OauthClientUpsertBulk) Update(set func(*OauthClientUpsert)) *OauthClientUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OauthClientUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OauthClientUpsertBulk) SetUpdatedBy(v int) *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OauthClientUpsertBulk) AddUpdatedBy(v int) *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OauthClientUpsertBulk) UpdateUpdatedBy() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OauthClientUpsertBulk) ClearUpdatedBy() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OauthClientUpsertBulk) SetUpdatedAt(v time.Time) *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OauthClientUpsertBulk) UpdateUpdatedAt() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OauthClientUpsertBulk) ClearUpdatedAt() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *OauthClientUpsertBulk) SetName(v string) *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OauthClientUpsertBulk) UpdateName() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateName()
	})
}

// SetGrantTypes sets the "grant_types" field.
func (u *OauthClientUpsertBulk) SetGrantTypes(v oauthclient.GrantTypes) *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetGrantTypes(v)
	})
}

// UpdateGrantTypes sets the "grant_types" field to the value that was provided on create.
func (u *OauthClientUpsertBulk) UpdateGrantTypes() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateGrantTypes()
	})
}

// SetUserID sets the "user_id" field.
func (u *OauthClientUpsertBulk) SetUserID(v int) *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OauthClientUpsertBulk) UpdateUserID() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateUserID()
	})
}

// SetLastAuthAt sets the "last_auth_at" field.
func (u *OauthClientUpsertBulk) SetLastAuthAt(v time.Time) *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetLastAuthAt(v)
	})
}

// UpdateLastAuthAt sets the "last_auth_at" field to the value that was provided on create.
func (u *OauthClientUpsertBulk) UpdateLastAuthAt() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateLastAuthAt()
	})
}

// ClearLastAuthAt clears the value of the "last_auth_at" field.
func (u *OauthClientUpsertBulk) ClearLastAuthAt() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.ClearLastAuthAt()
	})
}

// SetStatus sets the "status" field.
func (u *OauthClientUpsertBulk) SetStatus(v typex.SimpleStatus) *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OauthClientUpsertBulk) UpdateStatus() *OauthClientUpsertBulk {
	return u.Update(func(s *OauthClientUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *OauthClientUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OauthClientCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OauthClientCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OauthClientUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}