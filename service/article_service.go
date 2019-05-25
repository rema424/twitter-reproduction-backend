package service

import (
	"github.com/rema424/go-gae-blog-app-example/model"
	"github.com/rema424/go-gae-blog-app-example/repository"
)

type (
	// ArticleService ...
	ArticleService interface {
		Get(ID int) *model.Article
	}

	articleServiceImpl struct {
		Urepo repository.UserRepository
		Arepo repository.ArticleRepository
	}
)

// NewArticleService ...
func NewArticleService(urepo repository.UserRepository, arepo repository.ArticleRepository) ArticleService {
	return &articleServiceImpl{
		Urepo: urepo,
		Arepo: arepo,
	}
}

func (s *articleServiceImpl) Get(ID int) *model.Article {
	article := s.Arepo.Get(ID)
	user := s.Urepo.Get(article.AuthorID)
	article.Author = user

	return article
}
