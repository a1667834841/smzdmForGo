package smzdm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type result struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Data      data   `json:"data"`
}

type data struct {
	Rows  []Product `json:"rows"`
	Total int       `json:"total"`
}

type Product struct {
	ArticleTitle   string `json:"article_title"`
	ArticlePrice   string `json:"article_price"`
	ArticleWorthy  string `json:"article_worthy"`
	ArticleComment string `json:"article_comment"`
	ArticleId      string `json:"article_id"`
	ArticleDate    string `json:"article_date"`
	ArticlePic     string `json:"article_pic"`
	ArticleUrl     string `json:"article_url"`
}

// 配置文件
type Config struct {
	LowCommentNum int      `yaml:"lowCommentNum"`
	LowWorthyNum  int      `yaml:"lowWorthyNum"`
	SatisfyNum    int      `yaml:"satisfyNum"`
	FilterWords   []string `yaml:"filterWords"`
}

var globalConf = Config{}

// 获取值率大于80&且评论量大于10的商品
//  @return []product
func GetSatisfiedGoods(conf Config) []Product {
	globalConf = conf
	fmt.Println("开始爬取符合条件商品。。")

	// 符合条件的商品集合
	var satisfyGoodsList []Product

	page := 0
	for {

		// Get the good list
		goods := GetGoods(page)

		// add satisfy good
		if len(goods.Data.Rows) > 0 {
			rows := goods.Data.Rows
			for i := 0; i < len(rows); i++ {
				good := rows[i]

				// 过滤规则
				// 文章名称 包含特殊字符 一概不要
				var noNeed = false
				for j := 0; j < len(conf.FilterWords); j++ {
					if strings.Contains(good.ArticleTitle, conf.FilterWords[j]) {
						noNeed = true
						break
					}
				}
				if noNeed {
					continue
				}

				// 商品 包含 “k” 直接添加
				if strings.Contains(strings.ToLower(good.ArticleComment), "k") || strings.Contains(strings.ToLower(good.ArticleWorthy), "k") {
					fmt.Printf("appear satisfy good: %#v", good)
					satisfyGoodsList = append(satisfyGoodsList, good)
					continue
				}

				// 评论 和 值率 转int
				articleComment, err1 := strconv.Atoi(good.ArticleComment)
				articleWorthy, err2 := strconv.Atoi(good.ArticleWorthy)
				if err1 != nil || err2 != nil {
					fmt.Println("goods:", good)
					panic(err1)
				}

				// 评论，值率满足要求 则添加商品
				if articleComment >= conf.LowCommentNum && articleWorthy >= conf.LowWorthyNum {
					fmt.Printf("appear satisfy good: %#v", good)
					// fmt.Println("appear satisfy good:", good)
					satisfyGoodsList = append(satisfyGoodsList, good)
				}
			}
		}

		// 页数+1
		page++

		// 判断是否退出
		if len(satisfyGoodsList) > 0 && shouldStop(len(satisfyGoodsList), satisfyGoodsList[len(satisfyGoodsList)-1].ArticleDate) {
			break
		}
	}

	// 根据评论数排序
	sort.SliceStable(satisfyGoodsList, func(a, b int) bool {
		return strings.Compare(satisfyGoodsList[a].ArticleComment, satisfyGoodsList[b].ArticleComment) > 0
	})

	fmt.Println("结束爬取符合条件商品。。")

	return satisfyGoodsList
}

// GetGoods 获取商品集合
//  @param offset
//  @return result 商品集合
func GetGoods(page int) result {

	var res result

	params := url.Values{}
	Url, err := url.Parse("https://api.smzdm.com/v1/home/articles_new")
	if err != nil {
		return res
	}
	params.Set("f", "wxapp")
	params.Set("wxapp", "wxapp")
	params.Set("offset", strconv.Itoa(page*20))
	params.Set("limit", "20")

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	// fmt.Println(urlPath) //
	resp, err := http.Get(urlPath)
	if err != nil {
		return res
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	_ = json.Unmarshal(body, &res)
	// fmt.Printf("%#v", res)
	return res

}

// 根据条件 判断是否应该停止爬取
func shouldStop(length int, date string) bool {

	// 判断数量是否超过【符合商品个数】
	if length > globalConf.SatisfyNum {
		return true
	}

	// 判断文章日期是否超过昨天，超过昨天则退出
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	arDate, err1 := time.Parse("2006-01-02 15:04:05", date)

	if err1 != nil {
		panic(err1)
	}

	if err1 == nil && arDate.Before(yesTime) {
		return true
	}

	return false

}
