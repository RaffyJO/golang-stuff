package app

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	config := viper.New()
	config.SetConfigFile("config.json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.database"))

	db, err := sql.Open("postgres", connStr)
	// db, err := sql.Open("postgres", "user=postgres password=postgres dbname=go_migration sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

// migrate create -ext sql -dir db/migrations create_table_category
// migrate -database postgres://postgres:postgres@localhost:5432/go_migration?sslmode=disable -path db/migrations up
// migrate -database postgres://postgres:postgres@localhost:5432/go_migration?sslmode=disable -path db/migrations down
// migrate -database postgres://postgres:postgres@localhost:5432/go_migration?sslmode=disable -path db/migrations version
// migrate -database postgres://postgres:postgres@localhost:5432/go_migration?sslmode=disable -path db/migrations force version
