package dataBase

import (
	"context"
	"log"

	"github.com/azoma13/avito/models"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type DataBase interface {
	CreateAccount(*models.Employee) error
	CreateTable() error
}

var DB *pgx.Conn

func ConnectToDB() *pgx.Conn {
	var err error
	DB, err = pgx.Connect(context.Background(), "postgres://postgres:08082018@localhost:5432/avitoDB")
	if err != nil {
		log.Fatal("Unable to create connection dataBase: %w", err)
	}

	createTableDB()
	createShopMerchDB()
	return DB
}

// func (s *PostgresDB) CreateEmployee(*models.Employee) error {
// 	return s.CreateEmployeeTable()
// }

// func (s *PostgresDB) CreateEmployeeTable() error {
// 	query := `create table if not exists employee (
// 		id serial primary key,
// 		email varchar(50)
// 		)`
// 	_, err := s.db.Exec(query)
// 	return err
// }

// func ConnectToDB() *pgx.Conn {

// 	var err error
// 	connStr := "postgresql://postgres:08082018@localhost:5432/avito_db"

// 	DB, err = pgx.Connect(context.Background(), connStr)
// 	DB, err := gorm.Open(pgx., &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Unable to create connection pool: %v\n", err)
// 	}

// 	return DB
// }
