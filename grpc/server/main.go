package main

import (
	"StandardProject/grpc/server/pd"
	"StandardProject/grpc/server/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

//go:generate protoc --go-grpc_out=. --go_out=. ./proto/*
func main() {
	//creds, _ := credentials.NewServerTLSFromFile("", "")
	//server := grpc.NewServer(grpc.Creds(creds))
	var authInterceptor grpc.UnaryServerInterceptor
	authInterceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		err = Auth(ctx)
		if err != nil {
			return
		}
		return handler(ctx, req)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
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

func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credential")
	}
	var user string
	var password string

	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}
	if user != "admin" || password != "admin" {
		return status.Errorf(codes.Unauthenticated, "token 不合法")
	}
	return nil
}
