package collection

import (
	"golang.org/x/exp/constraints"
)

// Index returns the first index of the target `t`, or
// -1 if no match is found.
func Index[T constraints.Ordered](vs []T, t T) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns `true` if the target  t is in the
// slice.
func Include[T constraints.Ordered](vs []T, t T) bool {
	return Index(vs, t) >= 0
}

// Any returns `true` if one of the items in the slice
// satisfies the predicate `f`.
func Any[T any](vs []T, f func(T) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns `true` if all the items in the slice
// satisfy the predicate `f`.
func All[T any](vs []T, f func(T) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all values in the
// slice that satisfy the predicate.
func Filter[T any](values []T, predicate func(T) bool) []T {
	filteredValues := make([]T, 0)
	for _, value := range values {
		if predicate(value) {
			filteredValues = append(filteredValues, value)
		}
	}
	return filteredValues
}
func Find[T any](values []T, predicate func(T) bool) *T {
	var found T
	for _, value := range values {
		if predicate(value) {
			found = value
			break
		}
	}
	return &found
}

// Map returns a new slice containing the results of applying
// the function `f` to each item in the original slice.
func Map[T any, K any](vs []T, f func(T) K) []K {
	vsm := make([]K, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Unique returns a new slice containing all unique values
// contained in the original slice
func Unique[T constraints.Ordered](slice []T) []T {
	keys := make(map[T]bool)
	list := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
