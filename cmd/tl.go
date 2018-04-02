package cmd

import (
	"fmt"
	"context"
	"log"
	"os"
	"strings"

	m "github.com/mattn/go-mastodon"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

func init(){
	RootCmd.AddCommand(timelineCmd)
}

var timelineCmd = &cobra.Command{
	Use:	"tl",
	Short:	"Mastodon GetTimeline.",
	Run: func(cmd *cobra.Command, args []string){
		viper.SetConfigName(".cdg")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/")
		viper.SetConfigType("json")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot read config file: %v", err)
			os.Exit(1)
		}
		config := &m.Config{
			Server: viper.GetString("list.mastodon.server"),
			ClientID: viper.GetString("list.mastodon.clientid"),
			ClientSecret: viper.GetString("list.mastodon.clientsecret"),
		}
		//var email string
		//var pass string
		email := viper.GetString("list.mastodon.email")
		pass := viper.GetString("list.mastodon.pass")
		c := m.NewClient(config)
		c.Authenticate(context.Background(), email, pass)
		if err != nil {
			log.Println("Client:", err)
		}
		//	GetTimeline
		wsc := c.NewWSClient()
		q, err := wsc.StreamingWSPublic(context.Background(), true)
		for e := range q {
			if t, ok := e.(*m.UpdateEvent); ok {
				s := t.Status.Content
				s = strings.Replace(s, "<p>", "", -1)
				s = strings.Replace(s, "</p>", "", -1)
				blue := color.New(color.FgBlue).SprintFunc()
				magenta := color.New(color.FgMagenta).SprintFunc()
				red := color.New(color.FgRed).SprintFunc()
				//
				fmt.Printf("%s %s(%s) %s\n", blue(t.Status.CreatedAt.Local().Format("15:04:05")), magenta(t.Status.Account.DisplayName), magenta(t.Status.Account.Acct), red(s))
				//

			}
		}
	},

}
