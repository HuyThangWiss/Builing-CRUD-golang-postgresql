package Interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func DateLogClientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	today := time.Now()
	fmt.Println("client call services date: " + today.Format("05/12/2022 09:04:05"))
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
func MethodLogClientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("client call " + method + " service")
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}