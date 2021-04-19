package main

import (
	"fmt"
	"strings"

	"github.com/flagship-io/self-hosted-api/pkg/handlers"
	"github.com/flagship-io/self-hosted-api/pkg/httputils"
	"github.com/flagship-io/self-hosted-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/flagship-io/self-hosted-api/docs"
)

var mainLogger = logger.GetLogger("main")

// @title Flagship Decision Host
// @version 1.0
// @description This is the Flagship Decision Host API documentation

// @contact.name API Support
// @contact.url https://www.abtasty.com/solutions-product-teams/
// @contact.email support@flagship.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	r := gin.Default()

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		mainLogger.Warnf("Could not find config file: %v", err)
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.SetDefault("port", 8080)
	port := viper.GetInt("port")

	fsClient, err := initFsClient()

	if err != nil {
		mainLogger.Panicf("Error when initializing Flagship: %v", err)
	}

	r.POST("/v2/campaigns", handlers.Campaigns(fsClient))
	r.POST("/v2/campaigns/:id", handlers.Campaign(fsClient))
	r.POST("/v2/activate", handlers.Activate(fsClient))
	r.POST("/v2/flags", handlers.Flags(fsClient))
	r.POST("/v2/flags/:key", handlers.Flag(fsClient))
	r.POST("/v2/flags/:key/value", handlers.FlagValue(fsClient))
	r.POST("/v2/flags/:key/activate", handlers.FlagActivate(fsClient, fsClient.GetCacheManager() != nil))
	r.Any("/v2/hits/*proxyPath", httputils.Proxy("https://ariane.abtasty.com"))

	url := ginSwagger.URL("/v2/swagger/doc.json") // The url pointing to API definition
	r.GET("/v2/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(fmt.Sprintf("0.0.0.0:%v", port))
}
