# 📝 Expense Tracker with CSV

A simple **expense tracker** built in **Go**, powered by [Cobra](https://github.com/spf13/cobra).  
It allows you to add, update, delete, and summarize expense, persisting them in a local `expenses.csv` file.  

---

## 🚀 Features
- Add, update, delete expense
- List expenses
- Summarize total or a certain month's (Current Year) expenses 
- Stores expenses in `expenses.csv`

---

## ⚙️ Installation

```bash
git clone https://github.com/prateek-pradhan/expense-tracker-with-csv.git
cd expense-tracker-with-csv
go build -o expense-tracker-with-csv main.go
mv expense-tracker-with-csv /usr/local/bin/   # optional
```

---

## 📌 Usage

```bash
expense-tracker-with-csv [command]
```

### Commands

- **Add**: `expense-tracker-with-csv add --description "AWS" --amount 30000 `  
- **Update description**: `expense-tracker-with-csv update --id 1 --description "Buy vegetables"`  
- **Update amount**: `expense-tracker-with-csv update --id 1 --amount 200"`  
- **Update description and amount**: `expense-tracker-with-csv update --id 1 --amount "Buy vegetables" --amount 300`  
- **Delete**: `expense-tracker-with-csv delete --id 1`  
- **Summary**: `expense-tracker-with-csv summary`  
- **Summary by month**: `expense-tracker-with-csv summary --month 8`  
- **List**: `expense-tracker-with-csv list`

---

## 🛠 Notes
- Uses Go’s `encoding/csv` for reading and writing in csv  
- Auto-creates `tasks.json` if missing  
- Extendable with Cobra for more commands  

📖 This project is inspired by the [Expense Tracker project on roadmap.sh](https://roadmap.sh/projects/expense-tracker).

