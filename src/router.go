package src

import (
	"gin_example/src/model"
	"gin_example/src/service"
	"gin_example/src/service/repo"

	"github.com/gin-gonic/gin"
)

func MainRouter() *gin.Engine {
	var router = gin.Default()

	// Initialize repository and service
	userRepo := &repo.UserRepo{
		Repository: &repo.Repository[model.User]{},
	}
	userService := &service.UserService{
		Repo: *userRepo,
	}

	userRouter := router.Group("")
	{
		// Users routes
		users := userRouter.Group("/user")
		{
			users.POST("/", userService.Create)
			users.GET("/:id", userService.GetByID)
			users.GET("/find/:name", userService.FindByName)
			users.PUT("/:id", userService.Update)
			users.DELETE("/:id", userService.Delete)
		}
	}
	return router
}
