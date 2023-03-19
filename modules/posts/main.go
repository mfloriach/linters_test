package posts

import (
	"hoge/modules/posts/interfaces"
	"hoge/modules/posts/repositories"
	"hoge/modules/posts/services"
	"hoge/modules/posts/shared"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Allow conection to this module via gin restfull API
func GinServer(r *gin.Engine, db gorm.DB) {
	postService := dependencyInjection(db)
	interfaces.NewPostController(r, postService)
}

// Allow connection to this module via grpc protobuffers
func GrpcServer() {
}

func dependencyInjection(db gorm.DB) shared.PostServiceInterface {
	postRepo := repositories.NewPostMysqlRepository(db)
	postService := services.NewPostService(postRepo)

	return postService
}
