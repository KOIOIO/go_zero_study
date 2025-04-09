package logic

import (
	"context"
	"errors"
	"go-zeroframework/rpc_study/user_api_rpc/api/internal/types"
	"go-zeroframework/rpc_study/user_api_rpc/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zeroframework/rpc_study/user_api_rpc/api/internal/svc"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(request types.UserCreateRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user.UserCreateRequest{
		UserName: request.UserName,
		UserPwd:  request.Password,
	})
	if err != nil {
		return "", err
	}
	if response.Err != "" {
		return "", errors.New(response.Err)
	}
	return
}
