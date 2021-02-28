package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User ...
type User struct {
	gorm.Model
	Name string
	Age int
	Birthday time.Time
}

var db *gorm.DB
var err error

func init() {
	dsn := "host=sandbox-gin-db user=gorm dbname=gorm password=gorm sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func main()  {
	db.AutoMigrate(&User{})
	user1 := User{Name: "taro1", Age: 18, Birthday: time.Now()}

	db.Create(&user1)
	fmt.Println(user1.Name)

}
