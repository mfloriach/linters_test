
# Allways document endpoints

Contents:

* [Explanation](#Explanation)
* [Examples](#examples)
* [Notes](#notes)

## Explanation
Documenting endpoints are not consider important for some developers so they tend to forget to documented, this linter will help to not forget it.

## Examples
```go
// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (p postController) GetPosts(c *gin.Context) {
	...
}
```

## Notes
This linter is not prepared to check for updates on the endpoint.