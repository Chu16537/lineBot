package main

import (
	"fmt"
	"linebot/conf"
	"linebot/moudle/line"
	"linebot/moudle/mongoDB"
	"linebot/moudle/web"
)

func main() {

	err := conf.Init()

	if err != nil {
		fmt.Printf("env Fail: %v", err)
		return
	}

	web.Init()
	mongoDB.Init()

	line.Init(web.Get(), mongoDB.Get())
	web.Run()

}
