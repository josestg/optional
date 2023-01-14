package main

import (
	"context"

	"github.com/josestg/optional"
)

func RequestIDFromContext(ctx context.Context) optional.Value[string, bool] {
	return optional.New(
		optional.OK,
		func() (string, bool) {
			v, ok := ctx.Value("request-id").(string)
			return v, ok
		},
	)
}

func main() {
	ctx := context.Background()
	id := RequestIDFromContext(ctx).Or("generated-id")
	println(id) // generated-id
}
