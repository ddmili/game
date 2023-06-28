package sdk

import (
	"encoding/json"
	"fmt"
	"game/internal/lib/curl"
)

// HTTPResult http result
type HTTPResult struct {
	Code    int        `json:"code"`
	Data    resultData `json:"data"`
	Message string     `json:"message"`
}

type resultData []byte

// BaseSdk is the base SDK
type BaseSdk struct {
}

// RequestPost POST请求
func (b *BaseSdk) RequestPost(url string, data map[string]interface{}, v any) error {
	request := curl.NewRequest()

	//构建请求头信息
	headers := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}

	//请求接口
	res, err := request.SetUrl(url).SetPostData(data).SetHeaders(headers).Post()
	if err != nil {
		return fmt.Errorf("HttpRequestError: %+v", err)
	}

	err = json.Unmarshal([]byte(res.Body), v)
	if err != nil {
		return fmt.Errorf("ResponseUnmarshalError: %+v", err)
	}

	return nil
}

// RequestGet GET 请求
func (b *BaseSdk) RequestGet(url string, queries map[string]string, v any) error {
	request := curl.NewRequest()

	//构建请求头信息
	headers := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}
	//请求接口
	res, err := request.SetUrl(url).SetQueries(queries).SetHeaders(headers).Get()
	if err != nil {
		return fmt.Errorf("HttpRequestError: %+v", err)
	}

	err = json.Unmarshal([]byte(res.Body), v)
	if err != nil {
		return fmt.Errorf("ResponseUnmarshalError: %+v", err)
	}

	return nil
}
