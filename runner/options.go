package runner

import (
	"crypto/tls"
	"strings"

	"github.com/flagship-io/flagship-go-sdk/v2/pkg/cache"
	"github.com/flagship-io/self-hosted-api/internal/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Options struct {
	// The port the Self Hosted API will listen to (default to 8080)
	Port int
	// The Flagship Environment ID
	EnvID string
	// The Flagship API Key
	APIKey string
	// The polling interval for the API to synchronize with Flagship (default to 1mn)
	PollingInterval int
	// The log level for internal API logging (default to warn)
	LogLevel logrus.Level
	// The mode for gin API engine (default to debug)
	GinMode string
	// The cache option to pass to the Flagship SDK (for custom caching)
	CacheOptionsBuilder func(options *cache.Options)
}

// GetOptionsFromConfig will create the options from the config file or environment variables
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
	cacheRedisTls := viper.GetBool("cache.options.redisTls")

	logLevelLogrus, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.GetLogger().Warnf("Could not parse log level: %v, %v", logLevel, err)
		logLevelLogrus = logrus.WarnLevel
	}

	var cacheOptionsFunc cache.OptionBuilder
	switch cacheType {
	case "local":
		cacheOptionsFunc = cache.WithLocalOptions(cache.LocalOptions{DbPath: cacheLocalPath})
	case "redis":
		var tlsConfig *tls.Config
		if cacheRedisTls {
			tlsConfig = &tls.Config{}
		}
		cacheOptionsFunc = cache.WithRedisOptions(cache.RedisOptions{
			Host:      cacheRedisHost,
			Username:  cacheRedisUsername,
			Password:  cacheRedisPassword,
			Db:        cacheRedisDb,
			TLSConfig: tlsConfig,
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
