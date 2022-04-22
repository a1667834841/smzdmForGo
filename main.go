package main

import (
	"ggball.com/smzdm/cron"
	"ggball.com/smzdm/file"
	"ggball.com/smzdm/push"
	"ggball.com/smzdm/smzdm"
)

var conf = file.Config{}

func main() {

	// 定时任务开启
	requestSmzdm()
	tick := cron.NewMyTick(conf.TickTime, requestSmzdm)
	tick.Start()

}

// 推送商品任务
func requestSmzdm() {
	// 搜索商品
	products := smzdm.GetSatisfiedGoods(conf)
	// 推送商品
	push.PushDingDing(products, conf)
}

func init() {
	// 写入命令行
	// file.InputCmd()
	// 读取配置文件
	conf = file.ReadConf()
}
