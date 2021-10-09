package cmd

import (
	"fmt"

	"wgm/models"
	"wgm/services"

	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Manage Server",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	serverTitle      string
	serverAddress    string
	serverPort       int
	serverLanIP      string
	serverLanNetmask string
	serverMTU        string
	serverDNS        string
	serverAddCmd     = &cobra.Command{
		Use:   "create",
		Short: "Create Server",
		Run: func(cmd *cobra.Command, args []string) {
			info := models.Servers{
				Title:      serverTitle,
				Address:    serverAddress,
				Port:       serverPort,
				LanIP:      serverLanIP,
				LanNetmask: serverLanNetmask,
				MTU:        serverMTU,
				DNS:        serverDNS,
			}
			code := services.CreateServer(info)
			fmt.Println(code.String())
		},
	}
	serverUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update Server",
		Run: func(cmd *cobra.Command, args []string) {
			info := models.Servers{
				Title:      serverTitle,
				Address:    serverAddress,
				Port:       serverPort,
				LanIP:      serverLanIP,
				LanNetmask: serverLanNetmask,
				MTU:        serverMTU,
				DNS:        serverDNS,
			}
			serverID := services.GetServerID(serverTitle)
			code := services.UpdateServer(serverID, info)
			fmt.Println(code.String())
		},
	}
	serverDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete Server",
		Run: func(cmd *cobra.Command, args []string) {
			serverID := services.GetServerID(serverTitle)
			code := services.DeleteServer(serverID)
			fmt.Println(code.String())
		},
	}
)

func init() {
	serverAddCmd.Flags().StringVar(&serverTitle, "title", "", "Server Title")
	serverAddCmd.Flags().StringVar(&serverAddress, "address", "", "Server Address")
	serverAddCmd.Flags().IntVar(&serverPort, "port", 443, "Server Port")
	serverAddCmd.Flags().StringVar(&serverLanIP, "ip", "", "Server Lan IP")
	serverAddCmd.Flags().StringVar(&serverLanNetmask, "netmask", "", "Server Lan Netmask")
	serverAddCmd.Flags().StringVar(&serverMTU, "mtu", "", "Server MTU")
	serverAddCmd.Flags().StringVar(&serverDNS, "dns", "", "Server DNS")
	serverAddCmd.MarkFlagRequired("title")
	serverAddCmd.MarkFlagRequired("address")
	serverAddCmd.MarkFlagRequired("ip")
	serverAddCmd.MarkFlagRequired("netmask")

	serverUpdateCmd.Flags().StringVar(&serverTitle, "title", "", "Server Title")
	serverUpdateCmd.Flags().StringVar(&serverAddress, "address", "", "Server Address")
	serverUpdateCmd.Flags().IntVar(&serverPort, "port", 0, "Server Port")
	serverUpdateCmd.Flags().StringVar(&serverLanIP, "ip", "", "Server Lan IP")
	serverUpdateCmd.Flags().StringVar(&serverLanNetmask, "netmask", "", "Server Lan Netmask")
	serverUpdateCmd.Flags().StringVar(&serverMTU, "mtu", "", "Server MTU")
	serverUpdateCmd.Flags().StringVar(&serverDNS, "dns", "", "Server DNS")
	serverUpdateCmd.MarkFlagRequired("title")

	serverDelCmd.Flags().StringVar(&serverTitle, "title", "", "Server Title")
	serverDelCmd.MarkFlagRequired("title")

	serverCmd.AddCommand(serverAddCmd)
	serverCmd.AddCommand(serverUpdateCmd)
	serverCmd.AddCommand(serverDelCmd)
}
