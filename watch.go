package main

import (
	"io/ioutil"
	"log"
	"path"
)

func Watch(configuration Configuration) {
	// sync these configurations with signals
	folders := configuration.Folders.Watched
	to := configuration.Folders.To

	extensions := configuration.Predicates.Extensions

	// os.Rename to move files or use os.(Create|Copy|Remove) ...

	for {
		// Break the loop if no folders are being currently watched
		if len(folders) <= 0 {
			break
		}

		for _, dir := range folders {
			files, err := ioutil.ReadDir(dir)

			if err != nil {
				// Implement a try system and remove a folder if the system
				log.Printf("[!] Failed to read dir %s. Reason %v", dir, err)
			}

			// O(m*n)
			for _, file := range files {
				for _, extension := range extensions {
					log.Printf("[+] Verifying file %s", file.Name())
					if extension == path.Ext(file.Name()) {
						log.Printf("[+] Moving file from %s to %s", file.Name(), file.Name())
						Move(dir+"/"+file.Name(), to+"/"+file.Name())
					}
				}
			}

		}
	}
}
