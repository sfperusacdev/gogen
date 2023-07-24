package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

// bcryptCmd represents the uuid command
var bcryptCmd = &cobra.Command{
	Use:   "bcrypt",
	Short: "hash password with bcrypt",
	Run: func(cmd *cobra.Command, args []string) {
		num, err := cmd.Flags().GetInt("cost")
		if err != nil {
			log.Fatalln(err)
		}
		for _, v := range args {
			bytes, err := bcrypt.GenerateFromPassword([]byte(v), num)
			if err != nil {
				log.Printf("err [`%s`]: %v\n", v, err)
			}
			fmt.Printf("[`%s`]: %s\n", v, string(bytes))
		}
	},
}

func init() {
	rootCmd.AddCommand(bcryptCmd)
	bcryptCmd.Flags().IntP("cost", "c", 12, "a numeric cost")
}
