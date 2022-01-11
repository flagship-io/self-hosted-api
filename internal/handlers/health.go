package handlers

import (
	"net/http"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"

	// use for api docs
	_ "github.com/flagship-io/self-hosted-api/internal/httputils"
	"github.com/gin-gonic/gin"
)

type statusResponse struct {
	Status string `json:"status"`
}

// Health return a health handler
// @Summary Get health status
// @Tags Health
// @Description Get current health status for the API
// @ID health
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse{}
// @Router /health [get]
func Health(fsClient *client.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, statusResponse{
			Status: fsClient.GetStatus(),
		})
	}
}
