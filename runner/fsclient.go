package runner

import (
	"time"

	"github.com/flagship-io/flagship-go-sdk/v2"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/bucketing"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/logging"
)

func initFsClient(options Options) (*client.Client, error) {
	bucketingOptions := []func(r *bucketing.Engine){}

	if options.PollingInterval > 0 {
		bucketingOptions = append(bucketingOptions, bucketing.PollingInterval(time.Duration(options.PollingInterval)*time.Second))
	}

	optionsFunc := []client.OptionBuilder{
		client.WithBucketing(bucketingOptions...),
	}

	if options.CacheOptionsBuilder != nil {
		optionsFunc = append(optionsFunc, client.WithVisitorCache(options.CacheOptionsBuilder))
	}

	logging.SetLevel(options.LogLevel)
	fsClient, err := flagship.Start(options.EnvID, options.APIKey, optionsFunc...)

	return fsClient, err
}
