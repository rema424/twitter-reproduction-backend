package echoutil

import (
	"net/http"

	"github.com/rema424/twitter-reproduction-backend/lib/tmplutil"

	"github.com/labstack/echo"
)

// Render ...
func Render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := tmplutil.HTMLBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
