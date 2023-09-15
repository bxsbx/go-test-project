package main

import (
	"StandardProject/grpc/client/pd"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
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

	//单独调用
	//GetStudentByStuNumber(client)
	//客户端流
	//GetStudentByStuNumberClientStream(client)
	//服务端流
	//GetStudentByStuNumberServerStream(client)
	//双向流
	GetStudentByStuNumberStockStream(client)

}

func GetStudentByStuNumber(client pd.StudentServiceClient) {
	studentInfo, err := client.GetStudentByStuNumber(context.Background(), &pd.StudentRequest{
		Number: 1,
	})
	if err != nil {
		log.Fatal("调用GRPC失败", err)
	}
	fmt.Printf("RPC调用成功, 学生信息:%s\n", studentInfo.String())
}

func GetStudentByStuNumberClientStream(client pd.StudentServiceClient) {
	clientStream, err := client.GetStudentByStuNumberClientStream(context.Background())
	if err != nil {
		log.Fatal("调用GRPC失败", err)
	}
	count := 0
	for true {
		count++
		request := &pd.StudentRequest{Number: 1}
		log.Println("send StudentRequest:", request.Number, ", count:", count)

		err = clientStream.Send(request)
		if err != nil {
			log.Fatalln(err)
			return
		}

		//time.Sleep(time.Second)
		if count >= 10 {
			break
		}
	}

	response, err := clientStream.CloseAndRecv()
	if err != nil {

		log.Fatalln(err)
		return
	}
	log.Println("收到服务器的数据：", response)
}

func GetStudentByStuNumberServerStream(client pd.StudentServiceClient) {
	serverStream, err := client.GetStudentByStuNumberServerStream(context.Background(), &pd.StudentRequest{Number: 1})
	if err != nil {
		log.Fatal("调用GRPC失败", err)
	}
	count := 0
	// 接收多次
	for {
		count++
		response, err := serverStream.Recv()
		if err == io.EOF {
			log.Println("客户端接收完毕.")
			return
		}
		if err != nil {
			log.Fatalln(err)
			return
		}

		log.Println("客户端收到：", response, count)
	}
}

func GetStudentByStuNumberStockStream(client pd.StudentServiceClient) {
	stockStream, err := client.GetStudentByStuNumberStockStream(context.Background())
	if err != nil {
		log.Fatal("调用GRPC失败", err)
	}
	count := 0
	for {
		err := stockStream.Send(&pd.StudentRequest{Number: 1})
		if err != nil {
			log.Fatalln(err)
			return
		}
		//time.Sleep(time.Second)

		response, err := stockStream.Recv()
		if err != nil {
			log.Fatalln(err)
			return
		}
		log.Println("收到服务端数据：", response)
		count++
		if count > 2 {
			return
		}
	}
}
