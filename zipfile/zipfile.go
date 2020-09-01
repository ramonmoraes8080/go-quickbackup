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
		// ioWriter.Close()
		panic(err)
	}

	_, err = ioWriter.Write(fileContent)

	if err != nil {
		// ioWriter.Close()
		panic(err)
	}
}

func (zf *ZipFile) Save() {
	zf.ZipWriter.Close()
	osZipFile, err := os.Create(zf.FilePath)

	if err != nil {
		panic(err) // TODO do we really want to panic?
	}

	_, err = osZipFile.Write(zf.Buffer.Bytes())

	if err != nil {
		panic(err) // TODO do we really want to panic?
	}

	osZipFile.Close()
}
