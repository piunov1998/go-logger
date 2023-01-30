package main

import (
	logger "github.com/piunov1998/go-logger"
)

func init() {
	logger.BasicConfig.Colors = true
	logger.BasicConfig.LogLevel = "debug"
}

type AwesomeService struct {
	logger logger.Logger
}

func New() AwesomeService {
	service := AwesomeService{}
	service.logger = logger.New(service, nil)
	return service
}

func (a AwesomeService) Do() (int, error) {
	a.logger.Debug("doing...")
	// doing...
	return 0, nil
}

func main() {
	log := logger.New("main", &logger.Config{Colors: true, LogLevel: "info"})
	service := New()
	if result, err := service.Do(); err != nil {
		log.Fatalf("error during workflow -> %s", err)
	} else {
		log.Infof("result is %d", result)
	}
}
