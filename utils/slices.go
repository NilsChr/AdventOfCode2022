package utils

import "reflect"

func Contains[T comparable](s []T, i T) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func ContainsGeneric[T any](slice []T, item T) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, item) {
			return true
		}
	}
	return false
}
