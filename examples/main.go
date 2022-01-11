package main

import (
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/cache"
	"github.com/flagship-io/self-hosted-api/runner"
)

var inMemoryCache = map[string]map[string]*cache.CampaignCache{}

func main() {
	options := runner.Options{
		Port:   8080,
		EnvID:  "env_id",
		APIKey: "api_key",
		CacheOptionsBuilder: cache.WithCustomOptions(cache.CustomOptions{
			Getter: func(visitorID string) (map[string]*cache.CampaignCache, error) {
				return inMemoryCache[visitorID], nil
			},
			Setter: func(visitorID string, campaignCache map[string]*cache.CampaignCache) error {
				inMemoryCache[visitorID] = campaignCache
				return nil
			},
		}),
	}
	err := runner.RunAPI(options)

	panic(err)
}
