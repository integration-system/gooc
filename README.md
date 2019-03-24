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
    "b": [{"c": "c", "c1": "c2", "c2": "c2"}],
    "d": ["1", "2", "3"],
    "i": 111,
    "o": {}
  },
  "k": [{"c1": "c1", "c2": "c2", "c3": "c3"}, {"c1": "c1", "c2": "c2", "c3": "c3"}],
  "v": [{"c": false, "i": "i", "ops": "ops"}, {"c": true, "i": "i", "ops": "ops"}],
  "j": "j"
}
```
```go
wl := []string{"a.b.c", "a.b", "a.d", "k", "v.i", "j"}
c := gooc.NewCleaner(wl, nil)
result := c.Apply(m)
```
Result:
```json
{
  "a": {
    "b": [{"c": "c", "c1": "c2", "c2": "c2"}],
    "d": ["1", "2", "3"]
  },
  "j": "j",
  "k": [{"c1": "c1", "c2": "c2", "c3": "c3"}, {"c1": "c1", "c2": "c2", "c3": "c3"}],
  "v": [{"i": "i"}, {"i": "i"}]
}
```
## Notes
1) Cleaner object is thread safe
2) If no properties matches returns `nil`
3) Use `*` in whitelist to allow all properties from root
4) White and black list composition
* 'Black' list has higher priority than 'white'
* Some cases:

| Whitelist    | Blacklist    | Result                           |
|--------------|--------------|----------------------------------|
| `['*']`      | `['*']`      | No available fields              |
| `['a', 'b']` | `['*']`      | No available fields              |
| `['*']`      | `['a', 'b']` | All available except `a` and `b` |
| `['a', 'b']` | `['b, 'd']`  | Available only `a`               |



## TODO
* [ ] Full object copy option
