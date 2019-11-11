package main

import (
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
						shouldMove := false

						for i := 0; i < len(mask); i++ {
							l := string(filename[i])

							switch string(mask[i]) {
							case "D":
								_, err := strconv.Atoi(l)

								if err != nil {
									break
								}
							case "L":
								r := []rune(l)

								if !unicode.IsLetter(r[0]) {
									break
								}
							default:
								// Any other character has to be equal
								if string(mask[i]) != l {
									break
								}
							}

							shouldMove = true
						}

						if shouldMove {
							move(file, dir+"/"+file.Name(), to+"/"+file.Name())
						}
					}
				case "suffix":
					suffixes := configuration.Rules.Predicates.([]interface{})

					for _, suffix := range suffixes {
						suffix := suffix.(string)

						if strings.HasSuffix(filename, suffix) {
							move(file, dir+"/"+file.Name(), to+"/"+file.Name())
						}
					}
				case "prefix":
					prefixes := configuration.Rules.Predicates.([]interface{})

					for _, prefix := range prefixes {
						prefix := prefix.(string)

						if strings.HasPrefix(filename, prefix) {
							move(file, dir+"/"+file.Name(), to+"/"+file.Name())
						}
					}
				}
			}

		}

	}
}

func move(file os.FileInfo, dir, to string) {
	log.Printf("[+] Moving file from %s to %s", file.Name(), file.Name())
	Move(dir, to)
}
