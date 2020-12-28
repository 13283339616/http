package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//调取curl
func Curl(url, method string, data interface{}, headerMap map[string]string) (string, error) {

	//序列化数据 对象或者map
	jsonStr, _ := json.Marshal(data)
	//创建请求兑对象
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	//设置请求头
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	//封装客户端对象 设置超时为15秒
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	//请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	return content, nil
}

func CurlRes(url, method string, data interface{}, headerMap map[string]string, act interface{}) (string, *http.Response, error) {

	//序列化数据 对象或者map
	jsonStr, _ := json.Marshal(data)
	//创建请求兑对象
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	//设置请求头
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}
	if err != nil {
		return "", nil, err
	}
	defer req.Body.Close()

	//封装客户端对象 设置超时为15秒
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	//请求
	resp, err := client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		return "", nil, err
	}
	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	return content, resp, nil
}
