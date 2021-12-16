package fsclient

import (
	"errors"
	"time"

	"github.com/flagship-io/flagship-go-sdk/v2"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/bucketing"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/cache"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/flagship-io/self-hosted-api/pkg/config"
	"github.com/flagship-io/self-hosted-api/pkg/log"
)

func InitFsClient(clientOptions config.ClientOptions, options *config.CustomOptions) (*client.Client, error) {

	bucketingOptions := []func(r *bucketing.Engine){}

	if clientOptions.PollingInterval > 0 {
		bucketingOptions = append(bucketingOptions, bucketing.PollingInterval(time.Duration(clientOptions.PollingInterval)*time.Second))
	}

	optionsFunc := []client.OptionBuilder{
		client.WithBucketing(bucketingOptions...),
	}
	var cacheOptionsFunc cache.OptionBuilder

	hasCache := true
	switch clientOptions.CacheOptions.CacheType {
	case "local":
		cacheOptionsFunc = cache.WithLocalOptions(cache.LocalOptions{
			DbPath: clientOptions.CacheOptions.LocalPath,
		})
	case "redis":
		cacheOptionsFunc = cache.WithRedisOptions(cache.RedisOptions{
			Host:     clientOptions.CacheOptions.RedisHost,
			Username: clientOptions.CacheOptions.RedisUsername,
			Password: clientOptions.CacheOptions.RedisPassword,
			Db:       clientOptions.CacheOptions.RedisDb,
		})
	case "custom":
		if options == nil {
			return nil, errors.New("wrong cache option: when using custom cache, the CustomCacheOptions option is required")
		}
		cacheOptionsFunc = cache.WithCustomOptions(options.CustomCacheOptions)
	default:
		hasCache = false
	}

	if hasCache {
		optionsFunc = append(optionsFunc, client.WithVisitorCache(cacheOptionsFunc))
		log.GetLogger().Infof("Using cache of type %s", clientOptions.CacheOptions.CacheType)
	}

	fsClient, err := flagship.Start(clientOptions.EnvID, clientOptions.APIKey, optionsFunc...)

	return fsClient, err
}
