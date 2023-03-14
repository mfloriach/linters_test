package interfaces

import (
	"errors"
	"hoge/modules/comments/serializers"
	"hoge/modules/comments/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentsController struct {
	commentsService shared.CommentsServiceInterface
}

// NewCommentsController  constructor
func NewCommentsController(router *gin.Engine, commentsService shared.CommentsServiceInterface) error {
	if commentsService == nil {
		return errors.New("commentsService instance can not be null")
	}

	if router == nil {
		return errors.New("router instance can not be null")
	}

	ctr := commentsController{
		commentsService: commentsService,
	}

	router.GET("/comments", ctr.GetComments)

	return nil
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
func (p commentsController) GetComments(c *gin.Context) {
	comments := p.commentsService.GetComments(c)

	c.JSON(http.StatusOK, gin.H{
		"data": serializers.CommentsSerialize(comments),
	})
}
