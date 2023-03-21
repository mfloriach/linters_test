
# Traceavility

Contents:

* [Explanation](#Explanation)
* [Examples](#examples)
* [Notes](#notes)

## Explanation
Propagating the context allow to use log aggregation and improve readability and traceability of the request event, for example by using `request-id`.

But is necessary to use pass the context from the server across the services are injected to third parties (database, restfull API, ...).

This linters helps to not forget to pass the context and homoginaze the format.

## Examples
❌ Incorrect
```go
func (u commentsService) GetComments() string {
	...
}
```

✅ Correct
```go
func (u commentsService) GetComments(ctx context.Context) string {
	...
}
```

For queries to `Gorm`: 
❌ Incorrect
```go
func (u commentsService) GetComments(ctx context.Context) string {
	u.r.Find(1)
}
```

✅ Correct
```go
func (u commentsService) GetComments(ctx context.Context) string {
	u.r.WithContext(ctx).Find(1)
}
```

## Notes
