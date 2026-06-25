package common

import (
	"context"
	"sample/src/initializers"
)

type CommonRepository interface {
	Teapot(ctx context.Context) (string, error)
}

type repository struct{}

func NewCommonRepository() CommonRepository {
	return &repository{}
}

func (r *repository) Teapot(ctx context.Context) (string, error) {
	_, span := initializers.Tracer.Start(ctx, "repository.Teapot")
	defer span.End()

	return "i am a teapot", nil
}
