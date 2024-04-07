package connector

import (
	"fmt"
	"nestle/internal/config"
	"net/url"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func ConnectMSSQL(config *config.Database) *sqlx.DB {
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(config.User, config.Password),
		Host:   fmt.Sprintf("%s:%d", config.Host, config.Port),
	}
	db, err := sqlx.Connect("sqlserver", u.String())
	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(time.Second * 10)
	db.SetConnMaxLifetime(time.Minute)
	return db
}

func DisconnectMSSQL() {}
