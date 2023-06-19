package router

import (
	"fmt"
	"github.com/NubeIO/nrule/apirules"
	"github.com/NubeIO/nrule/config"
	"github.com/NubeIO/nrule/logger"
	"github.com/NubeIO/nrule/rules"
	"github.com/NubeIO/nrule/server/constants"
	"github.com/NubeIO/nrule/server/controller"
	"github.com/NubeIO/nrule/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"time"
)

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := fmt.Sprintf("%s %s [%d]: %s", ctx.Request.Method, ctx.Request.URL, 404, "api not found")
		ctx.JSON(http.StatusNotFound, controller.Message{Message: message})
	}
}

func Setup() *gin.Engine {
	engine := gin.New()
	// Set gin access logs
	if viper.GetBool("gin.log.store") {
		fileLocation := fmt.Sprintf("%s/rubix-rules.access.log", config.Config.GetAbsDataDir())
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

	eng := rules.NewRuleEngine()
	err := eng.Start()
	if err != nil {
		logrus.Error(err)
	}

	name := "Core"
	props := make(rules.PropertiesMap)
	props[name] = eng

	client := "Client"
	newClient := &apirules.Client{}
	props[client] = newClient
	logger.Logger.Infof("new db on location:%s", config.Config.GetAbsDatabaseFile())
	api := controller.Controller{
		Rules:   eng,
		Client:  newClient,
		Props:   props,
		Storage: storage.New(config.Config.GetAbsDatabaseFile()),
	}
	apiRoutes := engine.Group("/api")

	ping := apiRoutes.Group("/ping")
	{
		ping.GET("", api.Ping)
	}

	rule := apiRoutes.Group("/rules")
	{
		rule.GET("", api.SelectAllRules)
		rule.GET("/:uuid", api.SelectRule)
		rule.PATCH("/:uuid", api.UpdateRule)
		rule.DELETE("/:uuid", api.DeleteRule)
		rule.POST("", api.AddRule)
	}

	rulesRun := apiRoutes.Group("/rules/dry")
	{
		rulesRun.POST("", api.Dry)
	}

	return engine
}
