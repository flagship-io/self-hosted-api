package router

import (
	"fmt"
	"net/http"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/flagship-io/self-hosted-api/internal/handlers"
	"github.com/flagship-io/self-hosted-api/internal/httputils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(fsClient *client.Client) *gin.Engine {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/v2/health", handlers.Health(fsClient))
	r.POST("/v2/campaigns", handlers.CampaignMiddleware(fsClient), handlers.Campaigns(fsClient))
	r.POST("/v2/campaigns/:id", handlers.CampaignMiddleware(fsClient), handlers.Campaign(fsClient))
	r.POST("/v2/activate", handlers.Activate(fsClient))
	r.POST("/v2/flags", handlers.CampaignMiddleware(fsClient), handlers.Flags(fsClient))
	r.POST("/v2/flags/:key", handlers.CampaignMiddleware(fsClient), handlers.Flag(fsClient))
	r.POST("/v2/flags/:key/value", handlers.CampaignMiddleware(fsClient), handlers.FlagValue(fsClient))
	r.POST("/v2/flags/:key/activate", handlers.FlagActivate(fsClient, fsClient.GetCacheManager() != nil))
	r.POST("/v2/hits", httputils.Proxy("https://ariane.abtasty.com"))

	url := ginSwagger.URL("/v2/swagger/doc.json") // The url pointing to API definition
	r.GET("/v2/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
