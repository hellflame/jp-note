package api

import (
	"embed"
	"mime"
	"path"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func bindPages(pages *embed.FS) {
	router.GET("/", func(c echo.Context) error {
		content, e := pages.ReadFile("pages/index.html")
		if e != nil {
			return echo.NewHTTPError(404, e)
		}
		return c.HTMLBlob(200, content)
	})

	router.GET("/*", func(c echo.Context) error {
		target := path.Join("pages", c.Request().URL.Path[1:])
		f, e := pages.Open(target)
		if e != nil {
			return echo.NewHTTPError(404, e)
		}
		defer f.Close()
		contentType := mime.TypeByExtension(filepath.Ext(target))
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		return c.Stream(200, contentType, f)
	})
}
