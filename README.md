# ATAD_PFCLIManager
**ATAD_PFCLIManager** is a powerful, terminal-based Personal Finance Manager built with **Go**. It is designed for developers and power users who want to track their income and expenses, set budgets, and visualize their financial health—all without leaving the command line.


## Key Features
* **Transaction Management**: Manually add, list, and search through your financial history.
* **Automated Categorization**: Uses **Regex matching** to automatically assign categories to transactions based on their description (e.g., "Lidl" -> "Food").
* **Bulk Import**: Quickly import data from **CSV** files.
* **Budgeting & Alerts**: Set monthly limits per category and receive visual alerts if you exceed them.
* **Visual Reports**: Generate insightful **ASCII bar charts** directly in your terminal to see where your money goes.
* **Interactive TUI**: A modern **Terminal User Interface** for browsing transactions interactively using arrow keys.
* **Local Storage**: All data is stored locally in a **SQLite** database for privacy and speed.

## Project Structure

The project follows a clean, modular architecture:

* `/cmd`: Entry point of the application containing `main.go`.
* `/internal/cli`: Cobra-based command definitions and input handling.
* `/internal/core`: Core business logic including Transaction and Budget services.
* `/internal/database`: SQLite initialization and schema management.
* `/pkg`: Reusable utilities like CSV parsers and TUI components.

## Getting Started

### Prerequisites
* **Go** (version 1.21+)
* **GCC/MinGW** (required for the SQLite driver)

### Installation
1. **Clone the repository**:
    ```bash
    git clone [https://github.com/YourUsername/ATAD_PFCLIManager.git](https:github.com/YourUsername/ATAD_PFCLIManager.git)
    cd ATAD_PFCLIManager
2. **Install dependencies**:
    ```bash
    go mod tidy
3. **Build the executable**:
    ```bash
    go build -o pfcli.exe ./cmd
## Usage guide
| Command | Action | Example |
| :--- | :--- | :--- |
| `add` |  Manually add a transaction | `./pfcli add -a -45.0 -d "Dinner" -c "Food"` |
| `list` |  View all transactions in a table | `./pfcli list` |
| `search` |  Filter transactions by keyword | `./pfcli search Lidl` |
| `import` |  Import transactions from a CSV file | `./pfcli import --file statement.csv` |
| `budget set` |  Set a spending limit for a category | `./pfcli budget set Food 600` |
| `budget check` |  Compare spending vs. budgets | `./pfcli budget check` |
| `report` |  Show visual category breakdown | `./pfcli report` |
| `browse` |  Open interactive TUI mode | `./pfcli browse` |

## CSV Format for Import
To ensure a successful import, your CSV file should follow this structure:

    date,description,amount
    2026-01-20,GROCERIES STORE,-120.00
    2026-01-21,MONTHLY SALARY,4500.00

## Build with
* Cobra - CLI framework
* Bubble Tea - TUI framework
* SQLite3 - Database engine
* Lip Gloss - Terminal styling