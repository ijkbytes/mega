package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/config"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/server/pages"
	"go.uber.org/zap"
	"net/http"
	"runtime"
)

var logger *zap.Logger

func serveException(ctx *gin.Context)  {
	defer func(ctx *gin.Context) {
		if r := recover(); r != nil {
			pc := make([]uintptr, 10)
			length := runtime.Callers(2, pc)
			stack := ""
			for i := 0; i < length; i++ {
				f := runtime.FuncForPC(pc[i])
				file, line := f.FileLine(pc[i])
				stack = stack + fmt.Sprintf("%s\n\t%s:%d\n", f.Name(), file, line)
			}

			logger.Error(
				fmt.Sprint(r, "\n", stack),
			)

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"result": "error",
				"msg":    "Unknow Exception",
			})
		}
	}(ctx)

	ctx.Next()
}

func newRouter() *gin.Engine {
	if config.Mega.Http.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	engine.Use(serveException)

	return engine
}

func GetRouter() *gin.Engine {
	logger = log.Get("server")
	router := newRouter()

	router.LoadHTMLGlob("theme/sample/*")

	pagesGroup := router.Group("/")
	{
		pagesGroup.GET("/", pages.Index)
	}
	//
	//apiGroup := router.Group("/api")
	//{
	//	apiGroup.GET("articles")
	//}

	return router
}