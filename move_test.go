package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	var err error

	srcPath, err := os.MkdirTemp("", "src")

	assert.Nil(t, err)

	destPath, err := os.MkdirTemp("", "dst")

	assert.Nil(t, err)

	type args struct {
		src  string
		dest string
	}

	tests := []struct {
		name    string
		content string
		args    args
	}{
		{
			name:    "file_1.txt",
			content: "This is the content of file_1.txt",
			args:    args{src: srcPath, dest: destPath},
		},
		{
			name:    "file_2.txt",
			content: "Noam Chomsky",
			args:    args{src: srcPath, dest: destPath},
		},
		{
			name:    "file_3.txt",
			content: "Alan Turing",
			args:    args{src: srcPath, dest: destPath},
		},
		{
			name:    "file_4.txt",
			content: "Hello world from a go test.",
			args:    args{src: srcPath, dest: destPath},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.WriteFile(tt.args.src+"/"+tt.name, []byte(tt.content), os.ModePerm)

			if err != nil {
				t.Fatalf("Couldn't create file for testing. Reason %v", err)
			}

			// Move the file to the destination folder
			Move(tt.args.src+"/"+tt.name, tt.args.dest+"/"+tt.name)

			content, err := ioutil.ReadFile(tt.args.dest + "/" + tt.name)

			if err != nil {
				t.Fatalf("Couldn't read file contents. Reason %v", err)
			}

			if !reflect.DeepEqual(content, []byte(tt.content)) {
				t.Fatalf("Want %v, got %v", tt.content, content)
			} else {
				t.Logf("Want %v, got %v", tt.content, string(content))
			}
		})
	}
}
