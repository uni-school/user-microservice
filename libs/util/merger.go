package util

import (
	"github.com/imdario/mergo"
)

func MergeSlices[T comparable](slices ...[]T) []T {
	result := make([]T, 0)

	for _, s := range slices {
		result = append(result, s...)
	}

	return result
}

func MergeStructs[T comparable](structs ...T) T {
	dst := new(T)

	for _, s := range structs {
		mergo.Merge(dst, s, mergo.WithOverride)
	}

	return *dst
}
