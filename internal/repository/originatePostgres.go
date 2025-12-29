package repository

import (
	"asteriskAPI/internal/domain/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OriginatePostgres struct {
	db *sqlx.DB
}

func NewOriginatePostgres(db *sqlx.DB) *OriginatePostgres {
	return &OriginatePostgres{db: db}
}

func (o *OriginatePostgres) SaveOriginateCall(icr *entity.InitCallResponse, table string) error {

	query := fmt.Sprintf("INSERT INTO %s (channel_id, "+
		"name, "+
		"state, "+
		"protocol_id, "+
		"caller_name, "+
		"caller_number, "+
		"connected_name, "+
		"connected_number, "+
		"accountcode, "+
		"dialplan_context, "+
		"dialplan_exten, "+
		"dialplan_priority, "+
		"dialplan_app_name, "+
		"dialplan_app_data, "+
		"creationtime, "+
		"language) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", table)

	_, err := o.db.Exec(query,
		icr.ChannelId,
		icr.Name,
		icr.State,
		icr.ProtocolId,
		icr.Caller.Name,
		icr.Caller.Number,
		icr.Connected.Name,
		icr.Connected.Number,
		icr.AccountCode,
		icr.Dialplan.Context,
		icr.Dialplan.Exten,
		icr.Dialplan.Priority,
		icr.Dialplan.AppName,
		icr.Dialplan.AppData,
		icr.CreationTime,
		icr.Language)

	if err != nil {
		logrus.Errorf("error while inserting call info into database: %s", err)
		return err
	}

	return nil
}

func (o *OriginatePostgres) SelectByChannelId(table string, channelId string) (*entity.InitCallResponse, error) {

	//sl := []entity.InitCallResponse{}

	sl := []entity.InitCallResponseDB{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE channel_id='%s';", table, channelId)

	err := o.db.Select(&sl, query)
	if err != nil {
		logrus.Errorf("error while select call info from database: %s", err.Error())
		return nil, err
	}
	icr := sl[0].IcrDBToICR()

	return icr, nil
}

func (o *OriginatePostgres) DeleteByChannelId(table string, channelId string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE channel_id='%s';", table, channelId)

	_, err := o.db.Exec(query)

	if err != nil {
		logrus.Errorf("error while deleteing call info from database: %s", err.Error())
		return err
	}

	return nil
}
