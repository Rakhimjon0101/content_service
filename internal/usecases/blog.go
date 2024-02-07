package blog

import (
	"context"

	"projects/content_service/internal/domains"
	"projects/content_service/internal/errs"
	"projects/content_service/pkg/logger"
)

type BlogRepo interface {
	CreateBlog(blog *domains.Blog) error
}

type UseCase struct {
	repo BlogRepo
	lg   logger.Logger
}

func New(l logger.Logger, b BlogRepo) *UseCase {
	return &UseCase{repo: b, lg: l}
}

func (s UseCase) CreateBlog(ctx context.Context) error {
	err := s.repo.CreateBlog(nil)
	if err != nil {
		return errs.Errf(err, "error while getting regionCode ")
	}

	return nil
}
