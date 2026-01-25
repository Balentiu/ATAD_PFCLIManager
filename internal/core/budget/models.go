package budget

type Budget struct {
	Category string
	Amount   float64
}

type BudgetReport struct {
	Category string
	Budget   float64
	Spent    float64
	Status   string
}
