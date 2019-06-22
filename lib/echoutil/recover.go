package echoutil

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo"
)

// Recover ...
func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				var err error

				switch t := r.(type) {
				case error:
					err = t
				default:
					err = fmt.Errorf("%v", t)
				}

				stack := make([]byte, 4<<10) // 4KB
				length := runtime.Stack(stack, true)
				Errorf(c, "panic: %s %s\n", err, stack[:length])

				c.Error(err)
			}
		}()

		return next(c)
	}
}
