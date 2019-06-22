// +build appengine

package echoutil

import (
	"github.com/labstack/echo"
	"google.golang.org/appengine/log"
)

func Debugf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}

	log.Debugf(GetContext(c), format, args...)
}

func Infof(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}
	log.Infof(GetContext(c), format, args...)
}

func Warningf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}
	log.Warningf(GetContext(c), format, args...)
}

func Errorf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}
	// postSlack(c, fmt.Sprintf(format, args...))
	log.Errorf(GetContext(c), format, args...)
}
