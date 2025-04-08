package logic

import (
	"context"
	"go-zeroframework/mysql_study/user_gorm/model"

	"go-zeroframework/mysql_study/user_gorm/api/internal/svc"
	"go-zeroframework/mysql_study/user_gorm/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginRequest, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.DB.Create(&model.User{
		Username: req.UserName,
		Password: req.Password,
	}).Error
	if err != nil {
		return req, err
	}
	return req, nil
}
