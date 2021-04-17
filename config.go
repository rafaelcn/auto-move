package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

const (
	ActionNone   Actions = "nil"
	ActionMove   Actions = "move"
	ActionDelete Actions = "delete"
)

type (
	Actions string

	// Dirs represent a set of watched folders (what folders the program uses to
	// check and match against the defined rules) and to folders (to where the
	// program moves the files matched according to the defined predicates).
	Dirs struct {
		Watched []string `json:"watch"`
		To      string   `json:"to"`
	}

	// Rules represents a set of rules regarding a configuration file. A rule is
	// comprised of a type and predicates of that type, the known types for the
	// auto-move are: "filetype", "filename", "suffix", "prefix". Each type has
	// its own way to encode the predicates.
	// The filetype encodes the predicates in an array of string such as:
	// [".pdf", ".doc"].
	//
	// The filename acts kind of a mask on which the filename is matched with
	// the given mask. A mask can encode characters and digits, e.g., a file
	// with the name 02012-32120-21321.54654 can be encoded in the following
	// mask: DDDD-DDDDD-DDDDD.DDDDD
	//
	// The action tells what to do with the matched file, it can be one of the
	// actions specified by the type
	Rules struct {
		Type       string      `json:"type"`
		Predicates interface{} `json:"predicate"`
		Action     Actions     `json:"action"`
	}

	// Configuration represents a set of settings in a configuration file.
	Configuration struct {
		Folders Dirs  `json:"dirs"`
		Rules   Rules `json:"rule"`
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
		if path.Ext(file.Name()) == ".json" {
			var configuration Configuration

			content, err := ioutil.ReadFile(folder + file.Name())

			if err != nil {
				log.Fatalf("[!] Couldn't read file %s. Reason %v", file.Name(), err)
			}

			err = json.Unmarshal(content, &configuration)

			if err != nil {
				log.Printf("[!] Couldn't parse the configuration file. Reason: %v", err)
				return nil
			}

			configurations = append(configurations, configuration)
		}
	}

	return configurations
}
