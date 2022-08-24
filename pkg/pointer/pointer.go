package pointer

func Get[T any](v T) *T {
	return &v
}

func String(v string) *string {
	return &v
}
