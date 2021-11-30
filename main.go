package main

import (
	"jlb_shop_go/boot"
	"jlb_shop_go/global"
)

func newTask() {

}

func main() {

	global.Viper = boot.Viper()
	global.Db = boot.Gorm()

	if global.Db != nil {
		db, _ := global.Db.DB()
		defer db.Close()
	}
	boot.RunServer()
}
