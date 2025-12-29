package v1

import (
	"asteriskAPI/internal/handler"
	"asteriskAPI/internal/repository"
	"asteriskAPI/internal/server"
	"asteriskAPI/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Start() {
	Init("./.env", "configs/", "config")

	//logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := server.NewConfig()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})

	if err != nil {
		logrus.Fatalf("Cannot open connection to PostgresDB: %s", err.Error())
	}

	asteriskDB, err1 := repository.NewAsteriskDB(repository.ConfigAsteriskDB{
		Host:     os.Getenv("ASTERISK_DB_HOST"),
		Port:     os.Getenv("ASTERISK_DB_PORT"),
		Username: os.Getenv("ASTERISK_DB_USER"),
		Password: os.Getenv("ASTERISK_DB_PASSWORD"),
		DBName:   os.Getenv("ASTERISK_DB_NAME"),
	})

	if err1 != nil {
		logrus.Fatalf("Cannot open connection to AsteriskDB: %s", err.Error())
	}

	repos := repository.NewRepository(db, asteriskDB)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.New(cfg, handlers.InitRoutes())

	if err := srv.Start(); err != nil {
		logrus.Fatalf("Cannot start server due to: \n %s", err.Error())
	}
}
