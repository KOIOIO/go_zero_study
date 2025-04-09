package svc

import (
	"fmt"
	"go-zeroframework/commn/init_gorm"
	"go-zeroframework/rpc_study/user_gorm/model"
	"go-zeroframework/rpc_study/user_gorm/rpc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := init_gorm.InitGorm(c.Mysql.DataSource)
	err := db.AutoMigrate(&model.User{})
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
