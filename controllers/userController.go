package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sql-api/services"
)

func RunQuery() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		message, rowsAffected, err := services.QueryService(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":      err.Error(),
				"rowsAffected": rowsAffected,
				"err":          err,
			})
			return
		}

		if rowsAffected == 0 {
			ctx.JSON(200, gin.H{
				"message":      "Action Successful",
				"rowsAffected": rowsAffected,
				"err":          err,
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message":      string(message),
			"rowsAffected": rowsAffected,
			"err":          err,
		})
		return

	}
}
