package init_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm(MysqlSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlSource), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("gorm连接失败")
	} else {
		fmt.Println("gorm连接成功")
	}
	return db
}
