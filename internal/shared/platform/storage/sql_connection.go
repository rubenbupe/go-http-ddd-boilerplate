package storage

import (
	"database/sql"
	"fmt"
	"time"
)

type Connection struct {
	Db *sql.DB
}

var connections map[string]*Connection = map[string]*Connection{}

func CreateConnection(name string, cfg *Dbconfig) (*Connection, error) {
	connURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Database)
	db, err := sql.Open("postgres", connURI)
	if err != nil {
		return nil, err
	}

	connection := &Connection{
		Db: db,
	}

	connections[name] = connection

	return connection, nil
}

type Dbconfig struct {
	User     string        `default:"user"`
	Pass     string        `default:"password"`
	Host     string        `default:"localhost"`
	Port     uint          `default:"5432"`
	Database string        `default:"auth"`
	Timeout  time.Duration `default:"5s"`
}
