package main

import (
	pb "code.oldboyedu.com/gRPC/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// 1.需要监听
// 2.需要实例化gRPC服务端
// 3.在gRPC上注册微服务
// 4.启动服务端

// 定义空接口
type UserInfoService struct{}

var u = UserInfoService{}

// 实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	// 通过用户名查询用户信息
	//name := req.Name
	//// 数据库里查用户信息
	//if name == "zs"{
	//	resp = &pb.UserResponse{
	//		Id: 1,
	//		Name: name,
	//		Age: 22,
	//		Hobby: []string{"Sing", "Run"},
	//	}
	//}
	//return
	return &pb.UserResponse{Id: 1, Name: req.Name, Age: 22, Hobby: []string{"Sing"}}, nil
}

func main() {
	// 地址
	addr := "127.0.0.1:8080"
	// 1.监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口:%s\n", addr)
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3. 在gRPC上注册微服务
	pb.RegisterUserInfoServiceServer(s, &u)
	// 启动 服务端
	s.Serve(listener)
}
