package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ConvertToStruct[T any](obj any) (T, error) {
	c := new(T)
	bytes, err := json.Marshal(obj)
	if err != nil {
		return *c, err
	}
	err = json.Unmarshal(bytes, c)
	if err != nil {
		return *c, err
	}
	return *c, nil
}

func ConvertMapKeysToString[K comparable, V any](m map[K]V) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, fmt.Sprint(k))
	}
	return "[" + strings.Join(keys, ", ") + "]"
}
