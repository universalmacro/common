package dao

import (
	"github.com/universalmacro/common/singleton"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	*gorm.DB
}

func SingletonFactoryRepository[T any](db *gorm.DB) singleton.Singleton[Repository[T]] {
	return singleton.SingletonFactory(func() *Repository[T] {
		return NewRepository[T](db)
	}, singleton.Eager)
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		DB: db,
	}
}

func (r Repository[T]) GetById(id uint) (*T, *gorm.DB) {
	var dest T
	ctx := r.Find(&dest, id)
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
	ctx := r.Find(&dest, conds...)
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

func (r Repository[T]) List(options ...Option) ([]T, *gorm.DB) {
	var dests []T
	ctx := r.DB.Model(dests)
	for _, opt := range options {

		ctx = opt(ctx)
	}
	ctx = ctx.Find(&dests)
	return dests, ctx
}

func (r Repository[T]) Pagination(index, limit int64, options ...Option) (List[T], *gorm.DB) {
	var dests []T
	ctx := r.DB.Model(dests)
	for _, opt := range options {
		ctx = opt(ctx)
	}
	ctx = ctx.Limit(int(limit)).Offset(int(index) * int(limit)).Find(&dests)
	var count int64
	ctx.Count(&count)
	return List[T]{
		Items: dests,
		Pagination: Pagination{
			Index: index,
			Limit: limit,
			Total: count,
		},
	}, ctx
}

func (this Repository[T]) Begin() *Repository[T] {
	return NewRepository[T](this.DB.Begin())
}

func (this Repository[T]) Rollback() *Repository[T] {
	this.DB.Rollback()
	return &this
}

func (this Repository[T]) Commit() *Repository[T] {
	this.DB.Commit()
	return &this
}

func (this Repository[T]) Transaction(f func(tx *gorm.DB) error) {
	this.DB.Transaction(f)
}
