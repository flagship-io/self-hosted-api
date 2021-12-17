package fsclient

import (
	"os"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/cache"
	"github.com/flagship-io/self-hosted-api/pkg/config"
	"github.com/flagship-io/self-hosted-api/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var inMemoryCache = map[string]map[string]*cache.CampaignCache{}

func TestInitFsClient(t *testing.T) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	// Init logger with default Warn level
	log.InitLogger(logrus.WarnLevel)

	options := config.GetOptionsFromConfig()
	fsClient, err := InitFsClient(options)

	assert.NotNil(t, err)
	assert.Nil(t, fsClient)
	assert.Contains(t, err.Error(), "EnvID is required")

	os.Setenv("ENV_ID", "test_env_id")
	options = config.GetOptionsFromConfig()
	fsClient, err = InitFsClient(options)

	assert.Nil(t, fsClient)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "APIKey is required")

	os.Setenv("API_KEY", "test_api_key")
	options = config.GetOptionsFromConfig()
	fsClient, err = InitFsClient(options)

	assert.NotNil(t, fsClient)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "403 Forbidden")

	options.CacheOptionsBuilder = cache.WithLocalOptions(cache.LocalOptions{
		DbPath: "./test.db",
	})
	fsClient, err = InitFsClient(options)
	assert.NotNil(t, fsClient)
	assert.NotNil(t, err)
	assert.NotNil(t, fsClient.GetCacheManager())

	localCacheM := &cache.LocalDBManager{}
	assert.IsType(t, localCacheM, fsClient.GetCacheManager())
	os.RemoveAll("./test.db")

	s, err := miniredis.Run()
	assert.Nil(t, err)
	options.CacheOptionsBuilder = cache.WithRedisOptions(cache.RedisOptions{
		Host: s.Addr(),
	})
	fsClient, err = InitFsClient(options)
	assert.NotNil(t, fsClient)
	assert.NotNil(t, err)
	assert.NotNil(t, fsClient.GetCacheManager())

	redisCacheM := &cache.RedisManager{}
	assert.IsType(t, redisCacheM, fsClient.GetCacheManager())

	options.CacheOptionsBuilder = cache.WithCustomOptions(cache.CustomOptions{
		Getter: func(visitorID string) (map[string]*cache.CampaignCache, error) {
			return inMemoryCache[visitorID], nil
		},
		Setter: func(visitorID string, campaignCache map[string]*cache.CampaignCache) error {
			inMemoryCache[visitorID] = campaignCache
			return nil
		},
	})
	fsClient, err = InitFsClient(options)
	assert.NotNil(t, fsClient)
	assert.NotNil(t, err)
	assert.NotNil(t, fsClient.GetCacheManager())

	customCacheM := &cache.CustomManager{}
	assert.IsType(t, customCacheM, fsClient.GetCacheManager())
}
