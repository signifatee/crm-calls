package repository

import (
	"asteriskAPI/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type CallInfo interface {
	GetCallInfo(callId string) (*entity.FullCallInfo, error)
	GetCallByDst(dst string) (string, error)
}

type Originate interface {
	SaveOriginateCall(icr *entity.InitCallResponse, table string) error
	SelectByChannelId(table string, channelId string) (*entity.InitCallResponse, error)
	DeleteByChannelId(table string, channelId string) error
}

type Repository struct {
	CallInfo
	Originate
}

func NewRepository(db *sqlx.DB, asteriskdb *sqlx.DB) *Repository {
	return &Repository{
		CallInfo:  NewCallInfoAsterisk(asteriskdb),
		Originate: NewOriginatePostgres(db),
	}
}
