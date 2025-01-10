package database

import (
	"time"
)

func NewConfig(dbURL string) Config {
	return Config{
		URL:             dbURL,
		MaxOpenConns:    25,
		MaxIdleConns:    25,
		ConnMaxLifetime: 5 * time.Minute,
	}
}
