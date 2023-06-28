package message

// NewsMessage 新闻消息
const NewsMessage = 1

type msg struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

// FormatNewsMessage formats news message
func FormatNewsMessage(v any) any {
	return msg{Code: NewsMessage, Data: v}
}
