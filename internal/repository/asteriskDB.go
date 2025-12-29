package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type ConfigAsteriskDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewConfigAsteriskDB() *ConfigAsteriskDB {
	return &ConfigAsteriskDB{}
}

func NewAsteriskDB(cfg ConfigAsteriskDB) (*sqlx.DB, error) {

	asteriskDbHost := cfg.Host
	asteriskDbUser := cfg.Username
	asteriskDbPassword := cfg.Password
	asteriskDbName := cfg.DBName

	connStr := asteriskDbUser +
		":" +
		asteriskDbPassword +
		"@tcp(" + asteriskDbHost + ")/" +
		asteriskDbName

	db, err := sqlx.Open(os.Getenv("ASTERISK_DB_DRIVER"), connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
