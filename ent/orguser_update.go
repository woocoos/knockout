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
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/user"
)

// OrgUserUpdate is the builder for updating OrgUser entities.
type OrgUserUpdate struct {
	config
	hooks    []Hook
	mutation *OrgUserMutation
}

// Where appends a list predicates to the OrgUserUpdate builder.
func (ouu *OrgUserUpdate) Where(ps ...predicate.OrgUser) *OrgUserUpdate {
	ouu.mutation.Where(ps...)
	return ouu
}

// SetUpdatedBy sets the "updated_by" field.
func (ouu *OrgUserUpdate) SetUpdatedBy(i int) *OrgUserUpdate {
	ouu.mutation.ResetUpdatedBy()
	ouu.mutation.SetUpdatedBy(i)
	return ouu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableUpdatedBy(i *int) *OrgUserUpdate {
	if i != nil {
		ouu.SetUpdatedBy(*i)
	}
	return ouu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ouu *OrgUserUpdate) AddUpdatedBy(i int) *OrgUserUpdate {
	ouu.mutation.AddUpdatedBy(i)
	return ouu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ouu *OrgUserUpdate) ClearUpdatedBy() *OrgUserUpdate {
	ouu.mutation.ClearUpdatedBy()
	return ouu
}

// SetUpdatedAt sets the "updated_at" field.
func (ouu *OrgUserUpdate) SetUpdatedAt(t time.Time) *OrgUserUpdate {
	ouu.mutation.SetUpdatedAt(t)
	return ouu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableUpdatedAt(t *time.Time) *OrgUserUpdate {
	if t != nil {
		ouu.SetUpdatedAt(*t)
	}
	return ouu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ouu *OrgUserUpdate) ClearUpdatedAt() *OrgUserUpdate {
	ouu.mutation.ClearUpdatedAt()
	return ouu
}

// SetOrgID sets the "org_id" field.
func (ouu *OrgUserUpdate) SetOrgID(i int) *OrgUserUpdate {
	ouu.mutation.SetOrgID(i)
	return ouu
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableOrgID(i *int) *OrgUserUpdate {
	if i != nil {
		ouu.SetOrgID(*i)
	}
	return ouu
}

// SetUserID sets the "user_id" field.
func (ouu *OrgUserUpdate) SetUserID(i int) *OrgUserUpdate {
	ouu.mutation.SetUserID(i)
	return ouu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableUserID(i *int) *OrgUserUpdate {
	if i != nil {
		ouu.SetUserID(*i)
	}
	return ouu
}

// SetJoinedAt sets the "joined_at" field.
func (ouu *OrgUserUpdate) SetJoinedAt(t time.Time) *OrgUserUpdate {
	ouu.mutation.SetJoinedAt(t)
	return ouu
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableJoinedAt(t *time.Time) *OrgUserUpdate {
	if t != nil {
		ouu.SetJoinedAt(*t)
	}
	return ouu
}

// SetDisplayName sets the "display_name" field.
func (ouu *OrgUserUpdate) SetDisplayName(s string) *OrgUserUpdate {
	ouu.mutation.SetDisplayName(s)
	return ouu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableDisplayName(s *string) *OrgUserUpdate {
	if s != nil {
		ouu.SetDisplayName(*s)
	}
	return ouu
}

// SetUserType sets the "user_type" field.
func (ouu *OrgUserUpdate) SetUserType(ot orguser.UserType) *OrgUserUpdate {
	ouu.mutation.SetUserType(ot)
	return ouu
}

// SetNillableUserType sets the "user_type" field if the given value is not nil.
func (ouu *OrgUserUpdate) SetNillableUserType(ot *orguser.UserType) *OrgUserUpdate {
	if ot != nil {
		ouu.SetUserType(*ot)
	}
	return ouu
}

// SetOrg sets the "org" edge to the Org entity.
func (ouu *OrgUserUpdate) SetOrg(o *Org) *OrgUserUpdate {
	return ouu.SetOrgID(o.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ouu *OrgUserUpdate) SetUser(u *User) *OrgUserUpdate {
	return ouu.SetUserID(u.ID)
}

// AddOrgRoleIDs adds the "org_roles" edge to the OrgRole entity by IDs.
func (ouu *OrgUserUpdate) AddOrgRoleIDs(ids ...int) *OrgUserUpdate {
	ouu.mutation.AddOrgRoleIDs(ids...)
	return ouu
}

// AddOrgRoles adds the "org_roles" edges to the OrgRole entity.
func (ouu *OrgUserUpdate) AddOrgRoles(o ...*OrgRole) *OrgUserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouu.AddOrgRoleIDs(ids...)
}

// AddOrgRoleUserIDs adds the "org_role_user" edge to the OrgRoleUser entity by IDs.
func (ouu *OrgUserUpdate) AddOrgRoleUserIDs(ids ...int) *OrgUserUpdate {
	ouu.mutation.AddOrgRoleUserIDs(ids...)
	return ouu
}

// AddOrgRoleUser adds the "org_role_user" edges to the OrgRoleUser entity.
func (ouu *OrgUserUpdate) AddOrgRoleUser(o ...*OrgRoleUser) *OrgUserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouu.AddOrgRoleUserIDs(ids...)
}

// Mutation returns the OrgUserMutation object of the builder.
func (ouu *OrgUserUpdate) Mutation() *OrgUserMutation {
	return ouu.mutation
}

// ClearOrg clears the "org" edge to the Org entity.
func (ouu *OrgUserUpdate) ClearOrg() *OrgUserUpdate {
	ouu.mutation.ClearOrg()
	return ouu
}

// ClearUser clears the "user" edge to the User entity.
func (ouu *OrgUserUpdate) ClearUser() *OrgUserUpdate {
	ouu.mutation.ClearUser()
	return ouu
}

// ClearOrgRoles clears all "org_roles" edges to the OrgRole entity.
func (ouu *OrgUserUpdate) ClearOrgRoles() *OrgUserUpdate {
	ouu.mutation.ClearOrgRoles()
	return ouu
}

// RemoveOrgRoleIDs removes the "org_roles" edge to OrgRole entities by IDs.
func (ouu *OrgUserUpdate) RemoveOrgRoleIDs(ids ...int) *OrgUserUpdate {
	ouu.mutation.RemoveOrgRoleIDs(ids...)
	return ouu
}

// RemoveOrgRoles removes "org_roles" edges to OrgRole entities.
func (ouu *OrgUserUpdate) RemoveOrgRoles(o ...*OrgRole) *OrgUserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouu.RemoveOrgRoleIDs(ids...)
}

// ClearOrgRoleUser clears all "org_role_user" edges to the OrgRoleUser entity.
func (ouu *OrgUserUpdate) ClearOrgRoleUser() *OrgUserUpdate {
	ouu.mutation.ClearOrgRoleUser()
	return ouu
}

// RemoveOrgRoleUserIDs removes the "org_role_user" edge to OrgRoleUser entities by IDs.
func (ouu *OrgUserUpdate) RemoveOrgRoleUserIDs(ids ...int) *OrgUserUpdate {
	ouu.mutation.RemoveOrgRoleUserIDs(ids...)
	return ouu
}

// RemoveOrgRoleUser removes "org_role_user" edges to OrgRoleUser entities.
func (ouu *OrgUserUpdate) RemoveOrgRoleUser(o ...*OrgRoleUser) *OrgUserUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouu.RemoveOrgRoleUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ouu *OrgUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ouu.sqlSave, ouu.mutation, ouu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouu *OrgUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ouu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ouu *OrgUserUpdate) Exec(ctx context.Context) error {
	_, err := ouu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouu *OrgUserUpdate) ExecX(ctx context.Context) {
	if err := ouu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouu *OrgUserUpdate) check() error {
	if v, ok := ouu.mutation.UserType(); ok {
		if err := orguser.UserTypeValidator(v); err != nil {
			return &ValidationError{Name: "user_type", err: fmt.Errorf(`ent: validator failed for field "OrgUser.user_type": %w`, err)}
		}
	}
	if ouu.mutation.OrgCleared() && len(ouu.mutation.OrgIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrgUser.org"`)
	}
	if ouu.mutation.UserCleared() && len(ouu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrgUser.user"`)
	}
	return nil
}

func (ouu *OrgUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ouu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orguser.Table, orguser.Columns, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	if ps := ouu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouu.mutation.UpdatedBy(); ok {
		_spec.SetField(orguser.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ouu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(orguser.FieldUpdatedBy, field.TypeInt, value)
	}
	if ouu.mutation.UpdatedByCleared() {
		_spec.ClearField(orguser.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := ouu.mutation.UpdatedAt(); ok {
		_spec.SetField(orguser.FieldUpdatedAt, field.TypeTime, value)
	}
	if ouu.mutation.UpdatedAtCleared() {
		_spec.ClearField(orguser.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ouu.mutation.JoinedAt(); ok {
		_spec.SetField(orguser.FieldJoinedAt, field.TypeTime, value)
	}
	if value, ok := ouu.mutation.DisplayName(); ok {
		_spec.SetField(orguser.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := ouu.mutation.UserType(); ok {
		_spec.SetField(orguser.FieldUserType, field.TypeEnum, value)
	}
	if ouu.mutation.OrgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.OrgTable,
			Columns: []string{orguser.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouu.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.OrgTable,
			Columns: []string{orguser.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.UserTable,
			Columns: []string{orguser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.UserTable,
			Columns: []string{orguser.UserColumn},
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
	if ouu.mutation.OrgRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orguser.OrgRolesTable,
			Columns: orguser.OrgRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		createE := &OrgRoleUserCreate{config: ouu.config, mutation: newOrgRoleUserMutation(ouu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouu.mutation.RemovedOrgRolesIDs(); len(nodes) > 0 && !ouu.mutation.OrgRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orguser.OrgRolesTable,
			Columns: orguser.OrgRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &OrgRoleUserCreate{config: ouu.config, mutation: newOrgRoleUserMutation(ouu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouu.mutation.OrgRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orguser.OrgRolesTable,
			Columns: orguser.OrgRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &OrgRoleUserCreate{config: ouu.config, mutation: newOrgRoleUserMutation(ouu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouu.mutation.OrgRoleUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   orguser.OrgRoleUserTable,
			Columns: []string{orguser.OrgRoleUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouu.mutation.RemovedOrgRoleUserIDs(); len(nodes) > 0 && !ouu.mutation.OrgRoleUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   orguser.OrgRoleUserTable,
			Columns: []string{orguser.OrgRoleUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouu.mutation.OrgRoleUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   orguser.OrgRoleUserTable,
			Columns: []string{orguser.OrgRoleUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ouu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orguser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ouu.mutation.done = true
	return n, nil
}

// OrgUserUpdateOne is the builder for updating a single OrgUser entity.
type OrgUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgUserMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (ouuo *OrgUserUpdateOne) SetUpdatedBy(i int) *OrgUserUpdateOne {
	ouuo.mutation.ResetUpdatedBy()
	ouuo.mutation.SetUpdatedBy(i)
	return ouuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableUpdatedBy(i *int) *OrgUserUpdateOne {
	if i != nil {
		ouuo.SetUpdatedBy(*i)
	}
	return ouuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ouuo *OrgUserUpdateOne) AddUpdatedBy(i int) *OrgUserUpdateOne {
	ouuo.mutation.AddUpdatedBy(i)
	return ouuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ouuo *OrgUserUpdateOne) ClearUpdatedBy() *OrgUserUpdateOne {
	ouuo.mutation.ClearUpdatedBy()
	return ouuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouuo *OrgUserUpdateOne) SetUpdatedAt(t time.Time) *OrgUserUpdateOne {
	ouuo.mutation.SetUpdatedAt(t)
	return ouuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableUpdatedAt(t *time.Time) *OrgUserUpdateOne {
	if t != nil {
		ouuo.SetUpdatedAt(*t)
	}
	return ouuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ouuo *OrgUserUpdateOne) ClearUpdatedAt() *OrgUserUpdateOne {
	ouuo.mutation.ClearUpdatedAt()
	return ouuo
}

// SetOrgID sets the "org_id" field.
func (ouuo *OrgUserUpdateOne) SetOrgID(i int) *OrgUserUpdateOne {
	ouuo.mutation.SetOrgID(i)
	return ouuo
}

// SetNillableOrgID sets the "org_id" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableOrgID(i *int) *OrgUserUpdateOne {
	if i != nil {
		ouuo.SetOrgID(*i)
	}
	return ouuo
}

// SetUserID sets the "user_id" field.
func (ouuo *OrgUserUpdateOne) SetUserID(i int) *OrgUserUpdateOne {
	ouuo.mutation.SetUserID(i)
	return ouuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableUserID(i *int) *OrgUserUpdateOne {
	if i != nil {
		ouuo.SetUserID(*i)
	}
	return ouuo
}

// SetJoinedAt sets the "joined_at" field.
func (ouuo *OrgUserUpdateOne) SetJoinedAt(t time.Time) *OrgUserUpdateOne {
	ouuo.mutation.SetJoinedAt(t)
	return ouuo
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableJoinedAt(t *time.Time) *OrgUserUpdateOne {
	if t != nil {
		ouuo.SetJoinedAt(*t)
	}
	return ouuo
}

// SetDisplayName sets the "display_name" field.
func (ouuo *OrgUserUpdateOne) SetDisplayName(s string) *OrgUserUpdateOne {
	ouuo.mutation.SetDisplayName(s)
	return ouuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableDisplayName(s *string) *OrgUserUpdateOne {
	if s != nil {
		ouuo.SetDisplayName(*s)
	}
	return ouuo
}

// SetUserType sets the "user_type" field.
func (ouuo *OrgUserUpdateOne) SetUserType(ot orguser.UserType) *OrgUserUpdateOne {
	ouuo.mutation.SetUserType(ot)
	return ouuo
}

// SetNillableUserType sets the "user_type" field if the given value is not nil.
func (ouuo *OrgUserUpdateOne) SetNillableUserType(ot *orguser.UserType) *OrgUserUpdateOne {
	if ot != nil {
		ouuo.SetUserType(*ot)
	}
	return ouuo
}

// SetOrg sets the "org" edge to the Org entity.
func (ouuo *OrgUserUpdateOne) SetOrg(o *Org) *OrgUserUpdateOne {
	return ouuo.SetOrgID(o.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ouuo *OrgUserUpdateOne) SetUser(u *User) *OrgUserUpdateOne {
	return ouuo.SetUserID(u.ID)
}

// AddOrgRoleIDs adds the "org_roles" edge to the OrgRole entity by IDs.
func (ouuo *OrgUserUpdateOne) AddOrgRoleIDs(ids ...int) *OrgUserUpdateOne {
	ouuo.mutation.AddOrgRoleIDs(ids...)
	return ouuo
}

// AddOrgRoles adds the "org_roles" edges to the OrgRole entity.
func (ouuo *OrgUserUpdateOne) AddOrgRoles(o ...*OrgRole) *OrgUserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouuo.AddOrgRoleIDs(ids...)
}

// AddOrgRoleUserIDs adds the "org_role_user" edge to the OrgRoleUser entity by IDs.
func (ouuo *OrgUserUpdateOne) AddOrgRoleUserIDs(ids ...int) *OrgUserUpdateOne {
	ouuo.mutation.AddOrgRoleUserIDs(ids...)
	return ouuo
}

// AddOrgRoleUser adds the "org_role_user" edges to the OrgRoleUser entity.
func (ouuo *OrgUserUpdateOne) AddOrgRoleUser(o ...*OrgRoleUser) *OrgUserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouuo.AddOrgRoleUserIDs(ids...)
}

// Mutation returns the OrgUserMutation object of the builder.
func (ouuo *OrgUserUpdateOne) Mutation() *OrgUserMutation {
	return ouuo.mutation
}

// ClearOrg clears the "org" edge to the Org entity.
func (ouuo *OrgUserUpdateOne) ClearOrg() *OrgUserUpdateOne {
	ouuo.mutation.ClearOrg()
	return ouuo
}

// ClearUser clears the "user" edge to the User entity.
func (ouuo *OrgUserUpdateOne) ClearUser() *OrgUserUpdateOne {
	ouuo.mutation.ClearUser()
	return ouuo
}

// ClearOrgRoles clears all "org_roles" edges to the OrgRole entity.
func (ouuo *OrgUserUpdateOne) ClearOrgRoles() *OrgUserUpdateOne {
	ouuo.mutation.ClearOrgRoles()
	return ouuo
}

// RemoveOrgRoleIDs removes the "org_roles" edge to OrgRole entities by IDs.
func (ouuo *OrgUserUpdateOne) RemoveOrgRoleIDs(ids ...int) *OrgUserUpdateOne {
	ouuo.mutation.RemoveOrgRoleIDs(ids...)
	return ouuo
}

// RemoveOrgRoles removes "org_roles" edges to OrgRole entities.
func (ouuo *OrgUserUpdateOne) RemoveOrgRoles(o ...*OrgRole) *OrgUserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouuo.RemoveOrgRoleIDs(ids...)
}

// ClearOrgRoleUser clears all "org_role_user" edges to the OrgRoleUser entity.
func (ouuo *OrgUserUpdateOne) ClearOrgRoleUser() *OrgUserUpdateOne {
	ouuo.mutation.ClearOrgRoleUser()
	return ouuo
}

// RemoveOrgRoleUserIDs removes the "org_role_user" edge to OrgRoleUser entities by IDs.
func (ouuo *OrgUserUpdateOne) RemoveOrgRoleUserIDs(ids ...int) *OrgUserUpdateOne {
	ouuo.mutation.RemoveOrgRoleUserIDs(ids...)
	return ouuo
}

// RemoveOrgRoleUser removes "org_role_user" edges to OrgRoleUser entities.
func (ouuo *OrgUserUpdateOne) RemoveOrgRoleUser(o ...*OrgRoleUser) *OrgUserUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouuo.RemoveOrgRoleUserIDs(ids...)
}

// Where appends a list predicates to the OrgUserUpdate builder.
func (ouuo *OrgUserUpdateOne) Where(ps ...predicate.OrgUser) *OrgUserUpdateOne {
	ouuo.mutation.Where(ps...)
	return ouuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouuo *OrgUserUpdateOne) Select(field string, fields ...string) *OrgUserUpdateOne {
	ouuo.fields = append([]string{field}, fields...)
	return ouuo
}

// Save executes the query and returns the updated OrgUser entity.
func (ouuo *OrgUserUpdateOne) Save(ctx context.Context) (*OrgUser, error) {
	return withHooks(ctx, ouuo.sqlSave, ouuo.mutation, ouuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouuo *OrgUserUpdateOne) SaveX(ctx context.Context) *OrgUser {
	node, err := ouuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouuo *OrgUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ouuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouuo *OrgUserUpdateOne) ExecX(ctx context.Context) {
	if err := ouuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouuo *OrgUserUpdateOne) check() error {
	if v, ok := ouuo.mutation.UserType(); ok {
		if err := orguser.UserTypeValidator(v); err != nil {
			return &ValidationError{Name: "user_type", err: fmt.Errorf(`ent: validator failed for field "OrgUser.user_type": %w`, err)}
		}
	}
	if ouuo.mutation.OrgCleared() && len(ouuo.mutation.OrgIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrgUser.org"`)
	}
	if ouuo.mutation.UserCleared() && len(ouuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrgUser.user"`)
	}
	return nil
}

func (ouuo *OrgUserUpdateOne) sqlSave(ctx context.Context) (_node *OrgUser, err error) {
	if err := ouuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orguser.Table, orguser.Columns, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	id, ok := ouuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrgUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orguser.FieldID)
		for _, f := range fields {
			if !orguser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orguser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouuo.mutation.UpdatedBy(); ok {
		_spec.SetField(orguser.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ouuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(orguser.FieldUpdatedBy, field.TypeInt, value)
	}
	if ouuo.mutation.UpdatedByCleared() {
		_spec.ClearField(orguser.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := ouuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orguser.FieldUpdatedAt, field.TypeTime, value)
	}
	if ouuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(orguser.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ouuo.mutation.JoinedAt(); ok {
		_spec.SetField(orguser.FieldJoinedAt, field.TypeTime, value)
	}
	if value, ok := ouuo.mutation.DisplayName(); ok {
		_spec.SetField(orguser.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := ouuo.mutation.UserType(); ok {
		_spec.SetField(orguser.FieldUserType, field.TypeEnum, value)
	}
	if ouuo.mutation.OrgCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.OrgTable,
			Columns: []string{orguser.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouuo.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.OrgTable,
			Columns: []string{orguser.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.UserTable,
			Columns: []string{orguser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.UserTable,
			Columns: []string{orguser.UserColumn},
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
	if ouuo.mutation.OrgRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orguser.OrgRolesTable,
			Columns: orguser.OrgRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		createE := &OrgRoleUserCreate{config: ouuo.config, mutation: newOrgRoleUserMutation(ouuo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouuo.mutation.RemovedOrgRolesIDs(); len(nodes) > 0 && !ouuo.mutation.OrgRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orguser.OrgRolesTable,
			Columns: orguser.OrgRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &OrgRoleUserCreate{config: ouuo.config, mutation: newOrgRoleUserMutation(ouuo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouuo.mutation.OrgRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orguser.OrgRolesTable,
			Columns: orguser.OrgRolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &OrgRoleUserCreate{config: ouuo.config, mutation: newOrgRoleUserMutation(ouuo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouuo.mutation.OrgRoleUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   orguser.OrgRoleUserTable,
			Columns: []string{orguser.OrgRoleUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouuo.mutation.RemovedOrgRoleUserIDs(); len(nodes) > 0 && !ouuo.mutation.OrgRoleUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   orguser.OrgRoleUserTable,
			Columns: []string{orguser.OrgRoleUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouuo.mutation.OrgRoleUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   orguser.OrgRoleUserTable,
			Columns: []string{orguser.OrgRoleUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrgUser{config: ouuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orguser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouuo.mutation.done = true
	return _node, nil
}
