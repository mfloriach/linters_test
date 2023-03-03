
# Avoid nil dependency injection

Contents:

* [Explanation](#Explanation)
* [Examples](#examples)
* [Notes](#notes)

## Explanation
In order to improve testability is necessary to use dependency injections (DI) of the modules used for that object. To mock that dependencies is necessary to use interfaces.

Interfaces can be `nil` provoking errors on runtime, that are dificult find.

Therefore this linter helps to allways check dependencies injections to avoid this problem

## Examples
```go

func NewObject(serviceA ServiceAInterface) {
    if serviceA == nil {
        panic("serviceA dependency can not be nil")
    }

    ...
}
```

## Notes
