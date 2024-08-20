package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rubenbupe/go-auth-server/internal/users/application/create"
	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
	"github.com/rubenbupe/go-auth-server/kit/command"
)

type createRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, create.NewUserCommand(
			req.ID,
			req.Name,
		))

		if err != nil {
			switch {
			case errors.Is(err, usersdomain.ErrInvalidUserID),
				errors.Is(err, usersdomain.ErrEmptyUserName),
				errors.Is(err, usersdomain.ErrInvalidUserID),
				errors.Is(err, usersdomain.ErrUserAlreadyExists):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return

			default:
				println(err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
