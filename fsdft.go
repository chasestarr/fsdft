package fsdft

import (
	"io/ioutil"
	"log"
	"os"
)

// FileHandler is a callback function passed to the
// file system DFT. Each working directory and file
// it comes across will be passed to this function.
type FileHandler func(root string, file os.FileInfo)

// DFT will do a depth-first traveral through a file system.
// signature takes in a root working directory and a callback
// function that is run on each file
func DFT(root string, cb FileHandler) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		cb(root, file)
		if file.IsDir() {
			path := root + "/" + file.Name()
			DFT(path, cb)
		}
	}
}
