// Author: 枫梓亻
// Date: 2022-04-18 15:14:40
// LastEditTime: 2022-04-18 15:18:23
// LastEditors: 枫梓亻
// Description:
// FilePath: \grpc_demo\app\helloworld\app\service\helloworld.go

package service

import (
	"context"

	v1 "fengziren.top/grpc_demo/app/helloworld/api/pb/v1"
)

type HelloWorldServiceServer struct {
	v1.UnimplementedHelloWorldServiceServer
}

// Sends a greeting
func (s *HelloWorldServiceServer) SayHello(ctx context.Context, in *v1.HelloRequest) (out *v1.HelloReply, err error) {
	out = &v1.HelloReply{Message: "Hello " + in.GetName()}
	return
}

// Sends another greeting
func (s *HelloWorldServiceServer) SayHelloAgain(ctx context.Context, in *v1.HelloRequest) (out *v1.HelloReply, err error) {
	out = &v1.HelloReply{Message: "Hello again " + in.GetName()}
	return
}
