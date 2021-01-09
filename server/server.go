package main

import (
	"context"
	"fmt"
	"net"

	user "../proto"
	"google.golang.org/grpc"
)

// 定義服務端實現約定的接口
type UserInfoService struct {
}

var u = UserInfoService{}

// 實現服務端需要實現的接口
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	name := req.Name

	// 在數據庫查用戶信息
	if name == "ives" {
		resp = &user.UserResponse{
			Id:   1,
			Name: name,
			Age:  22,
			// 切片字段
			Hobby: []string{"Coding", "Photography"},
		}
	}

	err = nil
	return
}

func main() {
	// 1.監聽
	addr := "127.0.0.1:8080"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("監聽異常: %s\n", err)
	}
	fmt.Printf("開始監聽 %s \n", addr)

	// 2.實例化gRPC
	s := grpc.NewServer()

	// 3.在gRPC上註冊微服務
	// 在第二個參數要接口類型的變量
	user.RegisterUserInfoServiceServer(s, &u)

	// 4.啟動gRPC的服務端
	s.Serve(lis)
}
