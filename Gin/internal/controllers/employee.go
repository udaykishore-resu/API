package controllers

import (
	"gin-api/internal/models"
	"gin-api/internal/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	storage *storage.EmployeeStorage
}

func NewEmployeeController(storage *storage.EmployeeStorage) *EmployeeController {
	return &EmployeeController{storage: storage}
}

// Common middleware for ID param handling
func IDParamMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid employee ID"})
			return
		}
		c.Set("employeeID", id)
		c.Next()
	}
}

// Common handler for employee operations requiring ID
func (ctrl *EmployeeController) withEmployeeID(handler func(c *gin.Context, id int)) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get("employeeID")
		if !exists {
			c.AbortWithStatusJSON(500, gin.H{"error": "missing employee ID"})
			return
		}
		handler(c, id.(int))
	}
}

func (ctrl *EmployeeController) GetEmployees(c *gin.Context) {
	employees, err := ctrl.storage.GetEmployees()
	if handleError(c, err) {
		return
	}
	c.JSON(200, employees)
}

func (ctrl *EmployeeController) GetEmployee(c *gin.Context) {
	ctrl.withEmployeeID(func(c *gin.Context, id int) {
		employee, err := ctrl.storage.GetEmployee(id)
		if handleError(c, err) {
			return
		}
		c.JSON(200, employee)
	})(c)
}

func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request body"})
		return
	}

	if err := ctrl.storage.CreateEmployee(&emp); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, emp)
}

func (ctrl *EmployeeController) UpdateEmployee(c *gin.Context) {
	ctrl.withEmployeeID(func(c *gin.Context, id int) {
		var emp models.Employee
		if err := c.ShouldBindJSON(&emp); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid request body"})
			return
		}

		if err := ctrl.storage.UpdateEmployee(id, &emp); err != nil {
			handleError(c, err)
			return
		}

		emp.ID = id
		c.JSON(200, emp)
	})(c)
}

func (ctrl *EmployeeController) DeleteEmployee(c *gin.Context) {
	ctrl.withEmployeeID(func(c *gin.Context, id int) {
		if err := ctrl.storage.DeleteEmployee(id); err != nil {
			handleError(c, err)
			return
		}
		c.Status(204)
	})(c)
}

func handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	if err == storage.ErrEmployeeNotFound {
		c.AbortWithStatusJSON(404, gin.H{"error": err.Error()})
	} else {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}
	return true
}
