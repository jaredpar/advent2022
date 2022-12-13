package util

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func MinSlice[T constraints.Ordered](data []T) T {
	if len(data) == 0 {
		panic("must be at least one element")
	}

	min := data[0]
	for i := 1; i < len(data); i++ {
		min = Min(min, data[i])
	}

	return min
}

func MaxSlice[T constraints.Ordered](data []T) T {
	if len(data) == 0 {
		panic("must be at least one element")
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		max = Max(max, data[i])
	}

	return max
}

func InsertAt[T any](data []T, index int, value T) []T {
	length := len(data)
	if length == index {
		return append(data, value)
	}

	var dummy T
	data = append(data, dummy)

	for i := length; i > index; i-- {
		data[i] = data[i-1]
	}

	data[index] = value
	return data
}

func InsertSortedF[T any](data []T, value T, less func(left, right T) bool) []T {
	length := len(data)
	if length == 0 {
		return append(data, value)
	}

	index := sort.Search(length, func(i int) bool {
		return less(value, data[i])
	})

	return InsertAt(data, index, value)
}

func InsertSorted[S ~[]E, E constraints.Ordered](data S, value E) S {
	return InsertSortedF(data, value, func(left, right E) bool {
		return left < right
	})
}

func Project[T any, U any](data []T, project func(T) U) []U {
	projection := make([]U, len(data))
	for i, t := range data {
		projection[i] = project(t)
	}
	return projection
}
