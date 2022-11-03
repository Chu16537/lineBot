DBPull:
	docker pull mongo:4.4

DBRun:
	docker run --name mongo -v $(pwd)/data:/data/db -d -p 27017:27017 --rm mongo:4.4

DBStop:
	docker stop mongo

DBCheck:
	docker exec mongo mongo --eval "print(version())"

NgrokRun:
	docker run -it --name ngrok -e NGROK_AUTHTOKEN=2GyspSgx1Ddbtqmq4yO5aOwftCn_5jYc3AjoNqK3GmuGMWNRj ngrok/ngrok http 8888	
	
get:
	go get github.com/gin-gonic/gin
	go get  github.com/line/line-bot-sdk-go/v7/linebot
	go get gopkg.in/mgo.v2
	go get gopkg.in/mgo.v2/bson

run :
	go run main.go

