package local

import (
	"fmt"
	"os"
	"path/filepath"
)

import (
	"gitlab.com/velvetkeyboard/go-backup/utils"
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

// Copy file zip to an arbitrary path
func (b *BackendLocalFilesystem) Download(zipFileName string, dest string) {
}
