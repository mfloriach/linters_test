package interfaces

import (
	"hoge/modules/comments/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentControllerInterface interface {
	GetComments(c *gin.Context)
}

type commentsController struct {
	commentsService services.CommentsServiceInterface
}

func NewCommentsController(router *gin.Engine, commentsService services.CommentsServiceInterface) {
	ctr := commentsController{
		commentsService: commentsService,
	}

	router.GET("/comments", ctr.GetComments)
}

func (p commentsController) GetComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": p.commentsService.GetComments(c),
	})
}
