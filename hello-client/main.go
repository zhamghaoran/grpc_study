package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/hello-client/proto"
)

func main() {
	// 连接创建
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	//建立连接
	client := proto.NewSayHelloClient(conn)
	// 执行rpc的调用,远程的方法
	res, _ := client.SayHello(context.Background(), &proto.HelloRequest{RequestName: "123"})
	fmt.Println(res.GetResponseMsg())

}
