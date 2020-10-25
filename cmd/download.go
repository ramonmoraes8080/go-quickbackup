/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	googledrive "gitlab.com/velvetkeyboard/go-quickbackup/backends/googledrive"
	local "gitlab.com/velvetkeyboard/go-quickbackup/backends/local"
	"gitlab.com/velvetkeyboard/go-quickbackup/config"
	fs "gitlab.com/velvetkeyboard/go-quickbackup/utils/filesystem"
	"gitlab.com/velvetkeyboard/go-quickbackup/utils/console/input"
	"gitlab.com/velvetkeyboard/go-quickbackup/utils/console/output"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		schemaName, _ := cmd.Flags().GetString("schema")
		locationName, _ := cmd.Flags().GetString("location")
		configFilePath, _ := cmd.Flags().GetString("config")

		if _, err := fs.CheckFilePath(configFilePath); err != nil {
			output.LoggerError(err.Error())
			os.Exit(0)
		}

		config := new(config.Configuration)
		config.Init(configFilePath)

		// falling back to defaults
		if schemaName == "" {
			schemaName = config.GetDefaultSchemaName()
		}

		if locationName == "" {
			locationName = config.GetDefaultLocationName()
		}

		// Check if Location Name and the Backend Engine associated with it
		// exists

		if _, err := config.CheckLocationStatus(locationName); err != nil {
			output.LoggerError(err.Error())
			os.Exit(0)
		}

		output.LoggerInfo(fmt.Sprintf(
			"Which \"%s\" schema backup from location \"%s\" you want to download:\n",
			schemaName,
			locationName,
		))

		location := config.Locations[locationName]

		var backendConfig map[interface{}]interface{}
		if "filesystem" != location.Backend {
			// Go is not clever enough to resolve this one for me...
			backendConfig = config.Backends[location.Backend].(map[interface{}]interface{})
		}

		var fileNamesLen int
		var fileNames []string

		switch location.Backend {
		case "filesystem":
			backend := new(local.BackendLocalFilesystem)
			backend.Init(location.Path)
			fileNames = backend.List()

		case "google_drive":
			jsonCredentialPath, _ := backendConfig["json_credential"].(string)
			backend := new(googledrive.BackendGoogleDrive)
			backend.Init(location.Path, jsonCredentialPath)
			fileNames = backend.List(schemaName)

		default:
			output.LoggerError(fmt.Sprintf(
				"Backend \"%s\" is not implemented yet :'(",
				location.Backend,
			))
			os.Exit(0)
		}

		fileNamesLen = len(fileNames)

		for i, fileName := range fileNames {
			output.LoggerSuccess(fmt.Sprintf(
				"%d - %s",
				fileNamesLen-i,
				fileName,
			))
		}

		idx := fileNamesLen - input.ReadInt("\nType Number > ")

		if idx >= fileNamesLen || idx < 0 {
			output.LoggerError(fmt.Sprintf("%d is not a valid number\n", idx))
			os.Exit(0)
		}

		switch location.Backend {
		case "filesystem":
			backend := new(local.BackendLocalFilesystem)
			backend.Download(fileNames[idx], ".")
		case "google_drive":
			jsonCredentialPath, _ := backendConfig["json_credential"].(string)
			backend := new(googledrive.BackendGoogleDrive)
			backend.Init(location.Path, jsonCredentialPath)
			backend.Download(fileNames[idx], ".")
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	downloadCmd.Flags().StringP("schema", "s", "", "Schema Name")
	downloadCmd.Flags().StringP("location", "l", "", "Location Name")
}
