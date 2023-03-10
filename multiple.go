package testcraft

type MultipleAttribute[T any] func(i int) T

func Multiple[T any](num int, val MultipleAttribute[T]) []T {
	var attrs []T
	for i := 0; i < num; i++ {
		attrs = append(attrs, val(i))
	}
	return attrs
}
