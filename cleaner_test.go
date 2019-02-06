package gooc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	src = map[string]interface{}{
		"a": map[string]interface{}{
			"b": []interface{}{map[string]interface{}{"c": "c", "vvv": "vvv", "ops": "osp"}},
			"d": []interface{}{"1", "2", "3"},
			"i": 111,
			"o": map[string]interface{}{},
		},
		"k": []interface{}{
			map[string]interface{}{"c": "c", "vvv": "vvv", "ops": "osp"},
			map[string]interface{}{"c": "c", "vvv": "vvv", "ops": "osp"},
		},
		"v": &[]interface{}{
			map[string]interface{}{"c": "c", "i": "i", "ops": "osp"},
			map[string]interface{}{"c": "c", "i": "i", "ops": "osp"},
		},
		"j": "j",
	}
	dst = map[string]interface{}{
		"a": map[string]interface{}{
			"b": []interface{}{map[string]interface{}{"c": "c", "vvv": "vvv", "ops": "osp"}},
			"d": []interface{}{"1", "2", "3"},
		},
		"k": []interface{}{
			map[string]interface{}{"c": "c", "vvv": "vvv", "ops": "osp"},
			map[string]interface{}{"c": "c", "vvv": "vvv", "ops": "osp"},
		},
		"v": []interface{}{
			map[string]interface{}{"i": "i"},
			map[string]interface{}{"i": "i"},
		},
		"j": "j",
	}
)

func TestCompile(t *testing.T) {
	assert := assert.New(t)

	m := compile([]string{"a.b.c", "a.b", "a.d", "k", "v.i"})
	assert.EqualValues(map[string][]string{
		"a":     {"b.c", "b", "d"},
		"a.d":   {"*"},
		"a.b.c": {"*"},
		"a.b":   {"*"},
		"k":     {"*"},
		"v":     {"i"},
		"v.i":   {"*"},
	}, m)
}

func TestCleaner_Perform(t *testing.T) {
	assert := assert.New(t)

	c := NewCleaner([]string{"a.b.c", "a.b", "a.d", "k", "v.i", "j"})
	res := c.Apply(src)
	assert.Equal(dst, res)
}

func TestCleaner_PerformEmpty(t *testing.T) {
	assert := assert.New(t)

	c := NewCleaner([]string{})
	res := c.Apply(src)
	assert.Nil(res)
}

func TestCleaner_PerformAllMatch(t *testing.T) {
	assert := assert.New(t)

	c := NewCleaner(AllMatch)
	res := c.Apply(src)
	assert.EqualValues(src, res)
}

func BenchmarkCleaner_Perform(b *testing.B) {
	c := NewCleaner([]string{"a.b.c", "a.b", "a.d", "k", "v.i", "j"})
	var res interface{}
	for i := 0; i < b.N; i++ {
		res = c.Apply(src)
	}
	res = res
}
