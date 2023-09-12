package service

import (
	"StandardProject/grpc/server/pd"
	"context"
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

func (s *StudentService) GetStudentInfo(number int32) *pd.Student {
	//模拟数据
	if _, ok := data[number]; ok {
		return data[number]
	} else {
		return nil
	}
}
