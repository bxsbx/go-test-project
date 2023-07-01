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
	"strconv"
	"strings"
	"sync"
)

var (
	EbagClient = helper.GetClient(2)
)

// 电子书包试卷库导出题目格式
type EbagQuestion struct {
	ID       int    `json:"id,omitempty"`
	Type     int    `json:"type"`
	Subject  int    `json:"subject"`
	Content  string `json:"content"`
	Keypoint []struct {
		F_id   int    `json:"id"`   //知识点id
		F_name string `json:"name"` //知识点名称
	} `json:"keypoint"`
	Solution  string         `json:"solution"`
	Options   interface{}    `json:"options"`
	Answer    interface{}    `json:"answer"`
	Questions []EbagQuestion `json:"questions"`
	//Origin    int            `json:"origin,omitempty"`
}

// 获取智慧课堂题目详情
func EbagQuestionHttp(questionIds []string) ([]EbagQuestion, error) {

	v := url.Values{}
	v.Add("F_api_token", global.MyConfig.Paper2Token)
	v.Add("F_resource_ids", helper.JSONStringfy(helper.StringToIntArray(questionIds)))

	url := global.MyConfig.Paper2Domain + "/v1/resource/question/info" + "?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := EbagClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyByte, _ := ioutil.ReadAll(resp.Body)

	var response struct {
		F_responseMsg string `json:"F_responseMsg"`
		F_responseNo  int    `json:"F_responseNo"`
		F_Info        []struct {
			FResourceId int64        `json:"F_resource_id"`
			FTypeDetail int          `json:"F_type_detail"`
			FData       EbagQuestion `json:"F_data"`
		} `json:"F_info"`
	}

	err = json.Unmarshal(bodyByte, &response)
	if err != nil {
		return nil, err
	}

	if response.F_responseNo != 10000 {
		return nil, errors.New("获取智慧课堂题目详情失败,F_responseMsg:" + response.F_responseMsg)
	}

	var ebagQuestions []EbagQuestion
	for _, info := range response.F_Info {
		ebagQuestions = append(ebagQuestions, info.FData)
	}

	return ebagQuestions, nil
}

// 获取题目详情列表
func GetEbagQuestions(questionIds []string, coroutineNum int) {
	//最大批量数
	maxLimit := 50

	coroutineChan := make(chan int, coroutineNum)

	var wg sync.WaitGroup

	// 批量限制拆分
	for i := 0; i < len(questionIds); i += maxLimit {
		var list []string
		if i+maxLimit < len(questionIds) {
			list = questionIds[i : i+maxLimit]
		} else {
			list = questionIds[i:]
		}
		wg.Add(1)
		coroutineChan <- i
		go func(qIds []string) {
			defer wg.Done()
			questions, err := EbagQuestionHttp(qIds)
			x := <-coroutineChan
			if err != nil {
				redis.RedisObj.Set(redis.FAIL_EBAG_QUESTIONS+strconv.Itoa(x), strings.Join(qIds, ","))
				fmt.Println("EbagQuestionHttp——err:", err)
			} else {
				for _, qs := range questions {
					jsonData, _ := json.Marshal(qs)
					redis.RedisObj.Set(redis.EBAG_QUESTIONS+strconv.Itoa(qs.ID), string(jsonData))
				}
			}
		}(list)
	}
	wg.Wait()
}

// 重新获取失败的智慧课堂题目数据
func GetEbagQsAgain(coroutineNum int) {
	coroutineChan := make(chan int, coroutineNum)
	var wg sync.WaitGroup
	keys, _ := redis.RedisObj.GetKeys(redis.FAIL_EBAG_QUESTIONS + "*")
	for _, key := range keys {
		coroutineChan <- 1
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			qIds, _ := redis.RedisObj.GetString(k)
			qIdList := strings.Split(qIds, ",")
			questions, err := EbagQuestionHttp(qIdList)
			<-coroutineChan
			if err != nil {
				fmt.Println("EbagQuestionHttp——err:", err)
			} else {
				for _, qs := range questions {
					jsonData, _ := json.Marshal(qs)
					redis.RedisObj.Set(redis.EBAG_QUESTIONS+strconv.Itoa(qs.ID), string(jsonData))
				}
				redis.RedisObj.Remove(k)
			}
		}(key)
	}
	wg.Wait()
}
