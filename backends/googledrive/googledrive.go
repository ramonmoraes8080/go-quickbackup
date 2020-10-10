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
package googledrive

import (
	gdrive_svc "gitlab.com/velvetkeyboard/go-quickbackup/services/googledrive"
	"gitlab.com/velvetkeyboard/go-quickbackup/utils"
)

type BackendGoogleDrive struct {
	Path           string
	ConfigPath     string
	GoogleDriveSvc *gdrive_svc.GoogleDrive
}

func (b *BackendGoogleDrive) Init(path string, configPath string) {
	b.Path = utils.ExpandUser(path)
	b.ConfigPath = utils.ExpandUser(configPath)
	b.GoogleDriveSvc = new(gdrive_svc.GoogleDrive)
	b.GoogleDriveSvc.Init(b.ConfigPath)
}

func (b *BackendGoogleDrive) Upload(zipFilePath string) {
}

func (b *BackendGoogleDrive) Download(zipFileName string, dest string) {
	zipFileNameID := b.GoogleDriveSvc.GetFileId(zipFileName)
	println("zipFileNameID", zipFileNameID)
	b.GoogleDriveSvc.DownloadFile(zipFileNameID)
}

func (b *BackendGoogleDrive) List() []string {
	var ret []string
	if files := b.GoogleDriveSvc.ListFilesFromFolder(b.Path); len(files) > 0 {
		for _, file := range files {
			ret = append(ret, file.Name)
			b.Download(file.Name, "/foo/bar")
			break
		}
	}
	return ret
}
