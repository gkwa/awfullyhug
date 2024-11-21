package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/gkwa/awfullyhug/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of awfullyhug",
	Long:  `All software has versions. This is awfullyhug's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
