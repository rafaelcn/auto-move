package main

import (
	"log"
	"os"
)

// Move moves files from the source path to the destination path.
func Move(src string, dest string) {
	// FIXME: Might contain a bug regarding the move of files between devices.
	// It's best to change to use os.Create > os.Copy > os.Remove.
	err := os.Rename(src, dest)

	if err != nil {
		log.Printf("[!] Couldn't move file to the new folder. Reason %v", err)
	}
}
