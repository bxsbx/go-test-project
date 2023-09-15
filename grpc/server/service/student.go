package service

import (
	"StandardProject/grpc/server/pd"
	"context"
	"io"
	"log"
	"time"
)

var (
	data = map[int32]*pd.Student{
		1: {
			Name: "ERIC",
			Age:  18,
		},
		2: {
			Name: "JOHN",
			Age:  19,
		},
	}
)

type StudentService struct {
	pd.UnimplementedStudentServiceServer
}

func (s *StudentService) GetStudentByStuNumber(context context.Context, request *pd.StudentRequest) (*pd.StudentResponse, error) {
	student := s.GetStudentInfo(request.Number)
	return &pd.StudentResponse{
		Student: student,
	}, nil
}

func (s *StudentService) GetStudentByStuNumberClientStream(stream pd.StudentService_GetStudentByStuNumberClientStreamServer) error {
	count := 0
	for {
		count++
		request, err := stream.Recv() // 接收多次请求
		if err == io.EOF {
			log.Println("收到EOF，结束读取")

			err = stream.SendAndClose(&pd.StudentResponse{Student: s.GetStudentInfo(1)})
			if err != nil {
				log.Fatalln(err)
				return err
			}

			log.Println("发送完成，结束调用")
			return nil
		}
		if err != nil {
			log.Fatalln(err)
			return err
		}
		log.Println("stream id:", request.Number, ", count :", count)
	}
}

func (s *StudentService) GetStudentByStuNumberServerStream(request *pd.StudentRequest, stream pd.StudentService_GetStudentByStuNumberServerStreamServer) error {
	count := 0
	for {
		count++
		response := &pd.StudentResponse{Student: s.GetStudentInfo(request.Number)}
		log.Println("发送给客户端：", response, count)

		err := stream.Send(response)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		if count >= 10 {
			// 发送结束
			log.Println("发送结束.")
			return nil
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *StudentService) GetStudentByStuNumberStockStream(stream pd.StudentService_GetStudentByStuNumberStockStreamServer) error {
	count := 0
	for {
		count++
		request, err := stream.Recv()
		if err != nil {
			log.Fatalln(err)
			return err
		}
		log.Println("收到客户端请求：", request, "count:", count)

		//time.Sleep(time.Second)
		err = stream.Send(&pd.StudentResponse{Student: s.GetStudentInfo(request.Number)})
		if err != nil {
			log.Fatalln(err)
			return err
		}
		if count > 10 {
			return nil
		}
	}
}

func (s *StudentService) GetStudentInfo(number int32) *pd.Student {
	//模拟数据
	if _, ok := data[number]; ok {
		return data[number]
	} else {
		return nil
	}
}
