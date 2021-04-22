package handlers

import (
	"fmt"
	"net/http"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/model"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/pkg/httputils"
	"github.com/gin-gonic/gin"
)

// Flag returns a flag handler
// @Summary Get flag by name
// @Tags Flags
// @Description Get a single flag value and metadata for a visitor ID and context
// @ID get-flag
// @Accept  json
// @Produce  json
// @Param key path string true "Flag key"
// @Param request body campaignsBodySwagger true "Flag request body"
// @Success 200 {object} FlagInfos{}
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /flags/{key} [post]
func Flag(fsClient *client.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		modifications := c.MustGet("modifications").(map[string]model.FlagInfos)
		m, ok := modifications[c.Param("key")]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Flag %s not found", c.Param("key")),
			})
			return
		}
		flag := FlagInfos{
			Value: m.Value,
			Metadata: FlagMetadata{
				CampaignID:       m.Campaign.ID,
				VariationGroupID: m.Campaign.VariationGroupID,
				VariationID:      m.Campaign.Variation.ID,
			},
		}
		c.JSON(http.StatusOK, flag)
	}
}
