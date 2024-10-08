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
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/ent/oauthclient"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/user"
)

// OauthClientUpdate is the builder for updating OauthClient entities.
type OauthClientUpdate struct {
	config
	hooks    []Hook
	mutation *OauthClientMutation
}

// Where appends a list predicates to the OauthClientUpdate builder.
func (ocu *OauthClientUpdate) Where(ps ...predicate.OauthClient) *OauthClientUpdate {
	ocu.mutation.Where(ps...)
	return ocu
}

// SetUpdatedBy sets the "updated_by" field.
func (ocu *OauthClientUpdate) SetUpdatedBy(i int) *OauthClientUpdate {
	ocu.mutation.ResetUpdatedBy()
	ocu.mutation.SetUpdatedBy(i)
	return ocu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ocu *OauthClientUpdate) SetNillableUpdatedBy(i *int) *OauthClientUpdate {
	if i != nil {
		ocu.SetUpdatedBy(*i)
	}
	return ocu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ocu *OauthClientUpdate) AddUpdatedBy(i int) *OauthClientUpdate {
	ocu.mutation.AddUpdatedBy(i)
	return ocu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ocu *OauthClientUpdate) ClearUpdatedBy() *OauthClientUpdate {
	ocu.mutation.ClearUpdatedBy()
	return ocu
}

// SetUpdatedAt sets the "updated_at" field.
func (ocu *OauthClientUpdate) SetUpdatedAt(t time.Time) *OauthClientUpdate {
	ocu.mutation.SetUpdatedAt(t)
	return ocu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ocu *OauthClientUpdate) SetNillableUpdatedAt(t *time.Time) *OauthClientUpdate {
	if t != nil {
		ocu.SetUpdatedAt(*t)
	}
	return ocu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ocu *OauthClientUpdate) ClearUpdatedAt() *OauthClientUpdate {
	ocu.mutation.ClearUpdatedAt()
	return ocu
}

// SetName sets the "name" field.
func (ocu *OauthClientUpdate) SetName(s string) *OauthClientUpdate {
	ocu.mutation.SetName(s)
	return ocu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ocu *OauthClientUpdate) SetNillableName(s *string) *OauthClientUpdate {
	if s != nil {
		ocu.SetName(*s)
	}
	return ocu
}

// SetGrantTypes sets the "grant_types" field.
func (ocu *OauthClientUpdate) SetGrantTypes(ot oauthclient.GrantTypes) *OauthClientUpdate {
	ocu.mutation.SetGrantTypes(ot)
	return ocu
}

// SetNillableGrantTypes sets the "grant_types" field if the given value is not nil.
func (ocu *OauthClientUpdate) SetNillableGrantTypes(ot *oauthclient.GrantTypes) *OauthClientUpdate {
	if ot != nil {
		ocu.SetGrantTypes(*ot)
	}
	return ocu
}

// SetUserID sets the "user_id" field.
func (ocu *OauthClientUpdate) SetUserID(i int) *OauthClientUpdate {
	ocu.mutation.SetUserID(i)
	return ocu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ocu *OauthClientUpdate) SetNillableUserID(i *int) *OauthClientUpdate {
	if i != nil {
		ocu.SetUserID(*i)
	}
	return ocu
}

// SetLastAuthAt sets the "last_auth_at" field.
func (ocu *OauthClientUpdate) SetLastAuthAt(t time.Time) *OauthClientUpdate {
	ocu.mutation.SetLastAuthAt(t)
	return ocu
}

// SetNillableLastAuthAt sets the "last_auth_at" field if the given value is not nil.
func (ocu *OauthClientUpdate) SetNillableLastAuthAt(t *time.Time) *OauthClientUpdate {
	if t != nil {
		ocu.SetLastAuthAt(*t)
	}
	return ocu
}

// ClearLastAuthAt clears the value of the "last_auth_at" field.
func (ocu *OauthClientUpdate) ClearLastAuthAt() *OauthClientUpdate {
	ocu.mutation.ClearLastAuthAt()
	return ocu
}

