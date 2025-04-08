package logic

import (
	"context"
	"encoding/json"

	"go-zeroframework/api_study/user/api_jwt/internal/svc"
	"go-zeroframework/api_study/user/api_jwt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("user_id").(json.Number)
	username := l.ctx.Value("username").(string)
	u_id, _ := userId.Int64()
	return &types.UserInfoResponse{
		UserName: username,
		Addr:     "Beijing",
		Id:       uint(u_id),
	}, nil
}
