package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"personal-budget/models"
)

const insertBudgetQuery = `INSERT INTO personal_budget (name, amount) VALUES ($1, $2)`
const getBudgetQuery = `SELECT id, name, amount FROM personal_budget`
const getSingleBudgetQuery = `SELECT id, name, amount FROM personal_budget WHERE id = $1`
const addToBudgetQuery = `UPDATE personal_budget SET amount = amount + $1 WHERE id = $2`
const spendBudgetQuery = `UPDATE personal_budget SET amount = amount - $1 WHERE id = $2`

func errCheck(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// validate input
func validateBudget(budget models.Budget) error {
	var err error

	if budget.Name == "" {
		err = fmt.Errorf("budget name is required")
		return err
	}
	if budget.Amount <= 0 {
		err = fmt.Errorf("budget amount must be greater than 0")
		return err

	}
	return nil

}

// GetAllBudgetsDb retrieves all budgets from the database.
// It executes a query to fetch all rows from the personal_budget table
// and returns a slice of Budget models or an error if the operation fails.
func GetAllBudgetsDb(ctx context.Context, conn *pgx.Conn) ([]models.Budget, error) {

	rows, err := conn.Query(ctx, getBudgetQuery)
	errCheck(err)
	defer rows.Close()

	var budgets []models.Budget

	for rows.Next() {
		var budget models.Budget
		err = rows.Scan(&budget.ID, &budget.Name, &budget.Amount)
		errCheck(err)
		budgets = append(budgets, budget)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return budgets, nil

}

// GetSingleBudgetDb retrieves a single budget from the database by ID.
func GetSingleBudgetDb(ctx context.Context, conn *pgx.Conn, id int) (models.Budget, error) {
	log.Printf("Fetching budget with ID: %d\n", id)

	var budget models.Budget
	err := conn.QueryRow(ctx, getSingleBudgetQuery, id).Scan(&budget.ID, &budget.Name, &budget.Amount)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("Budget with ID %d not found\n", id)
			return models.Budget{}, fmt.Errorf("budget with ID %d not found: %w", id, err)
		}
		log.Printf("Error fetching budget with ID %d: %v\n", id, err)
		return models.Budget{}, fmt.Errorf("failed to fetch budget: %w", err)
	}

	log.Printf("Fetched budget: %+v\n", budget)
	return budget, nil
}

// PostBudgetDb inserts a new budget into the database.
func PostBudgetDb(ctx context.Context, conn *pgx.Conn, budget models.Budget) error {
	if err := validateBudget(budget); err != nil {
		return fmt.Errorf("invalid budget data: %w", err)
	}

	log.Printf("Inserting budget: Name=%s, Amount=%.2f\n", budget.Name, budget.Amount)

	_, err := conn.Exec(ctx, insertBudgetQuery, budget.Name, budget.Amount)
	if err != nil {
		return fmt.Errorf("failed to insert budget: %w", err)
	}

	log.Println("Budget inserted successfully!")
	return nil
}

// AddToBudgetDb adds an amount to an existing budget.
func AddToBudgetDb(ctx context.Context, conn *pgx.Conn, id int, amount float64) error {
	if id <= 0 {
		return fmt.Errorf("invalid budget ID: %d", id)
	}
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %.2f", amount)
	}

	log.Printf("Adding %.2f to budget with ID: %d\n", amount, id)

	_, err := conn.Exec(ctx, addToBudgetQuery, amount, id)
	if err != nil {
		log.Printf("Error updating budget: %v\n", err)
		return fmt.Errorf("failed to add to budget: %w", err)
	}

	log.Println("Added to budget successfully!")
	return nil
}

// SpendBudgetDb subtracts an amount from an existing budget.
func SpendBudgetDb(ctx context.Context, conn *pgx.Conn, id int, amount float64) error {
	if id <= 0 {
		return fmt.Errorf("invalid budget ID: %d", id)
	}
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %.2f", amount)
	}

	log.Printf("Spending %.2f from budget with ID: %d\n", amount, id)

	_, err := conn.Exec(ctx, spendBudgetQuery, amount, id)
	if err != nil {
		log.Printf("Error updating budget: %v\n", err)
		return fmt.Errorf("failed to spend from budget: %w", err)
	}

	log.Println("Spent from budget successfully!")
	return nil

}
