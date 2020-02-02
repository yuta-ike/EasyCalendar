build:
	go get github.com/aws/aws-lambda-go/events
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/line/line-bot-sdk-go/linebot
	go build -o ./dist/functions/calendar ./src/functions/calendar
	go build -o ./dist/functions/sample ./src/functions/sample
	go build -o ./dist/public/src ./src/public/...