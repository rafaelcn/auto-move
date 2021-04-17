package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"
)

// Watch watches for modifications on the listed folders set on the
// configuration file. It also watches modifications updates to that
// configuration file.
func Watch(configuration Configuration) {
	folders := configuration.Folders.Watched
	to := configuration.Folders.To

	predicate := configuration.Rules.Type
	action := configuration.Rules.Action

	for _, dir := range folders {
		files, err := ioutil.ReadDir(dir)

		if err != nil {
			// Implement a try system and remove a folder if the system
			log.Printf("[!] Failed to read dir %s. Reason %v", dir, err)
		}

		// O(m*n)
		for _, file := range files {
			_f := file.Name()
			filename := _f[0:(len(_f) - len(path.Ext(_f)))]

			if configuration.Rules.Predicates != nil {
				switch predicate {
				case "filetype":
					extensions := configuration.Rules.Predicates.([]interface{})

					for _, extension := range extensions {
						e := extension.(string)

						log.Printf("[+] Verifying file %s", file.Name())
						if e == path.Ext(file.Name()) {
							move(file, dir+"/"+file.Name(), to+"/"+file.Name())
						}
					}
				case "filename":
					mask := configuration.Rules.Predicates.(string)

					// D -> digit
					// L -> letter

					if len(mask) == len(filename) {
						// flag to verify the move after the for
						shouldMove := true
						var err error

						for i := 0; i < len(mask) && err == nil; i++ {
							l := string(filename[i])

							switch string(mask[i]) {
							case "D":
								_, err = strconv.Atoi(l)

								if err != nil {
									shouldMove = false
								}
							case "L":
								r := []rune(l)

								if !unicode.IsLetter(r[0]) {
									shouldMove = false
									msg := fmt.Sprintf("character at position %d is not a letter: %s", i, l)
									err = errors.New(msg)
								}
							default:
								// Any other character has to be equal
								if string(mask[i]) != l {
									shouldMove = false
									msg := fmt.Sprintf("character at position %d is not equal: %s", i, l)
									err = errors.New(msg)
								}
							}
						}

						if shouldMove {
							act(file, action, dir+"/"+file.Name(),
								to+"/"+file.Name())
						}
					}
				case "suffix":
					suffixes := configuration.Rules.Predicates.([]interface{})

					for _, suffix := range suffixes {
						suffix := suffix.(string)

						if strings.HasSuffix(filename, suffix) {
							act(file, action, dir+"/"+file.Name(),
								to+"/"+file.Name())
						}
					}
				case "prefix":
					prefixes := configuration.Rules.Predicates.([]interface{})

					for _, prefix := range prefixes {
						prefix := prefix.(string)

						if strings.HasPrefix(filename, prefix) {
							act(file, action, dir+"/"+file.Name(),
								to+"/"+file.Name())
						}
					}
				case "size":
					sizes := configuration.Rules.Predicates.([]interface{})

					for _, size := range sizes {
						size := size.(int64)
						fileSize := file.Size()

						log.Printf("[+] Dummy rule size %d %d", size, fileSize)
					}
				}
			}

		}

	}
}

func act(file os.FileInfo, action Actions, dir, to string) {
	switch action {
	case ActionMove:
		move(file, dir, to)
	case ActionDelete:
		delete(file)

	}
}

func move(file os.FileInfo, dir, to string) {
	log.Printf("[+] Moving file from %s to %s", file.Name(), file.Name())
	Move(dir, to)
}

func delete(file os.FileInfo) {
	os.Remove(file.Name())
}
