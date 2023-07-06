package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc/hello-server/proto"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//开启端口监听
	listen, _ := net.Listen("tcp", ":9090")
	// 创建一个grpc 服务
	grpcServer := grpc.NewServer()
	// 在grpc当中注册服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}

}
