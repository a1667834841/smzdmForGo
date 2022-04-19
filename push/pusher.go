package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"unsafe"
)

// 定义推送者，声明推送方法
type Pusher interface {
	Push(content string, contentType string)
}

type DingPusher struct {
	Token string
}

// 钉钉推送者实现推送方法
func (pusher DingPusher) Push(params DingParam) {
	Url, err := url.Parse("https://oapi.dingtalk.com/robot/send?access_token=" + pusher.Token)
	if err != nil {
		return
	}

	paramsJson, _ := json.Marshal(params)

	urlPath := Url.String()
	resp, err := http.Post(urlPath, "application/json;charset=utf-8", bytes.NewBuffer([]byte(string(paramsJson))))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	//fmt.Println(string(content))
	str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	fmt.Println(*str)

}
