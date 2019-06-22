// +build !appengine

package echoutil

import (
	"context"

	"github.com/labstack/echo"
)

// GetContext ...
func GetContext(c echo.Context) context.Context {
	return c.Request().Context()
}
