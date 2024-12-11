package services

import (
	"context"
	"github.com/jackc/pgx/v5"
	"personal-budget/models"
)

// helper function, error check

func errCheck(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// return all budgets
func GetAllBudgetsDb(ctx context.Context, conn *pgx.Conn) ([]models.Budget, error) {
	query := `SELECT id, name, amount FROM personal_budget`

	rows, err := conn.Query(ctx, query)
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
