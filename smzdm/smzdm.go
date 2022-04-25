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

	"ggball.com/smzdm/file"
)

type result struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Data      Data   `json:"data"`
}

type Data struct {
	Rows  []Product `json:"rows"`
	Total int       `json:"total"`
}

type Product struct {
	ArticleTitle   string `json:"article_title"`
	ArticlePrice   string `json:"article_price"`
	ArticleWorthy  string `json:"article_worthy"`
	ArticleComment string `json:"article_comment"`
	ArticleId      string `json:"article_id"`
	ArticleDate    string `json:"publish_date_lt"`
	ArticlePic     string `json:"article_pic"`
	ArticleUrl     string `json:"article_url"`
	Referral       string `json:"article_referrals"`
}

// 全局配置
var globalConf = file.Config{}

// 推送信息文件地址
var pushedPath = "./pushed.json"

// 获取商品
//  @return []product
func GetSatisfiedGoods(conf file.Config) []Product {
	globalConf = conf
	fmt.Println("开始爬取符合条件商品。。")

	// 获取已推送文章id
	pushedMap := file.ReadPusedInfo(pushedPath)

	// 符合条件的商品集合
	var satisfyGoodsList []Product

	page := 0
	for {

		var productList = []Product{}
		if len(globalConf.KeyWords) > 0 {
			for _, word := range globalConf.KeyWords {
				products := GetGoods(page, word).Data.Rows
				productList = append(productList, products...)
			}
		} else {
			// Get the good list
			productList = GetGoods(page, "").Data.Rows
		}

		// add satisfy good
		if len(productList) > 0 {
			rows := productList
			for i := 0; i < len(rows); i++ {
				good := rows[i]

				// 商品 包含 “k” 转换数字 默认给1000
				if strings.Contains(strings.ToLower(good.ArticleComment), "k") {
					good.ArticleComment = "1000"
				}

				if removeByFilterRules(good, pushedMap) {
					continue
				}

				if satisfy(good, satisfyGoodsList) {
					satisfyGoodsList = append(satisfyGoodsList, good)
				}

			}
		}

		// 页数+1
		page++
		// 延时2s
		time.Sleep(time.Duration(2) * time.Second)

		// 判断是否退出
		if shouldStop(len(satisfyGoodsList), page) || len(satisfyGoodsList) > 0 {
			break
		}
	}

	// 根据评论数排序
	sort.SliceStable(satisfyGoodsList, func(a, b int) bool {
		return strings.Compare(satisfyGoodsList[a].ArticleComment, satisfyGoodsList[b].ArticleComment) > 0
	})

	fmt.Println("结束爬取符合条件商品。。")

	// 保存推送商品，去重使用
	savePushed(pushedMap, pushedPath, satisfyGoodsList)

	return satisfyGoodsList
}

// GetGoods 获取商品集合
//  @param offset
//  @return result 商品集合
func GetGoods(page int, keword string) result {

	var res result

	params := url.Values{}
	Url, err := url.Parse("https://api.smzdm.com/v1/list")
	if err != nil {
		return res
	}
	params.Set("keyword", keword)
	params.Set("order", "time")
	params.Set("type", "good_price")
	params.Set("offset", strconv.Itoa(page*20))
	params.Set("limit", "200")

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
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
func shouldStop(length int, page int) bool {
	fmt.Println("length:" + strconv.Itoa(length) + "\n\r page:" + strconv.Itoa(page))
	//  判断数量是否超过【符合商品个数】 且 page > 20
	return length > globalConf.SatisfyNum || page > 20

}

// 根据过滤规则，去除商品
func removeByFilterRules(good Product, pushedMap map[string]interface{}) bool {
	var noNeed = false
	// 1. 文章名称 包含过滤字符 一概不要
	for j := 0; j < len(globalConf.FilterWords); j++ {
		if strings.Contains(good.ArticleTitle, globalConf.FilterWords[j]) || strings.Contains(good.ArticlePrice, globalConf.FilterWords[j]) {
			noNeed = true
			break
		}
	}

	// 2. 根据已推送文章id map 判断是否需要去除，如果已经推送过的，则去除
	_, b := pushedMap[good.ArticleId]
	if b {
		// fmt.Println(good.ArticleTitle + "文章已存在,不予添加")
		noNeed = true
	}

	// 3. 文章时间小于昨天 去除
	// var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
	nTime := time.Now()
	// 前天
	beforeYesDate := nTime.AddDate(0, 0, -2)
	dateInt64, err1 := strconv.ParseInt(good.ArticleDate, 10, 64)

	if err1 != nil {
		panic(err1)
	}

	arDate := time.Unix(dateInt64, 0)
	// fmt.Println("文章时间：" + arDate.Format(timeLayoutStr) + "昨天时间：" + beforeYesDate.Format(timeLayoutStr))
	if arDate.Before(beforeYesDate) {
		noNeed = true
	}

	return noNeed
}

// 根据规则判断符合规则的商品
func satisfy(good Product, satisfyGoodsList []Product) bool {

	// 文章名称,爆料人包含关键词 直接添加
	// if strings.Contains(good.ArticleTitle, globalConf.KeyWord) || strings.Contains(good.Referral, globalConf.KeyWord) {
	// 	// fmt.Printf("appear satisfy good: %#v", good)
	// 	return true
	// }

	// 评论 和 值率 转int
	articleComment, err1 := strconv.Atoi(good.ArticleComment)
	articleWorthy, err2 := strconv.Atoi(good.ArticleWorthy)
	if err1 != nil || err2 != nil {
		fmt.Println("goods:", good)
		panic(err1)
	}

	// 评论，值率满足要求 则添加商品
	if articleComment >= globalConf.LowCommentNum && articleWorthy >= globalConf.LowWorthyNum {
		// fmt.Printf("appear satisfy good: %#v", good)
		return true
	}

	return false
}

// 保存推送商品，去重使用
func savePushed(pushedMap map[string]interface{}, pushedPath string, satisfyGoodsList []Product) {
	tempMap := make(map[string]interface{})

	for index, value := range satisfyGoodsList {
		tempMap[value.ArticleId] = index
	}
	file.WritePushedInfo(tempMap, pushedMap, pushedPath)
}
