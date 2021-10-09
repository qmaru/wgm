package cmd

import (
	"wgm/apis"

	"github.com/spf13/cobra"
)

var (
	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Run Web API",
		Run: func(cmd *cobra.Command, args []string) {
			apis.Run()
		},
	}
)