// SetStatus sets the "status" field.
func (ocu *OauthClientUpdate) SetStatus(ts typex.SimpleStatus) *OauthClientUpdate {
	ocu.mutation.SetStatus(ts)
	return ocu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ocu *OauthClientUpdate) SetNillableStatus(ts *typex.SimpleStatus) *OauthClientUpdate {
	if ts != nil {
		ocu.SetStatus(*ts)
	}
	return ocu
}

// SetUser sets the "user" edge to the User entity.
func (ocu *OauthClientUpdate) SetUser(u *User) *OauthClientUpdate {
	return ocu.SetUserID(u.ID)
}

// Mutation returns the OauthClientMutation object of the builder.
func (ocu *OauthClientUpdate) Mutation() *OauthClientMutation {
	return ocu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ocu *OauthClientUpdate) ClearUser() *OauthClientUpdate {
	ocu.mutation.ClearUser()
	return ocu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ocu *OauthClientUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ocu.sqlSave, ocu.mutation, ocu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ocu *OauthClientUpdate) SaveX(ctx context.Context) int {
	affected, err := ocu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ocu *OauthClientUpdate) Exec(ctx context.Context) error {
	_, err := ocu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocu *OauthClientUpdate) ExecX(ctx context.Context) {
	if err := ocu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ocu *OauthClientUpdate) check() error {
	if v, ok := ocu.mutation.Name(); ok {
		if err := oauthclient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OauthClient.name": %w`, err)}
		}
	}
	if v, ok := ocu.mutation.GrantTypes(); ok {
		if err := oauthclient.GrantTypesValidator(v); err != nil {
			return &ValidationError{Name: "grant_types", err: fmt.Errorf(`ent: validator failed for field "OauthClient.grant_types": %w`, err)}
		}
	}
	if v, ok := ocu.mutation.Status(); ok {
		if err := oauthclient.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OauthClient.status": %w`, err)}
		}
	}
	if ocu.mutation.UserCleared() && len(ocu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OauthClient.user"`)
	}
	return nil
}

