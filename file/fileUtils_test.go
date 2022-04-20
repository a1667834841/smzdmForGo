package file

import (
	"fmt"
	"testing"
)

func TestReadPushed(t *testing.T) {
	readPusedInfo := ReadPusedInfo("../pushed.json")
	value, b := readPusedInfo["222"]
	if b {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}
}

func TestWritePushed(t *testing.T) {
	readPusedInfo := ReadPusedInfo("../pushed.json")

	pushedMap := make(map[string]interface{})

	pushedMap["222"] = "222"

	WritePushedInfo(pushedMap, readPusedInfo, "../pushed.json")
}
