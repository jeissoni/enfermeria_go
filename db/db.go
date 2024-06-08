package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int64
	User     string
	Password string
	DBName   string
}

// ConnString construye el string de conexión a partir de la configuración
func (config *Config) ConnString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
}

func NewPostgreSQLStorage(config Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", config.ConnString())

	if err != nil {
		return nil, err
	}

	return db, nil
}
