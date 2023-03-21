package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/gonzamedrano09/chat/pkg/adapter/controller"
	"github.com/gonzamedrano09/chat/pkg/entity"
	"github.com/gonzamedrano09/chat/pkg/usecase/service"
	"net/http"
	"strconv"
	"strings"
)

func Authorization(mdc controller.MiddlewareDatabaseConnection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader, ok := ctx.Request.Header["Authorization"]
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		auth := strings.Split(authHeader[0], " ")
		if auth[0] != "Bearer" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		js := service.NewJWTService()
		username, ok := js.CheckJWT(auth[1])
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user entity.User
		if result := mdc.Database.Find(&user, "username = ?", username); result.RowsAffected < 1 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		method := ctx.Request.Method
		idString := ctx.Param("id")
		if (method == "PUT" || method == "DELETE") && idString != "" {
			if id, err := strconv.Atoi(idString); err != nil {
				ctx.AbortWithStatus(http.StatusBadRequest)
				return
			} else {
				if uint(id) != user.ID {
					ctx.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			}
		}
		ctx.Set("user", &user)
	}
}
