package logic

import (
	"strconv"
	"strings"
	"stock-analyzer/models"
)

func RecommendBestStock(stocks []models.Stock) *models.Stock {
	var best models.Stock
	bestDiff := -9999.0

	for _, s := range stocks {
		fromStr := strings.ReplaceAll(strings.ReplaceAll(s.ObjetivoDesde, "$", ""), ",", ".")
		toStr := strings.ReplaceAll(strings.ReplaceAll(s.ObjetivoHasta, "$", ""), ",", ".")

		from, err1 := strconv.ParseFloat(fromStr, 64)
		to, err2 := strconv.ParseFloat(toStr, 64)

		if err1 == nil && err2 == nil {
			diff := to - from
			if diff > bestDiff {
				bestDiff = diff
				best = s
			}
		}
	}

	if best.Ticker == "" {
		return nil
	}
	return &best
}
