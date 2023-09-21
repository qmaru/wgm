package cmd

import (
	"fmt"
	"os"

	"wgm/cli/cmd/apis"
	"wgm/common/utils"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "wgm",
		Short:   "wgm is a simple Wireguard VPN management tool",
		Version: utils.BackendVersion,
	}
)

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(
		apis.ApiCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
