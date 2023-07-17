package utils

import (
	"github.com/patrickmn/go-cache"
	"time"
	"widget-sports/configurations"
)

var (
	cacheMinutes = configurations.LoadConfig().CacheMinutes
	c            = cache.New(time.Duration(cacheMinutes)*time.Minute, time.Duration(cacheMinutes)*time.Minute)
)

func GetFromCache(key string) ([]map[string]interface{}, bool) {
	value, found := c.Get(key)
	if found {
		if matches, ok := value.([]map[string]interface{}); ok {
			return matches, true
		}
	}
	return nil, false
}

func SetCache(cacheKey string, value interface{}) {
	c.Set(cacheKey, value, cache.DefaultExpiration)
}
