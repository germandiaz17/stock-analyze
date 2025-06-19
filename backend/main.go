package main

import (
	"context"
	"log"

	"stock-analyzer/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Cargar .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error cargando .env:", err)
	}

	InitDB() // ← Conecta a CockroachDB
	defer DB.Close(context.Background())

	// Configurar rutas
	r := gin.Default()

	handlers.DBConnection = func() *pgx.Conn { return DB }

	//GET
	r.GET("/stocks", handlers.GetStocks)
	r.GET("/stocks/recommend", handlers.GetRecommendation)
	
	//POST
	r.POST("/stocks/load", handlers.LoadMock)



	r.Run(":3000")
}
