package echoutil

import (
	"github.com/labstack/echo"
	"github.com/rema424/twitter-reproduction-backend/lib/logutil"
)

// Debugf ...
func Debugf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}

	logutil.Debugf(GetContext(c), format, args...)
}

// Infof ...
func Infof(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}
	logutil.Infof(GetContext(c), format, args...)
}

// Warningf ...
func Warningf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}
	logutil.Warningf(GetContext(c), format, args...)
}

// Errorf ...
func Errorf(c echo.Context, format string, args ...interface{}) {
	if c == nil {
		return
	}
	// postSlack(c, fmt.Sprintf(format, args...))
	logutil.Errorf(GetContext(c), format, args...)
}
