package routes

import (
	"github.com/gin-gonic/gin"
	"personal-budget/controllers"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.GET("/budgets", controllers.GetBudgets)
	r.GET("/budgets/:id", controllers.GetBudget)
	r.POST("/budgets/spend", controllers.SpendBudget)
	r.POST("/budgets/add", controllers.AddToBudget)
	r.GET("/health", controllers.Health)

	// db routes
	r.GET("/db/budgets", controllers.ControllerGetAllBudgetsDB)
	r.GET("/db/budgets/:id", controllers.ControllerGetSingleBudgetDb)
	r.POST("/db/budgets", controllers.ControllerPostBudgetDb)
	r.PUT("/db/budgets/add", controllers.ControllerAddToBudgetDb)

}
