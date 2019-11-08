package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestMove(t *testing.T) {
	srcPath, _ := ioutil.TempDir(os.TempDir(), "src")
	destPath, _ := ioutil.TempDir(os.TempDir(), "dest")

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
			err := ioutil.WriteFile(tt.args.src+"/"+tt.name, []byte(tt.content), os.ModePerm)

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
