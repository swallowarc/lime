package lime

import (
	"net/http"

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

func NewServer(env Env, ops ...APIServerOption) APIServer {
	opts := newOptions(ops...)

	mux := http.NewServeMux()
	mux.Handle(env.HandlePath, &webhookHandler{
		logger:         opts.logger,
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

	return s
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
