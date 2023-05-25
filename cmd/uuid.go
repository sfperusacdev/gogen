/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "gogenera uuid",
	Run: func(cmd *cobra.Command, args []string) {
		num, err := cmd.Flags().GetInt64("number")
		if err != nil {
			log.Fatalln(err)
		}
		for i := 0; i < int(num); i++ {
			uuid, err := uuid.NewV4()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(uuid.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	uuidCmd.Flags().Int64P("number", "n", 1, "numero de codigo")
}
