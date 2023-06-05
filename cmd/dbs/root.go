package dbs

import (
	"github.com/spf13/cobra"
)

var (
	DBCmd = &cobra.Command{
		Use:   "db",
		Short: "Database Manger",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)
