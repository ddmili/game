package sdk

import (
	"fmt"
	"game/internal/logger"
)

// JinseSdk is a Jinse
type JinseSdk struct {
	BaseSdk
	limit int
}

// JinseResult jinse news api result
type JinseResult struct {
	TopId int64 `json:"top_id"`
	List  []struct {
		Date  string `json:"date"`
		Lives []struct {
			Id            int64  `json:"id"`
			Content       string `json:"content"`
			ContextPrefix string `json:"content_prefix"`
			Link          string `json:"link"`
			Time          int    `json:"created_at"`
		} `json:"lives"`
	} `json:"list"`
}

// GetNews get news list
func (j *JinseSdk) GetNews() (*JinseResult, error) {
	j.limit = 20
	url := "https://api.jinse.cn/noah/v2/lives?limit=%s&reading=false&source=web&flag=down&id=0&category=0"
	url = fmt.Sprintf(url, j.limit)
	rs := &JinseResult{}
	query := map[string]string{}
	err := j.RequestGet(url, query, rs)
	if err != nil {
		logger.Errorf("get jinse news error:%s", err)
	}
	return rs, err
}
