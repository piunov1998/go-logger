package main

import (
	logger "github.com/piunov1998/go-logger"
)

func main() {
	log := logger.New("main", logger.BasicConfig())
	log.SetLevel(logger.DebugLevel)

	log.Debug("Starting process...")

	result, err := process()
	if err != nil {
		log.Errorf("error during process -> %s", err)
	}
	log.Infof("result is %d", result)
}

func process() (int, error) {
	// make some stuff
	return 0, nil
}
