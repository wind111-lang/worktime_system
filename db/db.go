package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local") //WIP
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func GetPerson() error {
	db := gormConnect()
	defer db.Close()
	// db.First()...WIP

	return nil
}
