package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const (
	ExecMain = "/usr/bin/wg-quick"
	EXecSub  = "/usr/bin/wg"
)

func RunShell(s string, args ...string) {
	cmd := exec.Command(s, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(strings.TrimSpace(string(out)))
}

func startShell(c string) {
	RunShell(ExecMain, "up", c)
}

func stopShell(c string) {
	RunShell(ExecMain, "down", c)
}

func statusShell() {
	RunShell(EXecSub)
}

var (
	mgrCmd = &cobra.Command{
		Use:   "mgr",
		Short: "wireguard [Start|Stop|Restart|Status]",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cfgPath  string
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Config",
		Run: func(cmd *cobra.Command, args []string) {
			startShell(cfgPath)
		},
	}
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop Config",
		Run: func(cmd *cobra.Command, args []string) {
			stopShell(cfgPath)
		},
	}
	restartCmd = &cobra.Command{
		Use:   "restart",
		Short: "Restart Config",
		Run: func(cmd *cobra.Command, args []string) {
			stopShell(cfgPath)
			startShell(cfgPath)
		},
	}
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Config Status",
		Run: func(cmd *cobra.Command, args []string) {
			statusShell()
		},
	}
)

func init() {
	startCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "Config Path")
	startCmd.MarkFlagRequired("config")

	stopCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "Config Path")
	stopCmd.MarkFlagRequired("config")

	restartCmd.Flags().StringVarP(&cfgPath, "config", "c", "", "Config Path")
	restartCmd.MarkFlagRequired("config")

	mgrCmd.AddCommand(startCmd)
	mgrCmd.AddCommand(stopCmd)
	mgrCmd.AddCommand(restartCmd)
	mgrCmd.AddCommand(statusCmd)
}
