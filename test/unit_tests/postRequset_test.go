package unit_tests

import (
	"asteriskAPI/internal/domain/dto"
	"asteriskAPI/internal/service"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type testCase struct {
	name   string
	ic     *dto.InitCall
	expect string
}

func TestPostRequest(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		logrus.Fatal("No .env file found")
	}

	testCases := []testCase{
		{
			name: "pass",
			ic: &dto.InitCall{
				Endpoint:  "PJSIP/00179875674432@infinity_outbound",
				Extension: "201",
				Context:   "from-internal",
				Priority:  1,
				Timeout:   30,
			},
			expect: "1",
		},
		{
			name: "no endpoint, fail on build url",
			ic: &dto.InitCall{
				Extension: "201",
				Context:   "from-internal",
				Priority:  1,
				Timeout:   30,
			},
			expect: "\"Key: 'InitCall.Endpoint' Error:Field validation for 'Endpoint' failed on the 'required' tag\"",
		},
	}

	expected1, err := strconv.ParseFloat(testCases[0].expect, 64)
	assert.Nil(t, err)

	testUrl1, errUrl1 := service.BuildUrl(testCases[0].ic)
	fmt.Println(testUrl1)
	assert.Nil(t, errUrl1)
	_, errUrl2 := service.BuildUrl(testCases[1].ic)
	assert.NotNil(t, errUrl2)

	result1, err1 := service.PostRequest(testUrl1)
	assert.Equal(t, expected1, result1.Dialplan.Priority)
	assert.Nil(t, err1)

}
