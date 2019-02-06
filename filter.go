package gooc

import (
	"reflect"
	"strings"
)

func filter(rv reflect.Value, m map[string][]string, currentPath []string, allowedPathList []string) interface{} {
	if !rv.IsValid() {
		return nil
	}

	wildcard := isWildcard(allowedPathList)
	if wildcard {
		return rv.Interface()
	}

	rt := rv.Type()
	switch rt.Kind() {
	case reflect.Map:
		keys := rv.MapKeys()
		var result map[string]interface{} = nil
		path := append(currentPath, "")
		lastIndex := len(path) - 1
		for _, key := range keys {
			keyString := key.String()
			path[lastIndex] = keyString
			if pathList, ok := m[strings.Join(path, ".")]; ok || wildcard {
				value := rv.MapIndex(key)
				if filtered := filter(value, m, path, pathList); filtered != nil {
					if result == nil {
						result = make(map[string]interface{}, len(keys)/2)
					}
					result[keyString] = filtered
				}
			}
		}
		return result
	case reflect.Slice, reflect.Array:
		l := rv.Len()
		arr := make([]interface{}, 0, l)
		for i := 0; i < l; i++ {
			value := rv.Index(i)
			if filtered := filter(value, m, currentPath, allowedPathList); filtered != nil {
				arr = append(arr, filtered)
			}
		}
		return arr
	case reflect.Ptr, reflect.Interface:
		return filter(rv.Elem(), m, currentPath, allowedPathList)
	default:
		return rv.Interface()
	}
}

func isWildcard(allowedPathList []string) bool {
	return len(allowedPathList) == 1 && allowedPathList[0] == WildcardMatching
}
