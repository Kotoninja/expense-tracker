/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	storage "github.com/Kotoninja/expense-tracker/internal"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var expenseList [][]string

		expenseList = storage.StorageIO.List()

		if len(expenseList) == 0 {
			fmt.Println("There are no expenses")
		} else {
			tabwriterData := [][]string{{"ID", "Description", "Amount", "Date"}}
			tabwriterData = append(tabwriterData, expenseList...)
			table := tablewriter.NewWriter(os.Stdout)
			table.Header(tabwriterData[0])
			table.Bulk(tabwriterData[1:])
			table.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
