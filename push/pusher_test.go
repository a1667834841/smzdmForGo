package push

import (
	"testing"
)

func TestDingPusher(t *testing.T) {
	dingPusher := DingPusher{
		Token: "106aef404757b5a5c7df598663a9590f7ad67a4edd82ed255faee5dbc986776a",
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

	dingPusher.PushWithFeedCard(params)

}
