package questions

import (
	"StandardProject/sync/ebag/global"
	"StandardProject/sync/ebag/helper"
	"StandardProject/sync/ebag/redis"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

var (
	TeachingAdminClient = helper.GetClient(2)
)

type response struct {
	No      int     `json:"F_responseNo"`
	Msg     string  `json:"F_responseMsg"`
	Student Student `json:"F_student_personal"`
}

type Student struct {
	Account    string `json:"F_account"`
	Name       string `json:"F_name"`
	Nickname   string `json:"F_nickname"`
	SchoolId   int    `json:"F_school_id"`
	SchoolName string `json:"F_school_name"`
	ClassId    string `json:"F_class_id"`
	ClassName  string `json:"F_class_name"`
}

// 获取学生的班级和学校
func GetStudentClassAndSchoolHttp(studentId string) (stu Student, err error) {
	v := url.Values{}
	v.Set("F_student_id", studentId)
	v.Set("F_accesstoken", global.MyConfig.TeachingAdminToken)

	url := global.MyConfig.TeachingAdminDomain + "/v1/teaching/info/student/personal?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")

	resp, err := TeachingAdminClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyByte, _ := ioutil.ReadAll(resp.Body)

	var data response
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		return
	}
	if data.No != 10000 {
		return stu, errors.New(fmt.Sprintf("err——>code:%v,msg:%v", data.No, data.Msg))
	}

	return data.Student, nil
}

func SaveStudentToRedis(studentIds []string, coroutineNum int) {
	coroutineChan := make(chan int, coroutineNum)

	var wg sync.WaitGroup

	for i := 0; i < len(studentIds); i++ {
		wg.Add(1)
		coroutineChan <- 1
		go func(stuId string) {
			defer wg.Done()
			student, err := GetStudentClassAndSchoolHttp(stuId)
			<-coroutineChan
			if err != nil {
				fmt.Println("GetStudentClassAndSchoolHttp——err:", err)
			} else {
				jsonData, _ := json.Marshal(student)
				redis.RedisObj.Set(redis.STUDENT_SCHOOL+stuId, string(jsonData))
			}
		}(studentIds[i])
	}
	wg.Wait()
}
