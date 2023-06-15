package router

import (
	"fmt"
	"github.com/NubeIO/nrule/config"
	"github.com/NubeIO/nrule/logger"
	"github.com/NubeIO/nrule/server/constants"
	"github.com/NubeIO/nrule/server/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"time"
)

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := fmt.Sprintf("%s %s [%d]: %s", ctx.Request.Method, ctx.Request.URL, 404, "rubix-edge-bios: api not found")
		ctx.JSON(http.StatusNotFound, controller.Message{Message: message})
	}
}

func Setup() *gin.Engine {
	engine := gin.New()
	// Set gin access logs
	if viper.GetBool("gin.log.store") {
		fileLocation := fmt.Sprintf("%s/rubix-edge-wires.access.log", config.Config.GetAbsDataDir())
		f, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, constants.Permission)
		if err != nil {
			logger.Logger.Errorf("Failed to create access log file: %v", err)
		} else {
			gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		}
	}
	gin.SetMode(viper.GetString("gin.log.level"))
	engine.NoRoute(NotFound())
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders: []string{
			"X-FLOW-Key", "Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
		},
		ExposeHeaders:          []string{"Content-Length"},
		AllowCredentials:       true,
		AllowAllOrigins:        true,
		AllowBrowserExtensions: true,
		MaxAge:                 12 * time.Hour,
	}))

	api := controller.Controller{}

	apiRoutes := engine.Group("/api")

	ping := apiRoutes.Group("/ping")
	{
		ping.GET("", api.Ping)

	}

	return engine
}
