package push

// 钉钉feedCard类型推送参数
type DingFeedCardParam struct {
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

// 钉钉text类型参数
type DingTextParam struct {
	MsgType string `json:"msgtype"`
	Texts   Text   `json:"text"`
}
type Text struct {
	Content string `json:"content"`
}
