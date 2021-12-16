package config

import (
	"strings"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/cache"
	"github.com/flagship-io/self-hosted-api/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Options struct {
	Port                int
	EnvID               string
	APIKey              string
	PollingInterval     int
	LogLevel            logrus.Level
	GinMode             string
	CacheOptionsBuilder func(options *cache.Options)
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
	viper.SetDefault("gin_mode", "debug")
	viper.SetDefault("log_level", "warn")

	port := viper.GetInt("port")

	envID := viper.GetString("env_id")
	apiKey := viper.GetString("api_key")
	pollingInterval := viper.GetInt("polling_interval")
	logLevel := viper.GetString("log_level")
	ginMode := viper.GetString("gin_mode")
	cacheType := viper.GetString("cache.type")
	cacheLocalPath := viper.GetString("cache.options.dbPath")
	cacheRedisHost := viper.GetString("cache.options.redisHost")
	cacheRedisUsername := viper.GetString("cache.options.redisUsername")
	cacheRedisPassword := viper.GetString("cache.options.redisPassword")
	cacheRedisDb := viper.GetInt("cache.options.redisDb")

	logLevelLogrus, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.GetLogger().Warnf("Could not parse log level: %v, %v", logLevel, err)
		logLevelLogrus = logrus.WarnLevel
	}

	var cacheOptionsFunc cache.OptionBuilder
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
	}

	return Options{
		Port:                port,
		EnvID:               envID,
		APIKey:              apiKey,
		LogLevel:            logLevelLogrus,
		GinMode:             ginMode,
		PollingInterval:     pollingInterval,
		CacheOptionsBuilder: cacheOptionsFunc,
	}
}
