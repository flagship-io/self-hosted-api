package main

import (
	"github.com/flagship-io/flagship-go-sdk/v2/pkg/cache"
	"github.com/flagship-io/self-hosted-api/pkg/config"
	"github.com/flagship-io/self-hosted-api/pkg/fsclient"
	"github.com/flagship-io/self-hosted-api/pkg/runner"
)

var inMemoryCache = map[string]map[string]*cache.CampaignCache{}

func main() {
	options := config.Options{}
	err := runner.RunAPI(fsclient.WithCustomCache(cache.CustomOptions{
		Getter: func(visitorID string) (map[string]*cache.CampaignCache, error) {
			return inMemoryCache[visitorID], nil
		},
		Setter: func(visitorID string, campaignCache map[string]*cache.CampaignCache) error {
			inMemoryCache[visitorID] = campaignCache
			return nil
		},
	}))

	panic(err)
}
