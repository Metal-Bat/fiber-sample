package common

import (
	"context"
	"sample/src/initializers"
)

type Service interface {
	Teapot(ctx context.Context) (string, error)
}

type service struct {
	repository CommonRepository
}

func NewCommonService(r CommonRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Teapot(ctx context.Context) (string, error) {
	_, span := initializers.Tracer.Start(ctx, "service.Teapot")
	defer span.End()

	return s.repository.Teapot(ctx)
}
