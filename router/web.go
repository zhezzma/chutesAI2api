package router

import (
	"chutesai2api/common"
	"chutesai2api/middleware"
	"embed"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetWebRouter(router *gin.Engine, buildFS embed.FS) {
	indexPageData, _ := buildFS.ReadFile(fmt.Sprintf("web/build/index.html"))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	//router.Use(middleware.GlobalWebRateLimit())
	router.Use(middleware.Cache())
	router.Use(static.Serve("/", common.EmbedFolder(buildFS, fmt.Sprintf("web/build"))))

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 处理 API 请求
		if strings.HasPrefix(path, "/v1") || strings.HasPrefix(path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "API endpoint not found",
				"path":  path,
				"code":  404,
			})
			return
		}

		// 处理静态资源请求
		if strings.Contains(path, ".") {
			// 可能是静态资源请求 (.js, .css, .png 等)
			c.Status(http.StatusNotFound)
			return
		}

		// 处理前端路由请求
		c.Header("Cache-Control", "no-cache")
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPageData)
	})
}
