package main

import (
	"context"
	"fmt"
	"log"

	m "github.com/mattn/go-mastodon"
)

func main() {
	config := &m.Config{
		Server:		"https://ichiji.social",
		ClientID:	"026a9bafbc6c83437fa81c56797ec177028bac19d683660fe31ff4237cfecc23",
		ClientSecret:	"c8feea25bf62ab6d1e447b49985601ab0201dc559a466362548bf975c17f37dc",
	}

	c := m.NewClient(config)
	err := c.Authenticate(context.Background(), "login email", "login pass")
	if err != nil {
		log.Fatal(err)
	}

	//	WebSocket用サブインスタンス
	wsc := c.NewWSClient()

	q, err := wsc.StreamingWSPublic(context.Background(), true)
	if err != nil {
		fmt.Printf("  ERR: %s\n", err)
	}
	for e := range q {
		if t, ok := e.(*m.UpdateEvent); ok {
			fmt.Printf("\x1b[37m[%s] \x1b[35m%-20s: \x1b[33m%s\n", t.Status.CreatedAt.Local().Format("15:04:05"), t.Status.Account.Acct, t.Status.Content)
		}
	}
}
