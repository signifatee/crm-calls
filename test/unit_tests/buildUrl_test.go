package unit_tests

import (
	"asteriskAPI/internal/domain/dto"
	"asteriskAPI/internal/service"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildURL(t *testing.T) {
	// #1 test case ->
	ic1 := dto.InitCall{
		Endpoint:       "PJSIP/testnum@testtrunk",
		Extension:      "testextension",
		Context:        "testcontext",
		Priority:       99,
		Label:          "testlabel",
		App:            "testapp",
		AppArgs:        "testappargs",
		CallerId:       "testcallerid",
		Timeout:        999,
		Variables:      "testvars",
		ChannelId:      "testchannellid",
		OtherChannelId: "testotherchannelid",
		Originator:     "testoriginator",
		Formats:        "testformats",
	}

	// #2 test case -> expected error
	ic2 := dto.InitCall{
		Extension: "testextension",
		Context:   "testcontext",
		Priority:  1,
	}

	result1, _ := service.BuildUrl(&ic1)
	result2, _ := service.BuildUrl(&ic2)

	base_url, _ := os.LookupEnv("ARI_URL")
	base_url = base_url + "channels?"

	api_key, _ := os.LookupEnv("ARI_KEY")

	expected1 := base_url + "api_key=" + api_key + "&app=testapp&appArgs=testappargs&" +
		"callerId=testcallerid&channelId=testchannellid&" +
		"context=testcontext&endpoint=PJSIP%2Ftestnum%40testtrunk&" +
		"extension=testextension&formats=testformats&label=testlabel&" +
		"originator=testoriginator&otherChannelId=testotherchannelid&" +
		"priority=99&timeout=999&variables=testvars"

	//expected2 := base_url + "api_key=" + api_key +
	//	"context=from-internal&" +
	//	"extension=307&" +
	//	"priority=1"

	assert.Equal(t, expected1, result1.String())
	assert.Nil(t, result2)
}
