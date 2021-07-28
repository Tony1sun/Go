package main

import (
	pb "code.oldboyedu.com/my-micro/02/proto"
	"context"
	log "github.com/Sirupsen/logrus"
	"github.com/go-micro/errors"
)
type Example struct {
	
}

type Foo struct {
	
}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, rep *pb.CallResponse) error {
	log.Print("收到Example.Call请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no name" )
	}
	rep.Message = "Exmaple.Call接收到了你的请求" + req.Name
	return nil
}


