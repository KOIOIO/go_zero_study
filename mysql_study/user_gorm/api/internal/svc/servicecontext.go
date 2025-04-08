package svc

import (
	"fmt"
	"go-zeroframework/commn/init_gorm"
	"go-zeroframework/mysql_study/user_gorm/api/internal/config"
	"go-zeroframework/mysql_study/user_gorm/model"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB := init_gorm.InitGorm(c.Mysql.DataSource)
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		panic("gorm迁移失败")
	} else {
		fmt.Println("gorm迁移成功")
	}
	return &ServiceContext{
		Config: c,
		DB:     DB,
	}
}
