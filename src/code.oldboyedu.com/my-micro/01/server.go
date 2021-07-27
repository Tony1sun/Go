package main

import (
	pb "code.oldboyedu.com/my-micro/01/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
)

type Hello struct{}

func (g *Hello) Info(ctx context.Context, req *pb.InfoRequest, rep *pb.InfoReponse) error {
	rep.Msg = "你好" + req.Username
	return nil
}

func main() {
	// 1.得到服务端实例
	service := micro.NewService(
		// 设置微服务的名，用来访问
		// micro call hello Hello.Info {\"username\":\"zhangsan\"}
		micro.Name("hello"),
		micro.Metadata(map[string]string{"protocol": "http"}),
	)
	// 2.初始化
	service.Init()
	// 3.服务注册
	err := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
