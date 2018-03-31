package cmd

import (

	"fmt"

	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use: "cdg",
	Short: "cd-bookmark for golang",
	Long: "cd-bookmark for golang by e10ulen",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "Version",
	Short: "Print the version number of cdg",
	Long: "All Software has versions. This is cdg",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("cd-bookmark for golang v1.0")
	},
}
