package utils

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
)

// NewCache 创建一个针对 MediaWarp 场景优化的 BigCache 实例
// 内存上限 32MB，适合低并发、中等条目的缓存场景
func NewCache(ttl time.Duration) (*bigcache.BigCache, error) {
	return bigcache.New(context.Background(), bigcache.Config{
		Shards:             2,
		LifeWindow:         ttl,
		CleanWindow:        1 * time.Second,
		MaxEntriesInWindow: 500,
		MaxEntrySize:       5 * 1024 * 1024, // 单条最大 5MB
		HardMaxCacheSize:   32,              // 总内存硬上限 32MB
		StatsEnabled:       false,
		Verbose:            false,
	})
}
