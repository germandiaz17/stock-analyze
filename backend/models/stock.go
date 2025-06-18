package models

import "time"

type Stock struct {
	Ticker            string    `json:"ticker"`
	ObjetivoDesde     string    `json:"objetivo_desde"`
	ObjetivoHasta     string    `json:"objetivo_hasta"`
	Empresa           string    `json:"empresa"`
	Accion            string    `json:"acción"`
	Corredor          string    `json:"corredor"`
	CalificacionDesde string    `json:"calificación_desde"`
	CalificacionHasta string    `json:"calificación_hasta"`
	Tiempo            time.Time `json:"tiempo"`
}
