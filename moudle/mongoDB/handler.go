package mongoDB

import (
	"linebot/dao"

	"gopkg.in/mgo.v2/bson"
)

// 查詢
// @param userID 會員id
// @param time  時間
func (h *Handler) Query(userID string, time int64) []dao.Msg {
	c := h.db.C(COL_MESSAGE)

	var result []dao.Msg

	query := bson.M{
		"name":      userID,
		"timestamp": bson.M{"$gte": time},
	}

	err := c.Find(query).All(&result)

	if err != nil {
		panic(err)
	}

	return result
}

// 儲存
func (h *Handler) Save(data dao.Msg) {
	c := h.db.C(COL_MESSAGE)

	err := c.Insert(data)
	if err != nil {
		panic(err)
	}
}
