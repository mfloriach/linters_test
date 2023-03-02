package main

import (
	"hoge/modules/comments"
	"hoge/modules/posts"
	"hoge/pkg"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const dbPath = "test.db"

var modes = map[string]func(db gorm.DB){
	"gin":  ginServer,
	"grpc": grpcServer,
}

func main() {
	db := pkg.GormSqliteSetup(dbPath)

	if startServer, ok := modes["gin"]; ok {
		startServer(db)
	} else {
		panic("this mode does not exist")
	}
}

func ginServer(db gorm.DB) {
	r := gin.Default()

	posts.GinServer(r, db)
	comments.GinServer(r, db)

	r.Run()
}

func grpcServer(db gorm.DB) {

}
