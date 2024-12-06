package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql(address, usr, passwd, dbname string, port int) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", usr, passwd, address, port, dbname)
	fmt.Println(dsn)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println("connect mysql success")
	} else {
		fmt.Println("connect mysql fail")
	}
}
