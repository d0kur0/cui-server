package gqlutils

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func AccessDeniedError(ctx context.Context) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "access denied",
		Extensions: map[string]interface{}{
			"code": "301",
		},
	}
}
