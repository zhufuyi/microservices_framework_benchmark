package logic

import (
	"context"

	helloworldV1 "helloworld/helloworld/api/helloworld/v1"
	"helloworld/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Sends a greeting
func (l *SayHelloLogic) SayHello(in *helloworldV1.HelloRequest) (*helloworldV1.HelloReply, error) {
	return &helloworldV1.HelloReply{Message: in.Name}, nil
}
