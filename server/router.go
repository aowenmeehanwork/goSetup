package server

import (

	"github.com/gin-gonic/gin"

	"github.com/aowenmeehanwork/goSetup/controllers"

)

func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	todo := new(controllers.TodoController)

	router.GET("/todo", todo.GetExample)
	router.POST("/todo/:number", todo.PostExample)
	router.PUT("/todo/:number", todo.PutExample)
	router.DELETE("/todo/:number", todo.DeleteExample)

	return router
}