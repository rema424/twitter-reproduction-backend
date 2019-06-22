// +build appengine

package echoutil

import (
	"context"

	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

// GetContext ...
func GetContext(c echo.Context) context.Context {
	return appengine.NewContext(c.Request())
}
