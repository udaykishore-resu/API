package routes

import (
	"gin-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl *controllers.EmployeeController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		employees := api.Group("/employees")
		{
			employees.GET("", ctrl.GetEmployees)
			employees.GET("/:id", ctrl.GetEmployee)
			employees.POST("", ctrl.CreateEmployee)
			employees.PUT("/:id", ctrl.UpdateEmployee)
			employees.DELETE("/:id", ctrl.DeleteEmployee)
		}
	}

	return router
}
