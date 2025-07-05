package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc_study/user/rpc/userclient"
	"grpc_study/video/api/internal/config"
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
