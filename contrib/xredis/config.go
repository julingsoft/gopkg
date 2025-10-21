package xredis

type Config struct {
	Default RedisConfig
	Cache   RedisConfig
}

type RedisConfig struct {
	Address     string
	Pass        string
	Db          int
	IdleTimeout int
}
