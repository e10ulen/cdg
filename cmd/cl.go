package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init(){
	RootCmd.AddCommand(colorCmd)
}

var colorCmd = &cobra.Command{
	Use:	"cl",
	Short:	"color.",
	Run: func(cmd *cobra.Command, args []string){
		color.Red("Print green")
	},
}
