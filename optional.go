package optional

// Predicate is a function that returns true if condition is met.
type Predicate[T any] func(t T) bool

var (
	// OK is a boolean predicate that returns true if b is true.
	OK = Predicate[bool](func(b bool) bool { return b })

	// ErrNil is a predicate that returns true if err is nil.
	ErrNil = Predicate[error](func(err error) bool { return err == nil })
)

// Supplier is a function that produces the value on request.
// T is the type of the value, and P is the type of the predicate input.
type Supplier[T any, P any] func() (T, P)

// Value is an optional value, which is either present or absent.
// The value is present if the predicate is returns true, otherwise absent.
// It also provides a method to set fallback value if the value is absent.
type Value[T any, P any] struct {
	predicate Predicate[P]
	supplier  Supplier[T, P]
}

// New creates a new optional value.
func New[T any, P any](p Predicate[P], s Supplier[T, P]) Value[T, P] {
	return Value[T, P]{
		predicate: p,
		supplier:  s,
	}
}

// Get returns the value and the predicate value.
func (v Value[T, P]) Get() (T, P) {
	return v.supplier()
}

// Or returns the value if the predicate is true, otherwise return the fallback.
func (v Value[T, P]) Or(fallback T) T {
	return v.OrFetch(func() T { return fallback })
}

// OrFetch is similar to Or, but it uses the fallback supplier to get the fallback value.
func (v Value[T, P]) OrFetch(fn func() T) T {
	t, p := v.supplier()
	if !v.predicate(p) {
		return fn()
	}

	return t
}
