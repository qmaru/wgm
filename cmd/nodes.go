package cmd

import (
	"fmt"

	"wgm/models"
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
	userServerTitle string
	userUsername    string
	userIP          string
	userAccess      bool
	userKeepalive   int
	userAddCmd      = &cobra.Command{
		Use:   "create",
		Short: "Create User",
		Run: func(cmd *cobra.Command, args []string) {
			isAccess := 0
			if userAccess {
				isAccess = 1
			}

			userServerID := services.GetServerID(userServerTitle)
			info := models.Users{
				ServerID:            userServerID,
				Username:            userUsername,
				IP:                  userIP,
				IsAccess:            isAccess,
				PersistentKeepalive: userKeepalive,
			}

			code := services.CreateUser(info)
			fmt.Println(code)
		},
	}
	userUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update User",
		Run: func(cmd *cobra.Command, args []string) {
			userID := services.GetUserID(userUsername)
			userServerID := services.GetServerID(userServerTitle)

			isAccess := 0
			if userAccess {
				isAccess = 1
			}

			info := models.Users{
				ServerID:            userServerID,
				Username:            userUsername,
				IP:                  userIP,
				IsAccess:            isAccess,
				PersistentKeepalive: userKeepalive,
			}

			code := services.UpdateUser(userServerID, userID, info)
			fmt.Println(code)
		},
	}
	userKeyUpdateCmd = &cobra.Command{
		Use:   "key",
		Short: "Update User key",
		Run: func(cmd *cobra.Command, args []string) {
			userID := services.GetUserID(userUsername)
			userServerID := services.GetServerID(userServerTitle)
			code := services.UpdateUserKey(userServerID, userID)
			fmt.Println(code)
		},
	}
	userDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete User",
		Run: func(cmd *cobra.Command, args []string) {
			userID := services.GetUserID(userUsername)
			userServerID := services.GetServerID(userServerTitle)
			code := services.DeleteUser(userServerID, userID)
			fmt.Println(code)
		},
	}
)

func init() {
	userAddCmd.Flags().StringVarP(&userServerTitle, "server", "s", "", "Server Title")
	userAddCmd.Flags().StringVarP(&userUsername, "user", "u", "", "User Name")
	userAddCmd.Flags().StringVarP(&userIP, "ip", "i", "", "User IP")
	userAddCmd.Flags().BoolVarP(&userAccess, "access", "a", false, "Access Lan")
	userAddCmd.Flags().IntVarP(&userKeepalive, "keepalive", "k", 25, "User Keepalive")

	userAddCmd.MarkFlagRequired("server")
	userAddCmd.MarkFlagRequired("user")
	userAddCmd.MarkFlagRequired("ip")

	userUpdateCmd.Flags().StringVarP(&userServerTitle, "server", "s", "", "Server Title")
	userUpdateCmd.Flags().StringVarP(&userUsername, "user", "u", "", "User Name")
	userUpdateCmd.Flags().StringVarP(&userIP, "ip", "i", "", "User IP")
	userUpdateCmd.Flags().BoolVarP(&userAccess, "access", "a", false, "Access Lan")
	userUpdateCmd.Flags().IntVarP(&userKeepalive, "keepalive", "k", 25, "User Keepalive")
	userUpdateCmd.MarkFlagRequired("server")
	userUpdateCmd.MarkFlagRequired("user")

	userKeyUpdateCmd.Flags().StringVarP(&userServerTitle, "server", "s", "", "Server Title")
	userKeyUpdateCmd.Flags().StringVarP(&userUsername, "user", "u", "", "User Name")
	userKeyUpdateCmd.MarkFlagRequired("server")
	userKeyUpdateCmd.MarkFlagRequired("user")

	userDelCmd.Flags().StringVarP(&userServerTitle, "server", "s", "", "Server Title")
	userDelCmd.Flags().StringVarP(&userUsername, "user", "u", "", "User Name")
	userDelCmd.MarkFlagRequired("server")
	userDelCmd.MarkFlagRequired("user")

	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userUpdateCmd)
	userCmd.AddCommand(userDelCmd)

	userUpdateCmd.AddCommand(userKeyUpdateCmd)
}
