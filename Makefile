build:
	go build -i -o ./dist/functions ./functions/sample/...
	go build -i -o ./dist/functions ./functions/sample2/...
	go build -i -o ./dist/public ./src/...