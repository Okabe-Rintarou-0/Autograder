package utils

func SendIfNotFull[T any](ch chan T, data T) bool {
	select {
	case ch <- data:
		return false
	default:
		return true
	}
}
