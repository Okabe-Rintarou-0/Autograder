package utils

func Map[T any, D any](slice []T, mapper func(v T) D) []D {
	var dst []D
	for _, v := range slice {
		dst = append(dst, mapper(v))
	}
	return dst
}
