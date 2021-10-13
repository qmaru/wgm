package cmd

import (
	"fmt"

	"wgm/services"

	"github.com/spf13/cobra"
)

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show configuration details or generate configuration.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	showUserServerID int
	showUserPlain    bool
	usersCmd         = &cobra.Command{
		Use:   "users",
		Short: "Show User List",
		Run: func(cmd *cobra.Command, args []string) {
			services.ShowUserList(showUserServerID, showUserPlain)
		},
	}
	serversCmd = &cobra.Command{
		Use:   "servers",
		Short: "Show Servers List",
		Run: func(cmd *cobra.Command, args []string) {
			services.ShowServerList()
		},
	}
	rulesCmd = &cobra.Command{
		Use:   "rules",
		Short: "Show Rules List",
		Run: func(cmd *cobra.Command, args []string) {
			services.ShowRuleList()
		},
	}
	cfgType        string
	cfgServerTitle string
	cfgNodename    string
	cfgExtraRule   bool
	cfgName        string
	cfgCmd         = &cobra.Command{
		Use:   "config",
		Short: "Show Config Details",
		Run: func(cmd *cobra.Command, args []string) {
			if TypeCheck(cfgType) {
				cfgInfo := services.ShowUserCfg(cfgType, cfgServerTitle, cfgNodename, cfgExtraRule, cfgName)
				fmt.Println(cfgInfo)
			} else {
				fmt.Println("Type Error")
			}
		},
	}
)

func init() {
	cfgCmd.Flags().StringVarP(&cfgType, "type", "t", "", "Config Type: s(erver) or c(lient)")
	cfgCmd.Flags().StringVarP(&cfgServerTitle, "server", "s", "", "Public Server Name")
	cfgCmd.Flags().StringVarP(&cfgNodename, "node", "c", "", "Node Name")
	cfgCmd.Flags().BoolVarP(&cfgExtraRule, "extra", "e", false, "Add Extra Rule")
	cfgCmd.Flags().StringVarP(&cfgName, "name", "", "", "Config Name (Create configuration if specified)")
	cfgCmd.MarkFlagRequired("type")
	cfgCmd.MarkFlagRequired("server")
	cfgCmd.MarkFlagRequired("node")

	usersCmd.Flags().IntVarP(&showUserServerID, "server", "s", 0, "Server ID")
	usersCmd.Flags().BoolVarP(&showUserPlain, "plain", "p", false, "Plain Text")

	showCmd.AddCommand(usersCmd)
	showCmd.AddCommand(serversCmd)
	showCmd.AddCommand(rulesCmd)
	showCmd.AddCommand(cfgCmd)
}

func TypeCheck(t string) bool {
	for _, i := range []string{"server", "client"} {
		if t == i {
			return true
		}
	}
	return false
}
