package schema

import (
	"os"
	"path/filepath"

	"gitlab.com/velvetkeyboard/go-backup/config"
	"gitlab.com/velvetkeyboard/go-backup/utils"
)

type Schema struct {
	Name  string
	Files []string
}

func (s *Schema) Init(cfg *config.Configuration, schemaName string) {
	s.Name = schemaName
	s.Files = expandFilesPaths(cfg.Schemas[schemaName])
}

// We pass a folder path and return a array of strings including all the
// relative file paths to it
func walkPath(dir_path string) []string {
	var paths []string

	err := filepath.Walk(
		dir_path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			paths = append(paths, path)
			return nil
		},
	)

	if err != nil {
		panic(err) // TODO do we really panic here?
	}

	return paths
}

// We pass a list of files and folders as an array of strings and return a new
// array where the folder paths are expanded by call for walkPath function
func expandFilesPaths(schemaFilePaths []string) []string {
	var ret []string
	for _, path := range schemaFilePaths {
		path := utils.ExpandUser(path)

		f, err := os.Stat(path)

		if err != nil {
			// Doesn't exist? That's Ok. Move on to the next file path
			continue
		}

		switch mode := f.Mode(); {
		case mode.IsDir():
			ret = append(ret, walkPath(path)...)
		case mode.IsRegular():
			ret = append(ret, path)
		}
	}
	return ret
}
