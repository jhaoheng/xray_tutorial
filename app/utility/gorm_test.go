package utility

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_Insert(t *testing.T) {

	username := "root"
	password := "root"
	host := getMySQLAddr()
	port := "3306"
	dbname := "myPrivateDatabase"
	charset := "utf8mb4"

	//
	DataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
	db, _ = gorm.Open(mysql.Open(DataSourceName), &gorm.Config{})
	//
	car := &CARS{
		Name: time.November.String(),
	}
	err := car.Insert(context.TODO())
	if err != nil {
		panic(err)
	}
}
