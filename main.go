package main

import (
	"linebot/moudle/line"
	"linebot/moudle/mongoDB"
	"linebot/moudle/web"
)

func main() {

	web.Init()
	mongoDB.Init()

	line.Init(web.Get(), mongoDB.Get())
	web.Run()

}
