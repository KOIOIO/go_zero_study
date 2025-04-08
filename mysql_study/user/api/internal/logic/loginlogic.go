package logic

import (
	"context"
	"go-zeroframework/mysql_study/user/model"

	"go-zeroframework/mysql_study/user/api/internal/svc"
	"go-zeroframework/mysql_study/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UsersModel.Insert(l.ctx, &model.User{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	logx.Infof("用户名:%s,密码:%s", req.UserName, req.Password)

	return "success", nil
}
