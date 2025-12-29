package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func ValidateToken(c *gin.Context) {

	if c.GetHeader("Authorization") != os.Getenv("API_TOKEN") {
		msg := fmt.Sprintf("%v %v invalid token", c.Request.Method, c.Request.URL)
		c.JSON(http.StatusForbidden, gin.H{"error": msg})
		logrus.Error(msg)
		c.Abort()
	}

	c.Next()
}
