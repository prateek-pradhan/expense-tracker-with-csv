package cmd

import (
	"github.com/spf13/cobra"
)

var deleteExpenseCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete Expense",
	Long: "Delete Expense from Expense Tracker",
	RunE: func(cmd *cobra.Command, args []string) error{
		id, err := cmd.Flags().GetString("id")
		if err != nil { 
			return err 
		}

		deleteExpense(id)
		return nil
	},	
}

func init() {
	deleteExpenseCmd.Flags().String("id", "", "expense Id")
	rootCmd.AddCommand(deleteExpenseCmd)
}
