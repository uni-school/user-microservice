package util

import (
	"github.com/imdario/mergo"
)

func MergeSlices[T comparable](slices ...[]T) []T {
	uniqueMap := make(map[T]bool)

	for _, slice := range slices {
		for _, val := range slice {
			uniqueMap[val] = true
		}
	}

	result := make([]T, 0, len(uniqueMap))

	for key := range uniqueMap {
		result = append(result, key)
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
