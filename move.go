package main

import (
	"log"
	"os"
)

// Move
func Move(src string, dest string) {
	err := os.Rename(src, dest)

	if err != nil {
		log.Printf("[!] Couldn't move file to the new folder. Reason %v", err)
	}
}
