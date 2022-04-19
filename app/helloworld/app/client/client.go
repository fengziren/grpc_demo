// Author: 枫梓亻
// Date: 2022-04-18 18:06:41
// LastEditTime: 2022-04-19 10:43:53
// LastEditors: 枫梓亻
// Description:
// FilePath: \grpc_demo\app\helloworld\app\client\client.go

package main

import (
	"context"
	"fmt"
	"log"

	v1 "fengziren.top/grpc_demo/app/helloworld/api/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// ALTS认证  只支持在GCP(Google Cloud Platform)上使用
	// altsTC := alts.NewClientCreds(alts.DefaultClientOptions())
	// opts = append(opts, grpc.WithTransportCredentials(altsTC))
	// SSL/TLS认证
	// opts = append(opts, grpc.WithTransportCredentials(credentials.NewClientTLSFromFile("", "")))
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	client := v1.NewHelloWorldServiceClient(conn)
	response, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "fengziren"})
	if err != nil {
		log.Fatalf("SayHello %v", err)
	}
	fmt.Println(response.Message)
	response, err = client.SayHelloAgain(context.Background(), &v1.HelloRequest{Name: "fengziren"})
	if err != nil {
		log.Fatalf("SayHelloAgain %v", err)
	}
	fmt.Println(response.Message)
	return
}
