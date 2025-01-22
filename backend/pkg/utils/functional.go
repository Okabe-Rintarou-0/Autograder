package utils

func Map[T any, D any](slice []T, mapper func(v T) D) []D {
	var dst []D
	for _, v := range slice {
		dst = append(dst, mapper(v))
	}
	return dst
}

func Reduce[T any, D any](slice []T, reducer func(sum D, v T) D) D {
	var dst D
	for _, v := range slice {
		dst = reducer(dst, v)
	}
	return dst
}

func Filter[T any](slice []T, filter func(v T) bool) []T {
	var dst []T
	for _, v := range slice {
		if filter(v) {
			dst = append(dst, v)
		}
	}
	return dst
}
