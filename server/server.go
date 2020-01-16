package server

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ijkbytes/mega/base/config"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/server/api"
	"github.com/ijkbytes/mega/server/middleware"
	"github.com/ijkbytes/mega/server/pages"
	"go.uber.org/zap"
	"net/http"
	"runtime"
)

var logger *zap.Logger

func serveException(ctx *gin.Context) {
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
				"code": http.StatusInternalServerError,
				"msg":  "Unknow Exception",
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

	store := cookie.NewStore([]byte(config.Mega.Session.Secret))
	router.Use(sessions.Sessions(config.Mega.Session.Name, store))

	router.LoadHTMLGlob("theme/simple/*.html")
	router.Static("/static/simple/css", "theme/simple/css")

	pagesGroup := router.Group("/")
	{
		pagesGroup.GET("/", pages.Index)
		pagesGroup.GET("/articles/:id", pages.Article)
	}

	apiGroup := router.Group("/api")
	{
		group1 := apiGroup.Group("/")
		{
			group1.GET("login", api.Login)
		}
		group2 := apiGroup.Group("/", middleware.NeedLogin)
		{
			group2.GET("/tags", api.GetTags)
			group2.POST("/tags", api.AddTag)
			group2.PUT("/tags/:id", api.EditTag)
			group2.DELETE("/tags/:id", api.DeleteTag)

			group2.GET("/articles", api.GetArticles)
			group2.GET("/articles/:id", api.GetArticle)
			group2.POST("/articles", api.AddArticle)
			group2.PUT("/articles/:id", api.EditArticle)
			group2.DELETE("/articles/:id", api.DeleteArticle)
		}
	}

	return router
}
