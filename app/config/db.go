package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPool(ctx context.Context,dsn string)(*pgxpool.Pool,error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("Parse config failed: %w",err)
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 15* time.Minute

	pool,err := pgxpool.NewWithConfig(ctx,config)
	if err != nil {
		return nil, fmt.Errorf("create pool failed: %w",err)
	}

	if err := pool.Ping(ctx); err != nil{
		pool.Close()
		return nil, fmt.Errorf("Ping failed: w%",err)
	}
	return pool,nil
}