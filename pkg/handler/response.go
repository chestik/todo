package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errors struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, msg string) {
	logrus.Error(msg)
	ctx.AbortWithStatusJSON(statusCode, errors{msg})
}
