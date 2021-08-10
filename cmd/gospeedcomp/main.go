package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/packr"
)

var root string

func main() {
	root, _ = os.Getwd()

	println("Extract files")
	println("-------------")
	extractFiles()

	testFiles := []string{"50B.txt", "100B.txt", "200B.txt", "500B.txt", "1KB.txt", "50KB.txt", "100KB.txt", "200KB.txt", "500KB.txt", "1MB.txt"}
	iterations := 20000

	println()
	println("Execute speedtest")
	println("-----------------")
	for _, testFile := range testFiles {
		println(testFile)
		start := time.Now()
		for i := 0; i < iterations; i++ {
			getFileFromBox(testFile)
		}
		t := time.Now()
		elapsed := t.Sub(start)
		println("  Box: " + elapsed.String())
		
		start = time.Now()
		for i := 0; i < iterations; i++ {
			getFileFromFS(testFile)
		}
		t = time.Now()
		elapsed = t.Sub(start)
		println("  FS:  " + elapsed.String())
		
	}
}

func getFileFromFS(fileName string) {
	fullpath := filepath.Join(root, "assets", fileName)
	_, err := os.Stat(fullpath)
	if err == nil {
		fileContentByte, err := ioutil.ReadFile(fullpath)
		if err != nil {
			panic(fmt.Errorf("error while reading file %s: %s", fullpath, err.Error()))
		}

		fileContent := string(fileContentByte)
		if len(fileContent) == 0 {
			panic(fmt.Errorf("error reading file %s, it is empty", fullpath))
		}
		return
	}

	panic(fmt.Errorf("file %s not found: %s", fullpath, err))
}

func getFileFromBox(fileName string) {
	box := packr.NewBox("../../assets")
	fileContent, err := box.FindString(fileName)
 	if err != nil {
 		panic(fmt.Errorf("file %s not found in box: %s", fileName, err.Error()))
 	}
	if len(fileContent) == 0 {
		panic(fmt.Errorf("error reading file %s from box, it is empty", fileName))
	}
}

func extractFiles() {
	box := packr.NewBox("../../assets/")
	// Extract folder name from path
	var assetPath = filepath.Base(box.Path)

	// Skip folder if it already exists
	println("Check for asset folder '" + assetPath + "'")
	if ok, _ := Exists(assetPath); ok {
		return
	}

	// Extract every file in the box
	println("Extract asset folder '" + assetPath + "'")
	box.Walk(func(path string, f packr.File) error {		
		// Make all folders
		os.MkdirAll(assetPath + string(filepath.Separator) + filepath.Dir(path), 0777)

		// Write data to files
		content, _ := box.FindString(path)
		var filename = assetPath + string(filepath.Separator) + path
		println("Extract asset file '" + filename + "'")
		ioutil.WriteFile(filename, []byte(content), 0644)
		return nil
	})
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}