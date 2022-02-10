package utility

/*
- 使用原生 sql.DB 進行 db 的操作
*/

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/aws/aws-xray-sdk-go/xray"
)

var dbWithXray *sql.DB

func init() {
	username := "root"
	password := "root"
	host := getMySQLAddr()
	port := "3306"
	dbname := "myPrivateDatabase"
	charset := "utf8mb4"

	DataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
	var err error
	dbWithXray, err = xray.SQLContext("mysql", DataSourceName)

	if err != nil {
		panic(err)
	}
}

/*
SQLByXrayWithSuccess -
*/
func SQLByXrayWithSuccess(ctx context.Context) (CARS, error) {
	return doSomething(ctx, "select * from cars limit 1")
}

/*
SQLByXrayWithError -
*/
func SQLByXrayWithError(ctx context.Context) (CARS, error) {
	return doSomething(ctx, "error query")
}

func doSomething(ctx context.Context, query string) (CARS, error) {
	row := dbWithXray.QueryRowContext(ctx, query)
	var id int
	var name string
	err := row.Scan(&id, &name)
	return CARS{
		ID:   id,
		Name: name,
	}, err
}
