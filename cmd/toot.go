package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	m "github.com/mattn/go-mastodon"
)

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
			Server:			viper.GetString("mastodon.server"),
			ClientID:		viper.GetString("mastodon.clientid"),
			ClientSecret:	viper.GetString("mastodon.clientsecret"),
		}
	  var email,pass string
	  email = viper.GetString("mastodon.email")
		pass = viper.GetString("mastodon.pass")
		c := m.NewClient(config)
		c.Authenticate(context.Background(), email, pass)
		//	ここから
		var toot string
		toot = args[0]
		c.PostStatus(context.Background(), &m.Toot{
			Status: toot,
		})
	},
}
