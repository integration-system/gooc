package gooc

import (
	"reflect"
	"strings"
)

func (c *cleaner) doClean(rv reflect.Value, currentPath []string, allowed []string, excluded []string) interface{} {
	if !rv.IsValid() {
		return nil
	}

	if isWildcard(excluded) {
		return nil
	}
	if isWildcard(allowed) {
		return rv.Interface()
	}

	rt := rv.Type()
	switch rt.Kind() {
	case reflect.Map:
		keys := rv.MapKeys()
		if len(keys) == 0 {
			return rv.Interface()
		}
		var result map[string]interface{} = nil
		path := append(currentPath, "")
		lastIndex := len(path) - 1
		for _, key := range keys {
			keyString := key.String()
			path[lastIndex] = keyString
			matched, wl, bl := c.match(strings.Join(path, "."))
			if matched {
				value := rv.MapIndex(key)
				if filtered := c.doClean(value, path, wl, bl); filtered != nil {
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
			if filtered := c.doClean(value, currentPath, allowed, excluded); filtered != nil {
				arr = append(arr, filtered)
			}
		}
		return arr
	case reflect.Ptr, reflect.Interface:
		return c.doClean(rv.Elem(), currentPath, allowed, excluded)
	default:
		return rv.Interface()
	}
}

func (c *cleaner) match(curPath string) (bool, []string, []string) {
	wl, inWl := c.compiledWl[curPath]
	bl, inBl := c.compiledBl[curPath]
	inBl = inBl && isWildcard(bl)
	if c.wlWildcard {
		return !inBl, wl, bl
	}
	if c.blWildcard {
		return inWl, wl, bl
	}
	return inWl && !inBl, wl, bl
}

func isWildcard(allowedPathList []string) bool {
	return len(allowedPathList) == 1 && allowedPathList[0] == WildcardMatching
}
