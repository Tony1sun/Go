package main

import (
	pb "code.oldboyedu.com/gRPC/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// 1.先连接服务端
// 2.实例&RPC客户端
// 3.调用

func main() {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常: %s\n", err)
	}
	defer conn.Close()
	// 2.实例化gRPC客户端
	client := pb.NewUserInfoServiceClient(conn)
	// 3. 组装请求参数
	req := new(pb.UserRequest)
	req.Name = "zs"
	// 4.调用接口
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常: %s\n", err)
	}
	fmt.Printf("响应结构: %v \n", response)
}
