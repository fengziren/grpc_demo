// Author: 枫梓亻
// Date: 2022-04-18 11:01:20
// LastEditTime: 2022-04-18 11:01:21
// LastEditors: 枫梓亻
// Description:
// FilePath: \grpc_demo\template\boot\boot.go

package main

import (
	"fmt"
	"log"
	"net"

	v1 "fengziren.top/grpc_demo/app/helloworld/api/pb/v1"
	"fengziren.top/grpc_demo/app/helloworld/app/service"
	"google.golang.org/grpc"
)

func main() {
	Start()
}

func Start() {
	// 启动grpc服务
	// flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	v1.RegisterHelloWorldServiceServer(grpcServer, &service.HelloWorldServiceServer{})
	grpcServer.Serve(lis)
}
