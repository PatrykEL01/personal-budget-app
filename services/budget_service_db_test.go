/*
Package services contains unit tests for the budget service functions that interact with the database.

Tests included:
- TestDbConnect: Tests the database connection function.
- TestGetAllBudgetsDb: Tests the function to retrieve all budgets from the database.
- TestGetSingleBudgetDb: Tests the function to retrieve a single budget by ID from the database.
- TestPostBudgetDb: Tests the function to insert a new budget into the database.
- TestAddToBudgetDb: Tests the function to add an amount to an existing budget.
- TestSpendBudgetDb: Tests the function to spend an amount from an existing budget.

*/

package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"personal-budget/models"
)

func TestDbConnect(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	defer conn.Close(ctx)
}

func TestGetAllBudgetsDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	budgets, err := GetAllBudgetsDb(ctx, conn)
	assert.NoError(t, err)
	assert.NotNil(t, budgets)
}

func TestGetSingleBudgetDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	budget, err := GetSingleBudgetDb(ctx, conn, 1)
	assert.NoError(t, err)
	assert.NotNil(t, budget)
}

func TestPostBudgetDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	budget := models.Budget{
		Name:   "Test Budget",
		Amount: 1000.0,
	}

	err = PostBudgetDb(ctx, conn, budget)
	assert.NoError(t, err)
}

func TestAddToBudgetDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	err = AddToBudgetDb(ctx, conn, 1, 500.0)
	assert.NoError(t, err)
}

func TestSpendBudgetDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	err = SpendBudgetDb(ctx, conn, 1, 500.0)
	assert.NoError(t, err)
}
