package config

import (
	"flag"
)

var (
	Environment       = flag.String("env", "development", "Environment")
	BrandID           = flag.Int("brand-id", 0, "Brand ID")
	Brand             = flag.String("brand", "", "Brand")
	ConsumerTags      = flag.String("consumer-tags", "", "Consumer Tags")
	LogLevel          = flag.String("log-level", "info", "Log level")
	APIURL            = flag.String("api-url", "0.0.0.0", "API URL")
	APIPort           = flag.String("api-port", "5001", "API Port")
	APIKey            = flag.String("api-key", "", "API Key")
	AMQPConnectionStr = flag.String("amqp-connection-str", "", "AMQP Connection String")
	AMQPQueue         = flag.String("amqp-queue", "", "AMQP Queue")
	AMQPPrefetchSize  = flag.Int("amqp-prefetch-size", 0, "AMQP Prefetch Size")
	DatabaseConnStr   = flag.String("db-conn-str", "", "Database Connection String")
	ServiceName       = flag.String("service-name", "", "Service Name")
)
