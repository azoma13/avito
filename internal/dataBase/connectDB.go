package dataBase

import (
	"context"
	"log"
	"time"

	"github.com/azoma13/avito/models"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type DataBase interface {
	CreateAccount(*models.Employee) error
	CreateTable() error
}

var DB *pgxpool.Pool

func ConnectToDB() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@db:5432/avitoDB")
	if err != nil {
		log.Fatalf("Unable to parse database config: %v", err)
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	time.Sleep(5 * time.Second)

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	createTableDB()
	createShopMerchDB()
	return DB
}
