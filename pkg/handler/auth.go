package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_todo"
)

// signUp Создание пользователя
func (h *Handler) signUp(ctx *gin.Context) {
	var input rest_todo.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {

}
