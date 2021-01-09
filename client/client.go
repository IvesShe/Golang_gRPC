package main

import (
	"fmt"

	user "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// 1.創建與gRPC服務端的連接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("連接異常: %s\n", err)
	}
	defer conn.Close()

	// 2.實例化gRPC客戶端
	client := user.NewUserInfoServiceClient(conn)

	// 3.組裝參數
	req := new(user.UserRequest)
	req.Name = "ives"

	// 4.調用接口
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("響應異常: %s\n", err)
	}

	fmt.Printf("響應結果: %v \n", resp)
}
