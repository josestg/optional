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
