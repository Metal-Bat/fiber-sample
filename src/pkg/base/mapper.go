package base

import "context"

type Mapper[E, L, D, C, U any] interface {
	ToList(*E) *L
	ToListSlice([]*E) []*L
	ToDetail(*E) *D
	ToEntity(*C) *E
	ToUpdateEntity(*U) *E
}

type Hooks[E any] struct {
	BeforeCreate func(ctx context.Context, entity *E) error
	BeforeUpdate func(ctx context.Context, entity *E) error
}
