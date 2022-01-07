package runner

import (
	"fmt"
	"reflect"

	"github.com/flagship-io/self-hosted-api/internal/log"
	"github.com/flagship-io/self-hosted-api/internal/router"
	"github.com/gin-gonic/gin"
)

// InitRouter creates a Gin engine for the self hosted API from the options
func InitRouter(options Options) (*gin.Engine, error) {

	// Init logger with default Warn level
	log.InitLogger(options.LogLevel)

	fsClient, err := initFsClient(options)

	if err != nil {
		return nil, fmt.Errorf("error when initializing Flagship: %v", err)
	}

	if fsClient.GetCacheManager() != nil {
		log.GetLogger().Infof("using cache of type: %v", reflect.TypeOf(fsClient.GetCacheManager()))
	}

	gin.SetMode(options.GinMode)

	return router.Init(fsClient), nil
}

// RunAPI will start Flagship Self-Hosted API with the selected options
func RunAPI(options Options) error {
	r, err := InitRouter(options)
	if err != nil {
		return fmt.Errorf("error when initializing router: %v", err)
	}

	err = r.Run(fmt.Sprintf("0.0.0.0:%v", options.Port))
	if err != nil {
		return fmt.Errorf("error when running HTTP engine: %v", err)
	}
	return nil
}
