package cmd

import (
//	"fmt"
//	"io/ioutil"
//	"strings"
	"github.com/spf13/cobra"
)

func init(){
	RootCmd.AddCommand(botCmd)
}

var botCmd = &cobra.Command{
	Use:	"bot",
	Short:	"Calculator of addition.",
	Long:	"Calculator to perform the addition",
	Run: func(cmd *cobra.Command, args []string){
		//var n1 int
		//n1, _ = (args[0])
		//fmt.Println(n1)
	},
}
