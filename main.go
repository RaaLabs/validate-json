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
	directory := getEnv("INPUT_DIRECTORY", ".")
	log.Printf("Using directory: %v\n ", directory)
	jsonFilepaths := getJsonFilePaths(directory)
	var failure int
	var validFiles int
	var invalidFiles int

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

			if len(js) < 1 {
				return
			}

			if strings.HasPrefix(string(js), "[") {
				data := []map[string]any{}
				err = json.Unmarshal(js, &data)
				if err != nil {
					log.Printf("Json format error: for file %v, failed to unmarshal data: %v\n", f, err)
					invalidFiles = invalidFiles + 1
					failure = 1
				} else {
					validFiles = validFiles + 1
				}
			} else {
				data := make(map[string]any)
				err = json.Unmarshal(js, &data)
				if err != nil {
					log.Printf("Json format error: for file %v, failed to unmarshal data: %v\n", f, err)
					invalidFiles = invalidFiles + 1
					failure = 1
				} else {
					validFiles = validFiles + 1
				}
			}
		}()
	}

	if failure != 0 {
		log.Printf("Check valid json format: There were json files with invalid format, please correct files before merging PR")
		log.Printf("Invalid json files: %v, valid json files: %v", invalidFiles, validFiles)
		os.Exit(failure)
	} else {
		log.Printf("Check valid json format: All json files have valid format")
		log.Printf("Invalid json files: %v, valid json files: %v", invalidFiles, validFiles)
	}
}

func getJsonFilePaths(directory string) []string {
	files := []string{}

	err := filepath.Walk(directory,
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

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
