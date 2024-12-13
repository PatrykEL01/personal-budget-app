package services

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"personal-budget/models"
)

var dbUrl string

func TestMain(m *testing.M) {
	err := LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	dbUrl = os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	code := m.Run()

	os.Exit(code)
}

func TestDbConnect(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx, dbUrl)
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	defer conn.Close(ctx)
}

func TestGetAllBudgetsDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx, dbUrl)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	budgets, err := GetAllBudgetsDb(ctx, conn)
	assert.NoError(t, err)
	assert.NotNil(t, budgets)
}

func TestGetSingleBudgetDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx, dbUrl)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	budget, err := GetSingleBudgetDb(ctx, conn, 1)
	assert.NoError(t, err)
	assert.NotNil(t, budget)
}

func TestPostBudgetDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx, dbUrl)
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
	conn, err := DbConnect(ctx, dbUrl)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	err = AddToBudgetDb(ctx, conn, 1, 500.0)
	assert.NoError(t, err)
}

func TestSpendBudgetDb(t *testing.T) {
	ctx := context.Background()
	conn, err := DbConnect(ctx, dbUrl)
	assert.NoError(t, err)
	defer conn.Close(ctx)

	err = SpendBudgetDb(ctx, conn, 1, 500.0)
	assert.NoError(t, err)
}
