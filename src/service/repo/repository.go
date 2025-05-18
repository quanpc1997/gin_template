package repo

import (
	"context"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(ctx context.Context, item *T) error {
	return r.DB.WithContext(ctx).Create(item).Error
}

func (r *Repository[T]) GetByID(ctx context.Context, id any) (T, error) {
	var result T
	err := r.DB.WithContext(ctx).First(&result, "id = ?", id).Error
	return result, err
}

func (r *Repository[T]) Update(ctx context.Context, item *T) error {
	return r.DB.WithContext(ctx).Save(item).Error
}

func (r *Repository[T]) Delete(ctx context.Context, id any) error {
	return r.DB.WithContext(ctx).Delete(new(T), "id = ?", id).Error
}

type IRepository[T any] interface {
	Create(ctx context.Context, item *T) error
	GetByID(ctx context.Context, id any) (T, error)
	Update(ctx context.Context, item *T) error
	Delete(ctx context.Context, id any) error
}
