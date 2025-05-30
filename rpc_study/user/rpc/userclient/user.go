// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: user.proto

package userclient

import (
	"context"

	"go-zeroframework/rpc_study/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserCreateRequest  = user.UserCreateRequest
	UserCreateResponse = user.UserCreateResponse
	UserInfoRequest    = user.UserInfoRequest
	UserInfoResponse   = user.UserInfoResponse

	User interface {
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUser) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserCreate(ctx, in, opts...)
}
