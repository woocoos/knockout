// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/woocoos/knockout/ent"
)

// The AppFunc type is an adapter to allow the use of ordinary
// function as App mutator.
type AppFunc func(context.Context, *ent.AppMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppMutation", m)
}

// The AppActionFunc type is an adapter to allow the use of ordinary
// function as AppAction mutator.
type AppActionFunc func(context.Context, *ent.AppActionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppActionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppActionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppActionMutation", m)
}

// The AppMenuFunc type is an adapter to allow the use of ordinary
// function as AppMenu mutator.
type AppMenuFunc func(context.Context, *ent.AppMenuMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppMenuFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppMenuMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppMenuMutation", m)
}

// The AppPolicyFunc type is an adapter to allow the use of ordinary
// function as AppPolicy mutator.
type AppPolicyFunc func(context.Context, *ent.AppPolicyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppPolicyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppPolicyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppPolicyMutation", m)
}

// The AppResFunc type is an adapter to allow the use of ordinary
// function as AppRes mutator.
type AppResFunc func(context.Context, *ent.AppResMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppResFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppResMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppResMutation", m)
}

// The AppRoleFunc type is an adapter to allow the use of ordinary
// function as AppRole mutator.
type AppRoleFunc func(context.Context, *ent.AppRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppRoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppRoleMutation", m)
}

// The AppRolePolicyFunc type is an adapter to allow the use of ordinary
// function as AppRolePolicy mutator.
type AppRolePolicyFunc func(context.Context, *ent.AppRolePolicyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppRolePolicyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AppRolePolicyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppRolePolicyMutation", m)
}

// The OrganizationFunc type is an adapter to allow the use of ordinary
// function as Organization mutator.
type OrganizationFunc func(context.Context, *ent.OrganizationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrganizationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrganizationMutation", m)
}

// The OrganizationAppFunc type is an adapter to allow the use of ordinary
// function as OrganizationApp mutator.
type OrganizationAppFunc func(context.Context, *ent.OrganizationAppMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationAppFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrganizationAppMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrganizationAppMutation", m)
}

// The OrganizationPolicyFunc type is an adapter to allow the use of ordinary
// function as OrganizationPolicy mutator.
type OrganizationPolicyFunc func(context.Context, *ent.OrganizationPolicyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationPolicyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrganizationPolicyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrganizationPolicyMutation", m)
}

// The OrganizationRoleFunc type is an adapter to allow the use of ordinary
// function as OrganizationRole mutator.
type OrganizationRoleFunc func(context.Context, *ent.OrganizationRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrganizationRoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrganizationRoleMutation", m)
}

// The OrganizationUserFunc type is an adapter to allow the use of ordinary
// function as OrganizationUser mutator.
type OrganizationUserFunc func(context.Context, *ent.OrganizationUserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationUserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrganizationUserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrganizationUserMutation", m)
}

// The PermissionFunc type is an adapter to allow the use of ordinary
// function as Permission mutator.
type PermissionFunc func(context.Context, *ent.PermissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PermissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PermissionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PermissionMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The UserDeviceFunc type is an adapter to allow the use of ordinary
// function as UserDevice mutator.
type UserDeviceFunc func(context.Context, *ent.UserDeviceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserDeviceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserDeviceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserDeviceMutation", m)
}

// The UserIdentityFunc type is an adapter to allow the use of ordinary
// function as UserIdentity mutator.
type UserIdentityFunc func(context.Context, *ent.UserIdentityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserIdentityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserIdentityMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserIdentityMutation", m)
}

// The UserLoginProfileFunc type is an adapter to allow the use of ordinary
// function as UserLoginProfile mutator.
type UserLoginProfileFunc func(context.Context, *ent.UserLoginProfileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserLoginProfileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserLoginProfileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserLoginProfileMutation", m)
}

// The UserPasswordFunc type is an adapter to allow the use of ordinary
// function as UserPassword mutator.
type UserPasswordFunc func(context.Context, *ent.UserPasswordMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserPasswordFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserPasswordMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserPasswordMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
