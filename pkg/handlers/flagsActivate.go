package handlers

import (
	"net/http"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/gin-gonic/gin"
)

type flagActivateBody struct {
	VisitorID string `json:"visitorId" binding:"required"`
}

type flagActivated struct {
	Status string `json:"status" binding:"required"`
}

// FlagActivate returns a flag activation handler
// @Summary Activate a flag key
// @Tags v2
// @Description Activate a flag by its key for a visitor ID
// @ID activate-flag
// @Accept  json
// @Produce  json
// @Param flagActivation body flagActivateBody true "Flag activation request body"
// @Success 200 {object} flagActivated
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Router /v2/flags/{key}/activate [post]
func FlagActivate(fsClient *client.Client, hasCache bool) func(*gin.Context) {
	return func(c *gin.Context) {
		if !hasCache {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "You need visitor caching to activate a flag",
			})
			return
		}
		aObj := &flagActivateBody{}
		err := c.BindJSON(aObj)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		v, err := fsClient.NewVisitor(aObj.VisitorID, map[string]interface{}{})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = v.ActivateCacheModification(c.Param("key"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Data(204, gin.MIMEJSON, nil)
	}
}
