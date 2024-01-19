package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var dbSingleton singleton.Singleton[gorm.DB] // singleton.NewSingleton[gorm.DB](CreateDBInstance, singleton.Lazy)

// func GetDBInstance() *gorm.DB {
// 	return dbSingleton.Get()
// }

// func Init(constructor func() *gorm.DB) {
// 	dbSingleton = singleton.NewSingleton[gorm.DB](constructor, singleton.Lazy)
// }

// func CreateDBInstance() *gorm.DB {
// 	db, err := NewConnection(
// 		config.GetString("database.user"),
// 		config.GetString("database.password"),
// 		config.GetString("database.host"),
// 		config.GetString("database.port"),
// 		config.GetString("database.database"),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

func NewConnection(user, password, host, port, database string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
