package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/laracom/demoservice/api"
	proto "github.com/laracom/demoservice/proto/demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	httpPort = ":9090"
	grpcPort = ":9999"
	address  = "localhost:9999"
)

type DemoService struct {
}

func (d *DemoService) SayHello(ctx context.Context, req *proto.DemoRequst) (*proto.DemoResponse, error) {

	return &proto.DemoResponse{
		Text: "你好" + req.Name,
	}, nil
}

func main() {

	mode := flag.String("mode", "grpc", "mode:grpc/http/client")
	flag.Parse()
	fmt.Printf("mode:%v\n", *mode)
	switch *mode {

	case "http":
		api.StartHttpServe(httpPort)
	case "client":
		client, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("连接服务失败:%v\n", err)
		}

		defer client.Close()

		obj := proto.NewDemoServiceClient(client)

		res, errResult := obj.SayHello(context.TODO(), &proto.DemoRequst{
			Name: "赵作武",
		})

		if errResult != nil {
			log.Fatalf("调用服务异常:%v\n", errResult)
		}

		fmt.Printf("结果是:%v\n", res.Text)

	case "grpc":
		fallthrough
	default:
		listen, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("监听端口失败:%v\n", err)
		}

		servive := grpc.NewServer()
		proto.RegisterDemoServiceServer(servive, &DemoService{})
		reflection.Register(servive)
		if err := servive.Serve(listen); err != nil {
			log.Fatalf("服务启动失败：%v\n", err)
		}

	}
}
