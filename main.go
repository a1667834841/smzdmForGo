package main

import (
	"ggball.com/smzdm/check_in"
	"ggball.com/smzdm/file"
	"ggball.com/smzdm/push"
	"ggball.com/smzdm/smzdm"
	"ggball.com/smzdm/trick"
	"github.com/robfig/cron"
)

var conf = file.Config{}

func main() {

	// 定时搜索商品任务开启
	requestSmzdm()
	tick := trick.NewMyTick(conf.TickTime, requestSmzdm)
	tick.Start()

	// 每天定时打卡任务开启
	c := cron.New()
	c.AddFunc(conf.Cron, func() {
		check_in.Run(conf)
	})
	c.Start()

}

// 推送商品任务
func requestSmzdm() {
	// 搜索商品
	products := smzdm.GetSatisfiedGoods(conf)
	// 推送商品
	push.PushProWithDingDing(products, conf)
}

func init() {
	// 写入命令行
	// file.InputCmd()
	// 读取配置文件
	conf = file.ReadConf()
}
