package dao

type Env struct {
	Port   string `yaml:"Port,omitempty"`
	DBAddr string `yaml:"DBAddr,omitempty"`
	Line   *Line  `yaml:"Line,omitempty"`
}

type Line struct {
	ChannelSecret      string `yaml:"ChannelSecret,omitempty"`
	ChannelAccessToken string `yaml:"ChannelAccessToken,omitempty"`
}
