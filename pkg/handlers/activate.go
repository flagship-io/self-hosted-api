package handlers

import (
	"net/http"
	"time"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/model"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/gin-gonic/gin"
)

type activateBody struct {
	VisitorID        string `json:"vid" binding:"required"`
	CampaignID       string `json:"cid" binding:"required"`
	VariationGroupID string `json:"caid" binding:"required"`
	VariationID      string `json:"vaid" binding:"required"`
}

// Activate returns a flag activation handler
// @Summary Activate a campaign
// @Tags v2
// @Description Activate a campaign for a visitor ID
// @ID activate
// @Accept  json
// @Produce  json
// @Param request body activateBody true "Campaign activation request body"
// @Success 204
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /v2/activate [post]
func Activate(fsClient *client.Client) func(*gin.Context) {
	return func(c *gin.Context) {

		aObj := &activateBody{}
		err := c.BindJSON(aObj)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = fsClient.SendHit(aObj.VisitorID, &model.ActivationHit{
			VisitorID:        aObj.VisitorID,
			EnvironmentID:    fsClient.GetEnvID(),
			VariationGroupID: aObj.VariationGroupID,
			VariationID:      aObj.VariationID,
			CreatedAt:        time.Now(),
			QueueTime:        0,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Data(204, gin.MIMEJSON, nil)
	}
}
