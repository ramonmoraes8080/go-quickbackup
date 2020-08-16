package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "flag"
)

import (
    "gitlab.com/velvetkeyboard/go-backup/config"
    "gitlab.com/velvetkeyboard/go-backup/zipfile"
    "gitlab.com/velvetkeyboard/go-backup/backends"
    "gitlab.com/velvetkeyboard/go-backup/utils"
)

func getFilePathsRecursively(dir_path string) []string {
    var paths []string

    err := filepath.Walk(
        dir_path,
        func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            paths = append(paths, path)
            return nil
        });

    if err != nil {
        panic(err)
    }

    return paths
}

// This is a simple encapsulation to expand the path if it's a dir path.
// Isolated it here for readability.
// Maybe add support to a channel here
func getAllSchemaFilesPaths(schemaFilePaths []string) []string {
    var ret []string
    for _, path := range schemaFilePaths {
        path := utils.ExpandUser(path)

        f, err := os.Stat(path)

        if err != nil {
            fmt.Println("[Zipping][Scanning]", path, err)
            continue
        } else {
        }

        switch mode := f.Mode(); {
            case mode.IsDir():
                ret = append(ret, getFilePathsRecursively(path)...)
            case mode.IsRegular():
                ret = append(ret, path)
        }
    }
    return ret
}

func createBackupZipFile(cfg *config.Configuration, schemaName string) string {
    currTimeStr := utils.GetCurrentISOTimeString()

    zipFileTitle := "backup-" + schemaName + "-" + currTimeStr
    zipFileNameFull := zipFileTitle + ".zip"

    zfile := new(zipfile.ZipFile)
    zfile.Init(zipFileNameFull)

    fmt.Println("[Zipping] ...",)

    for _, path := range getAllSchemaFilesPaths(cfg.Schemas[schemaName]) {
        fileContentBytes, _ := ioutil.ReadFile(path)
        zfile.AppendBytes(zipFileTitle + path, fileContentBytes)
    }

    zfile.Save()
    return zipFileNameFull
}

func main() {
    var configFilePathFlag string
    var modeFlag string
    var schemaFlag string
    var locationFlag string

    flag.StringVar(&configFilePathFlag, "c", "~/.backup.yaml", "config file path")
    flag.StringVar(&modeFlag, "m", "up", "modes: up, dw, ls")
    flag.StringVar(&schemaFlag, "s", "", "schema to backup")
    flag.StringVar(&locationFlag, "to", "", "location where to upload")

    flag.Parse()

    // gobackup -c=backup.yaml -m=up -s=fedora -to=local

    config := new(config.Configuration)
    config.Init(configFilePathFlag) // ("~/.backup.yaml")

    if schemaFlag == "" {
        schemaFlag = config.Defaults.Schema
    }

    if locationFlag == "" {
        locationFlag = config.Defaults.Location
    }

    // where should we move, copy, upload the .zip file
    location, ok := config.Locations[locationFlag]

    if !ok {
        fmt.Println(
            "Location", locationFlag, "does not exist in", configFilePathFlag);
        os.Exit(0)
    }

    zipFileNameFull := createBackupZipFile(config, schemaFlag)

    switch location.Backend {
        case "filesystem":
            _, ok := config.Backends[location.Backend]
            if ok {
                backendLocalFs := new(local.BackendLocalFilesystem)
                backendLocalFs.Init(location.Path)
                if modeFlag == "up" {
                    backendLocalFs.Upload(zipFileNameFull);
                }
            } else {
                fmt.Println("Backend", location.Backend, "config has errors")
                os.Exit(0)
            }
        case "google_drive":
            fmt.Println("Not Implemented yet u_u'")
            os.Exit(0)
        default:
            fmt.Println("Backend", location.Backend, "not supported")
            os.Exit(0)
    }
}
