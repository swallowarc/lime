package lime

import (
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"
)

type (
	APIServer interface {
		Start() error
	}

	apiServer struct {
		logger *zap.Logger

		server *http.Server
	}
)

var (
	defaultLogger = zap.NewExample()
)

func NewServer(env Env, ops ...APIServerOption) (APIServer, error) {
	opts := newOptions(ops...)

	lineBotCli, err := linebot.New(env.ChannelSecret, env.ChannelToken,
		linebot.WithHTTPClient(opts.client),
		linebot.WithEndpointBase(env.LineAPIEndpointBase),
		linebot.WithEndpointBaseData(env.LineAPIEndpointBaseData),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create linebot client: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle(env.HandlePath, &webhookHandler{
		logger:         opts.logger,
		client:         lineBotCli,
		channelSecret:  env.ChannelSecret,
		eventTimeout:   env.EventTimeout(),
		eventHandlers:  opts.eventHandlers,
		returnError:    env.EnableReturnErrorCode,
		enableEventLog: env.EnableEventLog,
	})
	mux.Handle("/healthz", opts.healthzHandler)
	mux.Handle("/readiness", opts.readinessHandler)
	mux.Handle("/", defaultHealthzHandler)

	s := &apiServer{
		logger: defaultLogger,
		server: &http.Server{
			Addr:              ":" + env.Port,
			ReadHeaderTimeout: env.ReadTimeout(),
			WriteTimeout:      env.WriteTimeout(),
			IdleTimeout:       env.IdleTimeout(),
			Handler:           mux,
		},
	}

	return s, nil
}

func (s *apiServer) Start() error {
	s.logger.Info("lime api server started")
	s.logger.Info("listening", zap.String("addr", s.server.Addr))

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error("failed to lime api server", zap.Error(err))
		return err
	}

	s.logger.Info("lime api server ended")
	return nil
}
