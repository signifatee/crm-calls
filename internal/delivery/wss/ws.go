package wss

import (
	v1 "asteriskAPI/internal/delivery/http/v1"
	"asteriskAPI/internal/domain/dto"
	"asteriskAPI/internal/domain/entity"
	"asteriskAPI/internal/domain/vocab"
	"asteriskAPI/internal/middleware"
	"asteriskAPI/internal/repository"
	"asteriskAPI/internal/service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func Start() {
	v1.Init("./.env", "configs/", "config")

	//db, err := repository.NewPostgresDB(repository.Config{
	//	Host:     viper.GetString("db.host"),
	//	Port:     viper.GetString("db.port"),
	//	Username: os.Getenv("DB_USER"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//	DBName:   os.Getenv("DB_NAME"),
	//	SSLMode:  viper.GetString("db.ssl_mode"),
	//})

	//if err != nil {
	//	logrus.Fatalf("Cannot open connection to PostgresDB: %s", err.Error())
	//}

	asteriskDB, err1 := repository.NewAsteriskDB(repository.ConfigAsteriskDB{
		Host:     os.Getenv("ASTERISK_DB_HOST"),
		Port:     os.Getenv("ASTERISK_DB_PORT"),
		Username: os.Getenv("ASTERISK_DB_USER"),
		Password: os.Getenv("ASTERISK_DB_PASSWORD"),
		DBName:   os.Getenv("ASTERISK_DB_NAME"),
	})

	if err1 != nil {
		logrus.Fatalf("Cannot open connection to AsteriskDB: %s", err1.Error())
	}

	repos := repository.NewRepository(nil, asteriskDB)
	services := service.NewService(repos)

	origin := "http://localhost/"
	url := fmt.Sprintf("ws://pbx.avecrm.work:8088/ari/events?app=hangup&subscribeAll=true&api_key=%s", os.Getenv("ARI_KEY"))

	for {
		ws, err := websocket.Dial(url, "", origin)
		if err != nil {
			logrus.Fatal("Ошибка при подключении по вебсокету:", err)
			return
		}
		logrus.Infof("Connected to websocket: %v", ws)

		for {
			// Получение ответа от сервера
			var response = make([]byte, 1024)
			_, err = ws.Read(response)
			if err != nil {
				logrus.Errorf("Ошибка при чтении ответа: %v", err)
				return
			}
			clearResponse := bytes.Trim(response, "\x00")

			cd := dto.ChannelDestroyed{}
			err = json.Unmarshal(clearResponse, &cd)
			if cd.Type == "ChannelDestroyed" {
				if len(cd.Channel.Caller.Number) > 6 {
					logrus.Infof("Call ended: %v", cd)
					go func() {
						time.Sleep(10 * time.Second)
						mci, err := getCallInfo(services, cd.Channel.Caller.Number)
						if err != nil {
							logrus.Errorf("Error while getting Call info by dst: %v", err)
						}
						logrus.Infof("DEBUG: TO POST")
						PostMainCallInfo(mci)
					}()
				}

			}
			if err != nil {
				logrus.Errorf("Cannot unmarshal: %v erros is: %v", string(clearResponse), err)
			}

		}

		logrus.Errorf("Connecton to websocket is lost")
	}

}

func PostMainCallInfo(mci *entity.MainCallInfo) {
	logrus.Infof("DEBUG MCI: %v", mci)
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("Panic while sending info: %v", err)
		}
	}()
	crmUrl := os.Getenv("CRM_URL")
	apiKey := os.Getenv("CRM_API_KEY")

	var param = url.Values{}
	param.Set("callId", mci.CallId)
	param.Set("trunk", mci.Trunk)
	param.Set("finishBy", strconv.Itoa(mci.FinishBy))
	param.Set("status", strconv.Itoa(middleware.SetStatus(mci.Status)))
	param.Set("finishReason", mci.FinishReason)
	param.Set("recordUrl", mci.RecordUrl)
	param.Set("callDuration", strconv.Itoa(mci.CallDuration))
	param.Set("answerDuration", strconv.Itoa(mci.AnswerDuration))

	var payload = bytes.NewBufferString(param.Encode())

	req, err := http.NewRequest("POST", crmUrl, payload)
	if err != nil {
		logrus.Errorf("Error while creating a http request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Error while sending call info to CRM: %v", err)
	}
	defer resp.Body.Close()

	logrus.Infof("CRM: sended info: %v", resp.Status)

}

func getCallInfo(service *service.Service, dst string) (*entity.MainCallInfo, error) {
	callId, err := service.GetCallIdByDst(dst)
	if err != nil {
		return nil, err
	}

	fci, err2 := service.GetCallInfo(callId)
	if err2 != nil {
		logrus.Errorf("Error while getting call info: %s", err2.Error())
		return nil, err
	}

	mci, err3 := service.ConvertToMainCallInfo(fci, dst)
	if err3 != nil {
		logrus.Errorf("Cannot convert to main call info")
		return nil, err
	}

	mci.FinishReason = getHangupCause(mci.FinishReason)

	logrus.Infof("got main call info with callid: %v", callId)
	return mci, nil
}

func getHangupCause(hangupcause string) string {
	switch hangupcause {
	case "16":
		return vocab.FINISH_REASON_16
	case "17":
		return vocab.FINISH_REASON_17
	case "18":
		return vocab.FINISH_REASON_18
	case "19":
		return vocab.FINISH_REASON_19
	case "21":
		return vocab.FINISH_REASON_21
	case "41":
		return vocab.FINISH_REASON_41
	case "34":
		return vocab.FINISH_REASON_34
	case "3":
		return vocab.FINISH_REASON_3
	case "0":
		return vocab.FINISH_REASON_0
	case "44":
		return vocab.FINISH_REASON_44
	case "127":
		return vocab.FINISH_REASON_127
	default:
		return vocab.FINISH_REASON_OTHER + hangupcause
	}
}
