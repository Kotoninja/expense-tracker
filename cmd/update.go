/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	storage "github.com/Kotoninja/expense-tracker/internal"
	"github.com/spf13/cobra"
	"time"
	"unicode/utf8"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		description, _ := cmd.Flags().GetString("description")
		amount, _ := cmd.Flags().GetInt("amount")
		date, _ := cmd.Flags().GetString("date")

		if description == "" && amount == 0 && date == "" {
			fmt.Println("Specify the fields to update")
			return
		}

		timestamp, err := time.Parse("2006-01-02", date)

		if (utf8.RuneCountInString(date) != 0) && (err != nil) {
			fmt.Println("Please enter the date in the format. (YYYY-MM-DD)")
		}

		if err := storage.StorageIO.Update(id, description, amount, timestamp); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().Int("id", 0, "Identifier of object")
	updateCmd.Flags().String("description", "", "Description of the item")
	updateCmd.Flags().Int("amount", 0, "The amount to be added")
	updateCmd.Flags().String("date", "", "The timestamp of the record (e.g., YYYY-MM-DD)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
