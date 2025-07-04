// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

// The AnnouncementFunc type is an adapter to allow the use of ordinary
// function as Announcement mutator.
type AnnouncementFunc func(context.Context, *generated.AnnouncementMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f AnnouncementFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.AnnouncementMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.AnnouncementMutation", m)
}

// The ContactFunc type is an adapter to allow the use of ordinary
// function as Contact mutator.
type ContactFunc func(context.Context, *generated.ContactMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ContactFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ContactMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ContactMutation", m)
}

// The EmailTemplateFunc type is an adapter to allow the use of ordinary
// function as EmailTemplate mutator.
type EmailTemplateFunc func(context.Context, *generated.EmailTemplateMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EmailTemplateFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EmailTemplateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EmailTemplateMutation", m)
}

// The FrontendTemplateFunc type is an adapter to allow the use of ordinary
// function as FrontendTemplate mutator.
type FrontendTemplateFunc func(context.Context, *generated.FrontendTemplateMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f FrontendTemplateFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.FrontendTemplateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.FrontendTemplateMutation", m)
}

// The GoodBenefitFunc type is an adapter to allow the use of ordinary
// function as GoodBenefit mutator.
type GoodBenefitFunc func(context.Context, *generated.GoodBenefitMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GoodBenefitFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GoodBenefitMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GoodBenefitMutation", m)
}

// The NotifFunc type is an adapter to allow the use of ordinary
// function as Notif mutator.
type NotifFunc func(context.Context, *generated.NotifMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f NotifFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.NotifMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.NotifMutation", m)
}

// The NotifChannelFunc type is an adapter to allow the use of ordinary
// function as NotifChannel mutator.
type NotifChannelFunc func(context.Context, *generated.NotifChannelMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f NotifChannelFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.NotifChannelMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.NotifChannelMutation", m)
}

// The NotifUserFunc type is an adapter to allow the use of ordinary
// function as NotifUser mutator.
type NotifUserFunc func(context.Context, *generated.NotifUserMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f NotifUserFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.NotifUserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.NotifUserMutation", m)
}

// The ReadAnnouncementFunc type is an adapter to allow the use of ordinary
// function as ReadAnnouncement mutator.
type ReadAnnouncementFunc func(context.Context, *generated.ReadAnnouncementMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ReadAnnouncementFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ReadAnnouncementMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ReadAnnouncementMutation", m)
}

// The SMSTemplateFunc type is an adapter to allow the use of ordinary
// function as SMSTemplate mutator.
type SMSTemplateFunc func(context.Context, *generated.SMSTemplateMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f SMSTemplateFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.SMSTemplateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.SMSTemplateMutation", m)
}

// The SendAnnouncementFunc type is an adapter to allow the use of ordinary
// function as SendAnnouncement mutator.
type SendAnnouncementFunc func(context.Context, *generated.SendAnnouncementMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f SendAnnouncementFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.SendAnnouncementMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.SendAnnouncementMutation", m)
}

// The UserAnnouncementFunc type is an adapter to allow the use of ordinary
// function as UserAnnouncement mutator.
type UserAnnouncementFunc func(context.Context, *generated.UserAnnouncementMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserAnnouncementFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserAnnouncementMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserAnnouncementMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, generated.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
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
	return func(ctx context.Context, m generated.Mutation) bool {
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
	return func(ctx context.Context, m generated.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op generated.Op) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
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
	return func(_ context.Context, m generated.Mutation) bool {
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
	return func(_ context.Context, m generated.Mutation) bool {
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
func If(hk generated.Hook, cond Condition) generated.Hook {
	return func(next generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(ctx context.Context, m generated.Mutation) (generated.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, generated.Delete|generated.Create)
func On(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, generated.Update|generated.UpdateOne)
func Unless(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) generated.Hook {
	return func(generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(context.Context, generated.Mutation) (generated.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []generated.Hook {
//		return []generated.Hook{
//			Reject(generated.Delete|generated.Update),
//		}
//	}
func Reject(op generated.Op) generated.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []generated.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...generated.Hook) Chain {
	return Chain{append([]generated.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() generated.Hook {
	return func(mutator generated.Mutator) generated.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...generated.Hook) Chain {
	newHooks := make([]generated.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
