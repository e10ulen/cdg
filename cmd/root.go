package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "cdg",
	Short:	"My CommandTools",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of cdg",
	Long: "All Software has versions. This is cdg",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("cd-bookmark for golang v1.0")
	},
}
