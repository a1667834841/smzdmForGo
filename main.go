package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ggball.com/smzdm/check_in"
	"ggball.com/smzdm/file"
	"ggball.com/smzdm/push"
	"ggball.com/smzdm/smzdm"
	"ggball.com/smzdm/trick"
	"github.com/robfig/cron"
)

var conf = file.Config{}
var checks = []file.CheckInfo{}

func main() {

	go cronForProduct()
	go cronForCheckIn()

	// 启动web服务，监听9090端口
	fmt.Println("启动web服务，监听9090端口")
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func cronForProduct() {

	// 定时搜索商品任务开启
	requestSmzdm()
	tick := trick.NewMyTick(conf.TickTime, requestSmzdm)
	tick.Start()
}

// 每天定时打卡任务开启
func cronForCheckIn() {

	c := cron.New()
	c.AddFunc(conf.Cron, func() {
		check_in.Run(conf, checks)
	})
	c.Start()
}

// 推送商品任务
func requestSmzdm() {
	// 搜索商品
	products := smzdm.GetSatisfiedGoods(conf)
	if len(products) == 0 {
		return
	}
	// 推送商品
	push.PushProWithDingDing(products, conf)
	time.Sleep(1 * time.Second)
}

func init() {

	// 读取配置文件
	conf = file.ReadConf()
	checks = file.ReadCheckInfoJsonToCheck()

	// 配置路由
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/conf", ReadCheckInfoHandler)
	http.HandleFunc("/addConf", AddCheckInfoHandler)
	http.HandleFunc("/check", CheckInHandler)
	http.HandleFunc("/html/", HtmlHandler)
}
