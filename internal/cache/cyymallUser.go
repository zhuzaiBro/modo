package cache

import (
	"context"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"

	"cyy/internal/model"
)

const (
	// cache prefix key, must end with a colon
	cyymallUserCachePrefixKey = "cyymallUser:"
	// CyymallUserExpireTime expire time
	CyymallUserExpireTime = 5 * time.Minute
)

var _ CyymallUserCache = (*cyymallUserCache)(nil)

// CyymallUserCache cache interface
type CyymallUserCache interface {
	Set(ctx context.Context, id uint64, data *model.CyymallUser, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.CyymallUser, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CyymallUser, error)
	MultiSet(ctx context.Context, data []*model.CyymallUser, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// cyymallUserCache define a cache struct
type cyymallUserCache struct {
	cache cache.Cache
}

// NewCyymallUserCache new a cache
func NewCyymallUserCache(cacheType *model.CacheType) CyymallUserCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.CyymallUser{}
		})
		return &cyymallUserCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.CyymallUser{}
		})
		return &cyymallUserCache{cache: c}
	}

	return nil // no cache
}

// GetCyymallUserCacheKey cache key
func (c *cyymallUserCache) GetCyymallUserCacheKey(id uint64) string {
	return cyymallUserCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *cyymallUserCache) Set(ctx context.Context, id uint64, data *model.CyymallUser, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCyymallUserCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *cyymallUserCache) Get(ctx context.Context, id uint64) (*model.CyymallUser, error) {
	var data *model.CyymallUser
	cacheKey := c.GetCyymallUserCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *cyymallUserCache) MultiSet(ctx context.Context, data []*model.CyymallUser, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCyymallUserCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *cyymallUserCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CyymallUser, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCyymallUserCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.CyymallUser)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.CyymallUser)
	for _, id := range ids {
		val, ok := itemMap[c.GetCyymallUserCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *cyymallUserCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCyymallUserCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *cyymallUserCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetCyymallUserCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
