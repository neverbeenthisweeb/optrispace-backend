package web

import (
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/labstack/echo/v4"
	"optrispace.com/work/pkg/clog"
	"optrispace.com/work/pkg/service"
)

const (
	anyMethod = "*"
	// Realm for Basic Auth header
	Realm = "Optrispace"
)

// exceptions endpoints
var exceptions = [][2]string{
	{anyMethod, "/login"},
	{anyMethod, "/signup"},
	{anyMethod, "/stop"},
	{anyMethod, "/info"},
	{http.MethodGet, "/jobs"},
	{http.MethodGet, "/jobs/*"},
	{anyMethod, "/notifications"},
}

func authSkip(c echo.Context) bool {
	for _, exc := range exceptions {
		method := c.Request().Method
		if method == http.MethodOptions { // preflight always enabled
			return true
		}

		match, err := path.Match(exc[1], c.Request().RequestURI)
		if err != nil {
			panic(fmt.Errorf("invalid pattern %s: %w", exc, err))
		}

		if match && (exc[0] == anyMethod || exc[0] == method) {
			return true
		}

	}
	return false
}

// Auth checks if there are user authenticated
func Auth(securitySvc service.Security) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if authSkip(c) {
				goto further
			}

			if _, e := securitySvc.FromEchoContext(c); e != nil {
				clog.Ectx(c).Warn().Err(e).Msg("Unable to authenticate user against bearer")
			} else {
				goto further
			}

			if _, e := securitySvc.FromEchoContextByBasicAuth(c, Realm); e != nil {
				clog.Ectx(c).Warn().Err(e).Msg("Unable to authenticate user against basic auth")
			} else {
				goto further
			}

			c.Response().Header().Set(echo.HeaderWWWAuthenticate, "Basic realm="+strconv.Quote(Realm))
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization required")

		further:
			return next(c)
		}
	}
}
