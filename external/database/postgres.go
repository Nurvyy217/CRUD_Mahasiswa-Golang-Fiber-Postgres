package database

import (
	"fmt"
	"github.com/Nurvyy217/crud-mahasiswa-go/internal/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DBConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		db = nil
		return
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdleTimeConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnection))
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdleConnection))

	return
}