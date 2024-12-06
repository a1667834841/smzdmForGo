package check_in

import (
	"testing"
	"ggball.com/smzdm/db"
	"ggball.com/smzdm/file"
)

func TestDoCheckIn(t *testing.T) {
	// 创建测试数据库连接
	c, err := NewCheckIn("../data/users.db")
	if err != nil {
		t.Fatalf("创建测试数据库失败: %v", err)
	}
	defer c.db.Close()

	// 设置配置
	conf := file.Config{
		DingdingToken: "test_token",
	}
	checks := []file.CheckInfo{}
	c.SetConfig(conf, checks)

	// 构造测试用户
	testUser := db.User{
		Name: "测试用户",
		Token: "test_cookie",
		Platform: "smzdm",
	}

	// 执行签到
	msg, err := c.doCheckIn(testUser)
	if err != nil {
		t.Errorf("签到失败: %v,msg:%s", err,msg)
	}
}
