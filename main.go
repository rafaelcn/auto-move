package main

import "time"

var (
	configuration Configuration
)

func main() {

	var rules []Configuration

	rules = Parse("./rules/")

	// Look for new rules on a period of 500ms
	go func() {
		rules = Parse("./rules/")
		time.Sleep(500 * time.Second)
	}()

	for _, rule := range rules {
		// For each config file init a goroutine to run the set up
		go Watch(rule)
	}

	// stop point
	done := make(chan bool, 1)

	<-done

}
