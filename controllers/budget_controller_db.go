package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-budget/models"
	"personal-budget/services"
	"strconv"
)

// Helper func to convert ID to int
func DbIdConversiontoInt(idParam string) (int, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Helper func to convert Amount to float
func DbAmountConversiontoFloat(amountParam string) (float64, error) {
	amount, err := strconv.ParseFloat(amountParam, 64)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

// Controller to get all budgets from the database
func ControllerGetAllBudgetsDB(c *gin.Context) {
	// Create context
	ctx := context.Background()

	conn, err := services.DbConnect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
		return
	}
	defer conn.Close(ctx)

	// Fetch budgets
	budgets, err := services.GetAllBudgetsDb(ctx, conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return budgets
	c.JSON(http.StatusOK, budgets)
}

// Controller to insert a budget into the database
func ControllerPostBudgetDb(c *gin.Context) {
	name := c.PostForm("name")
	amount := c.PostForm("amount")
	id := c.PostForm("id")

	// Amount conversion to float
	amountFloat, err := DbAmountConversiontoFloat(amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid amount format",
		})
		return
	}

	// ID conversion to int
	idInt, err := DbIdConversiontoInt(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	// Create context
	ctx := context.Background()

	conn, err := services.DbConnect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
		return
	}
	defer conn.Close(ctx)

	budget := models.Budget{
		ID:     idInt,
		Name:   name,
		Amount: amountFloat,
	}

	err = services.PostBudgetDb(ctx, conn, budget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Budget inserted successfully",
	})
}
