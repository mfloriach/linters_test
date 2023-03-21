package services_test

import (
	"context"
	"hoge/modules/posts/repositories"
	"hoge/modules/posts/services"
	"hoge/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

const dbPath = "../../../test.db"

func TestCommentService(t *testing.T) {
	db := pkg.GormSqliteSetup(dbPath)

	t.Run("successfull", func(t *testing.T) {
		commentsRepo := repositories.NewPostMysqlRepository(db)
		res := services.NewPostService(commentsRepo)

		assert.NotNil(t, res)
	})

	t.Run("avoid NilPointer", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		services.NewPostService(nil)
	})

}

func TestGetComments(t *testing.T) {
	db := pkg.GormSqliteSetup(dbPath)

	commentsRepo := repositories.NewPostMysqlRepository(db)
	srv := services.NewPostService(commentsRepo)

	res, err := srv.GetPosts(context.TODO())

	assert.NotNil(t, res)
	assert.Nil(t, err)
}
