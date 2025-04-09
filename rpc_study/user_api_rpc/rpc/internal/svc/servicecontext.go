package svc

import (
	"fmt"
	"go-zeroframework/commn/init_gorm"
	model2 "go-zeroframework/rpc_study/user_api_rpc/model"
	"go-zeroframework/rpc_study/user_api_rpc/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	err := db.AutoMigrate(&model2.User{})
	if err != nil {
		panic("gorm迁移失败")
	} else {
		fmt.Println("gorm迁移成功")
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
