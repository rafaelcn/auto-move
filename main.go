package main

import (
	"flag"
	"strings"
	"time"
)

var (
	configuration Configuration

	ConfigurationFolder = flag.String("rules-folder", "./rules/", "Change the directory used when reading .json configuration files")
	MoveUpdateTime      = flag.Int("move-refresh", 100, "Refresh time of the move algorithm in milliseconds")
	ConfigUpdateTime    = flag.Int("config-refresh", 500, "Refresh time of the configuration file parser in milliseconds")
)

func main() {
	flag.Parse()

	// Add a slash if the user didn't write one.
	if !strings.HasSuffix(*ConfigurationFolder, "/") {
		*ConfigurationFolder = *ConfigurationFolder + "/"
	}

	var rules []Configuration
	rules = Parse(*ConfigurationFolder)

	go func() {
		rules = Parse("./rules/")
		time.Sleep(time.Duration(*MoveUpdateTime) * time.Second)
	}()

	for _, rule := range rules {
		// For each config file init a goroutine to run the set up
		go Watch(rule)
	}

	// wait for a signal
	done := make(chan bool, 1)

	<-done

}
