package cmd

import (
	"fmt"

	"wgm/services"

	"github.com/spf13/cobra"
)

var (
	rulemapCmd = &cobra.Command{
		Use:   "rulemap",
		Short: "Manage User's rule",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	rmUserID      int
	rmRuleID      int
	rulemapAddCmd = &cobra.Command{
		Use:   "create",
		Short: "Create Rule For User",
		Run: func(cmd *cobra.Command, args []string) {
			code := services.CreateUserRule(rmUserID, rmRuleID)
			fmt.Println(code)
		},
	}
	rulemapDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete Rule For User",
		Run: func(cmd *cobra.Command, args []string) {
			code := services.DeleteUserRule(rmRuleID, rmUserID)
			fmt.Println(code)
		},
	}
)

func init() {
	rulemapAddCmd.Flags().IntVarP(&rmUserID, "user_id", "u", 0, "User ID")
	rulemapAddCmd.Flags().IntVarP(&rmRuleID, "rule_id", "r", 0, "Rule ID")
	rulemapAddCmd.MarkFlagRequired("user_id")
	rulemapAddCmd.MarkFlagRequired("rule_id")

	rulemapDelCmd.Flags().IntVarP(&rmUserID, "user_id", "u", 0, "User ID")
	rulemapDelCmd.Flags().IntVarP(&rmRuleID, "rule_id", "r", 0, "Rule ID")
	rulemapDelCmd.MarkFlagRequired("user_id")
	rulemapDelCmd.MarkFlagRequired("rule_id")

	rulemapCmd.AddCommand(rulemapAddCmd)
	rulemapCmd.AddCommand(rulemapDelCmd)
}
