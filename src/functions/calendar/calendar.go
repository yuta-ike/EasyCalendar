package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	accessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	bot, err := linebot.New(channelSecret, accessToken)

	if err != nil {
		log.Fatal(err)
	}

	eventList, err := UnmarshalLineRequest([]byte(request.Body))

	for _, event := range eventList.Events {
		replyToken := event.ReplyToken
		message := event.Message
		switch {
		case event.Message.Type == "text":
			if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
				log.Print(err)
			}
		}
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       request.Body,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	//
	lambda.Start(handler)
}
