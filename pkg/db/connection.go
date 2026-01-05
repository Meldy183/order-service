package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConfigDB struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-default:"postgres"`
	Database string `yaml:"name" env:"DB_NAME" env-default:"postgres"`
}
type DataBase struct {
	Pool *pgxpool.Pool
}

func NewDataBase(config ConfigDB) (*DataBase, error) {
	dataSource := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)
	pool, err := pgxpool.New(context.Background(), dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &DataBase{Pool: pool}, nil
}
