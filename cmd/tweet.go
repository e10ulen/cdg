package cmd

import (
	"fmt"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

func init(){
	RootCmd.AddCommand(tweetCmd)
}

var tweetCmd = &cobra.Command{
	Use:		"tweet",
	Short:		"BlueBird tweet",
	Run: func(cmd *cobra.Command, args []string){
		viper.SetConfigName(".zzz")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/")
		viper.SetConfigType("json")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Print(err)
		}
		var csKey, cssKey, acToken, asToken string
		csKey	= viper.GetString("list.twitter.consumerkey")
		cssKey	= viper.GetString("list.twitter.consumersecretkey")
		acToken	= viper.GetString("list.twitter.accesstoken")
		asToken	= viper.GetString("list.twitter.accesstokensecret")
		if csKey == "" || cssKey == "" || acToken == "" || asToken == "" {
			fmt.Print(os.Stderr, "config load missing")
			if csKey == "" {fmt.Println("No CsKey")}
			if cssKey == "" {fmt.Println("No CssKey")}
			if acToken == "" {fmt.Println("No AcToken")}
			if asToken == "" {fmt.Println("No AsToken")}
			os.Exit(1)
		}
		config := oauth1.NewConfig(csKey, cssKey)
		token := oauth1.NewToken(acToken, asToken)
		httpClient := config.Client(oauth1.NoContext, token)
		client := twitter.NewClient(httpClient)
		//	send tweet
		tweet, resp, err := client.Statuses.Update(args[0],nil)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(resp)
		fmt.Print(tweet)
	},
}
