package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ZyoGo/Backend-Challange/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getUnit(unit string) time.Duration {
	switch unit {
	case "second":
		return time.Second
	case "minute":
		return time.Minute
	case "hour":
		return time.Hour
	default:
		return time.Minute
	}
}

func NewDatabaseConnection(cfg *config.AppConfig) *pgxpool.Pool {
	var (
		maxConnLife        time.Duration = getUnit(config.GetConfig().Database.MaxConnLifeUnit) * time.Duration(config.GetConfig().Database.MaxConnLife)
		maxConnIdle        time.Duration = getUnit(config.GetConfig().Database.MaxConnIdleUnit) * time.Duration(config.GetConfig().Database.MaxConnIdle)
		connConnectTimeout time.Duration = getUnit(config.GetConfig().Database.ConnConnectTimeoutUnit) * time.Duration(config.GetConfig().Database.ConnConnectTimeout)
		healthCheckPeriod  time.Duration = getUnit(config.GetConfig().Database.HealthCheckPeriodUnit) * time.Duration(config.GetConfig().Database.HealthCheckPeriod)
		uri                string
	)

	uri = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	dbConfig, err := pgxpool.ParseConfig(uri)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
		panic(err)
	}

	dbConfig.MaxConns = cfg.Database.MaxConns
	dbConfig.MinConns = cfg.Database.MinConns
	dbConfig.MaxConnLifetime = maxConnLife
	dbConfig.MaxConnIdleTime = maxConnIdle
	dbConfig.HealthCheckPeriod = healthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = connConnectTimeout

	connPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	return connPool
}
