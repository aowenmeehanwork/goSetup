package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoController struct{}

func (TodoController TodoController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (Todocontroller TodoController) GetExample(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message":"This is the response from the get method",
	})
}

func (Todocontroller TodoController) PostExample(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message":"This is the response from the put method",
	})
}

func (Todocontroller TodoController) PutExample(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message":"This is the response from the put method",
	})
}

func (Todocontroller TodoController) DeleteExample(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message":"This is the response from the delete method",
	})
}