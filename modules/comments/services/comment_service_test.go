package services_test

import (
	"context"
	"hoge/modules/comments/repositories"
	"hoge/modules/comments/services"
	"hoge/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

const dbPath = "../../../test.db"

func TestCommentService(t *testing.T) {
	db := pkg.GormSqliteSetup(dbPath)

	t.Run("successfull", func(t *testing.T) {
		commentsRepo := repositories.NewCommentsMysqlRepository(db)
		res := services.NewCommentsService(commentsRepo)

		assert.NotNil(t, res)
	})

	t.Run("avoid NilPointer", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		services.NewCommentsService(nil)
	})

}

func TestGetComments(t *testing.T) {
	db := pkg.GormSqliteSetup(dbPath)

	commentsRepo := repositories.NewCommentsMysqlRepository(db)
	srv := services.NewCommentsService(commentsRepo)

	res, err := srv.GetComments(context.TODO())

	assert.NotNil(t, res)
	assert.Nil(t, err)
}
