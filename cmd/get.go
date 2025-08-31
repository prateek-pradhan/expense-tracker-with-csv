package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getAllExpenseCmd = &cobra.Command{
	Use: "list",
	Short: "List All Expenses",
	Long: "List All Expenses from Expense Tracker",
	Run: func(cmd *cobra.Command, args []string) {
		listExpenses()
	},	
}

var summaryCmd = &cobra.Command{
	Use: "summary",
	Short: "Get Expense Summary",
	Long: "Get Expense Summary from Expense Tracker",
	Run: func(cmd *cobra.Command, args []string) {
		month , err := cmd.Flags().GetString("month")
		if err != nil {
			fmt.Println(err)
			return
		}

		getExpenseSummary(month)
	},	
}

func init() {
	summaryCmd.Flags().String("month", "0", "month for which summary is required")
	rootCmd.AddCommand(getAllExpenseCmd, summaryCmd)
}
