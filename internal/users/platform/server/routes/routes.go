package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/rubenbupe/go-auth-server/internal/shared/platform/di"
	handlers "github.com/rubenbupe/go-auth-server/internal/shared/platform/server/handler"
)

func Register(router *gin.RouterGroup) {

	diContainer := di.Instance()

	postController := diContainer.Container.Get("users.infrastructure.controller.create").(handlers.Handler)
	router.POST("/", postController)

	getController := diContainer.Container.Get("users.infrastructure.controller.get").(handlers.Handler)
	router.GET("/:id", getController)

}
