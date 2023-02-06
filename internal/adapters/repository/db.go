package db

import (
	"database/sql"
	"fmt"

	"github.com/Nav1Cr0ss/s-event/config"
	"github.com/Nav1Cr0ss/s-lib/logger"
)

type DataBase struct {
}

func NewDB(c *config.Config, log *logger.Logger) *sql.DB {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
			c.Db.UserName, c.Db.Password, c.Db.Host, c.Db.Port, c.Db.Name,
		),
	)
	if err != nil {
		log.Fatalf("failed on create db conn : %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed on pinging conn : %s", err)
	}
	log.Info("db connection initialized")
	return db
}
