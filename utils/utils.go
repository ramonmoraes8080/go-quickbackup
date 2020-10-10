/*
Copyright © 2020 Ramon Moraes <ramonmoraes8080@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

// TODO Too much functions here. BREAK IT DOWN

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// https://gist.github.com/ik5/d8ecde700972d4378d87
const (
	SuccessColor = "\033[1;32m%s\033[0m"
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	MagentaColor = "\033[1;35m%s\033[0m"
)

func LoggerSuccess(msg string) {
	fmt.Println(fmt.Sprintf(SuccessColor, msg))
}

func LoggerInfo(msg string) {
	fmt.Println(fmt.Sprintf(InfoColor, msg))
}

func LoggerWarning(msg string) {
	fmt.Println(fmt.Sprintf(WarningColor, msg))
}

func LoggerError(msg string) {
	fmt.Println(fmt.Sprintf(ErrorColor, msg))
}

func LoggerMagenta(msg string) {
	fmt.Println(fmt.Sprintf(MagentaColor, msg))
}

func getHomePath() string {
	return os.Getenv("HOME")
}

func ExpandUser(path string) string {
	times := 1
	return strings.Replace(path, "~", getHomePath(), times)
}

func GetCurrentISOTimeString() string {
	currTime := time.Now()
	return currTime.Format("2006-01-02-150405")
}

func CheckFilePath(path string) (bool, error) {
	path = ExpandUser(path)
	if path == "" {
		return false, errors.New("Path is empty")
	} else {
		_, err := os.Stat(path)
		if err != nil {
			return false, errors.New(fmt.Sprintf(
				"File %s does not exist", path))
		}
	}
	return true, nil
}

// We pass a folder path and return a array of strings including all the
// relative file paths to it
func WalkDir(dirPath string) []string {
	var paths []string

	err := filepath.Walk(
		dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			paths = append(paths, path)
			return nil
		},
	)

	if err != nil {
		return []string{}
	}

	return paths
}

// We pass a list of files and folders as an array of strings and return a new
// array where the folder paths are expanded by call for WalkDir function
func WalkDirs(schemaFilePaths []string) []string {
	var ret []string
	for _, path := range schemaFilePaths {
		path := ExpandUser(path)

		f, err := os.Stat(path)
		// TODO Should we be using CheckFilePath here?

		if err != nil {
			// Doesn't exist? That's Ok. Move on to the next file path
			continue
		}

		switch mode := f.Mode(); {
		case mode.IsDir():
			ret = append(ret, WalkDir(path)...)
		case mode.IsRegular():
			ret = append(ret, path)
		}
	}
	return ret
}

func ReadInputInt(template string) int {
	var ret int
	print(template)
	fmt.Scanf("%d", &ret)
	return ret
}

// Source: https://opensource.com/article/18/6/copying-files-go
func Copy(src string, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
