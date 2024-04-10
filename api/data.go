package api

import (
	"embed"
	"mime"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func bindData(data *embed.FS) {
	router.GET("/data/*", func(c echo.Context) error {
		target := c.Request().URL.Path[1:]
		f, e := data.Open(target)
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
