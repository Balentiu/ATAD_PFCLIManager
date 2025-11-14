package transaction

import "time"

// Transaction reprezintă o singură tranzacție financiară
type Transaction struct {
    ID          int
    Date        time.Time
    Description string
    Amount      float64
    Category    string
}