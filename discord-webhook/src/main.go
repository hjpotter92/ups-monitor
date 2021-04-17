// main.go
package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
)

var discordWebhookURL = os.Getenv("DISCORD_WEBHOOK_URL")

var webhookData = `{
  "username": "UPS Monitor",
  "avatar_url": "https://i.imgur.com/4M34hi2.png",
  "content": "The bot detected a lower voltage level than recommended.",
  "embeds": [
    {
      "author": {
        "name": "UPS-bot",
        "url": "https://github.com/hjpotter92/ups-monitor",
        "icon_url": "https://thingspeak.com/apps/matlab_visualizations/403462"
      },
      "title": "Voltage fluctuation detected",
      "url": "https://thingspeak.com/channels/1229924/",
      "description": "The bot received a power fluctuation error.",
      "color": 15258703,
      "fields": [
        {
          "name": "Text",
          "value": "More text",
          "inline": true
        },
        {
          "name": "Even more text",
          "value": "Yup",
          "inline": true
        },
        {
          "name": "Use \"inline\": true parameter, if you want to display fields in the same line.",
          "value": "okay..."
        },
        {
          "name": "Thanks!",
          "value": "You're welcome :wink:"
        }
      ],
      "thumbnail": {
        "url": "https://upload.wikimedia.org/wikipedia/commons/3/38/4-Nature-Wallpapers-2014-1_ukaavUI.jpg"
      },
      "image": {
        "url": "https://upload.wikimedia.org/wikipedia/commons/5/5a/A_picture_from_China_every_day_108.jpg"
      },
      "footer": {
        "text": "Woah! So cool! :smirk:",
        "icon_url": "https://i.imgur.com/fKL31aD.jpg"
      }
    }
  ]
}`

func hello(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
	log.Println(ctx)
	log.Printf("Processing request '%s'.", request.RequestContext.RequestID)
	log.Printf("Request body: '%s'", request.Body)
	client := gorequest.New()
	resp, body, _ := client.Post(discordWebhookURL).
		Send(webhookData).
		End()
	log.Printf("%s, %s", resp, body)
	return `{"message": "Hello λ!"}`, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
