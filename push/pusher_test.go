package push

import (
	"testing"
)

// 图文推送
func TestDingPusher(t *testing.T) {
	dingPusher := DingPusher{
		Token: "9e4044952fe5c599afed3815ccaa387c650fd07bda96512648acfceb1b202ada",
	}

	// 需要提前申明数组的容量
	links := make([]Link, 2)

	links[0] = Link{
		Title:      "【什么值得买】时代的火车向前开1",
		MessageURL: "https://www.dingtalk.com/",
		PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
	}

	links[1] = Link{
		Title:      "【什么值得买】时代的火车向前开2",
		MessageURL: "https://www.dingtalk.com/",
		PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
	}

	feedCard := FeedCard{
		Links: links,
	}

	params := DingFeedCardParam{
		MsgType:  "feedCard",
		FeedCard: feedCard,
	}

	dingPusher.PushDingDing(params)

}

// @人员推送
func TestDingPusherWithMobiles(t *testing.T) {
	dingPusher := DingPusher{
		Token: "9e4044952fe5c599afed3815ccaa387c650fd07bda96512648acfceb1b202ada",
	}

	text := "【好物到了】 测试推送 https://www.dingtalk.com/"
	content := Text{Content: text}
	params := DingTextParam{
		MsgType: "text",
		Texts:   content,
		At:      At{AtMobiles: []string{"13217913287"}, IsAtAll: false},
	}

	dingPusher.PushDingDing(params)

}

func TestDingPusherMdWithMobiles(t *testing.T) {
	dingPusher := DingPusher{
		Token: "9e4044952fe5c599afed3815ccaa387c650fd07bda96512648acfceb1b202ada",
	}

	text := " [这是一个测试](https://www.runoob.com)"
	md := Markdown{Title: "【好物到了】", Text: text}
	params := DingMdParam{
		MsgType:  "markdown",
		Markdown: md,
		At:       At{AtMobiles: []string{"13217913287"}, IsAtAll: false},
	}

	dingPusher.PushDingDing(params)

}
