package unit_tests

import (
	v1 "asteriskAPI/internal/delivery/http/v1"
	"asteriskAPI/internal/domain/entity"
	"asteriskAPI/internal/repository"
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testCaseSaveCall struct {
	name string
	icr  *entity.InitCallResponse
}

func TestSaveOriginateCall(t *testing.T) {
	v1.Init("../../.env", "../../configs", "config")
	testCases := []testCaseSaveCall{
		{
			name: "pass",
			icr: &entity.InitCallResponse{
				ChannelId:  "test",
				Name:       "test",
				State:      "test",
				ProtocolId: "test",
				Caller: &entity.CallerIC{
					Name:   "test",
					Number: "test",
				},
				Connected: &entity.ConnectedIC{
					Name:   "test",
					Number: "test",
				},
				AccountCode: "test",
				Dialplan: &entity.DialplanIC{
					Context:  "test",
					Exten:    "test",
					Priority: 1,
					AppName:  "test",
					AppData:  "test",
				},
				CreationTime: "test",
				Language:     "ru",
			},
		},
	}

	table := "test_arichannelspostreply"

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})

	assert.Nil(t, err)
	repos := repository.NewRepository(db, nil)
	defer repos.DeleteByChannelId(table, testCases[0].icr.ChannelId)
	err = repos.Originate.SaveOriginateCall(testCases[0].icr, table)

	assert.Nil(t, err)

	testIcr := &entity.InitCallResponse{}
	testIcr, err = repos.SelectByChannelId(table, testCases[0].icr.ChannelId)

	fmt.Println("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	fmt.Println(testIcr.ChannelId)
	assert.Equal(t, testIcr.ChannelId, testCases[0].icr.ChannelId)
}
