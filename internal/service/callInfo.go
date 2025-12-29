package service

import (
	"asteriskAPI/internal/domain/entity"
	"asteriskAPI/internal/ftp"
	"asteriskAPI/internal/middleware"
	"asteriskAPI/internal/repository"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type CallInfoService struct {
	repo repository.CallInfo
}

func NewCallInfoService(repo repository.CallInfo) *CallInfoService {
	return &CallInfoService{repo: repo}
}

func (c *CallInfoService) GetCallInfo(callId string) (*entity.FullCallInfo, error) {
	fullCallInfo, err := c.repo.GetCallInfo(callId)
	return fullCallInfo, err
}

func (c *CallInfoService) GetCallIdByDst(dst string) (string, error) {
	callId, err := c.repo.GetCallByDst(dst)
	if err != nil {
		logrus.Errorf("Service cannot get callId by dst: %v:", err)
	}
	return callId, nil
}

func (c *CallInfoService) ConvertToMainCallInfo(fci *entity.FullCallInfo, rawCallId string) (*entity.MainCallInfo, error) {
	disposition := fci.Adcrd.Disposition
	duration, _ := strconv.Atoi(fci.Adcrd.Duration)
	if len(fci.Adcrd.Recordingfile) == 0 {
		str := fmt.Sprintf("No recoring file for callid: %v", fci.Adcrd.Uniqueid)
		logrus.Error(str)
		return nil, errors.New(str)
	}

	recordingUrl := c.ConvertSendToStorageGetFileUrl(fci.Adcrd.Recordingfile)
	if recordingUrl == "" {
		str := fmt.Sprint("RecordingUrl is empty")
		logrus.Error(str)
		return nil, errors.New(str)
	}

	extra := entity.Extra{}
	err := json.Unmarshal([]byte(fci.Adcel.Extra), &extra)
	if err != nil {
		logrus.Errorf("Cannot unmarshal extra info: %v", err)
		return nil, err
	}

	callId, trunk := middleware.GetCallIdAndTrunk(rawCallId)
	logrus.Infof("DEBUG: %s, %s", callId, trunk)

	finishBy, err := middleware.SetFinishedBy(extra.HangupSource)
	if err != nil {
		logrus.Errorf("Cannot set finishedBy: %v", err)
		return nil, err
	}
	answerDuration, err := strconv.Atoi(fci.Adcrd.Billsec)
	if err != nil {
		logrus.Errorf("Cannot convert string billsec to int: %v", err)
		return nil, err
	}

	return &entity.MainCallInfo{
		CallId:         callId,
		Trunk:          trunk,
		FinishBy:       finishBy,
		Status:         disposition,
		FinishReason:   strconv.Itoa(extra.HangupCause),
		RecordUrl:      recordingUrl,
		CallDuration:   duration,
		AnswerDuration: answerDuration,
	}, nil
}

func (c *CallInfoService) ConvertSendToStorageGetFileUrl(fileName string) (fileUrl string) {
	if len(fileName) == 0 {
		logrus.Errorf("Empty filename of recording")
		return ""
	}
	pathToFileWav, err := ftp.GetPathOfRecordingFile(fileName)
	if err != nil {
		logrus.Errorf("Error while getting path to record file: %v", err)
		return ""
	}
	logrus.Infof("Path to wav file: %s", pathToFileWav)
	pathToFileMp3, err := ftp.ConvertToMp3(pathToFileWav)
	if err != nil {
		logrus.Errorf("Error converting to MP3: %v", err)
		return ""
	}

	n, err := strconv.Atoi(os.Getenv("NUMBER_OF_SYMBOLS_TO_CUT"))
	if err != nil {
		logrus.Errorf("Error while converting string NUMBER_OF_SYMBOLS_TO_CUT: %v", err)
		return ""
	}

	mp3FileName := fileName[:len(fileName)-(n-1)]
	mp3FileName = fmt.Sprintf("%s.mp3", mp3FileName)

	logrus.Infof("Mp3 filepath: %s", pathToFileMp3)
	logrus.Infof("Mp3 filename: %s", mp3FileName)

	err = ftp.SendToColdStorageFtp(pathToFileMp3, mp3FileName)
	if err != nil {
		logrus.Errorf("Error while sending to cold storage: %v", err)
		return ""
	}

	fileUrl = ftp.BuildUrlToRecording(mp3FileName)
	return fileUrl
}
