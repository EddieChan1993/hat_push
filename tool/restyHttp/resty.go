package restyHttp

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

var Client *resty.Client

//InitResty url请求curl
func InitResty() {
	Client = resty.New()
}

func Post(url string, body []byte) ([]byte, error) {
	data, err := Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(body)).
		Post(url)
	if err != nil {
		return nil, err
	}
	if data.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("返回状态码异常 %d", data.StatusCode())
	}
	return data.Body(), nil
}
