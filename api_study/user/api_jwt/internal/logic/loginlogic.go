package logic

import (
	"context"
	"go-zeroframework/commn/jwt"

	"go-zeroframework/api_study/user/api_jwt/internal/svc"
	"go-zeroframework/api_study/user/api_jwt/internal/types"

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
	auth := l.svcCtx.Config.Auth
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   1212,
		Username: "admin",
		Role:     2,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, err
}
