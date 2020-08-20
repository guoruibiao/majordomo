package elasticsearch


type ClipboardItem struct {
	Content string
	Ext     string // 比如 Content 是一个链接，则 ext 则为抓取到的网页内容
	Timestamp int
}


type LogItem struct {
	Index string `json:"_index"`
	Type  string `json:"_type"`
	ID    string `json:"_id"`
	Score float64 `json:"_score"`
	Source ClipboardItem `json:"_source"`
}

type LogHitsContainer struct {
	Total struct{
		Value int `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	Hits []LogItem `json:"hits"`
}

type LogList struct {
	Took int `json:"took"`
	Timeout bool `json:"timeout"`
	HitsContainer LogHitsContainer `json:"hits"`
}
