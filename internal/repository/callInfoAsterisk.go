package repository

import (
	"asteriskAPI/internal/domain/entity"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CallInfoAsterisk struct {
	db *sqlx.DB
}

func NewCallInfoAsterisk(db *sqlx.DB) *CallInfoAsterisk {
	return &CallInfoAsterisk{db: db}
}

func (c *CallInfoAsterisk) GetCallInfo(callId string) (*entity.FullCallInfo, error) {

	celt, err1 := c.selectFromCEL(callId)
	if err1 != nil {
		logrus.Errorf("Cannot get data from Asterisk CEL table: %s", err1.Error())
		return nil, err1
	}

	cdrt, err2 := c.selectFromCDR(callId)
	if err2 != nil {
		logrus.Errorf("Cannot get data from Asterisk CDR table: %s", err2.Error())
		return nil, err2
	}

	fullInfo := &entity.FullCallInfo{
		Adcel: celt,
		Adcrd: cdrt,
	}

	if fullInfo.Adcel.Id == 0 {
		msg := "call Id not found in cel table"
		logrus.Error(msg)
		return nil, errors.New(msg)
	}

	return fullInfo, nil

}

func (c *CallInfoAsterisk) selectFromCEL(callId string) (*entity.CelTable, error) {
	celts := []entity.CelTable{}
	db := c.db

	query := "SELECT * FROM cel WHERE uniqueid=? AND eventtype='HANGUP';"
	err := db.Select(&celts, query, callId)
	if err != nil {
		logrus.Errorf("error occured while getting info from cel table: %s", err)
		return nil, err
	}

	if len(celts) == 0 {
		msg := fmt.Sprintf("no rows from cel table by id: %s", callId)
		logrus.Error(msg)
		return nil, errors.New(msg)
	}

	return &celts[0], nil
}

func (c *CallInfoAsterisk) selectFromCDR(callId string) (*entity.CdrTable, error) {
	cdrs := []entity.CdrTable{}
	db := c.db

	query := "SELECT * FROM cdr WHERE uniqueid=?;"
	err := db.Select(&cdrs, query, callId)
	if err != nil {
		logrus.Errorf("error occured while getting info from cdr table: %s", err)
		return nil, err
	}

	//cdr table can have 0 rows, if hangup cause was 0, so it's ok
	if len(cdrs) == 0 {
		logrus.Infof("no rows from cdr table: %s", err)
		return &entity.CdrTable{}, err
	}

	return &cdrs[0], nil
}

func (c *CallInfoAsterisk) GetCallByDst(dst string) (string, error) {
	query := "SELECT uniqueid FROM cdr WHERE dst=?;"
	var callId []string
	err := c.db.Select(&callId, query, dst)
	if err != nil {
		logrus.Errorf("Error while getting unique id by dst: %v", err)
		return "", err
	}
	logrus.Infof("Callid of dst %s is %s:", dst, callId)
	return callId[0], nil
}
