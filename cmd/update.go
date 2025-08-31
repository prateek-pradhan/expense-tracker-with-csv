package cmd

import (
	"github.com/spf13/cobra"
)

var updateExpenseCmd = &cobra.Command{
	Use: "update",
	Short: "Update Expense",
	Long: "Update Expense from Expense Tracker",
	RunE: func(cmd *cobra.Command, args []string) error{
		id, err := cmd.Flags().GetString("id")
		if err != nil { 
			return err 
		}
		description, err := cmd.Flags().GetString("description")
		amount, e := cmd.Flags().GetString("amount")
		if err != nil { 
			return err 
		}

		if e != nil { 
			return e
		}

		updateExpense(id, description, amount)
		return nil
	},	
}

func init() {
	updateExpenseCmd.Flags().String("id", "", "expense Id")
	updateExpenseCmd.Flags().String("description", "", "expense description")
	updateExpenseCmd.Flags().String("amount", "", "expense amount")
	rootCmd.AddCommand(updateExpenseCmd)
}
