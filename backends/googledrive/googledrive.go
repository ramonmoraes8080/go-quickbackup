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
	"gitlab.com/velvetkeyboard/go-quickbackup/utils"
)

type BackendGoogleDrive struct {
	Path string
}

func (b *BackendGoogleDrive) Init(path string) {
	b.Path = utils.ExpandUser(path)
}

func (b *BackendGoogleDrive) Upload(zipFilePath string) {
}

// Copy file zip to an arbitrary path
func (b *BackendGoogleDrive) Download(zipFileName string, dest string) {
}
