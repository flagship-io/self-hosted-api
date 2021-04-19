package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/model"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/pkg/httputils"
	"github.com/gin-gonic/gin"
)

// Campaign returns a campaign handler
// @Summary Get a single campaigns for the visitor
// @Tags v2
// @Description Get a single campaign value and metadata for a visitor ID and context
// @ID get-campaign
// @Accept  json
// @Produce  json
// @Param id path string true "Campaign ID"
// @Param request body campaignsBody true "Campaign request body"
// @Success 200 {object} model.Campaign
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /v2/campaigns/{id} [post]
func Campaign(fsClient *client.Client) func(*gin.Context) {
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

		decisionResponse := v.GetDecisionResponse()
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

		response := model.Campaign{}
		found := false
		cID := c.Param("id")
		log.Println(cID)
		for _, c := range decisionResponse.Campaigns {
			log.Println(c.ID)
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
