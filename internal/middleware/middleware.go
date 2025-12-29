package middleware

import (
	"asteriskAPI/internal/domain/vocab"
	"github.com/sirupsen/logrus"
	"regexp"
)

func SetFinishedBy(finishByStr string) (int, error) {

	matched, err := regexp.MatchString(`^SIP/1.*`, finishByStr)
	if err != nil {
		logrus.Errorf("Ошибка при компиляции регулярного выражения FinishBy: %v", err)
		return 0, err
	}

	if matched {
		return vocab.FINISHED_BY_OPERATOR, nil
	}

	matched, err = regexp.MatchString(`^PJSIP/.*`, finishByStr)
	if err != nil {
		logrus.Errorf("Ошибка при компиляции регулярного выражения FinishBy: %v", err)
		return 0, err
	}

	if matched {
		return vocab.FINISHED_BY_CLIENT, nil
	}

	return vocab.FINISHED_BY_SYSTEM, nil
}

func GetCallIdAndTrunk(full string) (callId string, trunk string) {
	logrus.Infof("Getting CallId and trunk from dst: %s", full)
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("Panic while getting callId and trunk: %v", err)
		}
	}()

	trunk = full[0:2]
	callId = full[2:10]

	return callId, trunk
}

func checkIfFirstCharacterZero(str string) string {
	if str[0:1] == "0" {
		return str[1:]
	}
	return str
}

func SetStatus(status string) int {
	switch status {
	case "ANSWERED":
		return vocab.STATUS_ANSWERED
	case "NO ANSWER":
		return vocab.STATUS_NOT_ANSWERED
	case "BUSY":
		return vocab.STATUS_FAILED
	default:
		return vocab.STATUS_UNKNOWN
	}
}
