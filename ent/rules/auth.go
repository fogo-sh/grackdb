package rules

import (
	"context"

	"github.com/fogo-sh/grackdb"
	"github.com/fogo-sh/grackdb/ent"
	"github.com/fogo-sh/grackdb/ent/privacy"
)

func DenyIfNotAuthenticated() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		ginCtx, err := grackdb.GinContextFromContext(ctx)
		if err != nil {
			return privacy.Denyf("unable to fetch context: %s", err)
		}
		user, _ := ginCtx.Value("user").(*ent.User)
		if user == nil {
			return privacy.Denyf("no valid bearer token provided")
		}
		return privacy.Skip
	})
}