func (ocu *OauthClientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ocu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauthclient.Table, oauthclient.Columns, sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeInt))
	if ps := ocu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocu.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthclient.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ocu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(oauthclient.FieldUpdatedBy, field.TypeInt, value)
	}
	if ocu.mutation.UpdatedByCleared() {
		_spec.ClearField(oauthclient.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := ocu.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthclient.FieldUpdatedAt, field.TypeTime, value)
	}
	if ocu.mutation.UpdatedAtCleared() {
		_spec.ClearField(oauthclient.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ocu.mutation.Name(); ok {
		_spec.SetField(oauthclient.FieldName, field.TypeString, value)
	}
	if value, ok := ocu.mutation.GrantTypes(); ok {
		_spec.SetField(oauthclient.FieldGrantTypes, field.TypeEnum, value)
	}
	if value, ok := ocu.mutation.LastAuthAt(); ok {
		_spec.SetField(oauthclient.FieldLastAuthAt, field.TypeTime, value)
	}
	if ocu.mutation.LastAuthAtCleared() {
		_spec.ClearField(oauthclient.FieldLastAuthAt, field.TypeTime)
	}
	if value, ok := ocu.mutation.Status(); ok {
		_spec.SetField(oauthclient.FieldStatus, field.TypeEnum, value)
	}
	if ocu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ocu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthclient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ocu.mutation.done = true
	return n, nil
}

// OauthClientUpdateOne is the builder for updating a single OauthClient entity.
type OauthClientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OauthClientMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (ocuo *OauthClientUpdateOne) SetUpdatedBy(i int) *OauthClientUpdateOne {
	ocuo.mutation.ResetUpdatedBy()
	ocuo.mutation.SetUpdatedBy(i)
	return ocuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ocuo *OauthClientUpdateOne) SetNillableUpdatedBy(i *int) *OauthClientUpdateOne {
	if i != nil {
		ocuo.SetUpdatedBy(*i)
	}
	return ocuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ocuo *OauthClientUpdateOne) AddUpdatedBy(i int) *OauthClientUpdateOne {
	ocuo.mutation.AddUpdatedBy(i)
	return ocuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ocuo *OauthClientUpdateOne) ClearUpdatedBy() *OauthClientUpdateOne {
	ocuo.mutation.ClearUpdatedBy()
	return ocuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ocuo *OauthClientUpdateOne) SetUpdatedAt(t time.Time) *OauthClientUpdateOne {
	ocuo.mutation.SetUpdatedAt(t)
	return ocuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ocuo *OauthClientUpdateOne) SetNillableUpdatedAt(t *time.Time) *OauthClientUpdateOne {
	if t != nil {
		ocuo.SetUpdatedAt(*t)
	}
	return ocuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ocuo *OauthClientUpdateOne) ClearUpdatedAt() *OauthClientUpdateOne {
	ocuo.mutation.ClearUpdatedAt()
	return ocuo
}

// SetName sets the "name" field.
func (ocuo *OauthClientUpdateOne) SetName(s string) *OauthClientUpdateOne {
	ocuo.mutation.SetName(s)
	return ocuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ocuo *OauthClientUpdateOne) SetNillableName(s *string) *OauthClientUpdateOne {
	if s != nil {
		ocuo.SetName(*s)
	}
	return ocuo
}

// SetGrantTypes sets the "grant_types" field.
func (ocuo *OauthClientUpdateOne) SetGrantTypes(ot oauthclient.GrantTypes) *OauthClientUpdateOne {
	ocuo.mutation.SetGrantTypes(ot)
	return ocuo
}

// SetNillableGrantTypes sets the "grant_types" field if the given value is not nil.
func (ocuo *OauthClientUpdateOne) SetNillableGrantTypes(ot *oauthclient.GrantTypes) *OauthClientUpdateOne {
	if ot != nil {
		ocuo.SetGrantTypes(*ot)
	}
	return ocuo
}

// SetUserID sets the "user_id" field.
func (ocuo *OauthClientUpdateOne) SetUserID(i int) *OauthClientUpdateOne {
	ocuo.mutation.SetUserID(i)
	return ocuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ocuo *OauthClientUpdateOne) SetNillableUserID(i *int) *OauthClientUpdateOne {
	if i != nil {
		ocuo.SetUserID(*i)
	}
	return ocuo
}

// SetLastAuthAt sets the "last_auth_at" field.
func (ocuo *OauthClientUpdateOne) SetLastAuthAt(t time.Time) *OauthClientUpdateOne {
	ocuo.mutation.SetLastAuthAt(t)
	return ocuo
}

// SetNillableLastAuthAt sets the "last_auth_at" field if the given value is not nil.
func (ocuo *OauthClientUpdateOne) SetNillableLastAuthAt(t *time.Time) *OauthClientUpdateOne {
	if t != nil {
		ocuo.SetLastAuthAt(*t)
	}
	return ocuo
}

// ClearLastAuthAt clears the value of the "last_auth_at" field.
func (ocuo *OauthClientUpdateOne) ClearLastAuthAt() *OauthClientUpdateOne {
	ocuo.mutation.ClearLastAuthAt()
	return ocuo
}

// SetStatus sets the "status" field.
func (ocuo *OauthClientUpdateOne) SetStatus(ts typex.SimpleStatus) *OauthClientUpdateOne {
	ocuo.mutation.SetStatus(ts)
	return ocuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ocuo *OauthClientUpdateOne) SetNillableStatus(ts *typex.SimpleStatus) *OauthClientUpdateOne {
	if ts != nil {
		ocuo.SetStatus(*ts)
	}
	return ocuo
}

// SetUser sets the "user" edge to the User entity.
func (ocuo *OauthClientUpdateOne) SetUser(u *User) *OauthClientUpdateOne {
	return ocuo.SetUserID(u.ID)
}

// Mutation returns the OauthClientMutation object of the builder.
func (ocuo *OauthClientUpdateOne) Mutation() *OauthClientMutation {
	return ocuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ocuo *OauthClientUpdateOne) ClearUser() *OauthClientUpdateOne {
	ocuo.mutation.ClearUser()
	return ocuo
}

// Where appends a list predicates to the OauthClientUpdate builder.
func (ocuo *OauthClientUpdateOne) Where(ps ...predicate.OauthClient) *OauthClientUpdateOne {
	ocuo.mutation.Where(ps...)
	return ocuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ocuo *OauthClientUpdateOne) Select(field string, fields ...string) *OauthClientUpdateOne {
	ocuo.fields = append([]string{field}, fields...)
	return ocuo
}

// Save executes the query and returns the updated OauthClient entity.
func (ocuo *OauthClientUpdateOne) Save(ctx context.Context) (*OauthClient, error) {
	return withHooks(ctx, ocuo.sqlSave, ocuo.mutation, ocuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ocuo *OauthClientUpdateOne) SaveX(ctx context.Context) *OauthClient {
	node, err := ocuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ocuo *OauthClientUpdateOne) Exec(ctx context.Context) error {
	_, err := ocuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocuo *OauthClientUpdateOne) ExecX(ctx context.Context) {
	if err := ocuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ocuo *OauthClientUpdateOne) check() error {
	if v, ok := ocuo.mutation.Name(); ok {
		if err := oauthclient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OauthClient.name": %w`, err)}
		}
	}
	if v, ok := ocuo.mutation.GrantTypes(); ok {
		if err := oauthclient.GrantTypesValidator(v); err != nil {
			return &ValidationError{Name: "grant_types", err: fmt.Errorf(`ent: validator failed for field "OauthClient.grant_types": %w`, err)}
		}
	}
	if v, ok := ocuo.mutation.Status(); ok {
		if err := oauthclient.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OauthClient.status": %w`, err)}
		}
	}
	if ocuo.mutation.UserCleared() && len(ocuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OauthClient.user"`)
	}
	return nil
}

