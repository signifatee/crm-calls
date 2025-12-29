package unit_tests

import (
	"asteriskAPI/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestValidateToken(t *testing.T) {
	router := gin.Default()
	router.GET("/", handler.ValidateToken)

	token := "valid_token"
	os.Setenv("API_TOKEN", token)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", token)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestValidateToken_InvalidToken(t *testing.T) {
	router := gin.Default()
	router.GET("/", handler.ValidateToken)

	invalidToken := "invalid_token"
	os.Setenv("API_TOKEN", "valid_token")

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", invalidToken)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusForbidden, resp.Code)
}
