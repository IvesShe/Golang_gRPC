# Golang gRPC

# 下載安裝

```bash
go get -u -v google.golang.org/grpc
go get -u -v github.com/golang/protobuf/protoc-gen-go

go get -u -v google.golang.org/grpc/codes
go get -u -v google.golang.org/grpc/status
go get -u -v google.golang.org/protobuf/reflect/protoreflect
go get -u -v google.golang.org/protobuf/runtime/protoimpl
```

![image](./images/20210109205042.png)

```bash



上述如果卡住，可以更換下載的鏡像倉庫後，再重新下載安裝

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

![image](./images/20210109210019.png)

再重新下載

![image](./images/20210109210052.png)



# user.proto

```c++
// 版本號
syntax = "proto3";

// 指定生成user.pb.go的包名
package user;

// 定義客戶端請求的數據格式
message UserRequest{
    // 定義請求參數
    string name = 1;
}

// 定義服務端響應的數據格式
message UserResponse{
    // 定義響應參數
    int32 id = 1;
    string name = 2;
    int32 age = 3;
    
    // 字段修飾符
    // repeated表示可變數組，類似於切片類型
    repeated string hobby = 4;
}

// 相當於接口
// service定義開放調用的服務
service UserInfoService{
    // 相當於接口內的方法
    // 定義請求參數為UserRequest，響應參數為UserResponse
    rpc GetUserInfo (UserRequest) returns (UserResponse){}
}

```

# 生成.go文件

```bash
protoc -I . --go_out=plugins=grpc:. ./user.proto
```

![image](./images/20210109202623.png)