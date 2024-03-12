package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Option func(*gorm.DB) *gorm.DB

func ApplyOptions(db *gorm.DB, opts ...Option) *gorm.DB {
	for _, opt := range opts {
		db = opt(db)
	}
	return db
}

func Where(query interface{}, args ...interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func Limit(limit int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func Offset(offset int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}

func OrderBy(orderBy string, desc bool) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: orderBy}, Desc: desc})
	}
}
