
# Modularization

Contents:

* [Explanation](#Explanation)
* [Examples](#examples)
* [Notes](#notes)

## Explanation
For large base code,  divide code by domain (modules) is common practices to allow teams to work independently without friction at the same time it allows to changes in one module don't effect others. To get all benefits is necessary:
- prevent coupling between modules
- use facade design pattern to prevent developers use internal code inside the modules, developers should only use public function/classes exposed by the module.

## Examples
❌ Incorrect
```go
package main

import (
	comments "hoge/modules/comments/repositories"
)

func main() {
	comments.GetComments()
}
```

✅ Correct
```go
package main

import (
	"hoge/modules/comments"
)

func main() {
	comments.GetComments()
}
```