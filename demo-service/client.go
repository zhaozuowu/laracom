package main

import (
	"context"
	"fmt"
	proto "github.com/laracom/demoservice/proto/demo"
	"google.golang.org/grpc"
	"log"
)

const (
	adress = "localhost:9999"
)

func main() {

	client, err := grpc.Dial(adress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接服务失败:%v\n", err)
	}

	defer client.Close()

	obj  := proto.NewDemoServiceClient(client)

	res, errResult := obj.SayHello(context.TODO(),&proto.DemoRequst{
		Name:                 "赵作武",
	})

	if errResult != nil {
		log.Fatalf("调用服务异常:%v\n",errResult)
	}

	fmt.Printf("结果是:%v\n",res.Text)

}
