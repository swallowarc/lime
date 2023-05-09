package lime

import (
	"context"
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	EventHandler interface {
		EventType() linebot.EventType
		Handle(ctx context.Context, event *linebot.Event) error
	}

	webhookHandler struct {
		logger        *zap.Logger
		channelSecret string

		eventTimeout   time.Duration
		eventHandlers  map[linebot.EventType]EventHandler
		returnError    bool
		enableEventLog bool
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

	if err != nil {
		wh.logger.Error("failed to create linebot client", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errored := false
	for _, event := range events {
		if isErr := wh.callEventHandler(event); isErr {
			errored = true
		}
	}

	if errored && wh.returnError {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (wh *webhookHandler) callEventHandler(event *linebot.Event) bool {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), wh.eventTimeout)
	defer cancel()

	handler, ok := wh.eventHandlers[event.Type]
	if !ok {
		wh.eventLog(zap.DebugLevel, event, "no handler found", start, nil)
		return false
	}

	if err := handler.Handle(ctx, event); err != nil {
		wh.eventLog(zap.ErrorLevel, event, "failed to handle event", start, err)
		return true
	}

	wh.eventLog(zap.InfoLevel, event, "success to handle event", start, nil)
	return false
}

func (wh *webhookHandler) eventLog(lv zapcore.Level, e *linebot.Event, msg string, start time.Time, err error) {
	if !wh.enableEventLog {
		return
	}

	l := wh.logger.With(
		zap.String("webhook_event_id", e.WebhookEventID),
		zap.String("event_source_type", string(e.Source.Type)),
		zap.String("event_type", string(e.Type)),
		zap.Duration("elapsed_time", time.Since(start)),
		zap.String("message", msg),
	)

	if err != nil {
		l = l.With(zap.Error(err))
	}

	l.Log(lv, "event")
}
