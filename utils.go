package webidl

func new2[T any](v T) *T {
	return &v
}
