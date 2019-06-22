package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/rema424/twitter-reproduction-backend/lib/echoutil"

	// "google.golang.org/appengine/memcache"

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
	us, err := h.userService.SelectAll(echoutil.GetContext(c))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, "データが見つかりませんでした。")
		}
		return c.JSON(http.StatusOK, "予期せぬエラーが発生しました。")
	}
	return c.JSONPretty(http.StatusOK, us, "  ")
}

func (h *userHandler) Show(c echo.Context) error {
	ctx := echoutil.GetContext(c)
	id, _ := strconv.Atoi(c.Param("id"))
	u, err := h.userService.Get(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, "データが見つかりませんでした。")
		}
		return c.JSON(http.StatusOK, "予期せぬエラーが発生しました。")
	}

	// eg, ctx := errgroup.WithContext(echoutil.GetContext(c))
	// var u1, u2, u3 *model.User

	// eg.Go(func() (err error) {
	// 	u1, err = h.userService.Get(ctx, 11)
	// 	return
	// })

	// eg.Go(func() (err error) {
	// 	u2, err = h.userService.Get(ctx, 2)
	// 	return
	// })

	// eg.Go(func() (err error) {
	// 	u3, err = h.userService.Get(ctx, 3)
	// 	return
	// })

	// if err := eg.Wait(); err != nil {
	// 	log.Errorf(ctx, "error: %s", err)
	// }

	// data := map[string]interface{}{
	// 	"u1": u1,
	// 	"u2": u2,
	// 	"u3": u3,
	// }

	return c.JSONPretty(http.StatusOK, u, "  ")
	// return c.Render(http.StatusOK, "hello", data)
}
