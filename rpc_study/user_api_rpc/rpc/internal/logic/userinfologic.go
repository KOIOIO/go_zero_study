package logic

import (
	"context"
	"errors"
	model2 "go-zeroframework/rpc_study/user_api_rpc/model"

	"go-zeroframework/rpc_study/user_api_rpc/rpc/internal/svc"
	"go-zeroframework/rpc_study/user_api_rpc/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	var model model2.User
	err := l.svcCtx.DB.Take(&model, in.UserId).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user.UserInfoResponse{
		UserId:   uint32(model.ID),
		UserName: model.Username,
	}, nil
}
