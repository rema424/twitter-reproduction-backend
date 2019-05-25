package database

import (
	"github.com/rema424/go-gae-blog-app-example/model"
)

type (
	// DB ...
	DB interface {
		Article() map[int]*model.Article
		User() map[int]*model.User
	}

	// DB ...
	dBImpl struct {
		Users    map[int]*model.User
		Articles map[int]*model.Article
	}
)

// NewDatabase ...
func NewDatabase() DB {

	user1 := &model.User{ID: 1, Name: "池田"}
	user2 := &model.User{ID: 2, Name: "工藤"}

	users := map[int]*model.User{
		1: user1,
		2: user2,
	}

	article1 := &model.Article{ID: 1, Title: "朝", Body: "おはよう", AuthorID: 1}
	article2 := &model.Article{ID: 2, Title: "昼", Body: "こんにちは", AuthorID: 2}
	article3 := &model.Article{ID: 3, Title: "晩", Body: "こんばんは", AuthorID: 2}

	articles := map[int]*model.Article{
		1: article1,
		2: article2,
		3: article3,
	}

	return &dBImpl{
		Users:    users,
		Articles: articles,
	}
}

func (d *dBImpl) Article() map[int]*model.Article {
	return d.Articles
}

func (d *dBImpl) User() map[int]*model.User {
	return d.Users
}
