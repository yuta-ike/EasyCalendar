build:
	go build -i -o ./dist/functions/calendar ./src/functions/calendar
	go build -i -o ./dist/functions/sample ./src/functions/sample
	go build -o ./dist/public/src ./src/public/...