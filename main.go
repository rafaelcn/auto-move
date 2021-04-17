package main

import (
	"flag"
	"strings"
	"time"
)

var (
	configuration Configuration

	rulesFolder      = flag.String("rules", "./rules/", "Change the directory used when reading .json configuration files")
	rulesRefreshTime = flag.Int("rules-refresh-time", 100, "Refresh time of the move algorithm in milliseconds")
	rulesReadTime    = flag.Int("rules-read-time", 200, "Refresh time that the program keeps reading for rules change")
)

func main() {
	flag.Parse()

	// Add a slash if the user didn't write one.
	if !strings.HasSuffix(*rulesFolder, "/") {
		*rulesFolder = *rulesFolder + "/"
	}

	rules := Parse(*rulesFolder)

	go func() {
		for {
			rules = Parse(*rulesFolder)
			time.Sleep(time.Duration(*rulesReadTime) * time.Millisecond)
		}
	}()

	for {
		for _, rule := range rules {
			Watch(rule)
		}
		time.Sleep(time.Duration(*rulesRefreshTime) * time.Millisecond)
	}
}
