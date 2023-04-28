package lime

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"
)

type (
	OptionType string

	APIServerOption func(s *apiOptions)

	apiOptions struct {
		logger           *zap.Logger
		client           *http.Client
		eventHandlers    map[linebot.EventType]EventHandler
		healthzHandler   http.HandlerFunc
		readinessHandler http.HandlerFunc
	}
)

func newOptions(ops ...APIServerOption) *apiOptions {
	opts := &apiOptions{
		logger:           defaultLogger,
		client:           http.DefaultClient,
		eventHandlers:    make(map[linebot.EventType]EventHandler),
		healthzHandler:   defaultHealthzHandler,
		readinessHandler: defaultReadinessHandler,
	}

	for _, op := range ops {
		op(opts)
	}

	return opts
}

func WithLogger(logger *zap.Logger) APIServerOption {
	return func(s *apiOptions) {
		s.logger = logger
	}
}

func WithHTTPClient(client *http.Client) APIServerOption {
	return func(s *apiOptions) {
		s.client = client
	}
}

func WithEventHandler(eventType linebot.EventType, handler EventHandler) APIServerOption {
	return func(s *apiOptions) {
		s.eventHandlers[eventType] = handler
	}
}

func WithHealthz(fn http.HandlerFunc) APIServerOption {
	return func(s *apiOptions) {
		s.healthzHandler = fn
	}
}

func WithReadiness(fn http.HandlerFunc) APIServerOption {
	return func(s *apiOptions) {
		s.readinessHandler = fn
	}
}
