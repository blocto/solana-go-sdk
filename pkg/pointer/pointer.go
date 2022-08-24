package pointer

func Get[T any](v T) *T {
	return &v
}
