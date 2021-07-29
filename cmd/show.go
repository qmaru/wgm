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
	cfgType      int
	cfgServerID  int
	cfgUsername  string
	cfgMainnode  string
	cfgExtraRule bool
	cfgName      string
	cfgCmd       = &cobra.Command{
		Use:   "config",
		Short: "Show Config Details",
		Run: func(cmd *cobra.Command, args []string) {
			if TypeCheck(cfgType) {
				cfgInfo := services.ShowUserCfg(cfgType, cfgServerID, cfgUsername, cfgMainnode, cfgExtraRule, cfgName)
				fmt.Println(cfgInfo)
			} else {
				fmt.Println("Type Error")
			}
		},
	}
)

func init() {
	cfgCmd.Flags().IntVarP(&cfgType, "type", "t", 0, "Config Type: 1 or 2")
	cfgCmd.Flags().IntVarP(&cfgServerID, "server", "s", 0, "Server ID")
	cfgCmd.Flags().StringVarP(&cfgUsername, "user", "u", "", "User Name")
	cfgCmd.Flags().StringVarP(&cfgMainnode, "node", "m", "", "Type 1 User Name (client only)")
	cfgCmd.Flags().BoolVarP(&cfgExtraRule, "extra", "e", false, "Add Extra Rule")
	cfgCmd.Flags().StringVarP(&cfgName, "name", "n", "", "Config Name (Create configuration if specified)")
	cfgCmd.MarkFlagRequired("type")
	cfgCmd.MarkFlagRequired("server")
	cfgCmd.MarkFlagRequired("user")

	usersCmd.Flags().IntVarP(&showUserServerID, "server", "s", 0, "Server ID")
	usersCmd.Flags().BoolVarP(&showUserPlain, "plain", "p", false, "Plain Text")

	showCmd.AddCommand(usersCmd)
	showCmd.AddCommand(serversCmd)
	showCmd.AddCommand(rulesCmd)
	showCmd.AddCommand(cfgCmd)
}

func TypeCheck(t int) bool {
	for _, i := range []int{1, 2} {
		if t == i {
			return true
		}
	}
	return false
}
