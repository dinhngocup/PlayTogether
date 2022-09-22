package utils

func Contains[T comparable](array []T, element T) bool {
	for _, elementInArray := range array {
		if elementInArray == element {
			return true
		}
	}
	return false
}
