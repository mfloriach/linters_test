package interfaces

import (
	"errors"
	"hoge/modules/comments/serializers"
	"hoge/modules/comments/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentsController struct {
	commentsService services.CommentsServiceInterface
}

// NewCommentsController  constructor
func NewCommentsController(router *gin.Engine, commentsService services.CommentsServiceInterface) error {
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

func (p commentsController) GetComments(c *gin.Context) {
	comments := p.commentsService.GetComments(c)

	c.JSON(http.StatusOK, gin.H{
		"data": serializers.CommentsSerialize(comments),
	})
}
