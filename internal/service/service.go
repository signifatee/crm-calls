package service

import (
	"asteriskAPI/internal/domain/dto"
	"asteriskAPI/internal/domain/entity"
	"asteriskAPI/internal/repository"
)

type CallInfo interface {
	GetCallInfo(callId string) (*entity.FullCallInfo, error)
	ConvertToMainCallInfo(fci *entity.FullCallInfo, rawCallId string) (*entity.MainCallInfo, error)
	GetCallIdByDst(dst string) (string, error)
	ConvertSendToStorageGetFileUrl(fileName string) (fileUrl string)
}

type Originate interface {
	OriginateCall(initCall *dto.InitCall, table string) (*entity.InitCallResponse, error)
}

type Service struct {
	CallInfo
	Originate
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		CallInfo:  NewCallInfoService(repos.CallInfo),
		Originate: NewOriginateService(repos.Originate),
	}
}
