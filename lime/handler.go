package lime

import (
	"context"
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"
)

type (
	EventHandler interface {
		Handle(ctx context.Context, event *linebot.Event, cli LineBotClient) error
	}

	webhookHandler struct {
		logger *zap.Logger
		client *http.Client

		channelSecret           string
		channelToken            string
		lineAPIEndpointBase     string
		lineAPIEndpointBaseData string

		eventTimeout  time.Duration
		eventHandlers map[linebot.EventType]EventHandler
		returnError   bool
	}
)

var (
	defaultHealthzHandler http.HandlerFunc = func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	defaultReadinessHandler http.HandlerFunc = func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
)

func (wh *webhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := linebot.ParseRequest(wh.channelSecret, r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			wh.logger.Info("invalid signature")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		wh.logger.Error("failed to parse request", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cli, err := linebot.New(wh.channelSecret, wh.channelToken,
		linebot.WithHTTPClient(wh.client),
		linebot.WithEndpointBase(wh.lineAPIEndpointBase),
		linebot.WithEndpointBaseData(wh.lineAPIEndpointBaseData),
	)
	if err != nil {
		wh.logger.Error("failed to create linebot client", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if errored := wh.callEventHandlers(events, cli); errored {
		if wh.returnError {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (wh *webhookHandler) callEventHandlers(events []*linebot.Event, cli LineBotClient) bool {
	ctx, cancel := context.WithTimeout(context.Background(), wh.eventTimeout)
	defer cancel()

	for _, event := range events {
		handler, ok := wh.eventHandlers[event.Type]
		if !ok {
			wh.logger.Debug("no handler found", zap.String("event_type", string(event.Type)))
			continue
		}

		if err := handler.Handle(ctx, event, cli); err != nil {
			wh.logger.Error("failed to handle event", zap.Error(err), zap.String("webhook_event_id", event.WebhookEventID))
			return true
		}
	}

	return false
}
