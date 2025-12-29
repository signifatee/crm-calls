package ftp

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
)

type CallRecordingsFtp interface {
	GetPathOfRecordingFile(fileName string) (pathToRecordingFile string, err error)
	ConvertToMp3(pathToRecordingFile string) (pathToNewFile string, err error)
	SendToColdStorageFtp(pathToRecording string) error
	BuildUrlToRecording(fileName string) error
}

func GetPathOfRecordingFile(fileName string) (pathToRecordingFile string, err error) {
	str := fmt.Sprintf("find %s -name \"%s\"", os.Getenv("CALL_RECORDINGS_DIR"), fileName)
	cmd := exec.Command("bash", "-c", str)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("Error finding recording file: %v", err)
		logrus.Errorf(string(bytes))
		return "", err
	}
	if len(bytes) == 0 {
		logrus.Errorf("Could not find recording %s: %v", pathToRecordingFile, err)
		return "", err
	}
	pathToRecordingFile = string(bytes)
	return pathToRecordingFile, nil
}

func ConvertToMp3(pathToRecordingFile string) (pathToNewFile string, err error) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("Panic while converting to Mp3: %v", err)
		}
	}()
	numSymbolsToCut, err := strconv.Atoi(os.Getenv("NUMBER_OF_SYMBOLS_TO_CUT"))
	if err != nil {
		logrus.Errorf("Error while converting NUMBER_OF_SYMBOLS_TO_CUT to int")
	}

	pathWithoutFileFormat := pathToRecordingFile[:len(pathToRecordingFile)-numSymbolsToCut]
	str := fmt.Sprintf("ffmpeg -i %s.wav -vn -ar 44100 -ac 2 -b:a 192k %s.mp3", pathWithoutFileFormat, pathWithoutFileFormat)

	cmd := exec.Command("bash", "-c", str)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("Error converting from wav to mp3: %v \n Error: %v", string(bytes), err)
		return "", err
	}

	pathToNewFile = fmt.Sprintf("%s.mp3", pathWithoutFileFormat)
	return pathToNewFile, nil
}

func SendToColdStorageFtp(pathToRecording string, fileName string) error {

	str := fmt.Sprintf("curl -T %s ftp://u384640:xbYhimFf5r2RiXDz@u384640.your-storagebox.de/records/%s", pathToRecording, fileName)
	cmd := exec.Command("bash", "-c", str)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Errorf("Error finding recording file: %v", err)
		return err
	}
	logrus.Infof("Successfully sended recording to cold storage: %v", string(bytes))
	return nil
}

func BuildUrlToRecording(fileNameMp3 string) string {
	basePath := os.Getenv("STORAGE_PATH_TO_RECORDINGS")
	fullpath := basePath + fileNameMp3

	return fullpath
}
