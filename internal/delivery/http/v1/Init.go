package v1

import (
	"asteriskAPI/pkg"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Init(envPath string, configPath string, configName string) {
	if err := godotenv.Load(envPath); err != nil {
		logrus.Fatalf("No .env file found: %v", err)
	}
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("cannot load config: %v", err)
	}
	err, msg := pkg.ValidateConfigs()
	if err != true {
		logrus.Fatal(msg)
	}

	file, err1 := os.OpenFile("/var/log/apiAsterisk/wss.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err1 == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Infof("Failed to log to file, using default stderr: %v", err1)
	}

}
