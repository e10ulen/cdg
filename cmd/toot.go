package cmd

import (
	"context"
	"log"
	"github.com/spf13/cobra"
	m "github.com/mattn/go-mastodon"
)

func init(){
	RootCmd.AddCommand(tootCmd)
}

var tootCmd = &cobra.Command{
	Use:	"toot",
	Short:	"Mastodon toot",
	Run: func(cmd *cobra.Command, args []string){


		config := &m.Config{
			Server:			"https://ichiji.social",
			ClientID:		"026a9bafbc6c83437fa81c56797ec177028bac19d683660fe31ff4237cfecc23",
			ClientSecret:	"c8feea25bf62ab6d1e447b49985601ab0201dc559a466362548bf975c17f37dc",
		}
		c := m.NewClient(config)
		err := c.Authenticate(context.Background(), "email", "pass")
		if err != nil {
			log.Fatal(err)
		}

		//	ここから
		var toot string
		toot = args[0]
		c.PostStatus(context.Background(), &m.Toot{
			Status: toot,
		})
	},
}
