package handlers

import (
	"fmt"
	"net/http"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/model"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/internal/httputils"
	"github.com/gin-gonic/gin"
)

// Campaign returns a campaign handler
// @Summary Get a single campaigns for the visitor
// @Tags Campaigns
// @Description Get a single campaign value and metadata for a visitor ID and context
// @ID get-campaign
// @Accept  json
// @Produce  json
// @Param id path string true "Campaign ID"
// @Param request body campaignsBodySwagger true "Campaign request body"
// @Success 200 {object} model.Campaign
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /campaigns/{id} [post]
func Campaign(fsClient *client.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		visitor := c.MustGet("visitor").(*client.Visitor)

		decisionResponse := visitor.GetDecisionResponse()

		response := model.Campaign{}
		found := false
		cID := c.Param("id")
		for _, c := range decisionResponse.Campaigns {
			if c.ID == cID || c.CustomID == cID {
				found = true
				response = c
				break
			}
		}
		if !found {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("The campaign %s is paused or doesn’t exist. Verify your customId or campaignId.", cID),
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
