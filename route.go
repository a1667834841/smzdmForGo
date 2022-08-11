package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"ggball.com/smzdm/check_in"
	"ggball.com/smzdm/file"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, err := template.ParseFiles("template/html/index.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, nil)
	}

}

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	t, err := template.ParseFiles("template/" + r.URL.Path + ".html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func ReadCheckInfoHandler(w http.ResponseWriter, r *http.Request) {

	// 读取本地checkInfo文件
	jsonByte := readCheckInfoJson()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(wrapDataWithResult(string(jsonByte))))
	// fmt.Println(wrapDataWithResult(string(jsonByte)))
}

func AddCheckInfoHandler(w http.ResponseWriter, r *http.Request) {

	// 读取本地checkInfo文件
	jsonByte := readCheckInfoJson()

	// 转struct
	checks := deserializeJson(string(jsonByte))

	// 读取添加的数据
	body, _ := ioutil.ReadAll(r.Body)
	newcheckInfos := deserializeJson("[" + string(body) + "]")
	newcheckInfo := newcheckInfos[len(newcheckInfos)-1]
	newcheckInfo.Id = checks[len(checks)-1].Id + 1
	// 写入
	checks = append(checks, newcheckInfo)
	file.WriteCheckInfoJson(checks)
	w.Write([]byte(wrapDataWithResult("\"" + "添加成功" + "\"")))
	// fmt.Println(checks)
}

func CheckInHandler(w http.ResponseWriter, r *http.Request) {
	// 读取添加的数据
	body, _ := ioutil.ReadAll(r.Body)
	checkInfo := deserializeJson(string(body))[0]
	fmt.Println("checkInfo:", checkInfo)
	conf = file.Config{}
	conf.DingdingToken = "106aef404757b5a5c7df598663a9590f7ad67a4edd82ed255faee5dbc986776a"

	check_in.Run(conf, []file.CheckInfo{checkInfo})
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(wrapDataWithResult("\"" + "签到结束！" + "\"")))

}

func readCheckInfoJson() []byte {
	// 打开json文件
	jsonFile, err := os.Open("template/json/checkInfo.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	jsonByte, _ := ioutil.ReadAll(jsonFile)
	return jsonByte
}

func deserializeJson(CheckInfoJson string) []file.CheckInfo {
	// fmt.Println("CheckInfoJson:", CheckInfoJson)
	jsonAsBytes := []byte(CheckInfoJson)
	checks := make([]file.CheckInfo, 0)
	err := json.Unmarshal(jsonAsBytes, &checks)
	// fmt.Printf("%#v", checks)
	if err != nil {
		panic(err)
	}
	return checks
}

func wrapDataWithResult(data string) string {

	result := `
	{"code":"0",
	"msg":   "",
	"count": "10",
	"data":  ` + data + `}`

	return result
}
