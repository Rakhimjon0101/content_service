package bootstrap

import (
	"projects/content_service/internal/config"
	"projects/content_service/internal/dbstore"
	blog "projects/content_service/internal/usecases"
	"projects/content_service/pkg/logger"
)

type UseCases struct { //nolint
	Blog *blog.UseCase
}

func buildUseCases(cfg config.Config, lg logger.Logger, db *dbstore.DBStore) UseCases {
	return UseCases{
		Blog: blog.New(lg, db.BlogRepo()),
	}
}
