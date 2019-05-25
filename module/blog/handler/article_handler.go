package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/rema424/go-gae-blog-app-example/service"
)

type (
	// ArticleHandler ...
	ArticleHandler interface {
		Show(c echo.Context) error
	}

	articleHandlerImpl struct {
		articleService service.ArticleService
	}
)

// NewArticleHandler ...
func NewArticleHandler(articleService service.ArticleService) ArticleHandler {
	return &articleHandlerImpl{
		articleService: articleService,
	}
}

func (h *articleHandlerImpl) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article := h.articleService.Get(id)
	return c.JSON(http.StatusOK, article)
}
