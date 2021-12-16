package config

import (
	"strings"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/cache"
	"github.com/flagship-io/self-hosted-api/pkg/log"
	"github.com/spf13/viper"
)

// OptionBuilder is a func type to set options to the FlagshipOption.
type OptionBuilder func(*CustomOptions)

// BuildOptions fill out the FlagshipOption struct from option builders
func (f *CustomOptions) BuildOptions(clientOptions ...OptionBuilder) {
	// extract options
	for _, opt := range clientOptions {
		opt(f)
	}
}

// WithBucketing enables the bucketing decision mode for the SDK
func WithCustomCache(customCacheOptions cache.CustomOptions) OptionBuilder {
	return func(f *CustomOptions) {
		f.CustomCacheOptions = customCacheOptions
	}
}

type Options struct {
	Port          int
	ClientOptions ClientOptions
}

type CacheOptions struct {
	CacheType     string
	LocalPath     string
	RedisHost     string
	RedisUsername string
	RedisPassword string
	RedisDb       int
}

type ClientOptions struct {
	EnvID           string
	APIKey          string
	PollingInterval int
	CacheOptions    CacheOptions
}

type CustomOptions struct {
	CustomCacheOptions cache.CustomOptions
}

func GetOptionsFromConfig() Options {

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.GetLogger().Warnf("Could not find config file: %v", err)
	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	viper.SetDefault("port", 8080)
	port := viper.GetInt("port")

	envID := viper.GetString("env_id")
	apiKey := viper.GetString("api_key")
	pollingInterval := viper.GetInt("polling_interval")
	cacheType := viper.GetString("cache.type")
	cacheLocalPath := viper.GetString("cache.options.dbPath")
	cacheRedisHost := viper.GetString("cache.options.redisHost")
	cacheRedisUsername := viper.GetString("cache.options.redisUsername")
	cacheRedisPassword := viper.GetString("cache.options.redisPassword")
	cacheRedisDb := viper.GetInt("cache.options.redisDb")

	return Options{
		Port: port,
		ClientOptions: ClientOptions{
			EnvID:           envID,
			APIKey:          apiKey,
			PollingInterval: pollingInterval,
			CacheOptions: CacheOptions{
				CacheType:     cacheType,
				LocalPath:     cacheLocalPath,
				RedisHost:     cacheRedisHost,
				RedisUsername: cacheRedisUsername,
				RedisPassword: cacheRedisPassword,
				RedisDb:       cacheRedisDb,
			},
		},
	}
}
