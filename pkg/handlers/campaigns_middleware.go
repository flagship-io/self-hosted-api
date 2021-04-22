package handlers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
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

		start = time.Now()
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
		elapsed = time.Since(start)
		log.Printf("Activate took %s", elapsed)

		c.Set("modifications", modifications)
		c.Set("visitor", v)
	}
}
