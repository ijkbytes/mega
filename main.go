package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ijkbytes/mega/base/config"
	"github.com/ijkbytes/mega/base/log"
	"github.com/ijkbytes/mega/server"
	"github.com/ijkbytes/mega/service"
	"go.uber.org/zap"
	"net/http"
)

type Application struct {
	logger *zap.Logger
	router *gin.Engine
}

func (app *Application) run() error {
	addr := fmt.Sprintf(":%d", config.Mega.Http.Port)
	app.logger.Sugar().Info("Listening at: ", addr)
	err := http.ListenAndServe(addr, app.router)
	if err != nil {
		app.logger.Sugar().Panic("run http server error: ", err)
		return err
	}
	return nil
}

func main() {
	defer log.Sync()
	service.ConnectDB()
	defer service.DisconnectDB()

	app := &Application{
		logger: log.Get("Application"),
		router: server.GetRouter(),
	}

	app.run()
}
