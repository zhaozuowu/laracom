package api

import (
	"context"
	"encoding/json"
	"fmt"
	proto "github.com/laracom/demoservice/proto/demo"
	"google.golang.org/grpc"
	"net/http"
)

func StartHttpServe(port string) {

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {

		client, err := grpc.Dial(":9999", grpc.WithInsecure())

		if err != nil {

			data := map[string]interface{}{
				"code":    -1,
				"message": fmt.Sprintf("服务调用超时:%v", err),
				"data":    []interface{}{},
			}
			result, _ := json.Marshal(data)
			writer.Write(result)
		}

		obj := proto.NewDemoServiceClient(client)

		res, err := obj.SayHello(context.TODO(), &proto.DemoRequst{
			Name: "赵作武",
		})

		if err != nil {
			data := map[string]interface{}{
				"code":    -1,
				"message": fmt.Sprintf("返回结果异常:%v", err),
				"data":    []interface{}{},
			}
			result, _ := json.Marshal(data)
			writer.Write(result)
		}

		data := map[string]interface{}{
			"code": 1,
			"data": map[string]string{"name": res.Text},
		}
		result, _ := json.Marshal(data)
		writer.Write(result)

	})

	if err := http.ListenAndServe(port, nil); err != nil {

		fmt.Printf("服务器连接失败")
	}

}
