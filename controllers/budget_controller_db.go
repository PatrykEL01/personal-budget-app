package controllers

import (
    "context"
    "github.com/gin-gonic/gin"
    "net/http"
    "personal-budget/models"
    "personal-budget/services"
    "strconv"
)

// Ping godoc
// @Summary Ping the server
// @Description Ping the server to check if it's running
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func Ping(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}

// Health godoc
// @Summary Check the health of the server
// @Description Check if the server is running and healthy
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func Health(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Health OK",
    })
}

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

// ControllerGetAllBudgetsDB godoc
// @Summary Get all budgets
// @Description Get all budgets from the database
// @Tags Budgets
// @Produce json
// @Success 200 {array} models.Budget
// @Router /budgets [get]
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

// ControllerGetSingleBudgetDb godoc
// @Summary Get a single budget
// @Description Get a single budget by ID from the database
// @Tags Budgets
// @Produce json
// @Param id path int true "Budget ID"
// @Success 200 {object} models.Budget
// @Router /budgets/{id} [get]
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

// ControllerPostBudgetDb godoc
// @Summary Insert a new budget
// @Description Insert a new budget into the database
// @Tags Budgets
// @Accept x-www-form-urlencoded
// @Produce json
// @Param name formData string true "Budget Name"
// @Param amount formData number true "Budget Amount"
// @Success 200 {object} map[string]string
// @Router /budgets [post]
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

// ControllerAddToBudgetDb godoc
// @Summary Add to a budget
// @Description Add an amount to an existing budget
// @Tags Budgets
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id formData int true "Budget ID"
// @Param amount formData number true "Amount to add"
// @Success 200 {object} map[string]interface{}
// @Router /budgets/add [post]
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

// ControllerSpendBudgetDb godoc
// @Summary Spend from a budget
// @Description Spend an amount from an existing budget
// @Tags Budgets
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id formData int true "Budget ID"
// @Param amount formData number true "Amount to spend"
// @Success 200 {object} map[string]interface{}
// @Router /budgets/spend [post]
func ControllerSpendBudgetDb(c *gin.Context) {

    idParam := c.PostForm("id")
    amountParam := c.PostForm("amount")

    id, err := DbIdConversiontoInt(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Invalid ID format",
            "idParam": idParam,
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