package config

import (
	"GoDBProject/pkg/storage/postgres"
)

func GetConfig() postgres.Config {
	return postgres.Config{
		Host:                     "127.0.0.1",
		Port:                     uint16(5433),
		User:                     "postgres",
		Password:                 pgPass,
		DBName:                   "testdb",
		Schema:                   "public",
		PoolConnectionCount:      1,
		ConnectionLifetimeMillis: 150000,
		ConnectionIdleTimeMillis: 150000,
		ApplicationName:          "...",
	}
}
