package apis

import (
	"log"

	"wgm/apis"

	"github.com/spf13/cobra"
)

var (
	listenAddr string
	ApiCmd     = &cobra.Command{
		Use:   "api",
		Short: "Run web api",
		Run: func(cmd *cobra.Command, args []string) {
			err := apis.Run(listenAddr, nil)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	ApiCmd.Flags().StringVarP(&listenAddr, "listen", "l", "127.0.0.1:8373", "Listen address [host:port]")
}
