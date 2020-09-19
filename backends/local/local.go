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
package local

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"gitlab.com/velvetkeyboard/go-quickbackup/constants"
	"gitlab.com/velvetkeyboard/go-quickbackup/utils"
)

type BackendLocalFilesystem struct {
	Path string
}

func (b *BackendLocalFilesystem) Init(path string) {
	b.Path = utils.ExpandUser(path)
}

func (b *BackendLocalFilesystem) Upload(zipFilePath string) {
	zipFileName := filepath.Base(zipFilePath)
	zipNewFilePath := filepath.Join(b.Path, zipFileName)
	utils.LoggerSuccess(fmt.Sprintf(
		"[Backend][Filesystem] copying %s to %s",
		zipFilePath,
		zipNewFilePath,
	))
	err := os.Rename(zipFilePath, zipNewFilePath)
	if err != nil {
		panic(err)
	}
}

func (b *BackendLocalFilesystem) List() []string {
	// TODO Maybe we could present this list grouped by date
	var ret []string
	for _, filePath := range utils.WalkDir(b.Path) {
		regexStr := fmt.Sprintf(
			"%s-\\w+-\\d{4}-\\d{2}-\\d{2}-\\d{6}\\.%s",
			constants.PREFIX, constants.EXTENSION,
		)
		if match, _ := regexp.MatchString(regexStr, filePath); match {
			ret = append(ret, filePath)
		}
	}
	return ret
}

// Copy file zip to an arbitrary path
func (b *BackendLocalFilesystem) Download(zipFilePath string, moveTo string) {
	// TODO Check if the upload file exists?
	_, zipFileName := filepath.Split(zipFilePath)
	// err := os.Rename(zipFilePath, filepath.Join(moveTo, zipFileName))
	_, err := utils.Copy(zipFilePath, filepath.Join(moveTo, zipFileName))
	if err != nil {
		panic(err) // TODO Should we panic here?
	}
}
