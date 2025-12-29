package handler

import (
	"asteriskAPI/internal/domain/dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getCallInfo(c *gin.Context) {

	cib := &dto.CallInfoBody{}
	err := c.ShouldBindJSON(cib)
	if err != nil {
		logrus.Errorf("Error while binding JSON to CallInfoBody: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	callId := cib.CallId

	if cib.Dst != "" {
		callId, err = h.services.GetCallIdByDst(cib.Dst)
		if err != nil {
			logrus.Errorf("Error while getting callId by Dst: %s", err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	}

	fci, err2 := h.services.GetCallInfo(callId)
	if err2 != nil {
		logrus.Errorf("Error while getting call info: %s", err2.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err2.Error()})
		return
	}

	mci, err3 := h.services.ConvertToMainCallInfo(fci, cib.Dst)
	if err3 != nil {
		logrus.Errorf("Cannot convert to main call info")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err3.Error()})
		return
	}

	logrus.Infof("got main call info: %v", mci)
	c.JSON(http.StatusOK, mci)

}
