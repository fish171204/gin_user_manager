package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler: handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", ur.handler.GetAllUsers)
		users.POST("", ur.handler.CreateUsers)
		users.GET("/:uuid", ur.handler.GetUserByUUID)
		users.GET("/:uuid", ur.handler.UpdateUser)
		users.GET("/:uuid", ur.handler.DeleteUser)
	}
}
