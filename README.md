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
- Exportación a **CSV**.
- Visualización de **tendencias** con íconos.

### Recomendaciones de inversión
- Algoritmo que sugiere la mejor acción para invertir.
- Visual atractivo y responsivo.

### Pruebas
- Cobertura de pruebas para endpoints principales en Go.

## Ejecutar con Docker

```bash
docker-compose up --build
