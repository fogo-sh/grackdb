// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/fogo-sh/grackdb/ent"

	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query a policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutate a  policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The DiscordAccountQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DiscordAccountQueryRuleFunc func(context.Context, *ent.DiscordAccountQuery) error

// EvalQuery return f(ctx, q).
func (f DiscordAccountQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DiscordAccountQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DiscordAccountQuery", q)
}

// The DiscordAccountMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DiscordAccountMutationRuleFunc func(context.Context, *ent.DiscordAccountMutation) error

// EvalMutation calls f(ctx, m).
func (f DiscordAccountMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DiscordAccountMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DiscordAccountMutation", m)
}

// The GithubAccountQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GithubAccountQueryRuleFunc func(context.Context, *ent.GithubAccountQuery) error

// EvalQuery return f(ctx, q).
func (f GithubAccountQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GithubAccountQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GithubAccountQuery", q)
}

// The GithubAccountMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GithubAccountMutationRuleFunc func(context.Context, *ent.GithubAccountMutation) error

// EvalMutation calls f(ctx, m).
func (f GithubAccountMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GithubAccountMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GithubAccountMutation", m)
}

// The GithubOrganizationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GithubOrganizationQueryRuleFunc func(context.Context, *ent.GithubOrganizationQuery) error

// EvalQuery return f(ctx, q).
func (f GithubOrganizationQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GithubOrganizationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GithubOrganizationQuery", q)
}

// The GithubOrganizationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GithubOrganizationMutationRuleFunc func(context.Context, *ent.GithubOrganizationMutation) error

// EvalMutation calls f(ctx, m).
func (f GithubOrganizationMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GithubOrganizationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GithubOrganizationMutation", m)
}

// The GithubOrganizationMemberQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GithubOrganizationMemberQueryRuleFunc func(context.Context, *ent.GithubOrganizationMemberQuery) error

// EvalQuery return f(ctx, q).
func (f GithubOrganizationMemberQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GithubOrganizationMemberQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GithubOrganizationMemberQuery", q)
}

// The GithubOrganizationMemberMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GithubOrganizationMemberMutationRuleFunc func(context.Context, *ent.GithubOrganizationMemberMutation) error

// EvalMutation calls f(ctx, m).
func (f GithubOrganizationMemberMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GithubOrganizationMemberMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GithubOrganizationMemberMutation", m)
}

// The ProjectQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProjectQueryRuleFunc func(context.Context, *ent.ProjectQuery) error

// EvalQuery return f(ctx, q).
func (f ProjectQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProjectQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProjectQuery", q)
}

// The ProjectMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProjectMutationRuleFunc func(context.Context, *ent.ProjectMutation) error

// EvalMutation calls f(ctx, m).
func (f ProjectMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProjectMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProjectMutation", m)
}

// The ProjectAssociationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProjectAssociationQueryRuleFunc func(context.Context, *ent.ProjectAssociationQuery) error

// EvalQuery return f(ctx, q).
func (f ProjectAssociationQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProjectAssociationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProjectAssociationQuery", q)
}

// The ProjectAssociationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProjectAssociationMutationRuleFunc func(context.Context, *ent.ProjectAssociationMutation) error

// EvalMutation calls f(ctx, m).
func (f ProjectAssociationMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProjectAssociationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProjectAssociationMutation", m)
}

// The ProjectContributorQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ProjectContributorQueryRuleFunc func(context.Context, *ent.ProjectContributorQuery) error

// EvalQuery return f(ctx, q).
func (f ProjectContributorQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ProjectContributorQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ProjectContributorQuery", q)
}

// The ProjectContributorMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ProjectContributorMutationRuleFunc func(context.Context, *ent.ProjectContributorMutation) error

// EvalMutation calls f(ctx, m).
func (f ProjectContributorMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ProjectContributorMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ProjectContributorMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}
