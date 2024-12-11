package services

import (
	"errors"
	"personal-budget/models"
)

var budgets = []models.Budget{
	{ID: 1, Amount: 2743, Name: "Bob Johnson"},
	{ID: 2, Amount: 584, Name: "John Doe"},
	{ID: 3, Amount: 3469, Name: "Jane Doe"},
	{ID: 4, Amount: 2200, Name: "Eve Adams"},
	{ID: 5, Amount: 782, Name: "Diana Prince"},
	{ID: 6, Amount: 2094, Name: "Frank Miller"},
	{ID: 7, Amount: 2365, Name: "Diana Prince"},
	{ID: 8, Amount: 4015, Name: "Grace Hopper"},
	{ID: 9, Amount: 4309, Name: "Alice Smith"},
	{ID: 10, Amount: 1328, Name: "Alice Smith"},
}

// helper function to check if ID is valid
func idCheck(id int) error {

	if id == 0 {
		return errors.New("ID is required")
	}

	if id < 0 {
		return errors.New("ID cannot be negative")
	}

	return nil
}

// GetAllBudgets returns all budgets

func GetAllBudgets() []models.Budget {
	return budgets
}

// GetBudget returns a budget by ID

func GetBudget(id int) (models.Budget, error) {
	idCheck(id)

	for _, b := range budgets {
		if b.ID == id {
			return b, nil
		}
	}

	return models.Budget{}, errors.New("Budget not found")
}

// AddToBudget adds an amount to a budget

func AddToBudget(id int, amount float64) (bool, error) {
	idCheck(id)

	for i, b := range budgets {
		if b.ID == id {
			budgets[i].Amount += amount
			return true, nil
		}
	}

	return false, errors.New("Budget not found")
}

// SpendBudget subtracts an amount from a budget

func SpendBudget(id int, amount float64) (bool, error) {
	idCheck(id)

	for i, b := range budgets {
		if b.ID == id {
			if b.Amount < amount {
				return false, errors.New("Insufficient funds")
			}
			budgets[i].Amount -= amount
			return true, nil
		}
	}

	return false, errors.New("Budget not found")
}
