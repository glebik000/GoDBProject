package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host                     string `mapstructure:"HOST" valid:"required"`
	Port                     uint16 `mapstructure:"PORT" valid:"required"`
	User                     string `mapstructure:"USER" valid:"required"`
	Password                 string `mapstructure:"PASSWORD" valid:"required"`
	DBName                   string `mapstructure:"DB_NAME" valid:"required"`
	Schema                   string `mapstructure:"SCHEMA" valid:"required"`
	PoolConnectionCount      int    `mapstructure:"POOL_CONNECTION_COUNT" valid:"required"`
	ConnectionLifetimeMillis int    `mapstructure:"CONNECTION_LIFETIME_MILLIS" valid:"required"`
	ConnectionIdleTimeMillis int    `mapstructure:"CONNECTION_IDLE_TIME_MILLIS" valid:"required"`
	ApplicationName          string `mapstructure:"APPLICATION_NAME" valid:"required"`
}

// DSN возвращает DSN строку с настройками для подключения к базе.
func (c Config) DSN() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%v dbname=%s application_name=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.ApplicationName,
	)
}

// Pool создаёт пул соединений к базе.
func (c Config) Pool() (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(c.DSN())
	if err != nil {
		return nil, fmt.Errorf("создание пула содинений: парсинг конфига: %w", err)
	}

	conf.MaxConns = int32(c.PoolConnectionCount)
	conf.MaxConnLifetime = time.Duration(c.ConnectionLifetimeMillis) * time.Millisecond
	conf.MaxConnIdleTime = time.Duration(c.ConnectionIdleTimeMillis) * time.Millisecond

	conf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		sql := fmt.Sprintf("SET search_path TO %s;", c.Schema)
		_, err := conn.Exec(ctx, sql)
		return err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		return nil, fmt.Errorf("создание пула соединений: подключение: %v", err)
	}

	return pool, nil
}
