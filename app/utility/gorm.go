package utility

/*
- 使用 gorm 套件, 進行 db 的操作

## 已知 issue
目前 gorm 不支援 xray 的 subsegment, 所以會出現
`panic: failed to begin subsegment named 'myPrivateDatabase': segment cannot be found.`

## 解決方法
就只能透過 xray.BeginSubsegment(ctx, "gorm") 的方式
在既有的 ctx 下, 建立 Subsegment
- 缺點非常明顯就是, code 要增加許多, 因為要關閉 Subsegment
*/

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

	//
	DataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
	var err error

	//
	db, err = gorm.Open(mysql.Open(DataSourceName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//
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
func (cars *CARS) Insert(ctx context.Context) error {
	/*
		- 無法用其他方式取得 query, 所以使用此方法取得
	*/
	query := db.Session(&gorm.Session{DryRun: true}).Create(cars).Statement.SQL.String()

	//
	tx := XrayGormWrap(ctx, func() *gorm.DB {
		return db.Create(cars)
	}, query)
	return tx.Error
}

/*
GetAll ...
*/
func (cars *CARS) GetAll(ctx context.Context) ([]CARS, error) {
	c := []CARS{}

	query := db.Session(&gorm.Session{DryRun: true}).Find(&c).Statement.SQL.String()
	tx := XrayGormWrap(ctx, func() *gorm.DB {
		return db.Find(&c)
	}, query)

	return c, tx.Error
}

/*
Delete ...
*/
func (cars *CARS) Delete(ctx context.Context) error {
	query := db.Session(&gorm.Session{DryRun: true}).Delete(cars).Statement.SQL.String()
	tx := XrayGormWrap(ctx, func() *gorm.DB {
		return db.Delete(cars)
	}, query)

	return tx.Error
}
