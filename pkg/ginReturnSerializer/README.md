
# Force to return serializers

Contents:

* [Explanation](#Explanation)
* [Examples](#examples)
* [Notes](#notes)

## Explanation
Force endpoints to return serializer objects that require some data manipulation. For example in order to shadow the data IDs.

## Examples
```go
func (p commentsController) GetComments(c *gin.Context) {
	comments := p.commentsService.GetComments(c)

	c.JSON(http.StatusOK, gin.H{
		"data": serializers.CommentsSerialize(comments),
	})
}
```

## Notes