func (ocuo *OauthClientUpdateOne) sqlSave(ctx context.Context) (_node *OauthClient, err error) {
	if err := ocuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(oauthclient.Table, oauthclient.Columns, sqlgraph.NewFieldSpec(oauthclient.FieldID, field.TypeInt))
	id, ok := ocuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OauthClient.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ocuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauthclient.FieldID)
		for _, f := range fields {
			if !oauthclient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != oauthclient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ocuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocuo.mutation.UpdatedBy(); ok {
		_spec.SetField(oauthclient.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ocuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(oauthclient.FieldUpdatedBy, field.TypeInt, value)
	}
	if ocuo.mutation.UpdatedByCleared() {
		_spec.ClearField(oauthclient.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := ocuo.mutation.UpdatedAt(); ok {
		_spec.SetField(oauthclient.FieldUpdatedAt, field.TypeTime, value)
	}
	if ocuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(oauthclient.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ocuo.mutation.Name(); ok {
		_spec.SetField(oauthclient.FieldName, field.TypeString, value)
	}
	if value, ok := ocuo.mutation.GrantTypes(); ok {
		_spec.SetField(oauthclient.FieldGrantTypes, field.TypeEnum, value)
	}
	if value, ok := ocuo.mutation.LastAuthAt(); ok {
		_spec.SetField(oauthclient.FieldLastAuthAt, field.TypeTime, value)
	}
	if ocuo.mutation.LastAuthAtCleared() {
		_spec.ClearField(oauthclient.FieldLastAuthAt, field.TypeTime)
	}
	if value, ok := ocuo.mutation.Status(); ok {
		_spec.SetField(oauthclient.FieldStatus, field.TypeEnum, value)
	}
	if ocuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OauthClient{config: ocuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ocuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauthclient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ocuo.mutation.done = true
	return _node, nil
}
