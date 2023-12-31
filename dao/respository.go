package dao

import (
	"github.com/universalmacro/common/singleton"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func NewSingletonRepository[T any]() singleton.Singleton[Repository[T]] {
	return singleton.NewSingleton[Repository[T]](NewRepository[T], singleton.Eager)
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{
		DB: GetDBInstance(),
	}
}

func (r Repository[T]) GetById(id uint) (*T, *gorm.DB) {
	var dest T
	ctx := r.DB.Find(&dest, id)
	if ctx.RowsAffected == 0 {
		return nil, ctx
	}
	return &dest, ctx
}

func (r Repository[T]) Create(dest *T) (*T, *gorm.DB) {
	ctx := r.DB.Create(dest)
	return dest, ctx
}

func (r Repository[T]) FindOne(conds ...any) (*T, *gorm.DB) {
	var dest T
	ctx := r.DB.Find(&dest, conds...)
	if ctx.RowsAffected == 0 {
		return nil, ctx
	}
	return &dest, ctx
}

func (r Repository[T]) FindMany(conds ...any) ([]T, *gorm.DB) {
	var dests []T
	ctx := r.DB.Find(&dests, conds...)
	return dests, ctx
}

func (r Repository[T]) Update(dest *T) (*T, *gorm.DB) {
	ctx := r.DB.Save(dest)
	return dest, ctx
}

func (r Repository[T]) Delete(dest *T) (*T, *gorm.DB) {
	ctx := r.DB.Delete(dest)
	return dest, ctx
}
