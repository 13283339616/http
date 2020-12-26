package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Curl(url, method string, data interface{}, headerMap map[string]string, act *interface{}) (*interface{}, error) {

	//序列化数据 对象或者map
	jsonStr, _ := json.Marshal(data)
	//创建请求兑对象
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	//设置请求头
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}
	if err != nil {
		panic(err)
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
		panic(err)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	//封装数据
	if err := json.Unmarshal([]byte(content), act); err == nil {
		return act, nil
	} else {
		return nil, err
	}
}
