package database

import (
	"nestle/internal/config"
	"nestle/internal/database/connector"
	"nestle/internal/database/repository/query"

	"go.uber.org/zap"
)

func InitRepository(config *config.Database, logger *zap.Logger) *query.Core {
	db := connector.ConnectMSSQL(config)
	repo := query.Core{
		DB:     db,
		Logger: logger,
	}
	return &repo
}
