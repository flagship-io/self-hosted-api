package handlers

import (
	"net/http"
	"sync"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/pkg/httputils"
	"github.com/gin-gonic/gin"
)

// FlagMetadata represents the metadata informations about a flag key
type FlagMetadata struct {
	CampaignID       string `json:"campaignId"`
	VariationGroupID string `json:"variationGroupID"`
	VariationID      string `json:"variationID"`
}

// FlagInfos represents the informations about a flag key
type FlagInfos struct {
	Value    interface{}  `json:"value"`
	Metadata FlagMetadata `json:"metadata"`
}

// Flags returns a flags handler
// @Summary Get all flags
// @Tags v2
// @Description Get all flags value and metadata for a visitor ID and context
// @ID get-flags
// @Accept  json
// @Produce  json
// @Param request body campaignsBodySwagger true "Flag request body"
// @Success 200 {object} map[string]FlagInfos{}
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /v2/flags [post]
func Flags(fsClient *client.Client) func(*gin.Context) {
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
			var wg sync.WaitGroup
			for k := range modifications {
				wg.Add(1)
				go func(k string) { // Decrement the counter when the goroutine completes.
					defer wg.Done()
					v.ActivateModification(k)
				}(k)
			}

			wg.Wait()
		}

		result := map[string]FlagInfos{}

		for k, m := range modifications {
			result[k] = FlagInfos{
				Value: m.Value,
				Metadata: FlagMetadata{
					CampaignID:       m.Campaign.ID,
					VariationGroupID: m.Campaign.VariationGroupID,
					VariationID:      m.Campaign.Variation.ID,
				},
			}
		}
		c.JSON(http.StatusOK, result)
	}
}
