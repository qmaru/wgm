package web

import (
	"log"

	"wgm/webview"

	"github.com/spf13/cobra"
)

var (
	WebCmd = &cobra.Command{
		Use:   "web",
		Short: "Run web",
		Run: func(cmd *cobra.Command, args []string) {
			err := webview.Run()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)
