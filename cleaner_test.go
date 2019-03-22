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

	c := NewCleaner([]string{"a.b.c", "a.b", "a.d", "k", "v.i", "j"}, nil)
	res := c.Apply(src)
	assert.EqualValues(dst, res)
}

func TestBlackList(t *testing.T) {
	assert := assert.New(t)

	dst := map[string]interface{}{
		"a": map[string]interface{}{
			"b": []interface{}{map[string]interface{}{"c": "c"}},
			"i": 111,
			"o": map[string]interface{}{},
		},
		"v": []interface{}{
			map[string]interface{}{"c": "c", "i": "i", "ops": "osp"},
			map[string]interface{}{"c": "c", "i": "i", "ops": "osp"},
		},
	}
	c := NewCleaner([]string{"*"}, []string{"j", "a.d", "a.b.vvv", "a.b.ops", "k"})
	res := c.Apply(src)
	assert.EqualValues(dst, res)
}

func TestAllMatchBlackList(t *testing.T) {
	assert := assert.New(t)

	c := NewCleaner([]string{"a.b.c", "a.b", "a.d", "k", "v.i", "j"}, []string{"*"})
	res := c.Apply(src)
	assert.Nil(res)

	c = NewCleaner([]string{"*"}, []string{"*", "ignore_it"})
	res = c.Apply(src)
	assert.Nil(res)
}

func TestCleaner_PerformEmpty(t *testing.T) {
	assert := assert.New(t)

	c := NewCleaner([]string{}, nil)
	res := c.Apply(src)
	assert.Nil(res)
}

func TestCleaner_PerformAllMatch(t *testing.T) {
	assert := assert.New(t)

	c := NewCleaner(AllMatch, nil)
	res := c.Apply(src)
	assert.EqualValues(src, res)
}

func BenchmarkCleaner_Perform(b *testing.B) {
	c := NewCleaner([]string{"a.b.c", "a.b", "a.d", "k", "v.i", "j"}, nil)
	for i := 0; i < b.N; i++ {
		_ = c.Apply(src)
	}

}
