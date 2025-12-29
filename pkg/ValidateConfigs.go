package pkg

import (
	"github.com/spf13/viper"
	"os"
)

func ValidateConfigs() (err bool, msg string) {

	_, err = os.LookupEnv("API_TOKEN")
	if err == false {
		return false, "No API_TOKEN specified"
	}

	_, err = os.LookupEnv("ARI_URL")
	if err == false {
		return false, "No ARI_URL specified"
	}

	_, err = os.LookupEnv("ARI_KEY")
	if err == false {
		return false, "No ARI_KEY specified"
	}

	_, err = os.LookupEnv("DB_NAME")
	if err == false {
		return false, "No DB_NAME specified"
	}

	_, err = os.LookupEnv("DB_USER")
	if err == false {
		return false, "No DB_USER specified"
	}

	_, err = os.LookupEnv("DB_PASSWORD")
	if err == false {
		return false, "No DB_PASSWORD specified"
	}

	_, err = os.LookupEnv("ASTERISK_DB_HOST")
	if err == false {
		return false, "No ASTERISK_DB_HOST specified"
	}

	_, err = os.LookupEnv("ASTERISK_DB_USER")
	if err == false {
		return false, "No ASTERISK_DB_USER specified"
	}

	_, err = os.LookupEnv("ASTERISK_DB_PASSWORD")
	if err == false {
		return false, "No ASTERISK_DB_PASSWORD specified"
	}

	_, err = os.LookupEnv("ASTERISK_DB_NAME")
	if err == false {
		return false, "No ASTERISK_DB_NAME specified"
	}

	//Test config.yml
	res := viper.InConfig("db.host")
	if res == false {
		return false, "No db.host specified in config.yml"
	}
	res = viper.InConfig("db.port")
	if res == false {
		return false, "No db.port specified in config.yml"
	}
	res = viper.InConfig("db.ssl_mode")
	if res == false {
		return false, "No db.ssl_mode specified in config.yml"
	}

	return err, ""

}
