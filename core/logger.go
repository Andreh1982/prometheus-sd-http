package logger

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
)

func New() (*zap.Logger, func()) {
	var cfg zap.Config
	rawJSON := []byte(`{
			"level": "` + "DEBUG" + `",
			"encoding": "json",
			"outputPaths": ["stdout"],
			"errorOutputPaths": ["stderr"],
			"encoderConfig": {
			  "messageKey": "message",
			  "levelKey": "level",
			  "levelEncoder": "lowercase"
			}
		  }`)

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		log.Fatal(err.Error())
	}
	logger, err := cfg.Build()

	if err != nil {
		log.Fatal(err.Error())
	}

	undo := zap.ReplaceGlobals(logger)

	return logger, func() {
		logger.Sync()
		undo()
	}
}
