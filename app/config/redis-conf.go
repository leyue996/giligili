package config

type redis struct {
	RedisDb     string `yaml:"redisDb"`
	RedisAddr   string `yaml:"redisAddr"`
	RedisPw     string `yaml:"redisPw"`
	RedisDbName string `yaml:"redisDbName"`
}
