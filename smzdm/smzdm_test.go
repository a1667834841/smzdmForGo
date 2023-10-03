package smzdm

import (
	"testing"

	"ggball.com/smzdm/file"
)

func TestGetGoods(t *testing.T) {
	conf := file.ReadConf("E:\\project\\go\\smzdmForGo")
	GetSatisfiedGoods(conf)
}
