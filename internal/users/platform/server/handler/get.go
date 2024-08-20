package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rubenbupe/go-auth-server/internal/users/application/get"
	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
	"github.com/rubenbupe/go-auth-server/kit/query"
)

func GetHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id string = ctx.Param("id")

		u, err := queryBus.Ask(ctx, get.NewUserQuery(
			id,
		))

		if err != nil {
			switch {
			case
				errors.Is(err, usersdomain.ErrInvalidUserID):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return

			default:
				println(err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		}

		user, ok := u.(*usersdomain.User)

		if user == nil {
			ctx.JSON(http.StatusNotFound, gin.H{})
			return
		}

		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"id":   user.Id.String(),
			"name": user.Name.String(),
		})
	}
}
