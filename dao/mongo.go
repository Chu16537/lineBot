package dao

// msg存入db格式
type Msg struct {
	UserID    string `json:"userId,omitempty"`
	Token     string `json:"token,omitempty"`
	Name      string `json:"name,omitempty"`
	Msg       string `json:"msg,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}
