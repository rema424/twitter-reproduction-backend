package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rema424/twitter-reproduction-backend/service"
)

// UserHandler ...
type UserHandler interface {
	Index(c echo.Context) error
	Show(c echo.Context) error
}

type userHandler struct {
	userService service.UserService
}

// NewUserHandler ...
func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c echo.Context) error {
	us, err := h.userService.SelectAll()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, "データが見つかりませんでした。")
		}
		return c.JSON(http.StatusOK, "予期せぬエラーが発生しました。")
	}
	return c.JSONPretty(http.StatusOK, us, "  ")
}

func (h *userHandler) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := h.userService.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, "データが見つかりませんでした。")
		}
		return c.JSON(http.StatusOK, "予期せぬエラーが発生しました。")
	}
	return c.JSONPretty(http.StatusOK, u, "  ")
}
