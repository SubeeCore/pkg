package redis

type Config struct {
	CacheHost string `env:"CACHE_HOST"`
	CachePort uint16 `env:"CACHE_PORT"`
}
