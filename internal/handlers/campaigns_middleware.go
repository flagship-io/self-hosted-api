package handlers

import (
	"net/http"
	"sync"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/flagship-io/self-hosted-api/internal/log"
	"github.com/gin-gonic/gin"
)

func CampaignMiddleware(fsClient *client.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		vObj := &campaignsBody{}
		err := c.BindJSON(vObj)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		hasAnonymous := vObj.AnonymousID != nil && *vObj.AnonymousID != ""

		// set initial visitor ID as anonymous ID if exists
		vID := vObj.VisitorID
		if hasAnonymous {
			vID = *vObj.AnonymousID
		}
		v, err := fsClient.NewVisitor(vID, vObj.Context)

		// If anonymous id is set, authenticate the visitor
		if hasAnonymous {
			err = v.Authenticate(vObj.VisitorID, nil, false)
		}

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
						err := v.ActivateModification(k)
						if err != nil {
							log.GetLogger().Warnf("error when activating modification : %v", err)
						}
					}(k)
				}

				wg.Wait()
			}()
		}

		c.Set("modifications", modifications)
		c.Set("visitor", v)
	}
}
