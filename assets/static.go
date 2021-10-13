package assets

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed build
var build embed.FS

func SPAIndex() http.FileSystem {
	fsys := fs.FS(build)
	buildStatic, _ := fs.Sub(fsys, "build")
	return http.FS(buildStatic)
}

// StaticHand 静态文件
func StaticHand() gin.HandlerFunc {
	return func(c *gin.Context) {
		upath := c.Request.URL.Path
		if !strings.HasPrefix(upath, "/api") {
			content := SPAIndex()
			http.FileServer(content).ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
