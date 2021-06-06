package cmd

import (
	"fmt"

	"wgm/services"

	"github.com/spf13/cobra"
)

var (
	userCmd = &cobra.Command{
		Use:   "user",
		Short: "Manage User",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	userServerID  int
	userUsername  string
	userIP        string
	userExtra     bool
	userKeepalive string
	userAddCmd    = &cobra.Command{
		Use:   "create",
		Short: "Create User",
		Run: func(cmd *cobra.Command, args []string) {
			isExtra := 0
			if userExtra {
				isExtra = 1
			}
			info := map[string]interface{}{
				"server_id": userServerID,
				"username":  userUsername,
				"ip":        userIP,
				"is_extra":  isExtra,
				"keepalive": userKeepalive,
			}
			code := services.CreateUser(info)
			fmt.Println(code)
		},
	}
	userUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update User",
		Run: func(cmd *cobra.Command, args []string) {
			info := map[string]interface{}{
				"server_id": userServerID,
				"username":  userUsername,
			}
			if userIP != "" {
				info["ip"] = userIP
			}
			if userKeepalive != "" {
				info["keepalive"] = userKeepalive
			}
			isExtra := 0
			if userExtra {
				isExtra = 1
			}
			info["is_extra"] = isExtra
			code := services.UpdateUser(info)
			fmt.Println(code)
		},
	}
	userKeyUpdateCmd = &cobra.Command{
		Use:   "key",
		Short: "Update User key",
		Run: func(cmd *cobra.Command, args []string) {
			code := services.UpdateUserKey(userServerID, userUsername)
			fmt.Println(code)
		},
	}
	userDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete User",
		Run: func(cmd *cobra.Command, args []string) {
			code := services.DeleteUser(userServerID, userUsername)
			fmt.Println(code)
		},
	}
)

func init() {
	userAddCmd.Flags().IntVarP(&userServerID, "server_id", "s", 0, "Server ID")
	userAddCmd.Flags().StringVarP(&userUsername, "user", "u", "", "User Name")
	userAddCmd.Flags().StringVarP(&userIP, "ip", "i", "", "User IP")
	userAddCmd.Flags().BoolVarP(&userExtra, "extra", "e", false, "Open Extra Rule")
	userAddCmd.Flags().StringVarP(&userKeepalive, "keepalive", "k", "25", "User Keepalive")

	userAddCmd.MarkFlagRequired("server_id")
	userAddCmd.MarkFlagRequired("user")
	userAddCmd.MarkFlagRequired("ip")

	userUpdateCmd.Flags().IntVarP(&userServerID, "server_id", "s", 0, "Server ID")
	userUpdateCmd.Flags().StringVarP(&userUsername, "user", "u", "", "User Name")
	userUpdateCmd.Flags().StringVarP(&userIP, "ip", "i", "", "User IP")
	userUpdateCmd.Flags().BoolVarP(&userExtra, "extra", "e", false, "Open Extra Rule")
	userUpdateCmd.Flags().StringVarP(&userKeepalive, "keepalive", "k", "25", "User Keepalive")
	userUpdateCmd.MarkFlagRequired("server_id")
	userUpdateCmd.MarkFlagRequired("user")

	userKeyUpdateCmd.Flags().IntVarP(&userServerID, "server_id", "s", 0, "Server ID")
	userKeyUpdateCmd.Flags().StringVarP(&userUsername, "user", "u", "", "User Name")
	userKeyUpdateCmd.MarkFlagRequired("server_id")
	userKeyUpdateCmd.MarkFlagRequired("user")

	userDelCmd.Flags().IntVarP(&userServerID, "server_id", "s", 0, "Server ID")
	userDelCmd.Flags().StringVarP(&userUsername, "user", "u", "", "User Name")
	userDelCmd.MarkFlagRequired("server_id")
	userDelCmd.MarkFlagRequired("user")

	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userUpdateCmd)
	userCmd.AddCommand(userDelCmd)

	userUpdateCmd.AddCommand(userKeyUpdateCmd)
}
