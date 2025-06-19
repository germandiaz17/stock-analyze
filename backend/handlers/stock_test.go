package handlers

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

// Crear una conexión a DB de test
var testPool *pgxpool.Pool
func setupRouter() *gin.Engine {
	_ = godotenv.Load("../.env") 
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	connStr := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	testPool = pool

	DBConnection = func() *pgx.Conn {
		var realConn *pgx.Conn
		err := pool.AcquireFunc(context.Background(), func(c *pgxpool.Conn) error {
			realConn = c.Conn()
			return nil
		})
		if err != nil {
			panic(err)
		}
		return realConn
	}

	r.GET("/stocks", GetStocks)
	r.POST("/stocks/load", LoadMock)

	return r
}

// Test para POST /stocks/load
func TestLoadMock(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("POST", "/stocks/load", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Se esperaba 200 OK, se obtuvo %d. Body: %s", resp.Code, string(body))
	}
}

// Test para GET /stocks
func TestGetStocks(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/stocks", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Esperaba 200 OK, pero recibió %d", resp.Code)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) < 10 {
		t.Errorf("La respuesta es demasiado corta, ¿datos vacíos? Body: %s", string(body))
	}
}
