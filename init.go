package main

import (
	"time"

	"github.com/flagship-io/flagship-go-sdk/v2"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/bucketing"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/cache"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/client"
	"github.com/flagship-io/self-hosted-api/pkg/log"
	"github.com/spf13/viper"
)

func initFsClient() (*client.Client, error) {

	envID := viper.GetString("env_id")
	apiKey := viper.GetString("api_key")
	pollingInterval := viper.GetInt("polling_interval")
	cacheType := viper.GetString("cache.type")
	cacheLocalPath := viper.GetString("cache.options.dbPath")
	cacheRedisHost := viper.GetString("cache.options.redisHost")
	cacheRedisUsername := viper.GetString("cache.options.redisUsername")
	cacheRedisPassword := viper.GetString("cache.options.redisPassword")
	cacheRedisDb := viper.GetInt("cache.options.redisDb")

	bucketingOptions := []func(r *bucketing.Engine){}

	if pollingInterval > 0 {
		bucketingOptions = append(bucketingOptions, bucketing.PollingInterval(time.Duration(pollingInterval)*time.Second))
	}

	optionsFunc := []client.OptionBuilder{
		client.WithBucketing(bucketingOptions...),
	}
	var cacheOptionsFunc cache.OptionBuilder

	hasCache := true
	switch cacheType {
	case "local":
		cacheOptionsFunc = cache.WithLocalOptions(cache.LocalOptions{
			DbPath: cacheLocalPath,
		})
	case "redis":
		cacheOptionsFunc = cache.WithRedisOptions(cache.RedisOptions{
			Host:     cacheRedisHost,
			Username: cacheRedisUsername,
			Password: cacheRedisPassword,
			Db:       cacheRedisDb,
		})
	default:
		hasCache = false
	}

	if hasCache {
		optionsFunc = append(optionsFunc, client.WithVisitorCache(cacheOptionsFunc))
		log.GetLogger().Infof("Using cache of type %s", cacheType)
	}

	fsClient, err := flagship.Start(envID, apiKey, optionsFunc...)

	return fsClient, err
}
