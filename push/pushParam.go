package push

// 钉钉推送参数
type DingParam struct {
	MsgType  string   `json:"msgtype"`
	FeedCard FeedCard `json:"feedCard"`
}

type FeedCard struct {
	Links []Link `json:"links"`
}

type Link struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}
