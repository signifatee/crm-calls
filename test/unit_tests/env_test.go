package unit_tests

import (
	"asteriskAPI/pkg"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnv(t *testing.T) {
	err := godotenv.Load("../../configs/.env")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	assert.Equal(t, err, nil, "No .env file found")
	result, message := pkg.ValidateConfigs()
	assert.Equal(t, result, true, fmt.Sprint(message))
}
