package cmd

import (
	"fmt"

	"wgm/models"

	"github.com/spf13/cobra"
)

var (
	dbCmd = &cobra.Command{
		Use:   "init",
		Short: "Create Database",
		Run: func(cmd *cobra.Command, args []string) {
			models.InitTable()
			fmt.Println("Database created successfully.")
		},
	}
)
