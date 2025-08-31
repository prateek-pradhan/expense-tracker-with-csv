package cmd

import (
	"github.com/spf13/cobra"
)

var addExpenseCmd = &cobra.Command{
	Use: "add",
	Short: "Add Expense ",
	Long: "Add Expense to Expense Tracker",
	RunE: func(cmd *cobra.Command, args []string) error{
		description, err := cmd.Flags().GetString("description")
		amount, e := cmd.Flags().GetString("amount")
		if err != nil { 
			return err 
		}

		if e != nil { 
			return e
		}

		addExpense(description, amount)
		return nil
	},	
}

func init() {
	addExpenseCmd.Flags().String("description", "", "expense description")
	addExpenseCmd.Flags().String("amount", "0", "expense amount")
	rootCmd.AddCommand(addExpenseCmd)
}