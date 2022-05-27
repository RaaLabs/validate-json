package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	jsonFilepaths := getJsonFilePaths()
	var failure int

	for _, f := range jsonFilepaths {
		func() {
			fh, err := os.Open(f)
			if err != nil {
				log.Printf("error: failed to open file for reading: %v\n", err)
				os.Exit(1)
			}
			defer fh.Close()

			js, err := io.ReadAll(fh)
			if err != nil {
				log.Printf("error: failed to read file: %v\n", err)
				os.Exit(1)
			}

			var data map[string]any

			if len(js) < 1 {
				return
			}

			err = json.Unmarshal(js, &data)
			if err != nil {
				log.Printf("Json format error: for file %v, failed to unmarshal data: %v\n", f, err)
				failure = 1
			}
		}()
	}

	if failure != 0 {
		log.Printf("Check valid json format: There were json files with invalid format, please correct files above before merging PR")
		os.Exit(failure)
	} else {
		log.Printf("Check valid json format: All json files have valid format")
	}
}

func getJsonFilePaths() []string {
	files := []string{}

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(path, ".json") {
				files = append(files, path)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return files
}
