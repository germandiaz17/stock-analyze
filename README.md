# Stock Analyzer

Una plataforma full-stack para analizar acciones de la bolsa y recibir recomendaciones de inversión, desarrollada como parte de un reto técnico.

## Tecnologías Utilizadas

- **Frontend**: Vue 3 + TypeScript + Tailwind CSS + Axios + Font Awesome
- **Backend**: Golang (Gin)
- **Base de datos**: CockroachDB
- **Infraestructura**: Docker & Docker Compose

## Funcionalidades

### Visualización de acciones
- Listado detallado de acciones.
- Filtro de búsqueda por empresa o ticker.
- Paginación del listado.
- Visualización de **tendencias** con íconos.

### Recomendaciones de inversión
- Algoritmo que sugiere la mejor acción para invertir.
- Visual atractivo y responsivo.

### Pruebas
- Cobertura de pruebas para endpoints principales en Go.

## Ejecutar con Docker

## DESPUES DE QUE LA BASE DE DATOS SE SUBA EJECUTAR ESTE SCRIPT
CREATE TABLE IF NOT EXISTS stocks (
  ticker STRING PRIMARY KEY,
  objetivo_desde STRING,
  objetivo_hasta STRING,
  empresa STRING,
  accion STRING,
  corredor STRING,
  calificacion_desde STRING,
  calificacion_hasta STRING,
  tiempo TIMESTAMP
);

```bash
docker-compose up --build

