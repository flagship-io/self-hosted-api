package runner

import (
	"fmt"
	"reflect"

	"github.com/flagship-io/self-hosted-api/pkg/config"
	"github.com/flagship-io/self-hosted-api/pkg/fsclient"
	"github.com/flagship-io/self-hosted-api/pkg/log"
	"github.com/flagship-io/self-hosted-api/pkg/router"
	"github.com/gin-gonic/gin"
)

func RunAPI(options config.Options) error {

	// Init logger with default Warn level
	log.InitLogger(options.LogLevel)

	fsClient, err := fsclient.InitFsClient(options)

	if err != nil {
		return fmt.Errorf("error when initializing Flagship: %v", err)
	}

	if fsClient.GetCacheManager() != nil {
		log.GetLogger().Infof("using cache of type: %v", reflect.TypeOf(fsClient.GetCacheManager()))
	}

	gin.SetMode(options.GinMode)

	r := router.Init(fsClient)

	err = r.Run(fmt.Sprintf("0.0.0.0:%v", options.Port))
	if err != nil {
		return fmt.Errorf("error when running HTTP engine: %v", err)
	}
	return nil
}
