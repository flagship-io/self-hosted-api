package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/pkg/httputils"
	"github.com/gin-gonic/gin"
)

// FlagValue returns a flag value handler
// @Summary Get flag value by name
// @Tags v2
// @Description Get a single flag value for a visitor ID and context
// @ID get-flag-value
// @Accept  json
// @Produce  json
// @Param key path string true "Flag key"
// @Param request body campaignsBodySwagger true "Flag request body"
// @Success 200 {object} interface{}
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /v2/flags/{key}/value [post]
func FlagValue(fsClient *client.Client) func(*gin.Context) {
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
