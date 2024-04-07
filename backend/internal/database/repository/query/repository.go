package query

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Core struct {
	DB     *sqlx.DB
	Logger *zap.Logger
}
