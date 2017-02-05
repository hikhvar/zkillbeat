// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	ZKillboardFeed string `config:"zkillboard_feed"`
	QueueID string `config:"zkillboard_queue_id"`
	SystemCacheFile string `config:"system_cache_file"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	ZKillboardFeed: "https://redisq.zkillboard.com/listen.php",
	QueueID: "zkillbeat",
	SystemCacheFile: "/tmp/zkill_system.db",
}
