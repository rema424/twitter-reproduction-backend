package echoutil

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SetCommonMiddleware ...
func SetCommonMiddleware(e *echo.Echo) {
	// Recover
	e.Use(Recover)

	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// CORS
	e.Use(middleware.CORS())

	// Trailing Slash
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	// Static
	e.Static("/static", "asset")
}
