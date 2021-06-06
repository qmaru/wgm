package cmd

import (
	"fmt"
	"os"

	"wgm/models"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "wgm",
		Short:   "wgm is a simple Wireguard VPN management tool",
		Version: "1.2-20210606-1",
		Run: func(cmd *cobra.Command, args []string) {
			if models.TableCheck() {
				cmd.Help()
			} else {
				fmt.Println("Please use [wgm init] to init database.")
			}
		},
	}
)

func Execute() {
	rootCmd.AddCommand(
		dbCmd,
		serverCmd,
		userCmd,
		ruleCmd,
		rulemapCmd,
		showCmd,
		mgrCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
