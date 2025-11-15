# ATAD_PFCLIManager
ATAD_PFCLIManager (Personal Finance CLI Manager) is a command-line interface (CLI) tool built in Go, designed for managing your income and expenses directly from the terminal.
This project is built using a clean architecture (separating business logic from the CLI) and stores all data locally in a CGO-free SQLite database.


# Features (Implemented)
Add Transactions: Manually add income (positive amounts) or expenses (negative amounts) using the add command.

List Transactions: Display all transactions in a clean, colorful, date-sorted table using the list command.

Local Storage: All data is saved in a local finance.db file.
