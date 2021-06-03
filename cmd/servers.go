package cmd

import (
	"fmt"

	"wgm/services"

	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Manage Server",
		Run:   func(cmd *cobra.Command, args []string) {
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
			info := map[string]interface{}{
				"title":       serverTitle,
				"address":     serverAddress,
				"port":        serverPort,
				"lan_ip":      serverLanIP,
				"lan_netmask": serverLanNetmask,
				"mtu":         serverMTU,
				"dns":         serverDNS,
			}
			code := services.CreateServer(info)
			fmt.Println(code)
		},
	}
	serverUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update Server",
		Run: func(cmd *cobra.Command, args []string) {
			info := make(map[string]interface{})
			info["title"] = serverTitle
			if serverAddress != "" {
				info["address"] = serverAddress
			}
			if serverPort != 0 {
				info["port"] = serverPort
			}
			if serverLanIP != "" {
				info["lan_ip"] = serverLanIP
			}
			if serverLanNetmask != "" {
				info["lan_netmask"] = serverLanNetmask
			}
			if serverMTU != "" {
				info["mtu"] = serverMTU
			}
			if serverDNS != "" {
				info["dns"] = serverDNS
			}
			code := services.UpdateServer(info)
			fmt.Println(code)
		},
	}
	serverDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete Server",
		Run: func(cmd *cobra.Command, args []string) {
			code := services.DeleteServer(serverTitle)
			fmt.Println(code)
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
