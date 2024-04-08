package server

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/zhufuyi/sponge/pkg/servicerd/registry"
)

// GrpcOption grpc settings
type GrpcOption func(*grpcOptions)

type grpcOptions struct {
	readTimeout  time.Duration
	writeTimeout time.Duration
	instance     *registry.ServiceInstance
	iRegistry    registry.Registry
}

func defaultGrpcOptions() *grpcOptions {
	return &grpcOptions{
		readTimeout:  time.Second * 3,
		writeTimeout: time.Second * 3,
		instance:     nil,
		iRegistry:    nil,
	}
}

func (o *grpcOptions) apply(opts ...GrpcOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithGrpcReadTimeout setting up read timeout
func WithGrpcReadTimeout(timeout time.Duration) GrpcOption {
	return func(o *grpcOptions) {
		o.readTimeout = timeout
	}
}

// WithGrpcWriteTimeout setting up writer timeout
func WithGrpcWriteTimeout(timeout time.Duration) GrpcOption {
	return func(o *grpcOptions) {
		o.writeTimeout = timeout
	}
}

// WithGrpcRegistry registration services
func WithGrpcRegistry(iRegistry registry.Registry, instance *registry.ServiceInstance) GrpcOption {
	return func(o *grpcOptions) {
		o.iRegistry = iRegistry
		o.instance = instance
	}
}

// -----------------------------------------------------------------------

// UnaryServerSimpleLog server-side log unary interceptor, only print response
func UnaryServerSimpleLog(log *zap.Logger) grpc.UnaryServerInterceptor {
	if log == nil {
		log, _ = zap.NewProduction()
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		resp, err := handler(ctx, req)

		fields := []zap.Field{
			zap.String("code", status.Code(err).String()),
			zap.Error(err),
			zap.String("type", "unary"),
			zap.String("method", info.FullMethod),
			zap.Int64("time_us", time.Since(startTime).Microseconds()),
		}

		log.Info("[GRPC]", fields...)

		return resp, err
	}
}

// StreamServerSimpleLog Server-side log stream interceptor, only print response
func StreamServerSimpleLog(log *zap.Logger) grpc.StreamServerInterceptor {
	if log == nil {
		log, _ = zap.NewProduction()
	}

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now()

		err := handler(srv, stream)

		fields := []zap.Field{
			zap.String("code", status.Code(err).String()),
			zap.String("type", "stream"),
			zap.String("method", info.FullMethod),
			zap.Int64("time_us", time.Since(startTime).Microseconds()),
		}
		log.Info("[GRPC]", fields...)

		return err
	}
}
