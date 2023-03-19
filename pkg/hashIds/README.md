
# Hash JSON Ids

Contents:

* [Explanation](#Explanation)
* [Examples](#examples)
* [Notes](#notes)

## Explanation
To prevent competitor to knowledge the number of clients that the company currently has is a common practice to shadow the `ID`s by using for example a Hash.

## Examples
❌ Incorrect
```go
type post struct {
	ID      uint
	Title   string
	Content string
}
```

✅ Correct
```go
type post struct {
	ID      string
	Title   string
	Content string
}
```

## Notes