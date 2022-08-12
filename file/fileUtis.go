package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// 配置文件
type Config struct {
	LowCommentNum int      `yaml:"lowCommentNum"`
	LowWorthyNum  int      `yaml:"lowWorthyNum"`
	SatisfyNum    int      `yaml:"satisfyNum"`
	TickTime      int      `yaml:"tickTime"`
	FilterWords   []string `yaml:"filterWords"`
	KeyWords      []string `yaml:"keyWords"`
	DingdingToken string   `yaml:"dingdingToken"`
	Cron          string   `yaml:"cron"`
}

// 签到信息
type CheckInfo struct {
	Id         int    `json:Id`
	LastTIme   string `json:LastTIme`
	Remark     string `json:Remark`
	LastMsg    string `json:LastMsg`
	LastResult string `json:LastResult`
	Cookie     string `json:Cookie`
}

// 读取已推送文章id 返回map
func ReadPusedInfo(path string) map[string]interface{} {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	bytesFile, _ := ioutil.ReadAll(jsonFile)
	// fmt.Println(string(bytesFile))

	pushedMap := make(map[string]interface{})
	err1 := json.Unmarshal(bytesFile, &pushedMap)
	if err1 != nil {
		panic(err1)
	}
	// fmt.Println("json to map ", pushedMap)
	return pushedMap
}

// 保存已推送文章id 到本地
func WritePushedInfo(temp map[string]interface{}, pushed map[string]interface{}, path string) {
	for key, value := range temp {
		pushed[key] = value
	}

	// json 序列化map
	data, _ := json.Marshal(pushed)

	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}

}

// 输入命令行 写入配置文件
// func InputCmd() {

// 	file, _ := exec.LookPath(os.Args[0])
// 	path, _ := filepath.Abs(file)
// 	index := strings.LastIndex(path, string(os.PathSeparator))
// 	path = path[:index]

// 	v := viper.New()
// 	// v.SetConfigName("config") //这里就是上面我们配置的文件名称，不需要带后缀名
// 	// v.AddConfigPath("../")    //文件所在的目录路径
// 	// v.SetConfigType("yml")    //这里是文件格式类型
// 	v.SetConfigFile(path + "\\config.yml")
// 	err := v.ReadInConfig()
// 	if err != nil {
// 		log.Fatal("读取配置文件失败：", err)
// 	}

// 	// 读取命令参数
// 	for _, args := range os.Args {
// 		// fmt.Println("参数"+strconv.Itoa(idx)+":", args)
// 		if !strings.Contains(args, "--") {
// 			continue
// 		}
// 		cmdInfos := strings.Split(args, "--")

// 		for _, cmdInfo := range cmdInfos {
// 			if !strings.Contains(cmdInfo, "=") {
// 				continue
// 			}
// 			cmds := strings.Split(cmdInfo, "=")
// 			// fmt.Printf("%v\n", cmds)
// 			if len(cmds) != 2 {
// 				fmt.Println("非法命令行参数" + cmdInfo)
// 				break
// 			}

// 			v.Set(cmds[0], cmds[1])

// 		}

// 	}
// 	v.WriteConfigAs("./config.yml") // 直接写入，有内容就覆盖，没有文件就新建
// }

// 读取配置文件
func ReadConf() Config {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cnf := Config{}
	c := &cnf
	v := viper.New()
	v.SetConfigName("config")       //这里就是上面我们配置的文件名称，不需要带后缀名
	v.AddConfigPath(wd + "/config") //文件所在的目录路径
	v.SetConfigType("yml")          //这里是文件格式类型

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

func ReadPathConf(path string) Config {

	_, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cnf := Config{}
	c := &cnf
	v := viper.New()
	v.SetConfigName("config") //这里就是上面我们配置的文件名称，不需要带后缀名
	v.AddConfigPath(path)     //文件所在的目录路径
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

func UpdateCheckInfoById(id int, resultCode string, resultMsg string) {

	// 读取checkInfo 数据
	wd, err := os.Getwd()
	// 打开json文件
	jsonFile, err := os.Open("" + wd + "/template/json/checkInfo.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	checksByte, _ := ioutil.ReadAll(jsonFile)

	// 转为数组checkInfo
	checks := DeserializeJson(string(checksByte))
	// 更新最近一次签到结果
	for index, info := range checks {
		if info.Id == id {
			checks[index].LastMsg = resultMsg
			checks[index].LastResult = resultCode
			checks[index].LastTIme = time.Now().Format("2006-01-02 15:04:05")
		}
	}
	fmt.Println(checks)
	// 保存
	WriteCheckInfoJson(checks)

}

func DeserializeJson(CheckInfoJson string) []CheckInfo {
	// fmt.Println("CheckInfoJson:", CheckInfoJson)
	jsonAsBytes := []byte(CheckInfoJson)
	checks := make([]CheckInfo, 0)
	err := json.Unmarshal(jsonAsBytes, &checks)
	// fmt.Printf("%#v", checks)
	if err != nil {
		panic(err)
	}
	return checks
}

func WriteCheckInfoJson(chekInfos []CheckInfo) {
	// 互斥锁
	var mutex sync.Mutex

	mutex.Lock()
	// 读取checkInfo 数据
	wd, error := os.Getwd()
	if error != nil {
		fmt.Println(error)
	}
	// file, e := os.OpenFile(""+wd+"/template/json/checkInfo.json", os.O_CREATE|os.O_WRONLY, 0666)
	// if e != nil {
	// 	fmt.Println("文件打开失败")
	// } else {
	// 	fmt.Println("文件打开成功")
	// }

	data, _ := json.Marshal(chekInfos)

	err := ioutil.WriteFile(""+wd+"/template/json/checkInfo.json", data, 0644)
	if err != nil {
		panic(err)
	}

	// // 创建编码器
	// encoder := json.NewEncoder(file)
	// err := encoder.Encode(chekInfos)
	// if err != nil {
	// 	fmt.Println("编码失败")
	// } else {
	// 	fmt.Println("编码成功")
	// }
	mutex.Unlock()
}

func ReadCheckInfoJsonToByte() []byte {

	wd, error := os.Getwd()
	if error != nil {
		fmt.Println(error)
	}
	// 打开json文件
	jsonFile, err := os.Open("" + wd + "/template/json/checkInfo.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	jsonByte, _ := ioutil.ReadAll(jsonFile)
	return jsonByte
}

func ReadCheckInfoJsonToCheck() []CheckInfo {

	wd, err := os.Getwd()

	// 打开json文件
	jsonFile, err := os.Open("" + wd + "/template/json/checkInfo.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	checksByte, _ := ioutil.ReadAll(jsonFile)
	return DeserializeJson(string(checksByte))

}
