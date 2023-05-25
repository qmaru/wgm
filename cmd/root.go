package cmd

import (
	"fmt"
	"os"

	"wgm/cmd/apis"
	"wgm/cmd/dbs"
	"wgm/cmd/web"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "wgm",
		Short:   "wgm is a simple Wireguard VPN management tool",
		Version: "2.0-20230525",
	}
)

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(
		apis.ApiCmd,
		dbs.DBCmd,
		web.WebCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
