package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

type (
	// Dirs represent a set of watched folders (what folders the program uses to
	// check and match against the defined rules) and to folders (to where the
	// program moves the files matched according to the defined predicates).
	Dirs struct {
		Watched []string `json:"watch"`
		To      string   `json:"to"`
	}

	// Rules represents a set of rules regarding a configuration file. 
	Rules struct {
		Type       string   `json:"type"`
		Extensions []string `json:"extensions"`
	}

	// Configuration represents a set of settings in a configuration file.
	Configuration struct {
		Folders    Dirs `json:"dirs"`
		Predicates Rules `json:"rules"`
	}
)

// Parse reads the confgiuration folder and parse every configuration file
// inside and returns an array of configuration files.
func Parse(folder string) []Configuration {
	dir, err := ioutil.ReadDir(folder)

	if err != nil {
		log.Fatalf("[!] Couldn't read configuration folder. Reason: %v", err)
	}

	var configurations []Configuration

	for _, file := range dir {
		if ".json" == path.Ext(file.Name()) {
			var configuration Configuration

			content, err := ioutil.ReadFile(folder + file.Name())

			if err != nil {
				log.Fatalf("[!] Couldn't read file %s. Reason %v", file.Name(), err)
			}

			err = json.Unmarshal(content, &configuration)

			if err != nil {
				log.Fatalf("[!] Couldn't parse the configuration file. Reason: %v", err)
			}
			
			configurations = append(configurations, configuration)
		}
	}

	return configurations
}
