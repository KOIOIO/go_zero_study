package logic

import (
	"context"
	model2 "go-zeroframework/rpc_study/user_gorm/model"

	"go-zeroframework/rpc_study/user_gorm/rpc/internal/svc"
	"go-zeroframework/rpc_study/user_gorm/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	// todo: add your logic here and delete this line
	var model model2.User
	err := l.svcCtx.DB.Take(&model, "username = ?", in.UserName).Error
	if err == nil {
		return &user.UserCreateResponse{
			Err: "用户已存在",
		}, nil
	}
	model = model2.User{
		Username: in.UserName,
		Password: in.UserPwd,
	}
	err = l.svcCtx.DB.Create(&model).Error
	if err != nil {
		return &user.UserCreateResponse{
			Err: "用户创建失败",
		}, nil
	}

	return &user.UserCreateResponse{
		UserId: uint32(model.ID),
		Err:    "200",
	}, nil
}
