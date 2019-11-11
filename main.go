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
)

func main() {
	flag.Parse()

	// Add a slash if the user didn't write one.
	if !strings.HasSuffix(*ConfigurationFolder, "/") {
		*ConfigurationFolder = *ConfigurationFolder + "/"
	}

	rules := Parse(*ConfigurationFolder)

	go func() {
		for {
			rules = Parse(*ConfigurationFolder)
			time.Sleep(20 * time.Millisecond)
		}
	}()

	for {
		for _, rule := range rules {
			Watch(rule)
		}
		time.Sleep(time.Duration(*MoveUpdateTime) * time.Millisecond)
	}
}
