package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zeroframework/demo/video/api/internal/config"
	"go-zeroframework/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
