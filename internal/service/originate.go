package service

import (
	"asteriskAPI/internal/domain/dto"
	"asteriskAPI/internal/domain/entity"
	"asteriskAPI/internal/repository"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type OriginateService struct {
	repo repository.Originate
}

func NewOriginateService(repo repository.Originate) *OriginateService {
	return &OriginateService{repo: repo}
}

func (o *OriginateService) OriginateCall(initCall *dto.InitCall, table string) (*entity.InitCallResponse, error) {

	requestUrl, err1 := BuildUrl(initCall)
	if err1 != nil {
		logrus.Errorf("error occured while building url: %s", err1)
		return nil, err1
	}
	icr, err2 := PostRequest(requestUrl)
	if err2 != nil {
		logrus.Errorf("error occured while posting request to ARI: %s", err2)
		return nil, err1
	}

	err3 := o.repo.SaveOriginateCall(icr, table)
	if err3 != nil {
		logrus.Errorf("error occured while saving response to DB: %s", err3)
		return nil, err3
	}

	return icr, nil
}

func BuildUrl(ic *dto.InitCall) (*url.URL, error) {
	baseUrl, _ := os.LookupEnv("ARI_URL")
	baseUrl = baseUrl + "channels?"

	apiKey, _ := os.LookupEnv("ARI_KEY")

	u, err := url.Parse(baseUrl)
	if err != nil {
		logrus.Errorf("error occured while parsing url: %s", err.Error())
	}
	q := url.Values{}

	if ic.Endpoint == "" {
		msg := "no endpoint specified"
		logrus.Error(msg)
		return nil, errors.New(msg)
	}
	q.Set("endpoint", ic.Endpoint)
	if ic.Extension != "" {
		q.Set("extension", ic.Extension)
	}
	if ic.CallerId != "" {
		q.Set("callerId", ic.CallerId)
	}
	if ic.Context != "" {
		q.Set("context", ic.Context)
	}
	if ic.Priority != 0 {
		q.Set("priority", strconv.Itoa(ic.Priority))
	}
	if ic.Label != "" {
		q.Set("label", ic.Label)
	}
	if ic.App != "" {
		q.Set("app", ic.App)
	}
	if ic.AppArgs != "" {
		q.Set("appArgs", ic.AppArgs)
	}
	if ic.Timeout != 0 {
		q.Set("timeout", strconv.Itoa(ic.Timeout))
	}
	if ic.Variables != "" {
		q.Set("variables", ic.Variables)
	}
	if ic.ChannelId != "" {
		q.Set("channelId", ic.ChannelId)
	}
	if ic.OtherChannelId != "" {
		q.Set("otherChannelId", ic.OtherChannelId)
	}
	if ic.Originator != "" {
		q.Set("originator", ic.Originator)
	}
	if ic.Formats != "" {
		q.Set("formats", ic.Formats)
	}

	q.Set("api_key", apiKey)

	u.RawQuery = q.Encode()

	return u, nil
}

func PostRequest(requestUrl *url.URL) (*entity.InitCallResponse, error) {

	icr := entity.NewICR()
	body := bytes.NewReader(nil)

	resp, err := http.Post(requestUrl.String(),
		"application/json",
		body)
	if err != nil {
		logrus.Println(err)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("Error while reading response:", err)
		return nil, err
	}

	err = json.Unmarshal(responseBody, icr)
	if err != nil {
		logrus.Fatal("Error while decoding:", err)
		return nil, err
	}

	return icr, nil
}
