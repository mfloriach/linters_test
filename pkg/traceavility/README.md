
# Traceavility

Contents:

* [Explanation](#Explanation)
* [Examples](#examples)
* [Notes](#notes)

## Explanation
Propagating the context allow to use log aggregation and improve readability and tracebility of the request event, for example by using `request-id`.

But is necessary to use pass the context from the server across the services are injected to third parties (database, restfull API, ...).

This linters helps to not forget to pass the context and homoginaze the format.

## Examples
```go
func (u commentsService) GetComments(ctx context.Context) string {
	...
}
```

## Notes
