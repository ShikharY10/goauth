package handlers

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type CacheHandler struct {
	redisClient *redis.Client
}

func CreateCacheHandler(redisClient *redis.Client) *CacheHandler {
	return &CacheHandler{
		redisClient: redisClient,
	}
}

// Return true if token is not expired and saved hash of part of token is match is supplied token hash.
func (c *CacheHandler) IsTokenValid(id string, token string, tokenType string) bool {
	var key string
	if tokenType == "access" {
		key = id + ".accessTokenExpiry"
	} else if tokenType == "refresh" {
		key = id + ".refreshTokenExpiry"
	}

	if key == "" {
		return false
	}

	hash := strings.Split(token, ".")[2]

	result := c.redisClient.Get(key)
	return result.Val() == hash
}

// Saves token's hash part and set a expiry as specified in Redis Cache.
func (c *CacheHandler) SetAccessTokenExpiry(id string, token string, accessTokenExpiry time.Duration) error {
	hash := strings.Split(token, ".")[2]
	result1 := c.redisClient.Set(id+".accessTokenExpiry", hash, accessTokenExpiry)
	if result1.Err() != nil {
		return result1.Err()
	}
	return nil
}

// Saves token's hash part and set a expiry as specified in Redis Cache.
func (c *CacheHandler) SetRefreshTokenExpiry(id string, token string, refreshTokenExpiry time.Duration) error {
	hash := strings.Split(token, ".")[2]
	result2 := c.redisClient.Set(id+".refreshTokenExpiry", hash, refreshTokenExpiry)
	if result2.Err() != nil {
		return result2.Err()
	}
	return nil
}

// Deletes both refresh and access token from redis cache
func (c *CacheHandler) DeleteTokenExpiry(id string) {
	c.redisClient.Del(id + ".accessTokenExpiry")
	c.redisClient.Del(id + ".refreshTokenExpiry")
}
