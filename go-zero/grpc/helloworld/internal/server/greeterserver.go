// Code generated by goctl. DO NOT EDIT.
// Source: greeter.proto

package server

import (
	"context"

	helloworldV1 "helloworld/helloworld/api/helloworld/v1"
	"helloworld/internal/logic"
	"helloworld/internal/svc"
)

type GreeterServer struct {
	svcCtx *svc.ServiceContext
	helloworldV1.UnimplementedGreeterServer
}

func NewGreeterServer(svcCtx *svc.ServiceContext) *GreeterServer {
	return &GreeterServer{
		svcCtx: svcCtx,
	}
}

// Sends a greeting
func (s *GreeterServer) SayHello(ctx context.Context, in *helloworldV1.HelloRequest) (*helloworldV1.HelloReply, error) {
	l := logic.NewSayHelloLogic(ctx, s.svcCtx)
	return l.SayHello(in)
}
