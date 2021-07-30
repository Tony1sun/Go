package main

import (
	pb "code.oldboyedu.com/my-micro/02/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
)

type Example struct {
}

type Foo struct {
}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rep *pb.CallResponse) error {
	fmt.Println("收到Example.Call请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no name")
	}
	rep.Message = "Exmaple.Call接收到了你的请求" + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *pb.EmptyRequest, rep *pb.EmptyResponse) error {
	fmt.Println("收到Foo.Bar请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	service.Init()
	err := pb.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		fmt.Println(err)
	}
	err = pb.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		fmt.Println(err)
	}
	if err := service.Run(); err != nil {
		fmt.Printf("服务启动错误: %s\n", err)
	}
}

// 1. micro api --handler=rpc
// 2. 运行服务端代码
