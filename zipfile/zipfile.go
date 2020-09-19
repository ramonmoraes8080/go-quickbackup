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
package zipfile

import (
	"archive/zip"
	"bytes"
	"os"
	"path"
)

type ZipFile struct {
	Name      string
	FilePath  string
	Files     []string
	Buffer    *bytes.Buffer
	ZipWriter *zip.Writer
}

func (zf *ZipFile) Init(file_path string) {
	zf.Name = path.Base(file_path)
	zf.FilePath = file_path
	zf.Buffer = new(bytes.Buffer)
	zf.ZipWriter = zip.NewWriter(zf.Buffer)
}

func (zf *ZipFile) AppendString(fileName string, fileContent string) {
	zf.AppendBytes(fileName, []byte(fileContent))
}

func (zf *ZipFile) AppendBytes(fileName string, fileContent []byte) {
	ioWriter, err := zf.ZipWriter.Create(fileName)

	if err != nil {
		panic(err) // TODO do we really want to panic?
	}

	_, err = ioWriter.Write(fileContent)

	if err != nil {
		panic(err) // TODO do we really want to panic?
	}
}

func (zf *ZipFile) Save() {
	zf.ZipWriter.Close()

	osZipFile, err := os.Create(zf.FilePath)

	if err != nil {
		panic(err) // TODO do we really want to panic?
	}

	defer osZipFile.Close()

	_, err = osZipFile.Write(zf.Buffer.Bytes())

	if err != nil {
		panic(err) // TODO do we really want to panic?
	}
}
