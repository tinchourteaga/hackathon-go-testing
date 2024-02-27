package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	repo := NewRepository()
	service := NewService(repo)
	prod := NewHandler(service)
	r := gin.Default()

	rg := r.Group("/api/v1")
	rg.GET("/products/:id", prod.GetProducts)

	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestGetAll(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/api/v1/products/FEX112AC", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
