package utils

type mapKey interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}

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

func IntoMap[T any, K mapKey](slice []T, mapper func(v T) K) map[K]T {
	dst := map[K]T{}
	for _, v := range slice {
		dst[mapper(v)] = v
	}
	return dst
}

func IntoSet[T any, K mapKey](slice []T, mapper func(v T) K) map[K]struct{} {
	dst := map[K]struct{}{}
	for _, v := range slice {
		dst[mapper(v)] = struct{}{}
	}
	return dst
}
