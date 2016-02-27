package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type Members struct {
	Id	int64
	Name    string
}

/**
mysql insert
 */
func main() {

	db, err := gorm.Open("mysql", "root:password@tcp(1.1.1.1)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("open error", err)
	}

	defer db.Close()

	db.DB()

	newMembers := Members{
		Name: "Mike",
	}

	db.Create(&newMembers)

}