package optional_test

import (
	"errors"
	"github.com/josestg/optional"
	"testing"
)

func TestValue(t *testing.T) {
	t.Run("using boolean predicate: should return the value if it is present", func(t *testing.T) {
		opt := optional.New(optional.OK, func() (int, bool) {
			return 1, true
		})

		got, pre := opt.Get()
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		if !pre {
			t.Errorf("got %t, want %t", pre, true)
		}

		// or with fallback when predicate is true will return the value.
		got = opt.Or(2)
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		// or with fallback supplier when predicate is true will return the value.
		got = opt.OrFetch(func() int {
			return 3
		})

		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}
	})

	t.Run("using boolean predicate: should return the value if it is absent", func(t *testing.T) {
		opt := optional.New(optional.OK, func() (int, bool) {
			return 1, false
		})

		// using get still returns the value, but the predicate is false
		// this is useful when you want to check the value, but not the predicate.
		got, pre := opt.Get()
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		if pre {
			t.Errorf("got %t, want %t", pre, false)
		}

		// using or returns the fallback value
		got = opt.Or(2)
		if got != 2 {
			t.Errorf("got %d, want %d", got, 2)
		}

		// using or fetch returns the fallback value
		got = opt.OrFetch(func() int {
			return 3
		})

		if got != 3 {
			t.Errorf("got %d, want %d", got, 3)
		}
	})

	t.Run("using error predicate: should return the value if it is present", func(t *testing.T) {
		opt := optional.New(optional.ErrNil, func() (int, error) {
			return 1, nil
		})

		got, pre := opt.Get()
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		if pre != nil {
			t.Errorf("got %v, want %v", pre, nil)
		}

		// or with fallback when predicate is true will return the value.
		got = opt.Or(2)
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		// or with fallback supplier when predicate is true will return the value.
		got = opt.OrFetch(func() int {
			return 3
		})

		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}
	})

	t.Run("using error predicate: should return the value if it is absent", func(t *testing.T) {
		someErr := errors.New("some error")
		opt := optional.New(optional.ErrNil, func() (int, error) {
			return 1, someErr
		})

		// using get still returns the value, but the predicate is false
		// this is useful when you want to check the value, but not the predicate.
		got, pre := opt.Get()
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		if pre != someErr {
			t.Errorf("got %v, want %v", pre, nil)
		}

		// using or returns the fallback value
		got = opt.Or(2)
		if got != 2 {
			t.Errorf("got %d, want %d", got, 2)
		}

		// using or fetch returns the fallback value
		got = opt.OrFetch(func() int {
			return 3
		})

		if got != 3 {
			t.Errorf("got %d, want %d", got, 3)
		}
	})

	t.Run("using custom predicate: should return the value if it is present", func(t *testing.T) {
		empty := func(s string) bool { return s == "" }

		opt := optional.New(empty, func() (int, string) {
			return 1, ""
		})

		got, pre := opt.Get()
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		if pre != "" {
			t.Errorf("got %v, want %v", pre, nil)
		}

		// or with fallback when predicate is true will return the value.
		got = opt.Or(2)
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		// or with fallback supplier when predicate is true will return the value.
		got = opt.OrFetch(func() int {
			return 3
		})

		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}
	})

	t.Run("using custom predicate: should return the value if it is absent", func(t *testing.T) {
		empty := func(s string) bool { return s == "" }

		opt := optional.New(empty, func() (int, string) {
			return 1, "not empty"
		})

		// using get still returns the value, but the predicate is false
		// this is useful when you want to check the value, but not the predicate.
		got, pre := opt.Get()
		if got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}

		if pre != "not empty" {
			t.Errorf("got %v, want %v", pre, nil)
		}

		// using or returns the fallback value
		got = opt.Or(2)
		if got != 2 {
			t.Errorf("got %d, want %d", got, 2)
		}

		// using or fetch returns the fallback value
		got = opt.OrFetch(func() int {
			return 3
		})
	})

}
