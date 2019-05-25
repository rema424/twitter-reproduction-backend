package repository

import (
	"github.com/rema424/twitter-reproduction-backend/database"
	"github.com/rema424/twitter-reproduction-backend/model"
)

type (
	// ArticleRepository ...
	ArticleRepository interface {
		Get(ID int) *model.Article
	}

	articleRepositoryImpl struct {
		DB database.DB
	}
)

// NewArticleRepository ...
func NewArticleRepository(DB database.DB) ArticleRepository {
	return &articleRepositoryImpl{
		DB: DB,
	}
}

func (r *articleRepositoryImpl) Get(ID int) *model.Article {
	return r.DB.Article()[ID]
}
