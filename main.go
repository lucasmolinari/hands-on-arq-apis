package main

import (
	"swaggo-demo/handlers"

	_ "swaggo-demo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Task API
// @version         1.0
// @description     API de exemplo para demonstrar geracao de documentacao com Swaggo.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Suporte
// @contact.email  suporte@example.com

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", handlers.ListTasks)
			tasks.GET("/:id", handlers.GetTask)
			tasks.POST("", handlers.CreateTask)
			tasks.PUT("/:id", handlers.UpdateTask)
			tasks.DELETE("/:id", handlers.DeleteTask)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
