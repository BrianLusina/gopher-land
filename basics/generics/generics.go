package generics

// S is a type with a parameterized method Identity.
type S struct{}

// Identity is a simple identity method that works for any type.
//func (S) Identity[T any](v T) T { return v }
func Identity[T any](v T) T { return v }

// HasIdentity is an interface that matches any type with a
// parameterized Identity method.
type HasIdentity interface {
	Identity(t any) any
}

// CheckIdentity checks the Identity method if it exists.
// Note that although this function calls a parameterized method,
// this function is not itself parameterized.
func CheckIdentity(v interface{}) {
	if _, ok := v.(HasIdentity); ok {
		if got := Identity[int](0); got != 0 {
			panic(got)
		}
	}
}

// CheckSIdentity passes an S value to CheckIdentity.
func CheckSIdentity() {
	CheckIdentity(S{})
}
