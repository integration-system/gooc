package main

import (
	"encoding/json"
	"fmt"
	"github.com/integration-system/gooc"
)

const src = `{
  "a": {
    "b": [{
        "c": "c",
        "c1": "c2",
        "c2": "c2"
      }],
    "d": [
      "1",
      "2",
      "3"
    ],
    "i": 111,
    "o": {}
  },
  "k": [{
      "c1": "c1",
      "c2": "c2",
      "c3": "c3"
    }, {
      "c1": "c1",
      "c2": "c2",
      "c3": "c3"
    }
  ],
  "v": [{
      "c": false,
      "i": "i",
      "ops": "ops"
    }, {
      "c": true,
      "i": "i",
      "ops": "ops"
    }
  ],
  "j": "j"
}
`

func main() {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		panic(err)
	}

	wl := []string{"a.b.c", "a.b", "a.d", "k", "v.i", "j"}
	c := gooc.NewCleaner(wl, nil)
	result := c.Apply(m)

	bytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
