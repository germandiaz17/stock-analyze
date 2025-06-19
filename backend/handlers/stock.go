package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"


	"stock-analyzer/models"
	"stock-analyzer/logic"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var DBConnection func() *pgx.Conn // se setea desde main

func GetStocks(c *gin.Context) {
	rows, err := DBConnection().Query(context.Background(), "SELECT * FROM stocks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos"})
		return
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var s models.Stock
		err := rows.Scan(&s.Ticker, &s.ObjetivoDesde, &s.ObjetivoHasta, &s.Empresa, &s.Accion, &s.Corredor, &s.CalificacionDesde, &s.CalificacionHasta, &s.Tiempo)
		if err == nil {
			stocks = append(stocks, s)
		}
	}
	c.JSON(http.StatusOK, stocks)
}

func GetRecommendation(c *gin.Context) {
	rows, err := DBConnection().Query(context.Background(), "SELECT ticker, objetivo_desde, objetivo_hasta, empresa, accion, corredor, calificacion_desde, calificacion_hasta, tiempo FROM stocks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar la base de datos"})
		return
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var s models.Stock
		err := rows.Scan(&s.Ticker, &s.ObjetivoDesde, &s.ObjetivoHasta, &s.Empresa, &s.Accion, &s.Corredor, &s.CalificacionDesde, &s.CalificacionHasta, &s.Tiempo)
		if err == nil {
			stocks = append(stocks, s)
		}
	}

	best := logic.RecommendBestStock(stocks)
	if best == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No se encontró una acción recomendable"})
		return
	}

	c.JSON(http.StatusOK, best)
}



func LoadMock(c *gin.Context) {
	file, err := os.Open(path + "mock.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo abrir el archivo mock"})
		return
	}
	defer file.Close()

	var data struct {
		Items []models.Stock `json:"items"`
	}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar el JSON"})
		return
	}

	for _, stock := range data.Items {
		_, err := DBConnection().Exec(context.Background(),
			`INSERT INTO stocks (
				ticker, objetivo_desde, objetivo_hasta, empresa, accion,
				corredor, calificacion_desde, calificacion_hasta, tiempo
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
			ON CONFLICT (ticker) DO NOTHING`,
			stock.Ticker, stock.ObjetivoDesde, stock.ObjetivoHasta, stock.Empresa,
			stock.Accion, stock.Corredor, stock.CalificacionDesde, stock.CalificacionHasta, stock.Tiempo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error insertando registros", "detalle": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datos cargados correctamente", "registros": len(data.Items)})
}

