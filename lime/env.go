package lime

import "time"

type (
	Env struct {
		Port       string `envconfig:"port" default:"8080"`
		HandlePath string `envconfig:"handle_path" default:"/callback"`

		ReadTimeoutSec  int `envconfig:"read_timeout_sec" default:"5"`
		WriteTimeoutSec int `envconfig:"write_timeout_sec" default:"10"`
		IdleTimeoutSec  int `envconfig:"idle_timeout_sec" default:"120"`
		EventTimeoutSec int `envconfig:"event_timeout_sec" default:"15"`

		ChannelSecret string `envconfig:"channel_secret" default:"mock-secret"`
		ChannelToken  string `envconfig:"channel_token" default:"mock-token"`

		LineAPIEndpointBase     string `envconfig:"line_api_endpoint_base" default:"https://api.line.me"`
		LineAPIEndpointBaseData string `envconfig:"line_api_endpoint_base_data" default:"https://api-data.line.me"`

		// EnableReturnErrorCode is a flag to enable re-delivery webhook.
		//   If True, returns a code other than 200 when an error occurs in EventHandler.
		//   If Webhook resend is enabled in LINE Messaging API, the same event will be resent.
		//   Please implement the following in event handlers.
		//
		//  - Do not re-execute events with WebhookEventIDs that have already been processed.
		//  - Or, the processing in the handler should idempotent.
		EnableReturnErrorCode bool `envconfig:"enable_re_delivery_webhook" default:"false"`
	}
)

func (c Env) ReadTimeout() time.Duration {
	return time.Duration(c.ReadTimeoutSec) * time.Second
}

func (c Env) WriteTimeout() time.Duration {
	return time.Duration(c.WriteTimeoutSec) * time.Second
}

func (c Env) IdleTimeout() time.Duration {
	return time.Duration(c.IdleTimeoutSec) * time.Second
}

func (c Env) EventTimeout() time.Duration {
	return time.Duration(c.EventTimeoutSec) * time.Second
}
