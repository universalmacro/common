package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Option func(*gorm.DB) *gorm.DB

func Where(conds ...any) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(conds)
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
