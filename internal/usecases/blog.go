package blog

import (
	"context"
	"fmt"

	"projects/content_service/internal/domains"
	"projects/content_service/internal/errs"
	"projects/content_service/pkg/logger"
)

type BlogRepo interface {
	CreateBlog(ctx context.Context, blog *domains.Blog) error
	GetBlogByID(ctx context.Context, id string) (domains.Blog, error)
}

type UseCase struct {
	repo BlogRepo
	lg   logger.Logger
}

func New(l logger.Logger, b BlogRepo) *UseCase {
	return &UseCase{repo: b, lg: l}
}

func (s UseCase) CreateBlog(ctx context.Context, d *domains.Blog) error {
	err := s.repo.CreateBlog(ctx, d)
	if err != nil {
		s.lg.Debug(fmt.Sprintf("blog error while creating %s", err))
		return errs.Errf(err, "error while creating blog")
	}

	return nil
}

func (s UseCase) GetBlogByID(ctx context.Context, id string) (domains.Blog, error) {
	b, err := s.repo.GetBlogByID(ctx, id)
	if err != nil {
		s.lg.Debug(fmt.Sprintf("blog error while getting %s", err))
		return b, errs.Errf(err, "error while getting blog by its id")
	}

	return b, nil
}
