package routes

import (
	"github.com/gin-gonic/gin"
	"personal-budget/controllers"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
    _ "personal-budget/docs"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.GET("/health", controllers.Health)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// db routes
	r.GET("/db/budgets", controllers.ControllerGetAllBudgetsDB)
	r.GET("/db/budgets/:id", controllers.ControllerGetSingleBudgetDb)
	r.POST("/db/budgets", controllers.ControllerPostBudgetDb)
	r.PUT("/db/budgets/add", controllers.ControllerAddToBudgetDb)
	r.PUT("/db/budgets/spend", controllers.ControllerSpendBudgetDb)

}
