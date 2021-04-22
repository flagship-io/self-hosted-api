package main

import (
	"fmt"
	"strings"

	"github.com/flagship-io/self-hosted-api/pkg/handlers"
	"github.com/flagship-io/self-hosted-api/pkg/httputils"
	"github.com/flagship-io/self-hosted-api/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/flagship-io/self-hosted-api/docs"
)

// @title Flagship Decision Host
// @version 2.0
// @BasePath /v2
// @description This is the Flagship Decision Host API documentation

// @contact.name API Support
// @contact.url https://www.abtasty.com/solutions-product-teams/
// @contact.email support@flagship.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	r := gin.Default()

	// Init logger with default Warn level
	log.InitLogger(logrus.WarnLevel)

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.GetLogger().Warnf("Could not find config file: %v", err)
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.SetDefault("port", 8080)
	port := viper.GetInt("port")

	fsClient, err := initFsClient()

	if err != nil {
		log.GetLogger().Panicf("Error when initializing Flagship: %v", err)
	}

	r.GET("/v2/health", handlers.Health(fsClient))
	r.POST("/v2/campaigns", handlers.CampaignMiddleware(fsClient), handlers.Campaigns(fsClient))
	r.POST("/v2/campaigns/:id", handlers.CampaignMiddleware(fsClient), handlers.Campaign(fsClient))
	r.POST("/v2/activate", handlers.Activate(fsClient))
	r.POST("/v2/flags", handlers.CampaignMiddleware(fsClient), handlers.Flags(fsClient))
	r.POST("/v2/flags/:key", handlers.CampaignMiddleware(fsClient), handlers.Flag(fsClient))
	r.POST("/v2/flags/:key/value", handlers.CampaignMiddleware(fsClient), handlers.FlagValue(fsClient))
	r.POST("/v2/flags/:key/activate", handlers.FlagActivate(fsClient, fsClient.GetCacheManager() != nil))
	r.Any("/v2/hits/*proxyPath", httputils.Proxy("https://ariane.abtasty.com"))

	url := ginSwagger.URL("/v2/swagger/doc.json") // The url pointing to API definition
	r.GET("/v2/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err = r.Run(fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		panic(err)
	}
}
