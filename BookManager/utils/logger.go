package utils

import (
	"os"

	mbinterfaces "gitlab.com/dotpe/mindbenders/interfaces"
	logger "gitlab.com/dotpe/mindbenders/logging"
)

const (
	timeFormat = "02/Jan/2006:15:04:05 -0700"
	key        = "AKIASB5FWRSTYW3EU7FV"
	secret     = "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"
)

var DLogger mbinterfaces.IDotpeLogger

// InitLogger ..
func InitLogger() error {
	url := "https://search-stage-dotpe-b2ereag6qexcrglpdaiufjzwjy.ap-south-1.es.amazonaws.com"
	if os.Getenv("ENV") == "prod" {
		url = "https://search-prod-dotpe-6ccj4pln63pns5d3sobiyagmxi.ap-south-1.es.amazonaws.com"
	}

	hostname, _ := os.Hostname()
	lops := &logger.LoggerOptions{
		APP:    os.Getenv("APP"),
		WD:     os.Getenv("CWD"),
		LOGENV: os.Getenv("ENV"),
		KibanaConfig: logger.KibanaConfig{
			Client:    url,
			AccessKey: key,
			SecretKey: secret,
			Hostname:  hostname,
		},
	}
	var err error = nil
	DLogger, err = logger.InitLogger(lops)
	return err
}
