package check_in

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"ggball.com/smzdm/file"
	"ggball.com/smzdm/push"
	"ggball.com/smzdm/db"
)

type CheckIn struct {
	db       *db.DB
	conf     file.Config
	checks   []file.CheckInfo
}

func NewCheckIn(dbPath string) (*CheckIn, error) {
	database, err := db.NewDB(dbPath)
	if err != nil {
		return nil, err
	}
	
	if err := database.InitTables(); err != nil {
		return nil, err
	}

	return &CheckIn{
		db: database,
		// 初始化其他字段...
	}, nil
}

func (c *CheckIn) CheckInAllUsers() error {
	users, err := c.db.GetAllUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		// 使用用户信息执行签到
		// 根据具体平台执行不同的签到逻辑
		if msg, err := c.doCheckIn(user); err != nil {
			log.Printf("Failed to check in for user %s: %v,msg:%s", user.Name, err,msg)
		}
	}
	return nil
}

func (c *CheckIn) doCheckIn(user db.User) (string, error) {
	client := &http.Client{}
	url := "https://zhiyou.smzdm.com/user/checkin/jsonp_checkin"
	
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "请求失败", err
	}

	reqest.Header.Add("Cookie", user.Token)
	reqest.Header.Add("Host", "zhiyou.smzdm.com")
	reqest.Header.Add("Referer", "https://www.smzdm.com/")

	response, err := client.Do(reqest)
	if err != nil {
		return "请求失败", err
	}
	defer response.Body.Close()

	resMap := TransResToMap(response)
	returnText := returnResult(resMap)
	log.Println(returnText)

	// 推送结果
	push.PushTextWithDingDing(returnText, c.conf)

	return returnText,nil
}

func (c *CheckIn) SetConfig(conf file.Config, checks []file.CheckInfo) error {
	c.conf = conf
	c.checks = checks
	return nil
}

func (c *CheckIn) Run() {
	// 先执行数据库中的用户签到
	if err := c.CheckInAllUsers(); err != nil {
		log.Printf("Failed to check in for database users: %v", err)
	}

}

func returnResult(resMap map[string]interface{}) string {

	// 返回文字
	var returnText string

	// 结果code
	errorCode := resMap["error_code"]

	if float64(0) == errorCode {
		// 连续到天数
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
