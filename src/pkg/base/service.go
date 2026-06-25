package base

import (
	"context"
	"fmt"

	"sample/src/initializers"
	"sample/src/pkg/dto"
)

type Service[L, D, C, U any] interface {
	Name() string
	GetAll(ctx context.Context, page dto.PaginationStructure) ([]*L, int64, error)
	GetOne(ctx context.Context, id uint) (*D, error)
	Create(ctx context.Context, data *C) (*D, error)
	Update(ctx context.Context, id uint, data *U) error
	Delete(ctx context.Context, id uint) error
}

type baseService[E, L, D, C, U any] struct {
	repo   Repository[E]
	mapper Mapper[E, L, D, C, U]
	hooks  Hooks[E]
	name   string
}

func NewService[E, L, D, C, U any](repo Repository[E], mapper Mapper[E, L, D, C, U], hooks Hooks[E], name string) Service[L, D, C, U] {
	return &baseService[E, L, D, C, U]{repo: repo, mapper: mapper, hooks: hooks, name: name}
}

func (s *baseService[E, L, D, C, U]) Name() string { return s.name }

func (s *baseService[E, L, D, C, U]) GetAll(ctx context.Context, page dto.PaginationStructure) ([]*L, int64, error) {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("service.%s.GetAll", s.name))
	defer span.End()

	items, err := s.repo.GetAll(ctx, page)
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}
	count, err := s.repo.GetCount(ctx, page)
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}
	return s.mapper.ToListSlice(items), count, nil
}

func (s *baseService[E, L, D, C, U]) GetOne(ctx context.Context, id uint) (*D, error) {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("service.%s.GetOne", s.name))
	defer span.End()

	item, err := s.repo.GetOne(ctx, id)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return s.mapper.ToDetail(item), nil
}

func (s *baseService[E, L, D, C, U]) Create(ctx context.Context, data *C) (*D, error) {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("service.%s.Create", s.name))
	defer span.End()

	entity := s.mapper.ToEntity(data)
	if s.hooks.BeforeCreate != nil {
		if err := s.hooks.BeforeCreate(ctx, entity); err != nil {
			span.RecordError(err)
			return nil, err
		}
	}
	created, err := s.repo.Create(ctx, entity)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return s.mapper.ToDetail(created), nil
}

func (s *baseService[E, L, D, C, U]) Update(ctx context.Context, id uint, data *U) error {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("service.%s.Update", s.name))
	defer span.End()

	if _, err := s.repo.GetOne(ctx, id); err != nil {
		span.RecordError(err)
		return err
	}
	entity := s.mapper.ToUpdateEntity(data)
	if s.hooks.BeforeUpdate != nil {
		if err := s.hooks.BeforeUpdate(ctx, entity); err != nil {
			span.RecordError(err)
			return err
		}
	}
	if err := s.repo.Update(ctx, id, entity); err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}

func (s *baseService[E, L, D, C, U]) Delete(ctx context.Context, id uint) error {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("service.%s.Delete", s.name))
	defer span.End()

	if _, err := s.repo.GetOne(ctx, id); err != nil {
		span.RecordError(err)
		return err
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		span.RecordError(err)
		return err
	}
	return nil
}
