package rest

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFile("/main.js", "/usr/share/pbi-auction/assets/main.js")
	r.NoRoute(func(ctx *gin.Context) {
		ctx.File("/usr/share/pbi-auction/assets/index.html")
	})
	return r
}
