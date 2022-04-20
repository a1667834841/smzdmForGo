package main

import (
	"fmt"
	"log"
	"os"

	"ggball.com/smzdm/cron"
	"ggball.com/smzdm/push"
	"ggball.com/smzdm/smzdm"
	"github.com/spf13/viper"
)

// 读取配置文件
var conf = readConf()

func main() {
	// 定时任务开启
	tick := cron.NewMyTick(conf.TickTime, requestSmzdm)
	tick.Start()

}

func requestSmzdm() {
	// 搜索商品
	products := smzdm.GetSatisfiedGoods(conf)
	// 推送商品
	pushDingDing(products)
}

// 推送钉钉
func pushDingDing(pro []smzdm.Product) {
	dingPusher := push.DingPusher{
		Token: "106aef404757b5a5c7df598663a9590f7ad67a4edd82ed255faee5dbc986776a",
	}

	// 需要提前申明数组的容量
	links := make([]push.Link, len(pro))

	for index, item := range pro {
		link := push.Link{
			Title:      item.ArticlePrice + "!【" + item.ArticleTitle + "】" + "【什么值得买】" + "\n\r",
			MessageURL: item.ArticleUrl,
			PicURL:     item.ArticlePic,
		}
		links[index] = link
	}

	feedCard := push.FeedCard{
		Links: links,
	}

	params := push.DingParam{
		MsgType:  "feedCard",
		FeedCard: feedCard,
	}

	dingPusher.Push(params)
}

// 读取配置文件
func readConf() smzdm.Config {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cnf := smzdm.Config{}
	c := &cnf
	v := viper.New()
	v.SetConfigName("config") //这里就是上面我们配置的文件名称，不需要带后缀名
	v.AddConfigPath(wd)       //文件所在的目录路径
	v.SetConfigType("yml")    //这里是文件格式类型
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置文件失败：", err)
		return cnf
	}
	configs := v.AllSettings()
	for k, val := range configs {
		v.SetDefault(k, val)
	}
	err = v.Unmarshal(c) //反序列化至结构体
	if err != nil {
		log.Fatal("读取配置错误：", err)
	}
	fmt.Print("读取配置文件成功。。")
	return cnf
}
