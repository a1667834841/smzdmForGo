package check_in

import (
	"testing"

	"ggball.com/smzdm/file"
)

func TestCheckIn(t *testing.T) {
	conf := file.ReadPathConf("d:\\project\\vscode\\smzdmForGo")
	conf.Cron = "0 0 9 ? * *"
	Run(conf)
}
