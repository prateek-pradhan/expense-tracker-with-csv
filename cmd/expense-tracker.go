package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	id int
	date time.Time
	description string
	amount int
}

func addExpense(description string, amount string) error {

	cost, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println(err)
	}

	expenses, err := readCSV()
	if err != nil {
		return err
	}

	newId := 0
	for expense := range expenses {
		if expenses[expense].id > newId {
			newId = expenses[expense].id
		}
	}


	newExpense := Expense{
		id: newId + 1,
		date: time.Now(),
		description: description,
		amount: int(cost),
	}

	expenses = append(expenses, newExpense)
	
	err = writeCSV(expenses)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func updateExpense(id string, description string, amount string) {
	
	if id == "" {
		fmt.Println("Invalid ID")
		return
	}
	expenseId, err := strconv.ParseInt(id, 10, 64)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := readCSV()
	if err != nil {
		fmt.Println(err)
		return
	}

	if description == "" && amount == "0" {
		fmt.Println("Nothing to update")
		return
	}

	if description == "" {
		description = expenses[expenseId-1].description
	}

	var cost float64
	if amount == "" {
		cost = float64(expenses[expenseId-1].amount)
	} else {
		cost, err = strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	

	var found = false
	for i := 0; i < len(expenses); i++ {
		if expenses[i].id == int(expenseId) {
			expenses[i].description = description
			expenses[i].amount = int(cost)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Invalid Id")
		return
	}

	err = writeCSV(expenses)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteExpense (id string) {

	if id == "" {
		fmt.Println("Invalid ID")
		return
	}

	expenseId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := readCSV()

	if err != nil {
		fmt.Println(err)
		return
	}

	var found = false
	for i := 0; i < len(expenses); i++ {
		if expenses[i].id == int(expenseId) {
			expenses = append(expenses[:i], expenses[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Invalid Id")
		return
	}

	err = writeCSV(expenses)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func listExpenses() {
	expenses, err := readCSV()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}
	
	fmt.Printf("%-5s %-25s %-50s %-10s\n", "ID", "Date", "Description", "Amount")
	for _, expense := range expenses {
		fmt.Printf("%-5d %-25s %-50s %-10d\n", expense.id, expense.date.Format("2006-01-02 15:04:05"), expense.description, expense.amount)
	}
}

func getExpenseSummary(month string) {

	monthInt, err := strconv.ParseInt(month, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := readCSV()
	if err != nil {
		fmt.Println(err)
		return
	}

	if monthInt < 0 || monthInt > 12 {
		fmt.Println("Invalid month")
		return
	}

	totalExpense := 0

	for i := range expenses {
		if monthInt == 0 || int64(expenses[i].date.Month()) == monthInt {
			totalExpense += expenses[i].amount
		}
	}

	if monthInt == 0 {
		fmt.Printf("Total Expense for all months is: %d\n", totalExpense)
		return
	} else {
	fmt.Printf("Total Expense for month %d is: %d\n", monthInt, totalExpense)
	}

}

func readCSV() ([]Expense, error) {
	file, err := os.Open("expenses.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var expenses []Expense
	
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		id, _ := strconv.Atoi(record[0])
		date, _ := time.Parse(time.RFC3339, record[1])
		description := record[2]
		amount, _ := strconv.Atoi(record[3])

		expense := Expense{
			id: id,
			date: date,
			description: description,
			amount: amount,
		}

		expenses = append(expenses, expense)
	}

	fmt.Println(expenses)
	return expenses, nil
}

func writeCSV(expenses []Expense) error {
	file, err := os.Create("expenses.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, expense := range expenses {
		record := []string{
			strconv.Itoa(expense.id),
			expense.date.Format(time.RFC3339),
			expense.description,
			strconv.Itoa(expense.amount),
		}
		if err := writer.Write(record); err != nil {
			return err
		}		
	}
	
	return nil
}