package comments

import (
	"hoge/modules/comments/interfaces"
	"hoge/modules/comments/repositories"
	"hoge/modules/comments/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Allow connection to this module via gin restfull API
func GinServer(r *gin.Engine, db gorm.DB) {
	commentService := dependencyInjection(db)
	if err := interfaces.NewCommentsController(r, commentService); err != nil {
		panic(err)
	}
}

// Allow connection to this module via grpc protobuffers
func GrpcServer() {
}

func dependencyInjection(db gorm.DB) services.CommentsServiceInterface {
	commentsRepo := repositories.NewCommentsMysqlRepository(db)
	return services.NewCommentsService(commentsRepo)
}
