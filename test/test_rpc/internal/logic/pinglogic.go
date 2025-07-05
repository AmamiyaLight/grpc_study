package logic

import (
	"context"
	"grpc_study/test/test_rpc/internal/svc"
	"grpc_study/test/test_rpc/test_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *test_rpc.Request) (*test_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &test_rpc.Response{}, nil
}
