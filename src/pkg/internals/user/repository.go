package user

import (
	"context"

	"gorm.io/gorm"

	"sample/src/initializers"
	"sample/src/pkg/entities"
	"sample/src/pkg/utils/customErrors"
)

type UserRepository interface {
	FindByMobileWithPermissions(ctx context.Context, mobile string) (*entities.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &repository{db: initializers.DB}
}

func (r *repository) FindByMobileWithPermissions(ctx context.Context, mobile string) (*entities.User, error) {
	_, span := initializers.Tracer.Start(ctx, "repository.User.FindByMobileWithPermissions")
	defer span.End()

	user, err := gorm.G[entities.User](r.db.WithContext(ctx)).
		Preload("Permissions", func(db gorm.PreloadBuilder) error { return nil }).
		Where("mobile = ?", mobile).
		First(ctx)

	if err != nil {
		span.RecordError(err)
		return nil, customErrors.MapError(err)
	}
	return &user, nil
}
