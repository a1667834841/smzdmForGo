package push

// 钉钉feedCard类型推送参数
type DingFeedCardParam struct {
	MsgType  string   `json:"msgtype"`
	FeedCard FeedCard `json:"feedCard"`
	At       At       `json:"at"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
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
	At      At     `json:"at"`
}
type Text struct {
	Content string `json:"content"`
}

// 钉钉md类型
type DingMdParam struct {
	MsgType  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
