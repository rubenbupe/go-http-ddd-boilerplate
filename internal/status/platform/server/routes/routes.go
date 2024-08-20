package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/rubenbupe/go-auth-server/internal/shared/platform/di"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/server/handler"
)

func Register(router *gin.RouterGroup) {

	diContainer := di.Instance()

	getController := diContainer.Container.Get("status.infrastructure.controller.check").(handlers.Handler)
	router.GET("/", getController)
}
