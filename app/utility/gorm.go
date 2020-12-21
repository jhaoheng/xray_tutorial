package utility

import (
	"context"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getMySQLAddr() string {
	MySQLAddr := os.Getenv("MYSQL_ADDR")
	if MySQLAddr != "" {
		return MySQLAddr
	}
	return "localhost"
}

var db *gorm.DB

func init() {
	username := "root"
	password := "root"
	host := getMySQLAddr()
	port := "3306"
	dbname := "myPrivateDatabase"
	charset := "utf8mb4"

	DataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
	var err error
	db, err = gorm.Open(mysql.Open(DataSourceName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&CARS{}) // 根據 module 建立 table
}

/*
CARS ...
*/
type CARS struct {
	ID   int    `gorm:"type:int"`
	Name string `gorm:"type:text"`
}

/*
TableName ...
*/
func (cars *CARS) TableName() string {
	return "cars"
}

/*
Insert ...
*/
func (cars *CARS) Insert(ctx context.Context, name string) error {
	cars.Name = name
	// return db.Create(cars).Error
	return XrayMiddle(ctx, "Insert", func() error {
		return db.Create(cars).Error
	})
}

/*
GetAll ...
*/
func (cars *CARS) GetAll(ctx context.Context) []CARS {
	c := []CARS{}
	// db.Find(&c)
	XrayMiddle(ctx, "GetAll", func() error {
		return db.Find(&c).Error
	})
	return c
}

/*
Delete ...
*/
func (cars *CARS) Delete(ctx context.Context, id int) error {
	cars.ID = id
	// return db.Delete(cars).Error
	return XrayMiddle(ctx, "Delete", func() error {
		return db.Delete(cars).Error
	})
}
