package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "qqw",
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
	Short: "Print the version number of qqw",
	Long: "All Software has versions. This is qqw",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("qqw for golang v2.0")
	},
}
