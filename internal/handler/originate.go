package handler

import (
	"asteriskAPI/internal/domain/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) originate(c *gin.Context) {

	initCall := &dto.InitCall{}

	if err := c.BindJSON(&initCall); err != nil {
		logrus.Errorf("Error while binding JSON to initCall: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	initCallResponse, err := h.services.OriginateCall(initCall, "arichannelspostreply")
	if err != nil {
		logrus.Errorf("error occured while originating call: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Printf("Originated call: %v", initCallResponse)
	c.JSON(http.StatusOK, gin.H{"message": initCallResponse})
}
