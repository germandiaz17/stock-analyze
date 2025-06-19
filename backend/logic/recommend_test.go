package logic

import (
	"testing"
	"time"

	"stock-analyzer/models"
)

func TestRecommendBestStock(t *testing.T) {
	mockStocks := []models.Stock{
		{
			Ticker:        "AAA",
			ObjetivoDesde: "$10,00",
			ObjetivoHasta: "$15,00",
			Empresa:       "Empresa A",
			Tiempo:        time.Now(),
		},
		{
			Ticker:        "BBB",
			ObjetivoDesde: "$20,00",
			ObjetivoHasta: "$23,00",
			Empresa:       "Empresa B",
			Tiempo:        time.Now(),
		},
		{
			Ticker:        "CCC",
			ObjetivoDesde: "$5,00",
			ObjetivoHasta: "$12,00", // mayor ganancia
			Empresa:       "Empresa C",
			Tiempo:        time.Now(),
		},
	}

	t.Run("Retorna la mejor acción según mayor diferencia", func(t *testing.T) {
		best := RecommendBestStock(mockStocks)
		if best == nil {
			t.Fatal("No se retornó ninguna acción")
		}
		if best.Ticker != "CCC" {
			t.Errorf("Esperaba CCC, pero retornó %s", best.Ticker)
		}
	})

	t.Run("Retorna nil si no hay datos válidos", func(t *testing.T) {
		empty := RecommendBestStock([]models.Stock{})
		if empty != nil {
			t.Errorf("Esperaba nil, pero retornó %+v", empty)
		}
	})
}
