package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)
const DateFormat = "2006/01/02 15:04"
func init(){
	RootCmd.AddCommand(htCmd)
}

var htCmd = &cobra.Command{
	Use:	"ht",
	Short:	"Auto Commit And Push.",
	Run:	func(cmd *cobra.Command, args[]string){
		tm := time.Now()
		dir, err := os.Getwd()
		if err == nil {
			fmt.Print(dir, "\n")
		}
		fmt.Print("DateTime:", tm.Format(DateFormat), "\n")
		//	Add Git UnStageFiles
		add, err := exec.Command("git", "add", "--all").CombinedOutput()
		if err != nil{
			log.Print("Error...Add\n")
		}
		//	Git Commit
		cmt,err := exec.Command("git", "commit", "-m", "Commit:"+tm.Format(DateFormat)).CombinedOutput()
		if err != nil {
			log.Print("Error...Commit\n")
		}
		push, err :=exec.Command("git", "push").CombinedOutput()
		if err != nil {
			log.Print("Error...Push\n")
		}
	fmt.Print("Git Add:", string(add), "\n")
	fmt.Print("Git Commit:", string(cmt), "\n")
	fmt.Print("Git Push:", string(push), "\n")
	},
}

