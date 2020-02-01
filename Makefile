build:
	go get github.com/aws/aws-lambda-go/events
	go get github.com/aws/aws-lambda-go/lambda
	go build -i -o ./dist/functions/calendar ./src/functions/calendar
	go build -i -o ./dist/functions/sample ./src/functions/sample
	go build -o ./dist/public/src ./src/public/...