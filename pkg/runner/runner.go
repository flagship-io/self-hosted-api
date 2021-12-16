package runner

import (
	"fmt"

	"github.com/flagship-io/self-hosted-api/pkg/config"
	"github.com/flagship-io/self-hosted-api/pkg/fsclient"
	"github.com/flagship-io/self-hosted-api/pkg/log"
	"github.com/flagship-io/self-hosted-api/pkg/router"
	"github.com/sirupsen/logrus"
)

func RunAPI(options config.Options, otherOptions ...config.OptionBuilder) error {
	fsClientOptions := &config.CustomOptions{}
	fsClientOptions.BuildOptions(otherOptions...)

	// Init logger with default Warn level
	log.InitLogger(logrus.WarnLevel)

	fsClient, err := fsclient.InitFsClient(options.ClientOptions, fsClientOptions)

	if err != nil {
		return fmt.Errorf("error when initializing Flagship: %v", err)
	}

	r := router.Init(fsClient)

	err = r.Run(fmt.Sprintf("0.0.0.0:%v", options.Port))
	if err != nil {
		return fmt.Errorf("error when running HTTP engine: %v", err)
	}
	return nil
}
