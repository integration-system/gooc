package gooc

import (
	"reflect"
	"strings"
)

const (
	WildcardMatching = "*"
)

var (
	AllMatch = []string{WildcardMatching}
)

type Cleaner interface {
	Apply(value interface{}) interface{}
}

//default cleaner implementation, thread-safe
type cleaner struct {
	wl           []string
	allAvailable bool
	m            map[string][]string
}

// sanitize value by compiled spec
// not mutate original object, but for better performance new objects can contains original pointers to maps and slices
// works only with maps and slices, return original value otherwise
func (c *cleaner) Apply(value interface{}) interface{} {
	if c.allAvailable {
		return value
	}

	return filter(reflect.ValueOf(value), c.m, make([]string, 0, 3), nil)
}

// creates new cleaner, whiteList must contains path to allowed object properties
func NewCleaner(whiteList []string) Cleaner {
	return &cleaner{
		wl:           whiteList,
		allAvailable: isWildcard(whiteList),
		m:            compile(whiteList),
	}
}

func compile(whiteList []string) map[string][]string {
	m := make(map[string][]string)
	for _, path := range whiteList {
		parts := strings.Split(path, ".")
		for i := 0; i < len(parts); i++ {
			isLast := i == len(parts)-1
			if isLast {
				path := strings.Join(parts, ".")
				m[path] = AllMatch
			} else {
				path := strings.Join(parts[0:i+1], ".")
				value := strings.Join(parts[i+1:], ".")
				if arr, ok := m[path]; ok {
					m[path] = append(arr, value)
				} else {
					m[path] = []string{value}
				}

			}
		}
	}
	return m
}
