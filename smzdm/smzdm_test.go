package smzdm

import (
	"fmt"
	"testing"

	"ggball.com/smzdm/file"
)

func TestGetGoods(t *testing.T) {
	conf := file.ReadConf("E:\\project\\go\\smzdmForGo")
	v1, v2 := GetSatisfiedGoods(conf)
	fmt.Println(v1)
	fmt.Println(v2)
}
