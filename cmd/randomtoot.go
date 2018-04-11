package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	m "github.com/mattn/go-mastodon"
)

const BUFSIZE = 1024
func init(){
	RootCmd.AddCommand(tootCmd)
}


//tootCmd = &cobra.Command{
var tootCmd = &cobra.Command{
	Use:	"toot",
	Short:	"Mastodon toot",
	Run: func(cmd *cobra.Command, args []string){

		viper.SetConfigName(".cdg")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/")
		viper.SetConfigType("json")
		err := viper.ReadInConfig()
		if err != nil{
			fmt.Fprintf(os.Stderr, "cannot read config file: %v", err)
			os.Exit(-1)
		}
		config := &m.Config{
			Server:			viper.GetString("list.mastodon.server"),
			ClientID:		viper.GetString("list.mastodon.clientid"),
			ClientSecret:	viper.GetString("list.mastodon.clientsecret"),
		}
		//var email,pass string
		email := viper.GetString("list.mastodon.email")
		pass := viper.GetString("list.mastodon.pass")
		c := m.NewClient(config)
		c.Authenticate(context.Background(), email, pass)
		//	ここから
		var toot string
		filepath := viper.GetString("list.mastodon.file")
		ans,err := readLine(filepath)
		if err != nil {
			panic(err)
		}


		toot = ans[1]
		c.PostStatus(context.Background(), &m.Toot{
			Status: toot,
		})
	},
}

func readLine(filename string)([]string,error){
	ans := make([]string,10)
	data, err := ioutil.ReadFile(filename)
	if err !=nil {
		return ans, fmt.Errorf(filename + "Can't be oppend")
	}
	ans = strings.Split(string(data), "\n")
	return ans, err
}
