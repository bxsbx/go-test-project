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
	ZsClient = helper.GetClient(2)
)

// 中山题目数据结构
type ZSQuestion struct {
	ID            int      `json:"id"`            //题目id
	Subject       int      `json:"subject"`       // 学科
	Content       string   `json:"content"`       // 题目内容
	Type          int      `json:"type"`          // 题型
	Category      int      `json:"category"`      // 题型（比type更细分）
	Difficulty    int      `json:"difficulty"`    // 难度
	CorrectAnswer []string `json:"correctAnswer"` // 正确答案
	Solution      string   `json:"solution"`      // 解析
	Answer        string   `json:"answer"`        // 答案(解答题的正确答案)
	Grade         int      `json:"grade"`         // 年级
	Accessory     []struct {
		Type    int      `json:"type"` // 表示附加内容是什么，101，102时取选择题
		Options []string `json:"options"`
	} `json:"accessory"` // 题干附加内容，主要为了提取选择题的选项内容
	// 中山返回数据有两种大小题的形式，哪种有就用哪个
	RelationData []ZSQuestion `json:"relationData"` // 小题
	SubQst       []ZSQuestion `json:"subQst"`       // 小题
	PageNum      string       `json:"pageNum"`      // 题目所在页码
	Keypoint     []struct {
		Status int    `json:"status"`
		ID     int    `json:"id"`   // 知识点id
		Name   string `json:"name"` // 知识点名称
	} `json:"keypoint"` // 题目对应知识点
}

// 访问中山数据get请求
func ZsQuestionHttp(questionIds []string) ([]ZSQuestion, error) {

	v := url.Values{}

	tiemstamp := fmt.Sprintf("%d", helper.StrToTimeStamp(helper.GetNowDateTime()))

	sn := helper.Md5(global.MyConfig.DeviceId + global.MyConfig.Appsec + tiemstamp)

	v.Set("sn", sn)
	v.Set("device_id", global.MyConfig.DeviceId)
	v.Set("t", tiemstamp)

	v.Set("ids", strings.Join(questionIds, ","))

	url := global.MyConfig.ZsDomain + "/questions" + "?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := ZsClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyByte, _ := ioutil.ReadAll(resp.Body)

	var response struct {
		Msg  string       `json:"msg"`
		Data []ZSQuestion `json:"data"`
		Ok   int          `json:"ok"`
	}

	if resp.StatusCode == 200 {
		dec := json.NewDecoder(strings.NewReader(string(bodyByte)))
		dec.UseNumber()
		dec.Decode(&response)
		return response.Data, nil
	}
	return nil, errors.New("status:" + resp.Status)
}

// 获取题目详情列表
func GetZsQuestions(questionIds []string, coroutineNum int) {
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
			questions, err := ZsQuestionHttp(qIds)
			x := <-coroutineChan
			if err != nil {
				redis.RedisObj.Set(redis.FAIL_ZS_QUESTIONS+strconv.Itoa(x), qIds)
				fmt.Println("ZsQuestionHttp——err:", err)
			} else {
				for _, qs := range questions {
					redis.RedisObj.Set(redis.ZS_QUESTIONS+strconv.Itoa(qs.ID), qs)
				}
			}
		}(list)
	}
	wg.Wait()
}

// 重新获取失败的中山题目数据
func GetZsQsAgain(coroutineNum int) {
	coroutineChan := make(chan int, coroutineNum)
	var wg sync.WaitGroup
	keys, _ := redis.RedisObj.GetKeys(redis.FAIL_ZS_QUESTIONS + "*")
	for _, key := range keys {
		coroutineChan <- 1
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			qIds, _ := redis.RedisObj.GetString(k)
			qIdList := strings.Split(qIds, ",")
			questions, err := ZsQuestionHttp(qIdList)
			<-coroutineChan
			if err != nil {
				fmt.Println("ZsQuestionHttp——err:", err)
			} else {
				for _, qs := range questions {
					redis.RedisObj.Set(redis.ZS_QUESTIONS+strconv.Itoa(qs.ID), qs)
				}
				redis.RedisObj.Remove(k)
			}
		}(key)
	}
	wg.Wait()
}
