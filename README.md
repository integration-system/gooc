# GOOC (Golang object cleaner)

## Features
* Clean object by specified properties list;
* Do not mutate original objects;
* Works with Golang maps and slices;
* Supports partially properties matching.

## Install

```
go get github.com/integration-system/gooc
```

## Example

Source data:
```json
{
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
```
```go
wl := []string{"a.b.c", "a.b", "a.d", "k", "v.i", "j"}
c := gooc.NewCleaner(wl)
result := c.Apply(m)
```
Result:
```json
{
  "a": {
    "b": [
      {
        "c": "c",
        "c1": "c2",
        "c2": "c2"
      }
    ],
    "d": [
      "1",
      "2",
      "3"
    ]
  },
  "j": "j",
  "k": [
    {
      "c1": "c1",
      "c2": "c2",
      "c3": "c3"
    },
    {
      "c1": "c1",
      "c2": "c2",
      "c3": "c3"
    }
  ],
  "v": [
    {
      "i": "i"
    },
    {
      "i": "i"
    }
  ]
}
```

## TODO
* [ ] Blacklist supporting
* [ ] Full object copy option
