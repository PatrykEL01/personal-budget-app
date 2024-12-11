package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-budget/services"
	"strconv"
)

// helper func
func idConversiontoInt(idParam string) (int, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// helper func
func amountConversiontoFloat(amountParam string) (float64, error) {
	amount, err := strconv.ParseFloat(amountParam, 64)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

// Ping is a simple ping/pong handler
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// health check
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Health OK",
	})
}

// GetBudget is a simple handler to return a budget
func GetBudgets(c *gin.Context) {
	budgets := services.GetAllBudgets()
	c.JSON(http.StatusOK, budgets)
}

// get budget by ID
func GetBudget(c *gin.Context) {
	idParam := c.Param("id") // Get ID from URL

	// id conversion to int
	id, err := idConversiontoInt(idParam)

	// call function getBudget from services
	budget, err := services.GetBudget(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, budget) // return budget
}

// AddToBudget is a simple handler to add to a budget
func AddToBudget(c *gin.Context) {
	idParam := c.PostForm("id")         // Get ID
	amountParam := c.PostForm("amount") // Get amount

	// id conversion to int
	id, err := idConversiontoInt(idParam)

	// amount conversion to float64
	amount, err := strconv.ParseFloat(amountParam, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid amount format",
		})
		return
	}

	// call function AddToBudget from services
	ok, err := services.AddToBudget(id, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Added to budget",
		"ammount": amount,
		"ok":      ok,
	})
}

// SpendBudget is a simple handler to spend from a budget
func SpendBudget(c *gin.Context) {
	idParam := c.PostForm("id")         // Get ID
	amountParam := c.PostForm("amount") // Get amount

	// id conversion to int
	id, err := idConversiontoInt(idParam)

	// amount conversion to float64
	amount, err := amountConversiontoFloat(amountParam)

	// call function SpendBudget from services
	ok, err := services.SpendBudget(id, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"message": "Spent from budget",
		"ammount": amount,
		"ok":      ok,
	})
}
