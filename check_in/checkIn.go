package check_in

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"ggball.com/smzdm/file"
	"ggball.com/smzdm/push"
)

func Run(conf file.Config, checks []file.CheckInfo) {

	for _, check := range checks {
		client := &http.Client{}
		//生成要访问的url
		url := "https://zhiyou.smzdm.com/user/checkin/jsonp_checkin"
		//提交请求
		reqest, err := http.NewRequest("GET", url, nil)

		//增加header选项
		reqest.Header.Add("Cookie", check.Cookie)
		reqest.Header.Add("Host", "zhiyou.smzdm.com")
		reqest.Header.Add("Referer", "https://www.smzdm.com/")

		if err != nil {
			panic(err)
		}
		//处理返回结果
		response, _ := client.Do(reqest)
		defer response.Body.Close()

		// 将json转为map
		resMap := TransResToMap(response)

		returnText := returnResult(resMap)

		// 推送
		push.PushTextWithDingDing(returnText, conf)

		// 修改用戶最近一次签到结果
		resultCode := ""
		if resMap["error_code"] == float64(0) {
			resultCode = "success"
		} else {
			resultCode = "faild"
		}
		file.UpdateCheckInfoById(check.Id, resultCode, returnText)
	}

}

func returnResult(resMap map[string]interface{}) string {

	// 返回文字
	var returnText string

	// 结果code
	errorCode := resMap["error_code"]

	if float64(0) == errorCode {
		// 连续签到天数
		data := resMap["data"].(map[string]interface{})
		continueCheckinDays := data["continue_checkin_days"]

		fmt.Println("continueCheckinDays:", reflect.TypeOf(continueCheckinDays))
		v, ok := continueCheckinDays.(float64)
		if ok {
			// 签到成功
			fmt.Println("nor:", v)
			returnText = "恭喜签到成功！您已连续签到" + strconv.FormatFloat(continueCheckinDays.(float64), 'f', 0, 64) + "天!"
		} else {
			fmt.Println("err:", v)
			returnText = "error" + resMap["error_msg"].(string)
		}

	} else {
		// 签到失败
		returnText = "很遗憾，签到失败！"
	}

	return returnText
}

func TransResToMap(res *http.Response) map[string]interface{} {
	// 返回json
	var resText map[string]interface{}

	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal([]byte(string(body)), &resText)
	fmt.Println(resText)
	return resText
}
