package cmd

import (
	"fmt"

	"wgm/services"

	"github.com/spf13/cobra"
)

var (
	ruleCmd = &cobra.Command{
		Use:   "rule",
		Short: "Manage Rule",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	ruleID     int
	ruleIP     string
	ruleAddCmd = &cobra.Command{
		Use:   "create",
		Short: "Create Rule",
		Run: func(cmd *cobra.Command, args []string) {
			code := services.CreateRule(ruleIP)
			fmt.Println(code)
		},
	}
	ruleUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update Rule",
		Run: func(cmd *cobra.Command, args []string) {
			code := services.UpdateRule(ruleID, ruleIP)
			fmt.Println(code)
		},
	}
	ruleDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete Rule",
		Run: func(cmd *cobra.Command, args []string) {
			code := services.DeleteRule(ruleID)
			fmt.Println(code)
		},
	}
)

func init() {
	ruleAddCmd.Flags().StringVar(&ruleIP, "ip", "", "AllowedIP eg: 192.168.1.1/24")
	ruleAddCmd.MarkFlagRequired("ip")

	ruleUpdateCmd.Flags().IntVar(&ruleID, "id", 0, "Rule ID")
	ruleUpdateCmd.Flags().StringVar(&ruleIP, "ip", "", "AllowedIP eg: 192.168.1.1/24")
	ruleUpdateCmd.MarkFlagRequired("id")
	ruleUpdateCmd.MarkFlagRequired("ip")

	ruleDelCmd.Flags().IntVar(&ruleID, "id", 0, "Rule ID")
	ruleDelCmd.MarkFlagRequired("id")

	ruleCmd.AddCommand(ruleAddCmd)
	ruleCmd.AddCommand(ruleUpdateCmd)
	ruleCmd.AddCommand(ruleDelCmd)
}
