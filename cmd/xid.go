/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/rs/xid"
	"github.com/spf13/cobra"
)

// xidCmd represents the xid command
var xidCmd = &cobra.Command{
	Use:   "xid",
	Short: "gogenera xid",
	Run: func(cmd *cobra.Command, args []string) {
		num, err := cmd.Flags().GetInt64("number")
		if err != nil {
			log.Fatalln(err)
		}
		for i := 0; i < int(num); i++ {
			fmt.Println(xid.New().String())
		}
	},
}

func init() {
	rootCmd.AddCommand(xidCmd)
	xidCmd.Flags().Int64P("number", "n", 1, "numero de codigo")
}
