package smzdm

import (
	"testing"

	"ggball.com/smzdm/file"
)

func TestGetGoods(t *testing.T) {
	conf := file.ReadConf("../")
	v1, v2 := GetSatisfiedGoods(conf)
	if len(v1) == 0 {
		t.Error("获取符合条件的商品数量为0")
	}

	if len(v2) == 0 {
		t.Error("获取符合自己条件的商品数量为0") 
	}
}
