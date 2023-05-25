package webview

import (
	"embed"
	"io/fs"
	"net/http"
	"os/exec"
	"strings"

	"wgm/apis"

	"github.com/gin-gonic/gin"
)

//go:embed assets
var build embed.FS

// SPAIndex 载入静态文件
func SPAIndex(build embed.FS) fs.FS {
	fsys := fs.FS(build)
	buildStatic, _ := fs.Sub(fsys, "assets")
	return buildStatic
}

// StaticHand 静态文件
func StaticHand(static fs.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		upath := c.Request.URL.Path
		if !strings.HasPrefix(upath, "/api/v1") {
			http.FileServer(http.FS(static)).ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func GenStatic(build embed.FS) gin.HandlerFunc {
	index := SPAIndex(build)
	return StaticHand(index)
}

// OpenWeb 调用系统浏览器
func OpenWeb(url string) {
	exec.Command(`cmd`, `/c`, `start`, `http://`+url).Start()
}

func Run() error {
	listenAddr := "127.0.0.1:30000"
	go OpenWeb(listenAddr)
	return apis.Run(listenAddr, GenStatic(build))
}
