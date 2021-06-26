package deps

import "github.com/labstack/echo"

// EchoAuthMiddleware func
func EchoAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request accepted")
		return next(c)
	}
}
