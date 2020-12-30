package aws

import (
	"time"

	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/config/aws"
)

type (
	// SQSConfig holds the info required to work with Amazon SQS
	SQSConfig struct {
		aws.Config

		QueueName string `envconfig:"AWS_SQS_NAME"`
		// MaxMessages will override the DefaultSQSMaxMessages.
		MaxMessages *int64 `envconfig:"AWS_SQS_MAX_MESSAGES"`
		// TimeoutSeconds will override the DefaultSQSTimeoutSeconds.
		TimeoutSeconds *int64 `envconfig:"AWS_SQS_TIMEOUT_SECONDS"`
		// SleepInterval will override the DefaultSQSSleepInterval.
		SleepInterval *time.Duration `envconfig:"AWS_SQS_SLEEP_INTERVAL"`
		// DeleteBufferSize will override the DefaultSQSDeleteBufferSize.
		DeleteBufferSize *int `envconfig:"AWS_SQS_DELETE_BUFFER_SIZE"`
		// ConsumeBase64 is a flag to signal the subscriber to base64 decode the payload
		// before returning it. If it is not set in the config, the flag will default
		// to 'true'.
		ConsumeBase64 *bool `envconfig:"AWS_SQS_CONSUME_BASE64"`
	}

	// SNSConfig holds the info required to work with Amazon SNS.
	SNSConfig struct {
		aws.Config

		Topic string `envconfig:"AWS_SNS_TOPIC"`
	}
)

// LoadSQSConfigFromEnv will attempt to load the SQSConfig struct
// from environment variables. If not populated, nil
// is returned.
func LoadSQSConfigFromEnv() SQSConfig {
	var cfg SQSConfig
	config.LoadEnvConfig(&cfg)
	return cfg
}

// LoadSNSConfigFromEnv will attempt to load the SNSConfig struct
// from environment variables. If not populated, nil
// is returned.
func LoadSNSConfigFromEnv() SNSConfig {
	var cfg SNSConfig
	config.LoadEnvConfig(&cfg)
	return cfg
}
