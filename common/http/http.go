package http

import (
	"StandardProject/common/errorz"
	"StandardProject/common/logz"
	"StandardProject/common/tracer"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type response1 struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	// 结构类型
	CUSTOM_RESPONSE = 0
	RESPONSE1       = 1
	RESPONSE2       = 2

	// 请求体头Content——Type
	FORM = 1
	JSON = 2

	// 请求方法
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"

	//请求头
	ACCEPT       = "Accept"
	CONTENT_TYPE = "Content-Type"

	BODY_FORM = "application/x-www-form-urlencoded"
	BODY_JSON = "application/json"
)

func selectStructToResponse(bytes []byte, respType int, resultData interface{}) error {
	decoder := json.NewDecoder(strings.NewReader(string(bytes)))
	decoder.UseNumber()
	var data interface{}
	switch respType {
	case RESPONSE1:
		var resp response1
		err := decoder.Decode(&resp)
		if err != nil {
			return errorz.CodeError(errorz.ERR_DECODE, err)
		}
		if resp.Code != 10000 {
			return errorz.CodeMsg(resp.Code, resp.Msg)
		}
		data = resp.Data
	case RESPONSE2:
	default:
		err := decoder.Decode(resultData)
		if err != nil {
			return errorz.CodeError(errorz.ERR_DECODE, err)
		}
		return nil
	}
	if resultData == nil || data == nil {
		return nil
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return errorz.CodeError(errorz.ERR_MARSHAL, err)
	}
	err = json.Unmarshal(marshal, resultData)
	if err != nil {
		return errorz.CodeError(errorz.ERR_UNMARSHAL, err)
	}
	return nil
}

func recordeLog(path, method string, header, query, body interface{}, err error) {
	data := make(map[string]interface{})
	data["url"] = path
	data["method"] = method
	data["header"] = header
	data["query"] = query
	data["body"] = body
	if err != nil {
		data["err"] = err
		logz.Error(data)
	} else {
		logz.Info(data)
	}
}

func request(path string, method string, header map[string]string, query map[string]string, body interface{}, resultData interface{}, respType int, client *http.Client, appCtx context.Context) error {

	if query != nil && len(query) > 0 {
		params := url.Values{}
		for k, v := range query {
			params.Set(k, v)
		}
		path += "?" + params.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		marshal, err := json.Marshal(body)
		if err != nil {
			return errorz.CodeError(errorz.ERR_MARSHAL, err)
		}
		if header != nil {
			if val, ok := header[CONTENT_TYPE]; ok {
				switch val {
				case BODY_FORM:
					var formData map[string]string
					json.Unmarshal(marshal, &formData)
					bodyForm := url.Values{}
					for k, v := range formData {
						bodyForm.Set(k, v)
					}
					reqBody = strings.NewReader(bodyForm.Encode())
				case BODY_JSON:
					reqBody = strings.NewReader(string(marshal))
				}
			}
		} else {
			reqBody = strings.NewReader(string(marshal))
		}
	}

	req, err := http.NewRequest(method, path, reqBody)
	if err != nil {
		return errorz.CodeError(errorz.NEW_REQUEST, err)
	}

	if header != nil && len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	// 链路跟踪记录请求头信息
	tracer.InjectTracerSpan(appCtx, req.Header)

	resp, err := client.Do(req)
	if err != nil {
		return errorz.CodeError(errorz.REQUEST_ERR, err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errorz.CodeError(errorz.IO_READ_ERR, err)
	}

	if resp.StatusCode != http.StatusOK {
		return errorz.CodeMsg(resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	err = selectStructToResponse(bytes, respType, resultData)
	if err != nil {
		return err
	}
	recordeLog(path, method, header, query, body, err)
	return nil
}

func ppd(path string, method string, query map[string]string, body interface{}, resultData interface{}, bodyType, respType int, client *http.Client, appCtx context.Context) error {
	header := make(map[string]string)
	header[ACCEPT] = BODY_JSON
	switch bodyType {
	case FORM:
		header[CONTENT_TYPE] = BODY_FORM
	case JSON:
		header[CONTENT_TYPE] = BODY_JSON
	default:

	}
	return request(path, method, header, query, body, resultData, respType, client, appCtx)
}

func Get(path string, query map[string]string, resultData interface{}, respType int, client *http.Client, appCtx context.Context) error {
	header := make(map[string]string)
	header[ACCEPT] = BODY_JSON
	return request(path, GET, header, query, nil, resultData, respType, client, appCtx)
}

func Post(path string, query map[string]string, body interface{}, resultData interface{}, bodyType, respType int, client *http.Client, appCtx context.Context) error {
	return ppd(path, POST, query, body, resultData, bodyType, respType, client, appCtx)
}

func PostQuery(path string, query map[string]string, resultData interface{}, respType int, client *http.Client, appCtx context.Context) error {
	return ppd(path, POST, query, nil, resultData, 0, respType, client, appCtx)
}

func PostBodyForm(path string, body map[string]interface{}, resultData interface{}, respType int, client *http.Client, appCtx context.Context) error {
	return ppd(path, POST, nil, body, resultData, FORM, respType, client, appCtx)
}

func PostBodyJson(path string, body interface{}, resultData interface{}, respType int, client *http.Client, appCtx context.Context) error {
	return ppd(path, POST, nil, body, resultData, JSON, respType, client, appCtx)
}

func PutForm(path string, body map[string]interface{}, respType int, client *http.Client, appCtx context.Context) error {
	return ppd(path, PUT, nil, body, nil, FORM, respType, client, appCtx)
}

func Delete(path string, query map[string]string, respType int, client *http.Client, appCtx context.Context) error {
	return ppd(path, DELETE, query, nil, nil, 0, respType, client, appCtx)
}
