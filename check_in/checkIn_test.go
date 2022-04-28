package check_in

import (
	"testing"

	"ggball.com/smzdm/file"
)

func TestCheckIn(t *testing.T) {
	conf := file.ReadPathConf("d:\\project\\vscode\\smzdmForGo")
	Run(conf)
}
