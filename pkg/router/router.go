package router

import (
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/flagship-io/self-hosted-api/pkg/handlers"
	"github.com/flagship-io/self-hosted-api/pkg/httputils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(fsClient *client.Client) *gin.Engine {
	r := gin.Default()

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

	return r
}
