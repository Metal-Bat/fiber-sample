package base

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"sample/src/initializers"
	"sample/src/pkg/dto"
	"sample/src/pkg/utils"
	"sample/src/pkg/utils/customErrors"
)

type Repository[E any] interface {
	GetAll(ctx context.Context, page dto.PaginationStructure) ([]*E, error)
	GetCount(ctx context.Context, page dto.PaginationStructure) (int64, error)
	GetOne(ctx context.Context, id uint) (*E, error)
	Create(ctx context.Context, entity *E) (*E, error)
	Update(ctx context.Context, id uint, entity *E) error
	Delete(ctx context.Context, id uint) error
}

type gormRepository[E any] struct {
	db       *gorm.DB
	name     string
	preloads []string
}

func NewRepository[E any](db *gorm.DB, name string, preloads ...string) Repository[E] {
	return &gormRepository[E]{db: db, name: name, preloads: preloads}
}

func (r *gormRepository[E]) GetAll(ctx context.Context, page dto.PaginationStructure) ([]*E, error) {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("repository.%s.GetAll", r.name))
	defer span.End()

	q := gorm.G[E](r.db.WithContext(ctx)).Order("id")
	q = utils.ApplyFilters(ctx, q, page)
	items, err := utils.ApplyPagination(ctx, q, page).Find(ctx)
	if err != nil {
		span.RecordError(err)
		return nil, customErrors.MapError(err)
	}

	result := make([]*E, len(items))
	for i := range items {
		result[i] = &items[i]
	}
	return result, nil
}

func (r *gormRepository[E]) GetCount(ctx context.Context, page dto.PaginationStructure) (int64, error) {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("repository.%s.GetCount", r.name))
	defer span.End()

	q := gorm.G[E](r.db.WithContext(ctx)).Order("id")
	count, err := utils.ApplyFilters(ctx, q, page).Count(ctx, "id")
	if err != nil {
		span.RecordError(err)
		return 0, customErrors.MapError(err)
	}
	return count, nil
}

func (r *gormRepository[E]) GetOne(ctx context.Context, id uint) (*E, error) {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("repository.%s.GetOne", r.name))
	defer span.End()

	gq := gorm.G[E](r.db.WithContext(ctx))
	var cq gorm.ChainInterface[E]
	if len(r.preloads) > 0 {
		cq = gq.Preload(r.preloads[0], func(db gorm.PreloadBuilder) error { return nil }).Where("id = ?", id)
	} else {
		cq = gq.Where("id = ?", id)
	}
	item, err := cq.First(ctx)
	if err != nil {
		span.RecordError(err)
		return nil, customErrors.MapError(err)
	}
	return &item, nil
}

func (r *gormRepository[E]) Create(ctx context.Context, entity *E) (*E, error) {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("repository.%s.Create", r.name))
	defer span.End()

	if err := gorm.G[E](r.db.WithContext(ctx)).Create(ctx, entity); err != nil {
		span.RecordError(err)
		return nil, customErrors.MapError(err)
	}
	return entity, nil
}

func (r *gormRepository[E]) Update(ctx context.Context, id uint, entity *E) error {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("repository.%s.Update", r.name))
	defer span.End()

	if _, err := gorm.G[E](r.db.WithContext(ctx)).Where("id = ?", id).Updates(ctx, *entity); err != nil {
		span.RecordError(err)
		return customErrors.MapError(err)
	}
	return nil
}

func (r *gormRepository[E]) Delete(ctx context.Context, id uint) error {
	_, span := initializers.Tracer.Start(ctx, fmt.Sprintf("repository.%s.Delete", r.name))
	defer span.End()

	if _, err := gorm.G[E](r.db.WithContext(ctx)).Where("id = ?", id).Delete(ctx); err != nil {
		span.RecordError(err)
		return customErrors.MapError(err)
	}
	return nil
}
