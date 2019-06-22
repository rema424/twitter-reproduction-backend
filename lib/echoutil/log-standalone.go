// +build !appengine

package echoutil

import (
	"log"

	"github.com/labstack/echo"
)

// Debugf ...
func Debugf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}

	log.Printf(format, args...)
}

// Infof ...
func Infof(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}

	log.Printf(format, args...)
}

// Warningf ...
func Warningf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}

	log.Printf(format, args...)
}

// Errorf ...
func Errorf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}

	// postSlack(c, fmt.Sprintf(format, args...))
	log.Printf(format, args...)
}
