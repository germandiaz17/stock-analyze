package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	var err error
	DB, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(fmt.Sprintf("❌ Error conectando a la base de datos: %v", err))
	}
	fmt.Println("✅ Conectado a CockroachDB")
}
