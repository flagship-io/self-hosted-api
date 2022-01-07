package handlers

import (
	"fmt"
	"net/http"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/model"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/internal/httputils"
	"github.com/gin-gonic/gin"
)

// FlagValue returns a flag value handler
// @Summary Get flag value by name
// @Tags Flags
// @Description Get a single flag value for a visitor ID and context
// @ID get-flag-value
// @Accept  json
// @Produce  json
// @Param key path string true "Flag key"
// @Param request body campaignsBodySwagger true "Flag request body"
// @Success 200 {object} interface{}
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /flags/{key}/value [post]
func FlagValue(fsClient *client.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		modifications := c.MustGet("modifications").(map[string]model.FlagInfos)
		flag, ok := modifications[c.Param("key")]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Flag %s not found", c.Param("key")),
			})
			return
		}
		c.JSON(http.StatusOK, flag.Value)
	}
}
