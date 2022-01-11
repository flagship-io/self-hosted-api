package handlers

import (
	"net/http"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/model"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/internal/httputils"
	"github.com/gin-gonic/gin"
)

//nolint
type campaignsBodyContextSwagger struct {
	KeyString string  `json:"key_string"`
	KeyNumber float64 `json:"key_number"`
	KeyBool   bool    `json:"key_bool"`
}

//nolint
type campaignsBodySwagger struct {
	VisitorID   string                      `json:"visitor_id" binding:"required"`
	AnonymousID *string                     `json:"anonymous_id"`
	Context     campaignsBodyContextSwagger `json:"context"`
	TriggerHit  bool                        `json:"trigger_hit"`
}

type campaignsBody struct {
	VisitorID   string                 `json:"visitor_id" binding:"required"`
	AnonymousID *string                `json:"anonymous_id"`
	Context     map[string]interface{} `json:"context"`
	TriggerHit  *bool                  `json:"trigger_hit"`
}

// CampaignsResponse represents the campaigns call response
type CampaignsResponse struct {
	VisitorID string           `json:"visitor_id"`
	Panic     bool             `json:"panic"`
	Campaigns []model.Campaign `json:"campaigns"`
}

// Campaigns returns a campaigns handler
// @Summary Get all campaigns for the visitor
// @Tags Campaigns
// @Description Get all campaigns value and metadata for a visitor ID and context
// @ID get-campaigns
// @Accept  json
// @Produce  json
// @Param request body campaignsBodySwagger true "Campaigns request body"
// @Success 200 {object} CampaignsResponse
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /campaigns [post]
func Campaigns(fsClient *client.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		modifications, _ := c.MustGet("modifications").(map[string]model.FlagInfos)
		visitor := c.MustGet("visitor").(*client.Visitor)

		response := CampaignsResponse{
			VisitorID: visitor.ID,
			Campaigns: []model.Campaign{},
		}

		campaignsMap := map[string]bool{}

		for _, m := range modifications {
			_, exists := campaignsMap[m.Campaign.ID]
			if !exists {
				response.Campaigns = append(response.Campaigns, m.Campaign)
				campaignsMap[m.Campaign.ID] = true
			}
		}
		c.JSON(http.StatusOK, response)
	}
}
