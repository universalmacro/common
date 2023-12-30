package dao

import (
	"github.com/universalmacro/common/singleton"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewSingletonRepository[T any]() singleton.Singleton[Repository[T]] {
	return singleton.NewSingleton[Repository[T]](NewRepository[T], singleton.Eager)
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{
		db: GetDBInstance(),
	}
}

func (r Repository[T]) GetById(id uint) (*T, *gorm.DB) {
	var dest T
	ctx := r.db.Find(&dest, id)
	if ctx.RowsAffected == 0 {
		return nil, ctx
	}
	return &dest, ctx
}

func (r Repository[T]) Create(dest *T) (*T, *gorm.DB) {
	ctx := r.db.Create(dest)
	return dest, ctx
}

func (r Repository[T]) FindOne(dest *T, conds ...any) (*T, *gorm.DB) {
	ctx := r.db.Find(dest, conds...)
	return dest, ctx
}

func (r Repository[T]) FindMany(dest []T, conds ...any) ([]T, *gorm.DB) {
	ctx := r.db.Find(dest, conds...)
	return dest, ctx
}

func (r Repository[T]) Update(dest *T) (*T, *gorm.DB) {
	ctx := r.db.Save(dest)
	return dest, ctx
}

func (r Repository[T]) Delete(dest *T) (*T, *gorm.DB) {
	ctx := r.db.Delete(dest)
	return dest, ctx
}
