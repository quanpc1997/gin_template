package src

import (
	"gin_example/configure"
	"gin_example/src/model"
	"gin_example/src/service"
	"gin_example/src/service/repo"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MainRouter() *gin.Engine {
	var router = gin.Default()

	db := configure.InitDB()

	baseRepo := &repo.Repository[model.User]{
		DB: db,
	}

	// Initialize repository and service
	var userRepo repo.IUser = &repo.UserRepo{
		Repository: baseRepo,
	}
	userService := &service.UserService{
		Repo: userRepo,
	}

	monitorRepo := &repo.MonitorRepo{}
	monitorService := &service.MonitorService{
		Repo: *monitorRepo,
	}

	mainRouter := router.Group("")
	{
		// Users routes
		users := mainRouter.Group("/user")
		{
			users.POST("/", userService.Create)
			users.GET("/:id", userService.GetByID)
			users.GET("/find/:name", userService.FindByName)
			users.PUT("/:id", userService.Update)
			users.DELETE("/:id", userService.Delete)
		}
		monitor := mainRouter.Group("/monitor")
		{
			monitor.GET("/health", monitorService.HealthCheck)
		}
	}

	zap.L().Info("======== HAHA =========")
	return router
}
