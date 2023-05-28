// エントリーポイントはmain関数に配置する
// なぜなら、このファイルはCLIからgo run migrate/migrate.goとして実行するため
// 上記を行うためには、メインパッケージに属している必要がある
// 実際に動くアプリケーション上では実行されない
package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

// GO_ENV=dev go run migrate/migrate.go
func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfulle Migrated")
	defer db.CloseDB(dbConn);
	dbConn.AutoMigrate(&model.User{})
}