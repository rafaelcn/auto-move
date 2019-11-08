package main

import (
	"io"
	"log"
	"os"
)

// Move moves files from the source path to the destination path.
func Move(src string, dest string) {
	destFile, err := os.Create(dest)

	if err != nil {
		log.Printf("[!] Couldn't create a file to copy. Reason: %v", err)
	}

	srcFile, err := os.Open(src)

	_, err = io.Copy(destFile, srcFile)

	if err != nil {
		log.Printf("[!] Couldn't copy contents of the %s file. Reason %v", src, err)
	}

	srcFile.Close()
	err = os.Remove(src)

	if err != nil {
		log.Printf("[!] Couldn't remove file %s. Reason %v", src, err)
	}
}
