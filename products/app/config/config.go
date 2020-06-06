package config

import (
	"fmt"
	"otus-auth/app/system/database/redis"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	DBSN         string
	HTTPBindPort int
	Redis        *redis.Config
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
