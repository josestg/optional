# Go Optional Value

A simple value to replace if-else statement for value initialization.

## Installation

```bash
go get github.com/josestg/optional
```

> Minimum Go version: 1.18, because it uses generics.

## Motivation

For example, we have function that gets account id from context, and if not found, we want to generate a new one.
Basically, we can do it using if-else statement like this:

```go
package main

import (
	"context"
)

func RequestIDFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value("request-id").(string)
	return v, ok
}

func main() {
	ctx := context.Background()
	id, ok := RequestIDFromContext(ctx)
	if !ok {
		id = "generated-id"
	}

	println(id) // generated-id
}

```

But, we can simplify it using `optional.Value` like this:

```go

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
```

How it works:

- `optional.New` takes 2 arguments, first is `optional.Predicate[P]` and the second one is `optional.Supplier[T, P]`.
- `optional.Supplier[T, P]` is a function that returns `T` (value) and `P` (predicate), the supplier the function that produces the value.
- `optional.Predicate[P]` is a function that takes `P` as argument to determine if the value is present or absent.
-  when predicate returns `false` or absent, the value will be replaced with the value from `optional.Value.Or` method or `optional.Value.OrFetch` method.


This is another example, by using custom predicate:

```go
package main

import (
	"github.com/josestg/optional"
)

func OddPredicate(v int) bool {
	return v%2 == 1
}

func OnlyOdd() optional.Value[string, int] {
	return optional.New(
		OddPredicate,
		func() (string, int) {
			return "only-odd", 112
		},
	)
}

func main() {
	id := OnlyOdd().Or("my-fallback")
	println(id) // my-fallback
}
```