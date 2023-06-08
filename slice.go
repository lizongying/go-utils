package utils

func InSlice[T comparable](a T, s []T) bool {
	for _, i := range s {
		if a == i {
			return true
		}
	}
	return false
}
