package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
	}
}
func main() {

}
