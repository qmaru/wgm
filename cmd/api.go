package cmd

import (
	"fmt"

	"wgm/apis"

	"github.com/spf13/cobra"
)

var (
	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Run Web API",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("You can now view web in the browser.\n\n")
			fmt.Printf("	http://127.0.0.1:8373\n\n")
			apis.Run()
		},
	}
)
