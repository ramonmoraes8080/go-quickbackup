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
	"gitlab.com/velvetkeyboard/go-quickbackup/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
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

		if _, err := utils.CheckFilePath(configFilePath); err != nil {
			utils.LoggerError(err.Error())
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
			utils.LoggerError(err.Error())
			os.Exit(0)
		}

		utils.LoggerInfo(fmt.Sprintf(
			"Listing current backups from \"%s\" schema at location set \"%s\"",
			schemaName,
			locationName,
		))
		location := config.Locations[locationName]

		// TODO Add AWS S3 Backend
		// TODO Add Google Drive Backend
		switch location.Backend {
		case "filesystem":
			backend := new(local.BackendLocalFilesystem)
			backend.Init(location.Path)
			length := len(backend.List())
			for i, fileName := range backend.List() {
				println(length-i, "-", fileName)
			}
		case "googledrive":
			backend := new(googledrive.BackendGoogleDrive)
			backend.Init(location.Path)
		default:
			utils.LoggerError(fmt.Sprintf(
				"Backend \"%s\" is not implemented yet :'(",
				location.Backend,
			))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().StringP("schema", "s", "", "Schema Name")
	listCmd.Flags().StringP("location", "l", "", "Location Name")
}
