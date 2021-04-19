package handlers

import (
	"net/http"
	"sync"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/model"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/pkg/httputils"
	"github.com/gin-gonic/gin"
)

type campaignsBody struct {
	VisitorID  string                 `json:"visitor_id" binding:"required"`
	Context    map[string]interface{} `json:"context" binding:"required"`
	TriggerHit bool                   `json:"trigger_hit"`
}

// CampaignsResponse represents the campaigns call response
type CampaignsResponse struct {
	VisitorID string           `json:"visitor_id"`
	Panic     bool             `json:"panic"`
	Campaigns []model.Campaign `json:"campaigns"`
}

// Campaigns returns a campaigns handler
// @Summary Get all campaigns for the visitor
// @Tags v2
// @Description Get all campaigns value and metadata for a visitor ID and context
// @ID get-campaigns
// @Accept  json
// @Produce  json
// @Param request body campaignsBody true "Campaigns request body"
// @Success 200 {object} CampaignsResponse
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /v2/campaigns [post]
func Campaigns(fsClient *client.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		vObj := &campaignsBody{}
		err := c.BindJSON(vObj)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		v, err := fsClient.NewVisitor(vObj.VisitorID, vObj.Context)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = v.SynchronizeModifications()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		modifications := v.GetAllModifications()

		if vObj.TriggerHit {
			go func() {
				var wg sync.WaitGroup
				for k := range modifications {
					wg.Add(1)
					go func(k string) { // Decrement the counter when the goroutine completes.
						defer wg.Done()
						v.ActivateModification(k)
					}(k)
				}

				wg.Wait()
			}()
		}

		response := CampaignsResponse{
			VisitorID: v.ID,
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
