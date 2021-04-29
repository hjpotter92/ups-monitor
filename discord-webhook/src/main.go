// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
)

var discordWebhookURL = os.Getenv("DISCORD_WEBHOOK_URL")

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty,default:false"`
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type Link struct {
	URL string `json:"url"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type Embed struct {
	Author      *Author  `json:"author"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Color       uint64   `json:"color,omitempty"`
	Fields      *[]Field `json:"fields,omitempty"`
	Thumbnail   *Link    `json:"thumbnail,omitempty"`
	Image       *Link    `json:"image,omitempty"`
	Footer      *Footer  `json:"footer,omitempty"`
}

type Webhookdata struct {
	Username  string   `json:"username,omitempty"`
	AvatarURL string   `json:"avatar_url,omitempty"`
	Content   string   `json:"content"`
	Embeds    *[]Embed `json:"embeds,omitempty"`
}

type Triggerdata struct {
	Triggervalue float64
	Channelid    uint64
	Datetime     string
}

func ReadBody(body string) (Triggerdata, error) {
	var trigger Triggerdata
	log.Printf("Reading the body: '%s'", body)
	json.Unmarshal([]byte(body), &trigger)
	return trigger, nil
}

func GetWebhookData(trigger Triggerdata) *Webhookdata {
	var color uint64
	if trigger.Triggervalue <= 50 {
		// Red #ff0000
		color = 16711680
	} else if trigger.Triggervalue <= 200 {
		// Yellow/orange #cccc00
		color = 13421568
	} else {
		// Green #00cc33
		color = 52275
	}
	webhookData := &Webhookdata{
		Username:  "UPS Monitor",
		AvatarURL: "https://i.imgur.com/T4JpHlp.gif",
		Content:   "The bot detected a voltage fluctuation.",
		Embeds: &[]Embed{
			{
				Author: &Author{
					Name: "UPS-bot",
					URL:  "https://thingspeak.com/apps/matlab_visualizations/403462",
				},
				Title:       "Voltage fluctuation detected",
				URL:         "https://thingspeak.com/channels/1229924/",
				Description: fmt.Sprintf("A voltage fluctuation was detected at %s", trigger.Datetime),
				Color:       color,
				Fields: &[]Field{
					{
						Name:   "Triggered value",
						Value:  fmt.Sprintf("%f", trigger.Triggervalue),
						Inline: true,
					},
					{
						Name:   "Thingspeak channel ID",
						Value:  fmt.Sprintf("%d", trigger.Channelid),
						Inline: true,
					},
					{
						Name:  "Triggered at",
						Value: trigger.Datetime,
					},
					{
						Name:  "GitHub",
						Value: "https://github.com/hjpotter92/ups-monitor",
					},
				},
				Footer: &Footer{
					Text:    "[GitHub](https://github.com/hjpotter92/ups-monitor)",
					IconURL: ":git:",
				},
			},
		},
	}
	return webhookData
}

func hello(ctx context.Context, snsEvent events.SNSEvent) (string, error) {
	body := snsEvent.Records[0].SNS.Message
	triggerData, _ := ReadBody(body)
	log.Printf("Trigger data: %f, Date: %s", triggerData.Triggervalue, triggerData.Datetime)
	webhookData := GetWebhookData(triggerData)
	str, _ := json.Marshal(webhookData)
	log.Println(string(str))
	client := gorequest.New()
	resp, body, _ := client.Post(discordWebhookURL).
		Send(webhookData).
		End()
	log.Println(resp)
	log.Println(body)
	return `{"message": "Hello λ!"}`, nil
}

func main() {
	lambda.Start(hello)
}
