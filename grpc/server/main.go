package main

import (
	"StandardProject/grpc/server/pd"
	"StandardProject/grpc/server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

//go:generate protoc --go-grpc_out=. --go_out=. ./proto/student.proto
func main() {
	//creds, _ := credentials.NewServerTLSFromFile("", "")
	//server := grpc.NewServer(grpc.Creds(creds))

	server := grpc.NewServer()
	studentService := service.StudentService{}

	//注册服务
	pd.RegisterStudentServiceServer(server, &studentService)
	//启动监听程序
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("启动监听失败", err)
	}
	err = server.Serve(listener)
	if err != nil {
		log.Fatal("启动服务失败", err)
	}
}
