package main

import (
	"StandardProject/grpc/client/pd"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

//go:generate protoc --go-grpc_out=./ --go_out=./ ./proto/student.proto

func main() {
	//creds, _ := credentials.NewServerTLSFromFile("", "")
	//conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(creds))
	//无认证 grpc http 2 https
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("连接服务端失败", err)
	}
	defer conn.Close()
	client := pd.NewStudentServiceClient(conn)
	studentInfo, err := client.GetStudentByStuNumber(context.Background(), &pd.StudentRequest{
		Number: 1,
	})
	if err != nil {
		log.Fatal("调用GRPC失败", err)
	}
	fmt.Printf("RPC调用成功, 学生信息:%s\n", studentInfo.String())
}
