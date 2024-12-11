package services

import (
	"personal-budget/models"
	"testing"
)

func TestAddToBudget(t *testing.T) {
	budgets = []models.Budget{
		{ID: 1, Amount: 1000.0, Name: "Groceries"},
		{ID: 2, Amount: 500.0, Name: "Entertainment"},
	}

	err := AddToBudget(1, 500.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if budgets[0].Amount != 1500.0 {
		t.Errorf("Expected amount 1500.0, got %v", budgets[0].Amount)
	}

	err = AddToBudget(3, 500.0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestSpendBudget(t *testing.T) {
	budgets = []models.Budget{
		{ID: 1, Amount: 1000.0, Name: "Groceries"},
		{ID: 2, Amount: 500.0, Name: "Entertainment"},
	}

	err := SpendBudget(1, 500.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if budgets[0].Amount != 500.0 {
		t.Errorf("Expected amount 500.0, got %v", budgets[0].Amount)
	}

	err = SpendBudget(1, 1500.0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	err = SpendBudget(3, 500.0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGetBudget(t *testing.T) {
	budgets = []models.Budget{
		{ID: 1, Amount: 1000.0, Name: "Groceries"},
		{ID: 2, Amount: 500.0, Name: "Entertainment"},
	}

	b, err := GetBudget(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if b.ID != 1 {
		t.Errorf("Expected ID 1, got %v", b.ID)
	}

	b, err = GetBudget(3)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGetAllBudgets(t *testing.T) {
	budgets = []models.Budget{
		{ID: 1, Amount: 1000.0, Name: "Groceries"},
		{ID: 2, Amount: 500.0, Name: "Entertainment"},
	}

	b := GetAllBudgets()
	if len(b) != 2 {
		t.Errorf("Expected 2 budgets, got %v", len(b))
	}
}
