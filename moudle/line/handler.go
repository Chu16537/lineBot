package line

import (
	"fmt"
	"linebot/dao"
	"strconv"
	"strings"
	"time"
)

// 文字分析
func (h *Handler) messageAnalyze(userID, token, msg string) string {

	text := strings.ToLower(msg)

	// 取得動作
	action := getAction(text)

	switch action {
	case ACT_QUERY:
		t := getQueryTime(text)

		queryData := h.db.Query(userID, t)

		if len(queryData) == 0 {
			return "not msg"
		}

		result := " "
		for _, data := range queryData {
			tm := time.Unix(data.Timestamp, 0).Format(time.RFC3339)
			result += fmt.Sprintf("time:%v , msg: %v \n", tm, data.Msg)
		}

		return result
	default:
		// 如果是其他 則把資料存入db
		data := dao.Msg{
			Name:      userID,
			UserID:    userID,
			Token:     token,
			Msg:       msg,
			Timestamp: time.Now().Unix(),
		}
		h.db.Save(data)
		return "save success"
	}
}

// 取得動作
func getAction(text string) string {

	if strings.Index(text, ACT_QUERY) == 0 {
		return ACT_QUERY
	}

	if strings.Index(text, ACT_ALL) == 0 {
		return ACT_ALL
	}

	return ""
}

func getQueryTime(text string) int64 {
	d := strings.Split(text, ACT_QUERY)[1]
	day, err := strconv.Atoi(d)

	if err != nil {
		return 0
	}

	if day < 0 {
		day = 0
	}
	t := time.Now().AddDate(0, 0, -day).Unix()
	return t
}
