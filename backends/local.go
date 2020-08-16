package local

import (
    "os"
    "fmt"
    "path/filepath"
)

import (
    "gitlab.com/velvetkeyboard/go-backup/utils"
)

type BackendLocalFilesystem struct {
    Path string
}

/*
func (b *BackendLocalFilesystem) Init(options interface{}) {
    b.Path = utils.ExpandUser(
        fmt.Sprintf(
            "%v", options.(map[interface{}]interface{})["path"]));
}
*/

func (b *BackendLocalFilesystem) Init(path string) {
    b.Path = utils.ExpandUser(path)
}

func (b *BackendLocalFilesystem) Upload(zipFilePath string) {
    zipFileName := filepath.Base(zipFilePath)
    zipNewFilePath := filepath.Join(b.Path, zipFileName)
    fmt.Println("[Backend][Filesystem] copying", zipFilePath, "to", zipNewFilePath)
    err := os.Rename(zipFilePath, zipNewFilePath)
    if err != nil {
        panic(err)
    }
}
