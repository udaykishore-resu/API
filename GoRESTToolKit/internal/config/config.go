package config

import "time"

type Config struct {
	Port      string        `envconfig:"PORT"	default:"8080"`
	JWTSecret string        `envconfig:"JWT_SECRET"	required:"true"`
	DBURL     string        `envconfig:"DATABASE_URL"	required:"true"`
	RateLimit int           `envconfig:"RATE_LIMIT"	default:"100"`
	Timeout   time.Duration `envconfig:"TIMEOUT"	default:"15s"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}
