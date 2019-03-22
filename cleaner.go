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
	wlWildcard bool
	blWildcard bool
	emptyWl    bool
	emptyBl    bool
	compiledWl map[string][]string
	compiledBl map[string][]string
}

// sanitize value by compiled spec
// not mutate original object, but for better performance new objects can contains original pointers to maps and slices
// works only with maps and slices, return original value otherwise
func (c *cleaner) Apply(value interface{}) interface{} {
	if c.wlWildcard {
		if c.emptyBl {
			return value
		}
		if c.blWildcard {
			return nil
		}
	}

	if c.blWildcard {
		return nil
	}

	return c.doClean(reflect.ValueOf(value), make([]string, 0, 3), nil, nil)
}

// creates new cleaner, whiteList must contains path to allowed object properties
func NewCleaner(whiteList, blackList []string) Cleaner {
	return &cleaner{
		wlWildcard: hasWildcard(whiteList),
		blWildcard: hasWildcard(blackList),
		emptyWl:    len(whiteList) == 0,
		emptyBl:    len(blackList) == 0,
		compiledWl: compile(whiteList),
		compiledBl: compile(blackList),
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

func hasWildcard(ss []string) bool {
	for _, path := range ss {
		if path == WildcardMatching {
			return true
		}
	}
	return false
}
