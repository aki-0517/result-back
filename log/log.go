package log

import "go.uber.org/zap"

var (
	logger *zap.logger
)

func NewLogger() (*zap.Logger, error) {
	if logger != nil {
		return logger, nil
	}

	var err error

	if config.IsDevelopment(){
		logger, error = zap.NewDevelopment()
	} else if config.IsTest() {
		logger = zap.NewNop()
	} else {
		logger, err = zap.NewProduction()
	}

	return logger, err
}