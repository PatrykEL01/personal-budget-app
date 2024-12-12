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

func ControllerGetSingleBudgetDb(c *gin.Context) {

	ctx := context.Background()

	conn, err := services.DbConnect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
		return
	}
	defer conn.Close(ctx)

	idParam := c.Param("id")
	id, err := DbIdConversiontoInt(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	budget, err := services.GetSingleBudgetDb(ctx, conn, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, budget)

}

// Controller to insert a budget into the database
func ControllerPostBudgetDb(c *gin.Context) {
	name := c.PostForm("name")
	amount := c.PostForm("amount")

	// Amount conversion to float
	amountFloat, err := DbAmountConversiontoFloat(amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid amount format",
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

func ControllerAddToBudgetDb(c *gin.Context) {
	idParam := c.PostForm("id")
	amountParam := c.PostForm("amount")

	id, err := DbIdConversiontoInt(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	amount, err := DbAmountConversiontoFloat(amountParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid amount format",
		})
		return
	}

	ctx := context.Background()

	conn, err := services.DbConnect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
		return
	}
	defer conn.Close(ctx)

	err = services.AddToBudgetDb(ctx, conn, id, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Added to budget successfully",
		"amount":  amount,
	})
}

func ControllerSpendBudgetDb(c *gin.Context) {

	idParam := c.PostForm("id")
	amountParam := c.PostForm("amount")

	id, err := DbIdConversiontoInt(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
			"idParam":   idParam,
		})
		return
	}

	amount, err := DbAmountConversiontoFloat(amountParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid amount format",
		})
		return
	}

	ctx := context.Background()

	conn, err := services.DbConnect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
		return
	}
	defer conn.Close(ctx)

	err = services.SpendBudgetDb(ctx, conn, id, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Spent from budget successfully",
		"amount":  amount,
	})

}
