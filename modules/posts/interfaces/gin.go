package interfaces

import (
	"hoge/modules/posts/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postController struct {
	postService shared.PostServiceInterface
}

func NewPostController(router *gin.Engine, postService shared.PostServiceInterface) {
	if postService == nil {
		panic("please pass the postService")
	}

	ctr := postController{
		postService: postService,
	}

	router.GET("/posts", ctr.GetPosts)
}

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
	posts := p.postService.GetPosts(c)

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
