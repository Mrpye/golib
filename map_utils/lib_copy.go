// Package of utilities for working with maps
package map_utils

// CopyMap takes a map of strings to interfaces and returns a copy of that map
// with all the values copied as well.
// - m: the map to copy
// - returns: a copy of the map
func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}
	return cp
}
