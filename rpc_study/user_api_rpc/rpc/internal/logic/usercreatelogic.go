package logic

import (
	"context"
	model2 "go-zeroframework/rpc_study/user_api_rpc/model"

	"go-zeroframework/rpc_study/user_api_rpc/rpc/internal/svc"
	"go-zeroframework/rpc_study/user_api_rpc/rpc/types/user"

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

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (pd *user.UserCreateResponse, err error) {

	pd = new(user.UserCreateResponse)
	var model model2.User
	err = l.svcCtx.DB.Take(&model, "username = ?", in.UserName).Error
	if err == nil {
		pd.Err = "该用户名已存在"
		return
	}
	model = model2.User{
		Username: in.UserName,
		Password: in.UserPwd,
	}
	err = l.svcCtx.DB.Create(&model).Error
	if err != nil {
		logx.Error(err)
		pd.Err = err.Error()
		err = nil
		return
	}
	pd.UserId = uint32(model.ID)
	return
}
