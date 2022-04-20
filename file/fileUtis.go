package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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
