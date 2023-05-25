package dbs

import (
	"log"

	"wgm/dbs"

	"github.com/spf13/cobra"
)

var (
	dbInit = &cobra.Command{
		Use:   "init",
		Short: "Init tables",
		Run: func(cmd *cobra.Command, args []string) {
			err := dbs.CreataTable()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)
