package wrappers

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	jetModel "github.com/Thiti-Dev/scrumerization-core-service/.gen/scrumerization/public/model"
	"github.com/Thiti-Dev/scrumerization-core-service/graph"
	"github.com/Thiti-Dev/scrumerization-core-service/graph/model"
	context_type "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/context"
	infraUtils "github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/utils"
	"github.com/Thiti-Dev/scrumerization-core-service/pkg/tokenizer"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func RegisterDirectives(c *graph.Config, config *infraUtils.Config) {
	c.Directives.RequiredMember = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		token := ctx.Value(context_type.TokenCtxKey)
		prebuiltErr := &gqlerror.Error{
			Message: fmt.Sprintf("need to have the role of \"%s\" to perform this action", role),
			Extensions: map[string]interface{}{
				"code": "Authentication",
			},
		}
		if token == "" || token == nil {
			// return nil, fmt.Errorf("need to have the role of \"%s\" to perform this action", role)
			return nil, prebuiltErr
		} else {
			// if token is provided
			payload, err := tokenizer.VerifyToken(token.(string), []byte(config.JwtSecret))
			if err != nil {
				return nil, prebuiltErr
			}

			if payload.Role == jetModel.UserRole(model.RoleAdmin) || payload.Role == jetModel.UserRole(role) {
				newCtx := context.WithValue(ctx, context_type.UserDataCtxKey, payload)
				return next(newCtx)
			}

			return nil, prebuiltErr

		}
	}
}
