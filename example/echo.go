package main

import (
	"context"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"

	"github.com/swallowarc/lime/lime"
)

func main() {
	var env lime.Env
	err := envconfig.Process("", &env)
	if err != nil {
		panic(err)
	}

	logOpt := lime.WithLogger(zap.NewExample())
	msgHandler := lime.WithEventHandler(&echoMessageEventHandler{})
	sv, err := lime.NewServer(env, logOpt, msgHandler)
	if err != nil {
		panic(err)
	}

	if err := sv.Start(); err != nil {
		panic(err)
	}
}

type echoMessageEventHandler struct{}

func (h *echoMessageEventHandler) EventType() linebot.EventType {
	return linebot.EventTypeMessage
}

func (h *echoMessageEventHandler) Handle(_ context.Context, event *linebot.Event, cli lime.LineBotClient) error {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		if _, err := cli.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
			log.Print(err)
		}
	}

	return nil
}
